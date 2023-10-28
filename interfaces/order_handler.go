package interfaces

import (
	"gadget-points/application"
	"gadget-points/domain/entity"
	"gadget-points/infrastructure/auth"
	"github.com/gin-gonic/gin"
	"net/http"
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

// CreateOrder 创建订单
func (t *Order) CreateOrder(c *gin.Context) {
	order := &entity.Order{}
	if err := c.ShouldBindJSON(order); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	validateErr := order.Validate("")
	if len(validateErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, validateErr)
		return
	}

	// 获取代理信息
	_, err := t.agentApp.GetAgent(order.AgentCode)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, validateErr)
		return
	}

	// 获取产品ID
	product, err := t.productApp.GetProduct(order.ProductId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, validateErr)
		return
	}

	originalPrice := product.OriginalPrice

	// 获取优惠活动
	activity, _ := t.activityApp.GetActivity(order.ProductId)

	// 计算优惠价格
	var finalPrice int32
	var activityId uint64
	if activity != nil {
		finalPrice = activity.CalcFinalPrice(originalPrice)
		activityId = activity.ID
	}

	// 创建订单
	order.ActivityId = activityId
	order.OriginalPrice = originalPrice
	order.FinalPrice = finalPrice
	err := t.orderApp.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}
