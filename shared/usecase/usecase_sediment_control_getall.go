package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetListSedimentControlReq struct {
}
type GetListSedimentControlResp struct {
	Total            int                 `json:"total"`
	SedimentControls SedimentControlData `json:"sediment_controls"`
}

type SedimentControlData struct {
	Type     string                   `json:"type"`
	Features []SedimentControlFeature `json:"features"`
}
type SedimentControlFeature struct {
	Type       string                    `json:"type"`
	Geometry   datatypes.JSON            `json:"geometry"`
	Properties SedimentControlProperties `json:"properties"`
}

type SedimentControlProperties struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetListSedimentControlUseCase = core.ActionHandler[GetListSedimentControlReq, GetListSedimentControlResp]

func ImplGetListSedimentControlUseCase(getListSedimentControl gateway.GetSedimentControlListGateway) GetListSedimentControlUseCase {
	return func(ctx context.Context, req GetListSedimentControlReq) (*GetListSedimentControlResp, error) {
		groundWaterData, err := getListSedimentControl(ctx, gateway.GetSedimentControlListReq{})
		if err != nil {
			return nil, err
		}

		features := make([]SedimentControlFeature, 0, len(groundWaterData.SedimentControls))
		for _, weir := range groundWaterData.SedimentControls {

			feature := SedimentControlFeature{
				Type:     "Feature",
				Geometry: weir.Geometry,
				Properties: SedimentControlProperties{
					ID:   weir.ID,
					Name: weir.Name,
					Type: weir.TypeName,
				},
			}
			features = append(features, feature)
		}

		return &GetListSedimentControlResp{
			Total: len(groundWaterData.SedimentControls),
			SedimentControls: SedimentControlData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}
