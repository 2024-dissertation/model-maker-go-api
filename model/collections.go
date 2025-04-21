package model

import (
	"time"
)

type Collection struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"type:text;not null"`
	UserID    *uint      `gorm:"not null;index"`
	CreatedAt time.Time  `gorm:"not null;default:now()"`
	UpdatedAt time.Time  `gorm:"not null;default:now()"`
	DeletedAt *time.Time `gorm:"index"`
	Tasks     []Task     `gorm:"many2many:collection_tasks;constraint:OnDelete:CASCADE"`
}

type CollectionTask struct {
	CollectionID uint `gorm:"primaryKey"`
	TaskID       uint `gorm:"primaryKey"`
}
