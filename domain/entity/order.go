// 订单

package entity

import (
	"time"
)

type Order struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	OrderId   string     `gorm:"size:100;not null;unique" json:"order_id"` // 订单ID
	AgentCode string     `gorm:"size:100;not null;" json:"agent_code"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
