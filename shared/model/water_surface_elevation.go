package model

type WaterSurfaceElevationList struct {
	WaterChannelDoorID    int     `json:"water_channel_id"`
	WaterSurfaceElevation float64 `json:"water_surface_elevation"`
	Status                bool    `json:"status"`
}
