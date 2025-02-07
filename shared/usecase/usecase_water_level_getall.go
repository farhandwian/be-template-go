package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"time"
)

type GetListWaterLevelReq struct {
}
type GetListWaterLevelResp struct {
	Total       int            `json:"total"`
	WaterLevels WaterLevelData `json:"rain_falls"`
}

type WaterLevelData struct {
	Type     string              `json:"type"`
	Features []WaterLevelFeature `json:"features"`
}
type WaterLevelFeature struct {
	Type       string               `json:"type"`
	Geometry   Geometry             `json:"geometry"`
	Properties WaterLevelProperties `json:"properties"`
}

type WaterLevelProperties struct {
	ID         string  `json:"id"`
	ExternalID int64   `json:"external_id"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Manual     float64 `json:"manual"`
	Telemetry  float64 `json:"telemetry"`
}

type GetListWaterLevelUseCase = core.ActionHandler[GetListWaterLevelReq, GetListWaterLevelResp]

func ImplGetListWaterLevelUseCase(
	getListWaterLevel gateway.GetWaterLevelPostListGateway,
	getListWaterLevelPostCalculationLatestManual gateway.GetListLatestManualWaterPostGateway,
	getListWaterLevelPostCalculationLatestTelemetry gateway.GetListLatestTelemetryWaterPostGateway) GetListWaterLevelUseCase {
	return func(ctx context.Context, req GetListWaterLevelReq) (*GetListWaterLevelResp, error) {
		waterlevelData, err := getListWaterLevel(ctx, gateway.GetWaterLevelPostListReq{})
		if err != nil {
			return nil, err
		}

		// waterLevelPostIDs := make([]int64, 0)
		// for _, waterLevelPost := range waterlevelData.WaterLevelPost {
		// 	waterLevelPostIDs = append(waterLevelPostIDs, waterLevelPost.ExternalID)
		// }

		waterLevelLatestManual, err := getListWaterLevelPostCalculationLatestManual(ctx, gateway.GetListLatestManualWaterPostReq{
			Date: time.Now().Format(time.DateOnly)})
		if err != nil {
			return nil, err
		}

		waterLevelLatestManualMap := make(map[int64]float64)
		for _, wlManual := range waterLevelLatestManual.LatestManualWaterPost {
			waterLevelLatestManualMap[wlManual.WaterLevelPostID] = wlManual.TMA
		}

		waterLevelLatestTelemetry, err := getListWaterLevelPostCalculationLatestTelemetry(ctx, gateway.GetListLatestTelemetryWaterPostReq{
			Date: time.Now().Format(time.DateOnly)})
		if err != nil {
			return nil, err
		}

		waterLevelLatestTelemetryMap := make(map[int64]float64)
		for _, wlManual := range waterLevelLatestTelemetry.LatestManualWaterPost {
			waterLevelLatestTelemetryMap[wlManual.WaterLevelPostID] = wlManual.WaterLevel
		}

		features := make([]WaterLevelFeature, 0, len(waterlevelData.WaterLevelPost))
		for _, waterLevel := range waterlevelData.WaterLevelPost {
			feature := WaterLevelFeature{
				Type: "Feature",
				Geometry: Geometry{
					Type:        "Point",
					Coordinates: []float64{waterLevel.Longitude, waterLevel.Latitude},
				},
				Properties: WaterLevelProperties{
					ID:         waterLevel.ID,
					ExternalID: waterLevel.ExternalID,
					Name:       waterLevel.Name,
					Type:       waterLevel.Type,
					Manual:     waterLevelLatestManualMap[waterLevel.ExternalID],
					Telemetry:  waterLevelLatestTelemetryMap[waterLevel.ExternalID],
				},
			}
			features = append(features, feature)
		}

		return &GetListWaterLevelResp{
			Total: len(waterlevelData.WaterLevelPost),
			WaterLevels: WaterLevelData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}

// // Helper function to handle nil pointers
// func getManualLevel(manualLevel *float64) float64 {
// 	if manualLevel != nil {
// 		return *manualLevel
// 	}
// 	return 0
// }
