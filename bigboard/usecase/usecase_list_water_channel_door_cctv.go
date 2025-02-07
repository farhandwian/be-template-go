package usecase

import (
	"bigboard/gateway"
	"context"
	"shared/core"
	"strconv"
)

type ListWaterChannelDoorCCTVReq struct{}

type ListWaterChannelDoorCCTVRes struct {
	GeoJSON GeoJSONFeatureWaterDoorCCTVListCollection `json:"geo_json"`
}

type GeoJSONFeatureWaterDoorCCTVListCollection struct {
	Type     string                            `json:"type"`
	Features []ListWaterChannelDoorCCTVFeature `json:"features"`
}

type ListWaterChannelDoorCCTVFeature struct {
	Type       string                         `json:"type"`
	Properties WaterChannelDoorCCTVProperties `json:"properties"`
	Geometry   Geometry                       `json:"geometry"`
}

type WaterChannelDoorCCTVProperties struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CCTVCount int    `json:"cctv_count"`
}

type ListWaterChannelDoorCCTVUseCase = core.ActionHandler[
	ListWaterChannelDoorCCTVReq,
	ListWaterChannelDoorCCTVRes,
]

func ImplListWaterChannelDoorCCTV(
	getListWaterChannelDoorGateway gateway.GetListWaterChannelDoorWithCCTVAndOfficerGateway,
) ListWaterChannelDoorCCTVUseCase {
	return func(ctx context.Context, request ListWaterChannelDoorCCTVReq) (*ListWaterChannelDoorCCTVRes, error) {
		// get water channel doors
		waterChannelDoors, err := getListWaterChannelDoorGateway(ctx, gateway.GetListWaterChannelDoorWithCCTVAndOfficerReq{})
		if err != nil {
			return nil, err
		}

		var features []ListWaterChannelDoorCCTVFeature
		for _, waterChannelDoor := range waterChannelDoors.WaterChannels {

			longitude, _ := strconv.ParseFloat(waterChannelDoor.Longitude, 64)
			latitude, _ := strconv.ParseFloat(waterChannelDoor.Latitude, 64)

			feature := ListWaterChannelDoorCCTVFeature{
				Type: "Feature",
				Properties: WaterChannelDoorCCTVProperties{
					ID:        waterChannelDoor.ExternalID,
					Name:      waterChannelDoor.Name,
					CCTVCount: waterChannelDoor.CCTVCount,
				},
				Geometry: Geometry{
					Type:        "Point",
					Coordinates: []float64{longitude, latitude},
				},
			}

			features = append(features, feature)
		}

		geoJSON := GeoJSONFeatureWaterDoorCCTVListCollection{
			Type:     "FeatureCollection",
			Features: features,
		}
		return &ListWaterChannelDoorCCTVRes{GeoJSON: geoJSON}, nil
	}
}
