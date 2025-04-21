package model

import (
	"time"
)

type Chat struct {
	ID        uint          `gorm:"primaryKey"`
	UserID    uint          `gorm:"not null"`
	Title     string        `gorm:"type:text"`
	CreatedAt time.Time     `gorm:"not null;default:now()"`
	UpdatedAt time.Time     `gorm:"not null;default:now()"`
	DeletedAt *time.Time    `gorm:"index"`
	Messages  []ChatMessage `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE"`
}

type ChatMessage struct {
	ID        uint      `gorm:"primaryKey"`
	ChatID    uint      `gorm:"not null;index"`
	Sender    string    `gorm:"type:text;not null;check:sender IN ('USER','AI')"`
	Message   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}
