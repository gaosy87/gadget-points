// 活动

package entity

import (
	"encoding/json"
	"time"
)

type Activity struct {
	ID          uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Type        int32      `gorm:"size:100;not null;unique" json:"type"`  // 活动类型, 1:打折活动
	Title       string     `gorm:"size:100;not null;unique" json:"title"` // 活动标题
	Description string     `gorm:"text;not null;" json:"description"`     // 活动描述
	Price       int32      `gorm:"size:100;not null;unique" json:"price"`
	Rule        string     `gorm:"text;not null;" json:"rule"`                 // 活动规则(json)
	ProductId   uint64     `gorm:"size:100;not null;unique" json:"product_id"` // 产品ID
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// DiscountRule 折扣规则
type DiscountRule struct {
	DiscountRate int32 `json:"discount_rate"` // 折扣率
}

// CalcFinalPrice 计算最终价格
func (t *Activity) CalcFinalPrice(originalPrice int32) int32 {
	if t.Type == 1 { // 打折活动
		discountRule := &DiscountRule{}
		_ = json.Unmarshal([]byte(t.Rule), discountRule)
		return originalPrice / discountRule.DiscountRate
	}

	return originalPrice
}
