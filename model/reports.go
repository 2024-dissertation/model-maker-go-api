package model

import "time"

// ReportType defines the type of report sent by a user.
type ReportType string

const (
	ReportTypeBug      ReportType = "BUG"
	ReportTypeFeedback ReportType = "FEEDBACK"
)

// Report represents a user-submitted report (bug or feedback).
type Report struct {
	ID         uint       `gorm:"primaryKey"`
	Title      string     `gorm:"type:text;not null"`
	Body       string     `gorm:"type:text;not null"`
	ReportType ReportType `gorm:"type:reporttype;not null"`
	Rating     int        `gorm:"type:integer"`
	UserID     *uint      `gorm:"index"` // Optional link to user
	CreatedAt  time.Time  `gorm:"not null;default:now()"`
	UpdatedAt  time.Time  `gorm:"not null;default:now()"`
}
