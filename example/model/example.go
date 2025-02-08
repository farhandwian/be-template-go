package model

import "time"

type Example struct {
	ExampleString string `json:"example_string"`
	ExampleInt    int    `json:"example_int"`
	ExampleBool   bool   `json:"example_bool"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
