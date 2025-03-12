package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Id          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Title       string
	Description string
	Completed   bool
	Status      TaskStatus `gorm:"type:TaskStatus"`
	UserId      uint
	Images      []AppFile              `gorm:"foreignKey:TaskId"`
	Mesh        *AppFile               `gorm:"foreignKey:TaskId"`
	Metadata    map[string]interface{} `gorm:"type:json" json:"metadata"`
}
