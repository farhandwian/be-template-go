package model

import "time"

type Config struct {
	ID        string
	Name      string
	Value     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
