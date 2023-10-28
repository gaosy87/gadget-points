package interfaces

import (
	"gadget-points/application"
	"gadget-points/domain/entity"
	"gadget-points/infrastructure/auth"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"time"
)

type Order struct {
	orderApp    application.OrderAppInterface
	productApp  application.ProductAppInterface
	agentApp    application.AgentAppInterface
	activityApp application.ActivityAppInterface
	tk          auth.TokenInterface
	rd          auth.AuthInterface
}

// NewOrder Order constructor
func NewOrder(fApp application.OrderAppInterface,
	productApp application.ProductAppInterface,
	agentApp application.AgentAppInterface,
	activityApp application.ActivityAppInterface,
	rd auth.AuthInterface,
	tk auth.TokenInterface) *Order {
	return &Order{
		orderApp:    fApp,
		agentApp:    agentApp,
		activityApp: activityApp,
		rd:          rd,
		tk:          tk,
	}
}

// OrderDetails 产品详情
type OrderDetails struct {
	ProductId uint64 `json:"product_id"`                            // 产品ID
	Price     int32  `gorm:"size:100;not null;unique" json:"price"` // 原始价格
}

// CreateOrderReq 创建订单请求
type CreateOrderReq struct {
	AgentCode    string          `json:"agent_code"`    // 代理ID
	OrderDetails []*OrderDetails `json:"order_details"` // 多个产品详情
}

// CreateOrderRsp 创建订单响应
type CreateOrderRsp struct {
	Code int32  `json:"code"` // 错误码
	Msg  string `json:"msg"`  // 错误消息
	Data struct {
		OrderId string `json:"order_id"` // 订单ID
	} `json:"data"` // 数据
}

// CreateOrder 创建订单
func (t *Order) CreateOrder(c *gin.Context) {
	// 返回成功
	rsp := &CreateOrderRsp{
		Code: 0, // 0:成功
		Msg:  "",
	}

	// 解析body里的json
	req := &CreateOrderReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		rsp.Code = 1 // 1:失败
		c.JSON(http.StatusOK, rsp)
		return
	}

	// 校验订单
	validateErr := req.Validate("")
	if len(validateErr) > 0 {
		rsp.Code = 1 // 1:失败
		c.JSON(http.StatusOK, rsp)
		return
	}

	// 获取代理信息
	_, err := t.agentApp.GetAgent(req.AgentCode)
	if err != nil {
		rsp.Code = 1 // 1:失败
		c.JSON(http.StatusOK, rsp)
		return
	}

	// TODO: 以下暂时忽略了对事务的处理、忽略缓存的处理

	// 创建订单
	order := &entity.Order{}
	order.AgentCode = req.AgentCode
	order.OrderId = utils.GenOrderId() // 通过分布式生成ID算法生成产品ID
	err = t.orderApp.CreateOrder(order)
	if err != nil {
		rsp.Code = 1 // 1:失败
		c.JSON(http.StatusOK, rsp)
		return
	}

	for _, item := range req.OrderDetails {
		// 获取产品ID
		product, err := t.productApp.GetProduct(item.ProductId)
		if err != nil {
			continue
		}

		originalPrice := product.OriginalPrice // 原始价格
		var activityId uint64                  // 活动ID
		var finalPrice int32                   // 最终价格

		// 通过ID获取优惠活动
		activity, _ := t.activityApp.GetActivityByID(item.ProductId)
		// 计算优惠价格
		if activity != nil {
			if activity.Type == 1 { // 解析折扣活动规则（销售点通常有单独的优惠，但全国性的营销活动也会进行）
				finalPrice = activity.ParseDiscountRule(originalPrice)
			}

			activityId = activity.ID
		}

		// 创建订单详情
		orderDetails := &entity.OrderDetails{}
		orderDetails.ProductId = item.ProductId
		orderDetails.ActivityId = activityId
		orderDetails.OriginalPrice = originalPrice
		orderDetails.FinalPrice = finalPrice
		orderDetails.AgentCode = req.AgentCode
		err = t.orderApp.CreateOrderDetails(orderDetails)
		if err != nil {
			continue
		}
	}

	// 通过活动类型获取优惠活动（购买10个配件即可获得1个免费配件）
	activity, _ := t.activityApp.GetActivityByType(2) // 这里的2为“满10个送1个配件”活动，可以优化成枚举类型
	if activity != nil {
		giveProducts, isTrigger := activity.ParseFullGiveRule(len(req.OrderDetails))
		if isTrigger { // 触发送免费配件
			for _, item := range giveProducts {
				// 获取产品ID
				product, err := t.productApp.GetProduct(item.ProductId)
				if err != nil {
					continue
				}

				// 创建订单详情
				orderDetails := &entity.OrderDetails{}
				orderDetails.ProductId = item.ProductId
				orderDetails.ActivityId = item.ActivityId
				orderDetails.OriginalPrice = decimal.NewFromInt32(product.Price)
				orderDetails.FinalPrice = decimal.Zero // 因为是送的免费产品，所以这里填0
				orderDetails.AgentCode = req.AgentCode
				err = t.orderApp.CreateOrderDetails(orderDetails)
				if err != nil {
					continue
				}
			}
		}
	}

	rsp.OrderId = order.OrderId

	c.JSON(http.StatusOK, rsp)
}
