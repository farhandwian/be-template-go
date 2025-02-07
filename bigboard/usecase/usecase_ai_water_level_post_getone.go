package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
)

type AiGetWaterLevelDetailReq struct {
	ID string
}

type AiGetWaterLevelDetailResp struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Officer string `json:"officer"`
	//OfficerWhatsapp string  `json:"officerWhatsapp"`
	LatestTelemetry float64 `json:"latest_telemetry"`
	Elevation       float64 `json:"elevation"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Location        string  `json:"location"`
}

type AiGetWaterLevelDetailUseCase = core.ActionHandler[AiGetWaterLevelDetailReq, AiGetWaterLevelDetailResp]

func ImplAiGetWaterLevelDetail(
	getWaterLevelDetail gateway.WaterLevelGetDetailGateway,
	getListPostOfficerGateway gateway.GetOfficerListGateway,
	getLatestWaterLevelPostTelemetry gateway.GetLatestTelemetryWaterPostByIDGateway,
	sendSSEGateway gateway.SendSSEMessage,
) AiGetWaterLevelDetailUseCase {
	return func(ctx context.Context, req AiGetWaterLevelDetailReq) (*AiGetWaterLevelDetailResp, error) {
		data, err := getWaterLevelDetail(ctx, gateway.WaterLevelGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		calculation, err := getLatestWaterLevelPostTelemetry(ctx, gateway.GetLatestTelemetryWaterPostByIDReq{
			WaterLevelPostID: data.WaterLevel.ExternalID,
		})
		if err != nil {
			return nil, err
		}

		response := &AiGetWaterLevelDetailResp{
			ID:              data.WaterLevel.ID,
			Name:            data.WaterLevel.Name,
			Type:            data.WaterLevel.Type,
			Officer:         data.WaterLevel.Officer,
			LatestTelemetry: calculation.LatestManualWaterPost.WaterLevel,
			//OfficerWhatsapp: phoneNumber,
			Elevation: data.WaterLevel.Elevation,
			Latitude:  data.WaterLevel.Latitude,
			Longitude: data.WaterLevel.Longitude,
			Location:  data.WaterLevel.Location,
		}

		_, _ = sendSSEGateway(ctx, gateway.SendSSEMessageReq{
			Subject:      "get-detail-from-water-level",
			FunctionName: "showDetailWaterLevel",
			Data:         response,
		})

		return response, err
	}
}
