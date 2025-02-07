package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetListDamReq struct {
}
type GetDamListResp struct {
	Total int     `json:"total"`
	Dam   DamData `json:"dams"`
}

type DamData struct {
	Type     string       `json:"type"`
	Features []DamFeature `json:"features"`
}

type DamFeature struct {
	Type       string         `json:"type"`
	Geometry   datatypes.JSON `json:"geometry"`
	Properties DamProperties  `json:"properties"`
}

type DamProperties struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type GetListDamUseCase = core.ActionHandler[GetListDamReq, GetDamListResp]

func ImplGetListDamUseCase(getListDam gateway.GetDamListGateway) GetListDamUseCase {
	return func(ctx context.Context, req GetListDamReq) (*GetDamListResp, error) {
		damData, err := getListDam(ctx, gateway.GetDamListReq{})
		if err != nil {
			return nil, err
		}

		features := make([]DamFeature, 0, len(damData.Dams))
		for _, lake := range damData.Dams {

			feature := DamFeature{
				Type:     "Feature",
				Geometry: lake.Geometry,
				Properties: DamProperties{
					ID:   lake.ID,
					Name: lake.Name,
				},
			}
			features = append(features, feature)
		}

		return &GetDamListResp{
			Total: len(damData.Dams),
			Dam: DamData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}
