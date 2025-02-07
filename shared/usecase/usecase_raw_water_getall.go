package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetListRawWaterReq struct {
}
type GetListRawWaterResp struct {
	Total     int          `json:"total"`
	RawWaters RawWaterData `json:"raw_waters"`
}

type RawWaterData struct {
	Type     string            `json:"type"`
	Features []RawWaterFeature `json:"features"`
}
type RawWaterFeature struct {
	Type       string             `json:"type"`
	Geometry   datatypes.JSON     `json:"geometry"`
	Properties RawWaterProperties `json:"properties"`
}

type RawWaterProperties struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetListRawWaterUseCase = core.ActionHandler[GetListRawWaterReq, GetListRawWaterResp]

func ImplGetListRawWaterUseCase(getListRawWater gateway.GetRawWaterListGateway) GetListRawWaterUseCase {
	return func(ctx context.Context, req GetListRawWaterReq) (*GetListRawWaterResp, error) {
		groundWaterData, err := getListRawWater(ctx, gateway.GetRawWaterListReq{})
		if err != nil {
			return nil, err
		}

		features := make([]RawWaterFeature, 0, len(groundWaterData.RawWaters))
		for _, weir := range groundWaterData.RawWaters {

			feature := RawWaterFeature{
				Type:     "Feature",
				Geometry: weir.Geometry,
				Properties: RawWaterProperties{
					ID:   weir.ID,
					Name: weir.Name,
					Type: weir.TypeName,
				},
			}
			features = append(features, feature)
		}

		return &GetListRawWaterResp{
			Total: len(groundWaterData.RawWaters),
			RawWaters: RawWaterData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}
