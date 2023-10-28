// 公告

package entity

import (
	"strings"
	"time"
)

type Notice struct {
	ID          uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Title       string     `gorm:"size:100;not null;unique" json:"title"`
	Description string     `gorm:"text;not null;" json:"description"`
	ImgUrl      string     `gorm:"size:2048;not null;" json:"img_url"`
	Link        string     `gorm:"size:2048;not null;" json:"link"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
