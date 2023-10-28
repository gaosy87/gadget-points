// 代理

package entity

import (
	"time"
)

type Agent struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	AgentCode string     `gorm:"size:100;not null;" json:"agent_code"`
	AgentName string     `gorm:"size:100;not null;" json:"agent_name"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (t *Agent) Prepare() {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}
