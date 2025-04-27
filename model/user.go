package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id          uint      `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	DeletedAt   *time.Time
	Email       string
	FirebaseUid string `gorm:"uniqueIndex"`
}
