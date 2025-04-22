package model

import (
	"time"
)

type ChatMessage struct {
	Id        uint      `gorm:"primaryKey"`
	TaskId    uint      `gorm:"not null;index"`
	Sender    string    `gorm:"type:text;not null;check:sender IN ('USER','AI')"`
	Message   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}
