package model

import "time"

type Employee struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Role       string         `json:"role"`
	JoinedDate time.Time      `json:"joined_date"`
	Status     EmployeeStatus `json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

type EmployeeStatus string

const (
	EmployeeStatusNonActive EmployeeStatus = "non_active"
	EmployeeStatusActive    EmployeeStatus = "active"
)
