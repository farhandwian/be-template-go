package model

type RainfallHourlyCalculation struct {
	RainPostID             int64
	AvailabilityPercentage float64
	MorningData            float64
	AfternoonData          float64
	EveningData            float64
	DawnData               float64
}
