package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetListIntakeReq struct {
}
type GetListIntakeResp struct {
	Total   int        `json:"total"`
	Intakes IntakeData `json:"intakes"`
}

type IntakeData struct {
	Type     string          `json:"type"`
	Features []IntakeFeature `json:"features"`
}
type IntakeFeature struct {
	Type       string           `json:"type"`
	Geometry   datatypes.JSON   `json:"geometry"`
	Properties IntakeProperties `json:"properties"`
}

type IntakeProperties struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetListIntakeUseCase = core.ActionHandler[GetListIntakeReq, GetListIntakeResp]

func ImplGetListIntakeUseCase(getListIntake gateway.GetIntakeListGateway) GetListIntakeUseCase {
	return func(ctx context.Context, req GetListIntakeReq) (*GetListIntakeResp, error) {
		intakesData, err := getListIntake(ctx, gateway.GetIntakeListReq{})
		if err != nil {
			return nil, err
		}

		features := make([]IntakeFeature, 0, len(intakesData.Intakes))
		for _, weir := range intakesData.Intakes {

			feature := IntakeFeature{
				Type:     "Feature",
				Geometry: weir.Geometry,
				Properties: IntakeProperties{
					ID:   weir.ID,
					Name: weir.Name,
					Type: weir.TypeName,
				},
			}
			features = append(features, feature)
		}

		return &GetListIntakeResp{
			Total: len(intakesData.Intakes),
			Intakes: IntakeData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}
