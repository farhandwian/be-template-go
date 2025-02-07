package usecase

import (
	"bigboard/gateway"
	"context"
	"shared/core"
	"strconv"
)

type ListWaterChannelDoorOfficerReq struct{}

type ListWaterChannelDoorOfficerRes struct {
	GeoJSON GeoJSONFeatureWaterDoorOfficerListCollection `json:"geo_json"`
}

type GeoJSONFeatureWaterDoorOfficerListCollection struct {
	Type     string                               `json:"type"`
	Features []ListWaterChannelDoorOfficerFeature `json:"features"`
}

type ListWaterChannelDoorOfficerFeature struct {
	Type       string                            `json:"type"`
	Properties WaterChannelDoorOfficerProperties `json:"properties"`
	Geometry   Geometry                          `json:"geometry"`
}

type WaterChannelDoorOfficerProperties struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	OfficerCount int    `json:"officer_count"`
}

type ListWaterChannelDoorOfficerUseCase = core.ActionHandler[ListWaterChannelDoorOfficerReq, ListWaterChannelDoorOfficerRes]

func ImplListWaterChannelDoorOfficer(
	getListWaterChannelDoorGateway gateway.GetListWaterChannelDoorWithCCTVAndOfficerGateway,
) ListWaterChannelDoorOfficerUseCase {
	return func(ctx context.Context, request ListWaterChannelDoorOfficerReq) (*ListWaterChannelDoorOfficerRes, error) {
		// get water channel doors
		waterChannelDoors, err := getListWaterChannelDoorGateway(ctx, gateway.GetListWaterChannelDoorWithCCTVAndOfficerReq{})
		if err != nil {
			return nil, err
		}

		var features []ListWaterChannelDoorOfficerFeature
		for _, waterChannelDoor := range waterChannelDoors.WaterChannels {

			longitude, _ := strconv.ParseFloat(waterChannelDoor.Longitude, 64)
			latitude, _ := strconv.ParseFloat(waterChannelDoor.Latitude, 64)

			feature := ListWaterChannelDoorOfficerFeature{
				Type: "Feature",
				Properties: WaterChannelDoorOfficerProperties{
					ID:           waterChannelDoor.ExternalID,
					Name:         waterChannelDoor.Name,
					OfficerCount: waterChannelDoor.OfficerCount,
				},
				Geometry: Geometry{
					Type:        "Point",
					Coordinates: []float64{longitude, latitude},
				},
			}
			features = append(features, feature)
		}

		geoJSON := GeoJSONFeatureWaterDoorOfficerListCollection{
			Type:     "FeatureCollection",
			Features: features,
		}

		return &ListWaterChannelDoorOfficerRes{GeoJSON: geoJSON}, nil
	}
}
