package usecase

import (
	"context"
	"shared/core"
	sharedGateway "shared/gateway"
	"shared/model"
	"time"
)

type DetailRainfallPostCalculationReq struct {
	StartDate string
	EndDate   string
	ID        string
}

type DetailRainfallPostCalculationRes struct {
	StartDate                     string                        `json:"start_date"`
	EndDate                       string                        `json:"end_date"`
	ID                            string                        `json:"id"`
	ExternalID                    int64                         `json:"external_id"`
	Name                          string                        `json:"name"`
	DetailRainfallPostCalculation DetailRainfallPostCalculation `json:"rainfall_post"`
}

type DetailRainfallPostCalculation struct {
	DetailRainfallPostTelemetry DetailRainfallPostTelemetry `json:"telemetry"`
	DetailRainfallPostManual    DetailRainfallPostManual    `json:"manual"`
	DetailRainfallGraph         DetailRainfallGraph         `json:"graph"`
}

type DetailRainfallPostTelemetry struct {
	Vendor       string  `json:"vendor"`
	Rainfall     float64 `json:"rainfall"`
	LatestUpdate string  `json:"latest_update"`
}

type DetailRainfallPostManual struct {
	Officer      string   `json:"officer"`
	Rainfall     *float64 `json:"rainfall"`
	LatestUpdate string   `json:"latest_update"`
}

type DetailRainfallGraph struct {
	Data []CompressedRainfallHourlies `json:"detail"`
}

type CompressedRainfallHourlies struct {
	Hour      int       `json:"hour"`
	Rain      float64   `json:"rain"`
	Timestamp time.Time `json:"timestamp"`
}

type RainfallDetailWithCalculationUseCase = core.ActionHandler[DetailRainfallPostCalculationReq, DetailRainfallPostCalculationRes]

func ImplRainfallDetailWithCalculationUseCase(
	getLatestTelemetryRainPostByID sharedGateway.GetLatestTelemetryRainPostByIDGateway,
	getDetailRainfallPost sharedGateway.RainfallGetDetailGateway,
	getListRainfallHourlyGateway sharedGateway.GetRainfallHourlyGateway,
) RainfallDetailWithCalculationUseCase {
	return func(ctx context.Context, req DetailRainfallPostCalculationReq) (*DetailRainfallPostCalculationRes, error) {

		rainfallPostsAdditionalData, err := getDetailRainfallPost(ctx, sharedGateway.RainfallGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		latestRainfallTelemetry, err := getLatestTelemetryRainPostByID(ctx, sharedGateway.GetLatestTelemetryRainPostByIDReq{
			PostIDs: []int64{rainfallPostsAdditionalData.RainFall.ExternalID},
			Date:    req.EndDate,
		})
		if err != nil {
			return nil, err
		}
		detailRainfallPost := model.HydrologyRainHourly{}
		if latestRainfallTelemetry != nil && len(latestRainfallTelemetry.LatestTelemetry) > 0 {
			detailRainfallPost = latestRainfallTelemetry.LatestTelemetry[0]
		}

		RainfallHourly, err := getListRainfallHourlyGateway(ctx, sharedGateway.GetRainfallHourlyReq{
			PostID:    rainfallPostsAdditionalData.RainFall.ExternalID,
			StartDate: req.StartDate,
			EndDate:   req.EndDate,
		})
		if err != nil {
			return nil, err
		}

		rainFallHourlyGraphData := make([]CompressedRainfallHourlies, 0)
		for _, rainData := range RainfallHourly.RainfallHourly {
			samplingDate, _ := time.Parse(time.RFC3339, rainData.Sampling)
			rainFallHourlyGraphData = append(rainFallHourlyGraphData, CompressedRainfallHourlies{
				Hour:      rainData.Hour,
				Rain:      float64(rainData.Rain),
				Timestamp: samplingDate,
			})
		}
		//compressedRainfallHourlyData := CompressRainfallHourlyData(rainFallHourlyGraphData, 50)

		detailRainfallDaily := DetailRainfallPostTelemetry{
			Vendor:       rainfallPostsAdditionalData.RainFall.Vendor,
			Rainfall:     float64(detailRainfallPost.Rain),
			LatestUpdate: detailRainfallPost.UpdatedAt.Format(time.DateTime),
		}

		detailRainfallManual := DetailRainfallPostManual{
			Officer: rainfallPostsAdditionalData.RainFall.Officer,
			// Rainfall:     detailRainfallPost.RainFallDetail.Manual,
			// LatestUpdate: detailRainfallPost.RainFallDetail.Sampling,
		}

		return &DetailRainfallPostCalculationRes{
			StartDate:  req.StartDate,
			EndDate:    req.EndDate,
			Name:       rainfallPostsAdditionalData.RainFall.Name,
			ID:         rainfallPostsAdditionalData.RainFall.ID,
			ExternalID: rainfallPostsAdditionalData.RainFall.ExternalID,
			DetailRainfallPostCalculation: DetailRainfallPostCalculation{
				DetailRainfallPostTelemetry: detailRainfallDaily,
				DetailRainfallPostManual:    detailRainfallManual,
				DetailRainfallGraph: DetailRainfallGraph{
					Data: rainFallHourlyGraphData,
				},
			},
		}, nil
	}
}

func CompressRainfallHourlyData(data []model.HydrologyRainHourly, maxPoints int) []CompressedRainfallHourlies {
	dataLen := len(data)
	if dataLen <= maxPoints {
		result := make([]CompressedRainfallHourlies, dataLen)
		for i, d := range data {
			samplingDate, _ := time.Parse("2006-01-02", d.Sampling)
			timestamp := samplingDate.Add(time.Duration(d.Hour) * time.Hour)

			result[i] = CompressedRainfallHourlies{
				Hour:      d.Hour,
				Rain:      float64(d.Rain),
				Timestamp: timestamp,
			}
		}
		return result
	}

	// Calculate compression ratio
	compressionFactor := float64(dataLen) / float64(maxPoints)
	compressedData := make([]CompressedRainfallHourlies, maxPoints)

	for i := 0; i < maxPoints; i++ {
		startIdx := int(float64(i) * compressionFactor)
		endIdx := int(float64(i+1) * compressionFactor)

		if endIdx > dataLen {
			endIdx = dataLen
		}

		// Calculate average for the segment
		var sumRain float64
		var pointCount int

		for j := startIdx; j < endIdx; j++ {
			sumRain += float64(data[j].Rain)
			pointCount++
		}

		midIdx := (startIdx + endIdx) / 2

		// Parse sampling date and combine with hour for the middle point
		samplingDate, _ := time.Parse("2006-01-02", data[midIdx].Sampling)
		timestamp := samplingDate.Add(time.Duration(data[midIdx].Hour) * time.Hour)

		compressedData[i] = CompressedRainfallHourlies{
			Hour:      data[midIdx].Hour,
			Rain:      sumRain / float64(pointCount),
			Timestamp: timestamp,
		}
	}

	return compressedData
}
