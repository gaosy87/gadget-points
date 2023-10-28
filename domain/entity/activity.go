// 活动

package entity

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"time"
)

type Activity struct {
	ID          uint64          `gorm:"primary_key;auto_increment" json:"id"`
	Type        int32           `gorm:"size:100;not null;unique" json:"type"`  // 活动类型，1:打折活动，2:满10个送1个配件
	Title       string          `gorm:"size:100;not null;unique" json:"title"` // 活动标题
	Description string          `gorm:"text;not null;" json:"description"`     // 活动描述
	Price       decimal.Decimal `gorm:"size:100;not null;unique" json:"price"`
	Rule        string          `gorm:"text;not null;" json:"rule"`                 // 活动规则(json)
	ProductId   uint64          `gorm:"size:100;not null;unique" json:"product_id"` // 产品ID
	CreatedAt   time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time      `json:"deleted_at,omitempty"`
}

// DiscountRule 折扣活动规则
type DiscountRule struct {
	DiscountRate int32 `json:"discount_rate"` // 折扣率
}

// GiveProduct 送的产品
type GiveProduct struct {
	ProductId    uint64 `json:"product_id"`    // 送的产品ID
	ProductCount int32  `json:"product_count"` // 送的产品数量
}

// FullGiveRule 满送活动规则
type FullGiveRule struct {
	FullCount    int32          `json:"full_count"`   // 满的数量
	GiveProducts []*GiveProduct `json:"give_product"` // 送的产品
}

// ParseDiscountRule 解析折扣活动规则（销售点通常有单独的优惠，但全国性的营销活动也会进行）
func (t *Activity) ParseDiscountRule(originalPrice int32) int32 {
	discountRule := &DiscountRule{}
	_ = json.Unmarshal([]byte(t.Rule), discountRule)
	return originalPrice / discountRule.DiscountRate
}

// ParseFullGiveRule 解析满送活动规则（购买10个配件即可获得1个免费配件）
// fullCount: 满的数量
func (t *Activity) ParseFullGiveRule(fullCount int32) (giveProducts []*GiveProduct, isTrigger bool) {
	rule := &FullGiveRule{}
	_ = json.Unmarshal([]byte(t.Rule), rule)
	if fullCount < rule.FullCount {
		return nil, false
	}

	return rule.GiveProducts, true
}
