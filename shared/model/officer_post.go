package model

import "time"

type PostOfficer struct {
	ID          string `gorm:"primaryKey;size:36"`
	Name        string `gorm:"size:255"`
	Username    string `gorm:"size:100;uniqueIndex:idx_username_post,length:100"`
	Post        string `gorm:"size:100;uniqueIndex:idx_username_post,length:100"`
	PhoneNumber string
	Address     string
	RtRw        string
	Village     string
	Subdistrict string
	District    string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
