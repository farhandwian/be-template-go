package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetListLakeReq struct {
}
type GetListLakeResp struct {
	Total int      `json:"total"`
	Lake  LakeData `json:"lakes"`
}

type LakeData struct {
	Type     string        `json:"type"`
	Features []LakeFeature `json:"features"`
}
type LakeFeature struct {
	Type       string         `json:"type"`
	Geometry   datatypes.JSON `json:"geometry"`
	Properties LakeProperties `json:"properties"`
}

type LakeProperties struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetListLakeUseCase = core.ActionHandler[GetListLakeReq, GetListLakeResp]

func ImplGetListLakeUseCase(getListLake gateway.GetLakeListGateway) GetListLakeUseCase {
	return func(ctx context.Context, req GetListLakeReq) (*GetListLakeResp, error) {
		lakesData, err := getListLake(ctx, gateway.GetLakeListReq{})
		if err != nil {
			return nil, err
		}

		features := make([]LakeFeature, 0, len(lakesData.Lakes))
		for _, lake := range lakesData.Lakes {

			feature := LakeFeature{
				Type:     "Feature",
				Geometry: lake.Geometry,
				Properties: LakeProperties{
					ID:   lake.ID,
					Name: lake.Name,
					Type: lake.TypeName,
				},
			}
			features = append(features, feature)
		}

		return &GetListLakeResp{
			Total: len(lakesData.Lakes),
			Lake: LakeData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}
