package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetListPahAbsahReq struct {
}
type GetListPahAbsahResp struct {
	Total     int          `json:"total"`
	PahAbsahs PahAbsahData `json:"pah_absahs"`
}

type PahAbsahData struct {
	Type     string            `json:"type"`
	Features []PahAbsahFeature `json:"features"`
}
type PahAbsahFeature struct {
	Type       string             `json:"type"`
	Geometry   datatypes.JSON     `json:"geometry"`
	Properties PahAbsahProperties `json:"properties"`
}

type PahAbsahProperties struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetListPahAbsahUseCase = core.ActionHandler[GetListPahAbsahReq, GetListPahAbsahResp]

func ImplGetListPahAbsahUseCase(getListPahAbsah gateway.GetPahAbsahListGateway) GetListPahAbsahUseCase {
	return func(ctx context.Context, req GetListPahAbsahReq) (*GetListPahAbsahResp, error) {
		pahAbsahData, err := getListPahAbsah(ctx, gateway.GetPahAbsahListReq{})
		if err != nil {
			return nil, err
		}

		features := make([]PahAbsahFeature, 0, len(pahAbsahData.PahAbsahs))
		for _, weir := range pahAbsahData.PahAbsahs {

			feature := PahAbsahFeature{
				Type:     "Feature",
				Geometry: weir.Geometry,
				Properties: PahAbsahProperties{
					ID:   weir.ID,
					Name: weir.Name,
					Type: weir.TypeName,
				},
			}
			features = append(features, feature)
		}

		return &GetListPahAbsahResp{
			Total: len(pahAbsahData.PahAbsahs),
			PahAbsahs: PahAbsahData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}
