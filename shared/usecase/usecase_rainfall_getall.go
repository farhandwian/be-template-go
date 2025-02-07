package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"shared/model"
	"time"
)

type GetListRainFallReq struct {
}
type GetListRainFallResp struct {
	Total     int          `json:"total"`
	RainFalls RainFallData `json:"rain_falls"`
}

type RainFallData struct {
	Type     string            `json:"type"`
	Features []RainFallFeature `json:"features"`
}
type RainFallFeature struct {
	Type       string             `json:"type"`
	Geometry   Geometry           `json:"geometry"`
	Properties RainFallProperties `json:"properties"`
}

type RainFallProperties struct {
	ID           string  `json:"id"`
	ExternalID   int64   `json:"external_id"`
	Name         string  `json:"name"`
	Manual       float64 `json:"manual"`
	Telemetry    float64 `json:"telemetry"`
	RainCategory string  `json:"rain_category"`
	Type         string  `json:"type"`
}

type GetListRainFallUseCase = core.ActionHandler[GetListRainFallReq, GetListRainFallResp]

func ImplGetListRainFallUseCase(
	getListRainFall gateway.GetRainFallPostListGateway,
	getLatestTelemetryRainPost gateway.GetLatestTelemetryRainPostByIDGateway,
) GetListRainFallUseCase {
	return func(ctx context.Context, req GetListRainFallReq) (*GetListRainFallResp, error) {
		rainFallData, err := getListRainFall(ctx, gateway.GetRainfallPostListReq{})
		if err != nil {
			return nil, err
		}

		rainfallPostIDs := make([]int64, 0)

		for _, rainfallPost := range rainFallData.RainFalls {
			rainfallPostIDs = append(rainfallPostIDs, rainfallPost.ExternalID)
		}

		// get and map rainfall telemetry
		rainfallPostTelemetries, err := getLatestTelemetryRainPost(ctx,
			gateway.GetLatestTelemetryRainPostByIDReq{
				PostIDs: rainfallPostIDs,
				Date:    time.Now().Format(time.DateOnly),
			},
		)
		if err != nil {
			return nil, err
		}
		rainfallPostTelemetryMap := make(map[int64]model.HydrologyRainHourly)

		for _, telemetry := range rainfallPostTelemetries.LatestTelemetry {
			rainfallPostTelemetryMap[int64(telemetry.RainPostID)] = telemetry
		}

		features := make([]RainFallFeature, 0, len(rainFallData.RainFalls))
		for _, rainfall := range rainFallData.RainFalls {

			feature := RainFallFeature{
				Type: "Feature",
				Geometry: Geometry{
					Type:        "Point",
					Coordinates: []float64{rainfall.Longitude, rainfall.Latitude},
				},
				Properties: RainFallProperties{
					ID:         rainfall.ID,
					ExternalID: rainfall.ExternalID,
					Name:       rainfall.Name,
					// no need in map
					// Manual:       parseManualLevel(rainfallPostTelemetryMap[rainfall.ExternalID].Manual),
					Telemetry:    float64(rainfallPostTelemetryMap[rainfall.ExternalID].Rain),
					RainCategory: parseRainCategory(float64(rainfallPostTelemetryMap[rainfall.ExternalID].Rain)),
					Type:         rainfall.Type,
				},
			}
			features = append(features, feature)
		}

		return &GetListRainFallResp{
			Total: len(rainFallData.RainFalls),
			RainFalls: RainFallData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}

func parseManualLevel(manualLevel *float64) float64 {
	if manualLevel != nil {
		return *manualLevel
	}
	return 0 // Return 0 if the pointer is nil
}

func parseRainCategory(rainTelemetryValue float64) string {
	if rainTelemetryValue >= 0.5 && rainTelemetryValue <= 20 {
		return "Hujan Ringan"
	} else if rainTelemetryValue >= 20 && rainTelemetryValue <= 50 {
		return "Hujan Sedang"
	} else if rainTelemetryValue >= 50 && rainTelemetryValue <= 100 {
		return "Hujan Lebat"
	} else if rainTelemetryValue >= 100 && rainTelemetryValue <= 150 {
		return "Hujan Sangat Lebat"
	} else if rainTelemetryValue > 150 {
		return "Hujan Ekstrim"
	}
	return ""
}
