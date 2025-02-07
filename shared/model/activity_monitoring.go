package model

import "time"

type ActivityMonitor struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	UserName     string    `gorm:"size:255;not null" json:"user_name"` // Name of the user performing the activity
	Category     string    `gorm:"size:50;not null" json:"category"`   // Type of activity (e.g., 'alarm', 'user', 'pintu air')
	ActivityTime time.Time `gorm:"not null" `                          // Timestamp for when the activity occurred
	Description  string    `gorm:"type:text;not null"`                 // Description of the activity
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
