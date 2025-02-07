package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetListWellReq struct {
}
type GetListWellResp struct {
	Total int      `json:"total"`
	Wells WellData `json:"wells"`
}

type WellData struct {
	Type     string        `json:"type"`
	Features []WellFeature `json:"features"`
}
type WellFeature struct {
	Type       string         `json:"type"`
	Geometry   datatypes.JSON `json:"geometry"`
	Properties WellProperties `json:"properties"`
}

type WellProperties struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetListWellUseCase = core.ActionHandler[GetListWellReq, GetListWellResp]

func ImplGetListWellUseCase(getListWell gateway.GetWellListGateway) GetListWellUseCase {
	return func(ctx context.Context, req GetListWellReq) (*GetListWellResp, error) {
		weirData, err := getListWell(ctx, gateway.GetWellListReq{})
		if err != nil {
			return nil, err
		}

		features := make([]WellFeature, 0, len(weirData.Wells))
		for _, weir := range weirData.Wells {

			feature := WellFeature{
				Type:     "Feature",
				Geometry: weir.Geometry,
				Properties: WellProperties{
					ID:   weir.ID,
					Name: weir.Name,
					Type: weir.TypeName,
				},
			}
			features = append(features, feature)
		}

		return &GetListWellResp{
			Total: len(weirData.Wells),
			Wells: WellData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}
