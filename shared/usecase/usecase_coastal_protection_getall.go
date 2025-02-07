package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetListCoastalProtectionReq struct {
}
type GetListCoastalProtectionResp struct {
	Total              int                   `json:"total"`
	CoastalProtections CoastalProtectionData `json:"coastal_protections"`
}

type CoastalProtectionData struct {
	Type     string                     `json:"type"`
	Features []CoastalProtectionFeature `json:"features"`
}
type CoastalProtectionFeature struct {
	Type       string                      `json:"type"`
	Geometry   datatypes.JSON              `json:"geometry"`
	Properties CoastalProtectionProperties `json:"properties"`
}

type CoastalProtectionProperties struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetListCoastalProtectionUseCase = core.ActionHandler[GetListCoastalProtectionReq, GetListCoastalProtectionResp]

func ImplGetListCoastalProtectionUseCase(getListCoastalProtection gateway.GetCoastalProtectionListGateway) GetListCoastalProtectionUseCase {
	return func(ctx context.Context, req GetListCoastalProtectionReq) (*GetListCoastalProtectionResp, error) {
		coastalProtectionData, err := getListCoastalProtection(ctx, gateway.GetCoastalProtectionListReq{})
		if err != nil {
			return nil, err
		}

		features := make([]CoastalProtectionFeature, 0, len(coastalProtectionData.CoastalProtections))
		for _, weir := range coastalProtectionData.CoastalProtections {

			feature := CoastalProtectionFeature{
				Type:     "Feature",
				Geometry: weir.Geometry,
				Properties: CoastalProtectionProperties{
					ID:   weir.ID,
					Name: weir.Name,
					Type: weir.TypeName,
				},
			}
			features = append(features, feature)
		}

		return &GetListCoastalProtectionResp{
			Total: len(coastalProtectionData.CoastalProtections),
			CoastalProtections: CoastalProtectionData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}
