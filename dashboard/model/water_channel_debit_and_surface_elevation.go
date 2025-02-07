package model

type WaterChannelDoorDebitAndWaterSurfaceElevation struct {
	WaterChannelDoorID    int     `json:"water_channel_id"`
	Debit                 float64 `json:"debit"`
	WaterSurfaceElevation float64 `json:"water_surface_elevation"`
}
