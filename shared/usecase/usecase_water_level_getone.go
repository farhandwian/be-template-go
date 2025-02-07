package usecase

import (
	"context"
	"shared/core"
	sharedGateway "shared/gateway"
)

type GetWaterLevelDetailReq struct {
	ID string
}

type GetWaterLevelDetailResp struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Officer         string  `json:"officer"`
	OfficerWhatsapp string  `json:"officerWhatsapp"`
	Elevation       float64 `json:"elevation"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Location        string  `json:"location"`
}

type GetWaterLevelDetailUseCase = core.ActionHandler[GetWaterLevelDetailReq, GetWaterLevelDetailResp]

func ImplGetWaterLevelDetail(getWaterLevelDetail sharedGateway.WaterLevelGetDetailGateway, getListPostOfficerGateway sharedGateway.GetOfficerListGateway) GetWaterLevelDetailUseCase {
	return func(ctx context.Context, req GetWaterLevelDetailReq) (*GetWaterLevelDetailResp, error) {
		data, err := getWaterLevelDetail(ctx, sharedGateway.WaterLevelGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		officer, err := getListPostOfficerGateway(ctx, sharedGateway.GetOfficerListByNameReq{
			Name: []string{data.WaterLevel.Officer},
		})
		if err != nil {
			return nil, err
		}

		phoneNumber := ""
		for _, ofc := range officer.RainFalls {
			phoneNumber = ofc.PhoneNumber
		}

		return &GetWaterLevelDetailResp{
			ID:              data.WaterLevel.ID,
			Name:            data.WaterLevel.Name,
			Type:            data.WaterLevel.Type,
			Officer:         data.WaterLevel.Officer,
			OfficerWhatsapp: phoneNumber,
			Elevation:       data.WaterLevel.Elevation,
			Latitude:        data.WaterLevel.Latitude,
			Longitude:       data.WaterLevel.Longitude,
			Location:        data.WaterLevel.Location,
		}, err
	}
}
