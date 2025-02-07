package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"time"
)

type GetRainFallDetailReq struct {
	ID string
}

type GetRainFallDetailResp struct {
	ID                 string  `json:"id"`
	ExternalID         int64   `json:"external_id"`
	Name               string  `json:"name"`
	Type               string  `json:"type"`
	Officer            string  `json:"officer"`
	OfficerWhatsApp    string  `json:"officer_whatsapp"`
	HighestLevelDate   string  `json:"highest_level_date"`
	HighestLevelLevel  float64 `json:"highest_level_level"`
	CurrentLevel       float64 `json:"current_level"`
	CurrentLevelStatus string  `json:"current_level_status"`
	Elevation          float64 `json:"elevation"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	City               string  `json:"city"`
	Vendor             string  `json:"vendor"`
}

type GetRainFallDetailUseCase = core.ActionHandler[GetRainFallDetailReq, GetRainFallDetailResp]

func ImplGetRainFallDetail(getRainFallDetail gateway.RainfallGetDetailGateway,
	getListPostOfficerGateway gateway.GetOfficerListGateway,
	getLatestRainfallTelemetryGateway gateway.GetLatestTelemetryRainPostByIDGateway,
) GetRainFallDetailUseCase {
	return func(ctx context.Context, req GetRainFallDetailReq) (*GetRainFallDetailResp, error) {
		data, err := getRainFallDetail(ctx, gateway.RainfallGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		officer, err := getListPostOfficerGateway(ctx, gateway.GetOfficerListByNameReq{
			Name: []string{data.RainFall.Officer},
		})
		if err != nil {
			return nil, err
		}

		phoneNumber := ""
		for _, ofc := range officer.RainFalls {
			phoneNumber = ofc.PhoneNumber
		}

		// highestRainPost, err := getHighestRainPostGateway(ctx, gateway.RainfallPostGetHighestReq{
		// 	ID: data.RainFall.ExternalID,
		// })
		// if err != nil {
		// 	return nil, err
		// }

		rainfallPostTelemetries, err := getLatestRainfallTelemetryGateway(ctx,
			gateway.GetLatestTelemetryRainPostByIDReq{
				PostIDs: []int64{data.RainFall.ExternalID},
				Date:    time.Now().Format(time.DateOnly),
			},
		)
		if err != nil {
			return nil, err
		}

		var currentRainfallDailyLevel float64
		if rainfallPostTelemetries.LatestTelemetry != nil {
			currentRainfallDailyLevel = float64(rainfallPostTelemetries.LatestTelemetry[0].Rain)
		}

		return &GetRainFallDetailResp{
			ID:                 data.RainFall.ID,
			ExternalID:         data.RainFall.ExternalID,
			Name:               data.RainFall.Name,
			Type:               data.RainFall.Type,
			Officer:            data.RainFall.Officer,
			OfficerWhatsApp:    phoneNumber,
			CurrentLevel:       currentRainfallDailyLevel,
			CurrentLevelStatus: parseRainCategory(currentRainfallDailyLevel),
			Elevation:          data.RainFall.Elevation,
			Latitude:           data.RainFall.Latitude,
			Longitude:          data.RainFall.Longitude,
			City:               data.RainFall.City,
			Vendor:             data.RainFall.Vendor,
			// HighestLevelDate:   highestRainPost.RainfallPostHighest.Sampling,
			// HighestLevelLevel:  highestRainPost.RainfallPostHighest.Rain24,
		}, err
	}
}
