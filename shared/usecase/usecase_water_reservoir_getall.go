package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetListWaterReservoirReq struct {
}
type GetListWaterReservoirResp struct {
	Total           int                `json:"total"`
	WaterReservoirs WaterReservoirData `json:"water_reservoirs"`
}

type WaterReservoirData struct {
	Type     string                  `json:"type"`
	Features []WaterReservoirFeature `json:"features"`
}
type WaterReservoirFeature struct {
	Type       string                   `json:"type"`
	Geometry   datatypes.JSON           `json:"geometry"`
	Properties WaterReservoirProperties `json:"properties"`
}

type WaterReservoirProperties struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetListWaterReservoirUseCase = core.ActionHandler[GetListWaterReservoirReq, GetListWaterReservoirResp]

func ImplGetListWaterReservoirUseCase(getListWaterReservoir gateway.GetWaterReservoirListGateway) GetListWaterReservoirUseCase {
	return func(ctx context.Context, req GetListWaterReservoirReq) (*GetListWaterReservoirResp, error) {
		waterReservoirData, err := getListWaterReservoir(ctx, gateway.GetWaterReservoirListReq{})
		if err != nil {
			return nil, err
		}

		features := make([]WaterReservoirFeature, 0, len(waterReservoirData.WaterReservoirs))
		for _, weir := range waterReservoirData.WaterReservoirs {

			feature := WaterReservoirFeature{
				Type:     "Feature",
				Geometry: weir.Geometry,
				Properties: WaterReservoirProperties{
					ID:   weir.ID,
					Name: weir.Name,
					Type: weir.TypeName,
				},
			}
			features = append(features, feature)
		}

		return &GetListWaterReservoirResp{
			Total: len(waterReservoirData.WaterReservoirs),
			WaterReservoirs: WaterReservoirData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}
