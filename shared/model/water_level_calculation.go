package model

type WaterLevelCalculation struct {
	ID                  int64
	Name                string
	WaterLevelTelemetry float64
	WaterLevelManual    *float64
	LatestUpdate        string
}
