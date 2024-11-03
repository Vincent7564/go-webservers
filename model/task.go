package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"type:longtext"`
	Description string         `json:"description" gorm:"type:longtext"`
	IsActive    bool           `json:"Is_active" gorm:"column:is_active"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP"` // No fractional seconds
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime"`
}
