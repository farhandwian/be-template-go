package model

type SpeedtestResponse struct {
	Message string          `json:"message"`
	Data    SpeedtestResult `json:"data"`
}

type SpeedtestResult struct {
	Ping     float64 `json:"ping"`
	Download float64 `json:"download"`
	Upload   float64 `json:"upload"`
}
