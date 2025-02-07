package model

import "time"

type HydrologyRainDaily struct {
	RainPostID int     `gorm:"uniqueIndex:rain_post_sampling"`
	Sampling   string  `gorm:"uniqueIndex:rain_post_sampling"`
	Count      int     // count24
	Rain       float32 // rain24
	Manual     *float64
	UpdatedAt  time.Time
}
