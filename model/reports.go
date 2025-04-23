package model

import "time"

type Report struct {
	Id         uint       `gorm:"primaryKey"`
	Title      string     `gorm:"type:text;not null"`
	Body       string     `gorm:"type:text;not null"`
	ReportType ReportType `gorm:"type:ReportType;not null"`
	Rating     int        `gorm:"type:integer"`
	UserID     *uint      `gorm:"index"`
	CreatedAt  time.Time  `gorm:"not null;default:now()"`
	UpdatedAt  time.Time  `gorm:"not null;default:now()"`
}
