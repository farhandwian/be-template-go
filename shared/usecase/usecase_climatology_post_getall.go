package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
)

type GetListClimatologyPostReq struct {
}

type GetListClimatologyPostResp struct {
	Total         int             `json:"total"`
	Climatologies ClimatologyData `json:"climatologies"`
}

type ClimatologyData struct {
	Type     string               `json:"type"`
	Features []ClimatologyFeature `json:"features"`
}
type ClimatologyFeature struct {
	Type       string                `json:"type"`
	Geometry   Geometry              `json:"geometry"`
	Properties ClimatologyProperties `json:"properties"`
}

type ClimatologyProperties struct {
	ID         string `json:"id"`
	ExternalID int64  `json:"external_id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
}

type GetClimatologyPostUseCase = core.ActionHandler[GetListClimatologyPostReq, GetListClimatologyPostResp]

func ImplGetClimatologyPostUseCase(getListClimatologyPost gateway.GetClimatologyListGateway) GetClimatologyPostUseCase {
	return func(ctx context.Context, req GetListClimatologyPostReq) (*GetListClimatologyPostResp, error) {
		climatologyData, err := getListClimatologyPost(ctx, gateway.GetClimatologyListReq{})
		if err != nil {
			return nil, err
		}

		features := make([]ClimatologyFeature, 0, len(climatologyData.Climatology))
		for _, climatology := range climatologyData.Climatology {

			feature := ClimatologyFeature{
				Type: "Feature",
				Geometry: Geometry{
					Type:        "Point",
					Coordinates: []float64{climatology.Longitude, climatology.Latitude},
				},
				Properties: ClimatologyProperties{
					ID:         climatology.ID,
					ExternalID: climatology.ExternalID,
					Name:       climatology.Name,
					Type:       climatology.Type,
				},
			}
			features = append(features, feature)
		}

		return &GetListClimatologyPostResp{
			Total: len(climatologyData.Climatology),
			Climatologies: ClimatologyData{
				Type:     "FeatureCollection",
				Features: features,
			},
		}, nil
	}
}
