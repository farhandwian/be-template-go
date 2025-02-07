package usecase

import (
	"context"
	"fmt"
	"shared/core"
	"shared/gateway"
)

type AiGetRainFallDetailReq struct {
	ID string
}

type AiGetRainFallDetailResp struct {
	//ID         string `json:"id"`
	ExternalID int64  `json:"external_id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Officer    string `json:"officer"`
	//OfficerWhatsApp    string  `json:"officer_whatsapp"`
	CurrentLevel       float64 `json:"current_level"`
	CurrentLevelStatus string  `json:"current_level_status"`
	Elevation          float64 `json:"elevation"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	City               string  `json:"city"`
	//Vendor             string  `json:"vendor"`
}

type AiGetRainFallDetailUseCase = core.ActionHandler[AiGetRainFallDetailReq, AiGetRainFallDetailResp]

func ImplAiGetRainFallDetail(getRainFallDetail gateway.RainfallGetDetailGateway,
	getListPostOfficerGateway gateway.GetOfficerListGateway,
	getLatestTelemetryRainfallPostGateway gateway.GetLatestTelemetryRainPostByIDGateway,
	sendSSEGateway gateway.SendSSEMessage) AiGetRainFallDetailUseCase {
	return func(ctx context.Context, req AiGetRainFallDetailReq) (*AiGetRainFallDetailResp, error) {
		data, err := getRainFallDetail(ctx, gateway.RainfallGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		//officer, err := getListPostOfficerGateway(ctx, gateway.GetOfficerListByNameReq{
		//	Name: []string{data.RainFall.Officer},
		//})
		//if err != nil {
		//	return nil, err
		//}

		//phoneNumber := ""
		//for _, ofc := range officer.RainFalls {
		//	phoneNumber = ofc.PhoneNumber
		//}

		rainfallPostTelemetries, err := getLatestTelemetryRainfallPostGateway(ctx,
			gateway.GetLatestTelemetryRainPostByIDReq{
				PostIDs: []int64{data.RainFall.ExternalID},
			},
		)
		if err != nil {
			return nil, err
		}

		if len(rainfallPostTelemetries.LatestTelemetry) == 0 {
			return nil, fmt.Errorf("no telemetry data found for today")
		}
		rainfallPostTelemetry := rainfallPostTelemetries.LatestTelemetry[0]

		response := &AiGetRainFallDetailResp{
			//ID:         data.RainFall.ID,
			ExternalID: data.RainFall.ExternalID,
			Name:       data.RainFall.Name,
			Type:       data.RainFall.Type,
			Officer:    data.RainFall.Officer,
			//OfficerWhatsApp:    phoneNumber,
			CurrentLevel:       float64(rainfallPostTelemetry.Rain),
			CurrentLevelStatus: parseRainCategory(float64(rainfallPostTelemetry.Rain)),
			Elevation:          data.RainFall.Elevation,
			Latitude:           data.RainFall.Latitude,
			Longitude:          data.RainFall.Longitude,
			City:               data.RainFall.City,
			//Vendor:             data.RainFall.Vendor,
		}

		_, _ = sendSSEGateway(ctx, gateway.SendSSEMessageReq{
			Subject:      "get-detail-rainfall-post",
			FunctionName: "showDetailRainfallPost",
			Data:         response,
		})

		return response, err
	}
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
