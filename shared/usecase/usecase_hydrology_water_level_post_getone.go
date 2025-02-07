package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"sort"
	"strings"
	"time"
)

type WaterLevelPostDetailReq struct {
	StartDate string
	EndDate   string
	ID        string
}

type WaterLevelPostDetailResp struct {
	StartDate       string                        `json:"start_date"`
	EndDate         string                        `json:"end_date"`
	Name            string                        `json:"name"`
	AffectedArea    []string                      `json:"affected_area"`
	Telemetry       WaterLevelPostDetailTelemetry `json:"telemetry"`
	Manual          WaterLevelPostDetailManual    `json:"manual"`
	WaterLevelGraph WaterLevelGraph               `json:"graph"`
}

type WaterLevelPostDetailTelemetry struct {
	TmaMaximum   float64 `json:"tma_maximum"`
	TmaMinimum   float64 `json:"tma_minimum"`
	LatestUpdate string  `json:"latest_update"`
}

type WaterLevelPostDetailManual struct {
	Officer      string `json:"officer"`
	LatestUpdate string `json:"latest_update"`
}

type WaterLevelGraph struct {
	//ManualGraph    []ManualGraph    `json:"manual"`
	//TelemetryGraph []TelemetryGraph `json:"telemetry"`
	//RainGraph      []RainGraph      `json:"rain"`
	CombinedDataPoint []CombinedManualAndTelemetry `json:"combined_data_point"`
}
type CombinedManualAndTelemetry struct {
	Date             time.Time `json:"date"`
	Manual           float64   `json:"manual,omitempty"`
	Telemetry        float64   `json:"telemetry,omitempty"`
	WaterLevelPostID int64     `json:"water_level_post_id"`
}

type ManualGraph struct {
	Date             string  `json:"date"`
	Tma              float64 `json:"tma"`
	WaterLevelPostID int64   `json:"water_level_post_id"`
}

type TelemetryGraph struct {
	Date             string  `json:"date"`
	Tma              float64 `json:"tma"`
	WaterLevelPostID int64   `json:"water_level_post_id"`
}

type RainGraph struct{}

type GetWaterLevelDetailSihkaPostUseCase = core.ActionHandler[WaterLevelPostDetailReq, WaterLevelPostDetailResp]

func ImplDetailWaterLevelPost(
	getDetailWaterLevelSummaryTelemetry gateway.GetWaterLevelDetailSummaryTelemetryGateway,
	getWaterLevelDetail gateway.WaterLevelGetDetailGateway,
	getDetailWaterLevelManual gateway.GetWaterLevelDetailManualGateway,
	getDetailWaterLevelTelemetry gateway.GetWaterLevelDetailTelemetryGateway,
) GetWaterLevelDetailSihkaPostUseCase {
	return func(ctx context.Context, request WaterLevelPostDetailReq) (*WaterLevelPostDetailResp, error) {

		waterLevelPost, err := getWaterLevelDetail(ctx, gateway.WaterLevelGetDetailReq{
			ID: request.ID,
		})
		if err != nil {
			return nil, err
		}

		var latestUpdateManual string
		if waterLevelPost.WaterLevel.UpdatedAt != nil {
			latestUpdateManual = waterLevelPost.WaterLevel.UpdatedAt.Format("2006-01-02 15:04:05")
		}

		waterLevelSummaryTelemetry, err := getDetailWaterLevelSummaryTelemetry(ctx, gateway.GetWaterLevelDetailSummaryTelemetryReq{
			ID:        waterLevelPost.WaterLevel.ExternalID,
			StartDate: request.StartDate,
			EndDate:   request.EndDate,
		})
		if err != nil {
			return nil, err
		}

		waterLevelManual, err := getDetailWaterLevelManual(ctx, gateway.GetWaterLevelDetailManualReq{
			ID:        waterLevelPost.WaterLevel.ExternalID,
			StartDate: request.StartDate,
			EndDate:   request.EndDate,
		})
		if err != nil {
			return nil, err
		}

		waterLevelManualGraph := make([]ManualGraph, 0)
		for _, manual := range waterLevelManual.WaterLevelDetailTelemetry {
			waterLevelManualGraph = append(waterLevelManualGraph, ManualGraph{
				Date:             manual.Sampling,
				Tma:              manual.TMA,
				WaterLevelPostID: manual.WaterLevelPostID,
			})
		}

		waterLevelTelemetry, err := getDetailWaterLevelTelemetry(ctx, gateway.GetWaterLevelDetailTelemetryReq{
			ID:        waterLevelPost.WaterLevel.ExternalID,
			StartDate: request.StartDate,
			EndDate:   request.EndDate,
		})
		if err != nil {
			return nil, err
		}
		waterLevelTelemetryGraph := make([]TelemetryGraph, 0)
		for _, telemetry := range waterLevelTelemetry.WaterLevelDetailTelemetry {
			waterLevelTelemetryGraph = append(waterLevelTelemetryGraph, TelemetryGraph{
				Date:             telemetry.Sampling,
				Tma:              telemetry.WaterLevel,
				WaterLevelPostID: telemetry.WaterLevelPostID,
			})
		}

		compressedWaterLevelPost := combineManualAndTelemetryWaterLevel(waterLevelManualGraph, waterLevelTelemetryGraph)

		return &WaterLevelPostDetailResp{
			StartDate:    request.StartDate,
			EndDate:      request.EndDate,
			Name:         waterLevelPost.WaterLevel.Name,
			AffectedArea: parseAffectedArea(waterLevelPost.WaterLevel.AffectedArea),
			Manual: WaterLevelPostDetailManual{
				Officer:      waterLevelPost.WaterLevel.Officer,
				LatestUpdate: latestUpdateManual,
			},
			Telemetry: WaterLevelPostDetailTelemetry{
				TmaMaximum:   waterLevelSummaryTelemetry.WaterLevelDetailTelemetry.TMAMaximum,
				TmaMinimum:   waterLevelSummaryTelemetry.WaterLevelDetailTelemetry.TMAMinimum,
				LatestUpdate: waterLevelSummaryTelemetry.WaterLevelDetailTelemetry.LatestUpdate,
			},
			WaterLevelGraph: WaterLevelGraph{CombinedDataPoint: compressedWaterLevelPost},
		}, nil
	}
}

