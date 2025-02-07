package model

type Project struct {
	ID     string        `json:"id" gorm:"primaryKey"`
	Name   string        `json:"name"`
	Budget float64       `json:"budget"`
	Status ProjectStatus `json:"status"`
}

type ProjectStatus string

const (
	ProjectStatusNotStarted ProjectStatus = "not_started"
	ProjectStatusOnGoing    ProjectStatus = "on_going"
	ProjectStatusFinish     ProjectStatus = "finish"
)
