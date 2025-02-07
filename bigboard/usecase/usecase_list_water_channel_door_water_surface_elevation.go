package usecase

import (
	"context"
	"shared/core"
	sharedGateway "shared/gateway"
	sharedModel "shared/model"
	"strconv"
)

type ListWaterChannelDoorWithWaterSurfaceElevationReq struct{}
type ListWaterChannelDoorWithWaterSurfaceElevationResp struct {
	GeoJSON GeoJSONFeatureWaterDoorWaterElevationCollection `json:"geo_json"`
}

type GeoJSONFeatureWaterDoorWaterElevationCollection struct {
	Type     string                                      `json:"type"`
	Features []ListWaterChannelDoorWaterElevationFeature `json:"features"`
}

type ListWaterChannelDoorWaterElevationFeature struct {
	Type       string                                              `json:"type"`
	Properties WaterChannelDoorWithWaterSurfaceElevationProperties `json:"properties"`
	Geometry   Geometry                                            `json:"geometry"`
}

type WaterChannelDoorWithWaterSurfaceElevationProperties struct {
	ID                    int     `json:"id"`
	Name                  string  `json:"name"`
	WaterSurfaceElevation float64 `json:"water_surface_elevation"`
}

type ListWaterChannelDoorWithWaterSurfaceElevationUseCase = core.ActionHandler[
	ListWaterChannelDoorWithWaterSurfaceElevationReq,
	ListWaterChannelDoorWithWaterSurfaceElevationResp,
]

func ImplListWaterChannelDoorWithWaterSurfaceElevation(
	getListWaterChannelDoorGateway sharedGateway.GetListWaterChannelDoorGateway,
	getListWaterChannelWaterSurfaceElevationGateway sharedGateway.GetListWaterChannelTMAGateway,
) ListWaterChannelDoorWithWaterSurfaceElevationUseCase {

	return func(ctx context.Context, request ListWaterChannelDoorWithWaterSurfaceElevationReq) (*ListWaterChannelDoorWithWaterSurfaceElevationResp, error) {

		waterChannelDoors, err := getListWaterChannelDoorGateway(ctx, sharedGateway.GetListWaterChannelDoorReq{})
		if err != nil {
			return nil, err
		}

		waterChannelDoorsMap := make(map[int]sharedModel.WaterChannelDoor)
		for _, waterChannelDoor := range waterChannelDoors.WaterChannels {
			waterChannelDoorsMap[waterChannelDoor.ExternalID] = waterChannelDoor
		}

		waterSurfaceElevations, err := getListWaterChannelWaterSurfaceElevationGateway(ctx, sharedGateway.GetListWaterChannelTMAReq{})
		if err != nil {
			return nil, err
		}
		waterSurfaceElevationMap := make(map[int]sharedModel.WaterSurfaceElevationList)
		for _, waterSurfaceElevation := range waterSurfaceElevations.WaterChannelTMA {
			waterSurfaceElevationMap[waterSurfaceElevation.WaterChannelDoorID] = waterSurfaceElevation
		}

		var features []ListWaterChannelDoorWaterElevationFeature
		for _, waterChannelDoor := range waterChannelDoors.WaterChannels {
			waterChannelDetail := waterChannelDoorsMap[waterChannelDoor.ExternalID]
			waterSurfaceElevation := waterSurfaceElevationMap[waterChannelDetail.ExternalID]

			longitude, _ := strconv.ParseFloat(waterChannelDoor.Longitude, 64)
			latitude, _ := strconv.ParseFloat(waterChannelDoor.Latitude, 64)

			feature := ListWaterChannelDoorWaterElevationFeature{
				Type: "Feature",
				Properties: WaterChannelDoorWithWaterSurfaceElevationProperties{
					ID:                    waterChannelDoor.ExternalID,
					Name:                  waterChannelDoor.Name,
					WaterSurfaceElevation: waterSurfaceElevation.WaterSurfaceElevation,
				},
				Geometry: Geometry{
					Type:        "Point",
					Coordinates: []float64{longitude, latitude},
				},
			}
			features = append(features, feature)

		}

		geoJSON := GeoJSONFeatureWaterDoorWaterElevationCollection{
			Type:     "FeatureCollection",
			Features: features,
		}

		return &ListWaterChannelDoorWithWaterSurfaceElevationResp{
			GeoJSON: geoJSON,
		}, nil
	}
}