func parseAffectedArea(area string) []string {
	if area == "" {
		return []string{}
	}
	return strings.Split(area, ", ")
}

// MergeAndCompress merges manual and telemetry data
//func combineManualAndTelemetryWaterLevel(manualData []ManualGraph, telemetryData []TelemetryGraph, maxPoints int) []CombinedManualAndTelemetry {
//	parsedManual := make([]CombinedManualAndTelemetry, len(manualData))
//	for i, d := range manualData {
//		timestamp, _ := time.Parse("2006-01-02 15:04:05.000", d.Date)
//		parsedManual[i] = CombinedManualAndTelemetry{
//			Date:             timestamp,
//			Manual:           d.Tma,
//			WaterLevelPostID: d.WaterLevelPostID,
//		}
//	}
//
//	// Parse telemetry data
//	parsedTelemetry := make([]CombinedManualAndTelemetry, len(telemetryData))
//	for i, d := range telemetryData {
//		timestamp, _ := time.Parse(time.RFC3339, d.Date)
//		parsedTelemetry[i] = CombinedManualAndTelemetry{
//			Date:             timestamp,
//			Telemetry:        d.Tma,
//			WaterLevelPostID: d.WaterLevelPostID,
//		}
//	}
//
//	// Combine the two datasets
//	combinedData := append(parsedManual, parsedTelemetry...)
//	// Sort the combined data by date
//	sort.Slice(combinedData, func(i, j int) bool {
//		return combinedData[i].Date.Before(combinedData[j].Date)
//	})
//
//	// Create a map to store combined data points
//	dataMap := make(map[string]CombinedManualAndTelemetry)
//	for _, point := range combinedData {
//		dateKey := point.Date.Format(time.RFC3339) // Use RFC3339 format for consistency
//		if existingPoint, found := dataMap[dateKey]; found {
//			if point.Manual != 0 {
//				existingPoint.Manual = point.Manual
//			}
//			if point.Telemetry != 0 {
//				existingPoint.Telemetry = point.Telemetry
//			}
//			dataMap[dateKey] = existingPoint
//		} else {
//
//			if point.Telemetry != 0 {
//				point.Manual = 0 // Set manual to 0 if it’s not present
//			}
//			dataMap[dateKey] = point
//		}
//	}
//
//	// Convert map to slice
//	var result []CombinedManualAndTelemetry
//	for _, point := range dataMap {
//		result = append(result, point)
//	}
//
//	if len(result) > maxPoints {
//		result = result[:maxPoints]
//	}
//
//	sort.Slice(result, func(i, j int) bool {
//		return result[i].Date.Before(result[j].Date)
//	})
//
//	return result
//}

func combineManualAndTelemetryWaterLevel(manualData []ManualGraph, telemetryData []TelemetryGraph) []CombinedManualAndTelemetry {
	combinedData := make([]CombinedManualAndTelemetry, 0, len(manualData)+len(telemetryData))

	// Parse manual data
	for _, d := range manualData {
		timestamp, _ := time.Parse("2006-01-02 15:04:05.000", d.Date)
		combinedData = append(combinedData, CombinedManualAndTelemetry{
			Date:             timestamp,
			Manual:           d.Tma,
			WaterLevelPostID: d.WaterLevelPostID,
		})
	}

	// Parse telemetry data
	for _, d := range telemetryData {
		timestamp, _ := time.Parse(time.RFC3339, d.Date)
		combinedData = append(combinedData, CombinedManualAndTelemetry{
			Date:             timestamp,
			Telemetry:        d.Tma,
			WaterLevelPostID: d.WaterLevelPostID,
		})
	}

	// Sort the combined data by date
	sort.Slice(combinedData, func(i, j int) bool {
		return combinedData[i].Date.Before(combinedData[j].Date)
	})

	return combinedData
}
