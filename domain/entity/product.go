// 产品

package entity

import (
	"strings"
	"time"
)

type Product struct {
	ID          uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Title       string     `gorm:"size:100;not null;unique" json:"title"`
	Description string     `gorm:"text;not null;" json:"description"`
	Price       int32      `gorm:"size:100;not null;unique" json:"price"`
	AgentCode   string     `gorm:"size:100;not null;" json:"agent_code"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
