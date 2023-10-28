// 订单详情

package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderDetails struct {
	ID            uint64          `gorm:"primary_key;auto_increment" json:"id"`
	ProductId     uint64          `gorm:"size:100;not null;unique" json:"product_id"`     // 产品ID
	ActivityId    uint64          `gorm:"size:100;not null;unique" json:"activity_id"`    // 活动ID
	OriginalPrice decimal.Decimal `gorm:"size:100;not null;unique" json:"original_price"` // 原始价格
	FinalPrice    decimal.Decimal `gorm:"size:100;not null;unique" json:"final_price"`    // 最终价格
	AgentCode     string          `gorm:"size:100;not null;" json:"agent_code"`
	CreatedAt     time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     *time.Time      `json:"deleted_at,omitempty"`
}
