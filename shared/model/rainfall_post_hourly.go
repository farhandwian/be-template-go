package model

import "time"

type HydrologyRainHourly struct {
	RainPostID int
	Hour       int
	Count      int
	Rain       float32
	Sampling   string
	UpdatedAt  time.Time
}
