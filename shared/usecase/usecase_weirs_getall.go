package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetListWeirReq struct {
}
type GetListWeirResp struct {
	Total int      `json:"total"`
	Weirs WeirData `json:"weirs"`
}

type WeirData struct {
	Type     string        `json:"type"`
	Features []WeirFeature `json:"features"`
}
type WeirFeature struct {
	Type       string         `json:"type"`
	Geometry   datatypes.JSON `json:"geometry"`
	Properties WeirProperties `json:"properties"`
}

type WeirProperties struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetListWeirUseCase = core.ActionHandler[GetListWeirReq, GetListWeirResp]

func ImplGetListWeirUseCase(getListWeir gateway.GetWeirListGateway) GetListWeirUseCase {
	return func(ctx context.Context, req GetListWeirReq) (*GetListWeirResp, error) {
		weirData, err := getListWeir(ctx, gateway.GetWeirListReq{})
		if err != nil {
			return nil, err
		}

		features := make([]WeirFeature, 0, len(weirData.Weirs))
		for _, weir := range weirData.Weirs {

			feature := WeirFeature{
				Type:     "Feature",
				Geometry: weir.Geometry,
				Properties: WeirProperties{
					ID:   weir.ID,
					Name: weir.Name,
					Type: weir.TypeName,
				},
			}
			features = append(features, feature)
		}

		return &GetListWeirResp{
			Total: len(weirData.Weirs),
			Weirs: WeirData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}
