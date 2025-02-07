package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetListGroundWaterReq struct {
}
type GetListGroundWaterResp struct {
	Total        int             `json:"total"`
	GroundWaters GroundWaterData `json:"ground_waters"`
}

type GroundWaterData struct {
	Type     string               `json:"type"`
	Features []GroundWaterFeature `json:"features"`
}
type GroundWaterFeature struct {
	Type       string                `json:"type"`
	Geometry   datatypes.JSON        `json:"geometry"`
	Properties GroundWaterProperties `json:"properties"`
}

type GroundWaterProperties struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetListGroundWaterUseCase = core.ActionHandler[GetListGroundWaterReq, GetListGroundWaterResp]

func ImplGetListGroundWaterUseCase(getListGroundWater gateway.GetGroundWaterListGateway) GetListGroundWaterUseCase {
	return func(ctx context.Context, req GetListGroundWaterReq) (*GetListGroundWaterResp, error) {
		groundWaterData, err := getListGroundWater(ctx, gateway.GetGroundWaterListReq{})
		if err != nil {
			return nil, err
		}

		features := make([]GroundWaterFeature, 0, len(groundWaterData.GroundWaters))
		for _, weir := range groundWaterData.GroundWaters {

			feature := GroundWaterFeature{
				Type:     "Feature",
				Geometry: weir.Geometry,
				Properties: GroundWaterProperties{
					ID:   weir.ID,
					Name: weir.Name,
					Type: weir.TypeName,
				},
			}
			features = append(features, feature)
		}

		return &GetListGroundWaterResp{
			Total: len(groundWaterData.GroundWaters),
			GroundWaters: GroundWaterData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}
