package usecase

import (
	"context"
	"shared/core"
	sharedGateway "shared/gateway"
	sharedModel "shared/model"
	"strconv"
)

type ListWaterChannelDoorWithActualDebitReq struct{}

type ListWaterChannelDoorWithActualDebitResp struct {
	GeoJSON GeoJSONFeatureWaterDoorDebitListCollection `json:"geo_json"`
}

type GeoJSONFeatureWaterDoorDebitListCollection struct {
	Type     string                      `json:"type"`
	Features []ListWaterDoorDebitFeature `json:"features"`
}

type ListWaterDoorDebitFeature struct {
	Type       string                                    `json:"type"`
	Properties WaterChannelDoorWithActualDebitProperties `json:"properties"`
	Geometry   Geometry                                  `json:"geometry"`
}

type WaterChannelDoorWithActualDebitProperties struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ActualDebit float64 `json:"actual_debit"`
}

type ListWaterChannelDoorWithActualDebitUseCase = core.ActionHandler[
	ListWaterChannelDoorWithActualDebitReq,
	ListWaterChannelDoorWithActualDebitResp,
]

func ImplListWaterChannelDoorWithActualDebit(
	getListWaterChannelDoorGateway sharedGateway.GetListWaterChannelDoorGateway,
	getListWaterChannelDoorDebit sharedGateway.GetListWaterChannelActualDebitGateway,
) ListWaterChannelDoorWithActualDebitUseCase {

	return func(ctx context.Context, request ListWaterChannelDoorWithActualDebitReq) (*ListWaterChannelDoorWithActualDebitResp, error) {

		waterChannelDoors, err := getListWaterChannelDoorGateway(ctx, sharedGateway.GetListWaterChannelDoorReq{})
		if err != nil {
			return nil, err
		}

		waterChannelDoorsMap := make(map[int]sharedModel.WaterChannelDoor)
		for _, waterChannelDoor := range waterChannelDoors.WaterChannels {
			waterChannelDoorsMap[waterChannelDoor.ExternalID] = waterChannelDoor
		}

		debits, err := getListWaterChannelDoorDebit(ctx, sharedGateway.GetListWaterChannelActualDebitReq{})
		if err != nil {
			return nil, err
		}
		debitMap := make(map[int]sharedModel.ActualDebitData)
		for _, waterSurfaceElevation := range debits.WaterChannelActualDebit {
			debitMap[waterSurfaceElevation.WaterChannelDoorID] = waterSurfaceElevation
		}

		var features []ListWaterDoorDebitFeature
		for _, waterChannelDoor := range waterChannelDoors.WaterChannels {
			waterChannelDetail := waterChannelDoorsMap[waterChannelDoor.ExternalID]
			waterSurfaceElevationAndDebit := debitMap[waterChannelDetail.ExternalID]

			longitude, _ := strconv.ParseFloat(waterChannelDoor.Longitude, 64)
			latitude, _ := strconv.ParseFloat(waterChannelDoor.Latitude, 64)

			feature := ListWaterDoorDebitFeature{
				Type: "Feature",
				Properties: WaterChannelDoorWithActualDebitProperties{
					ID:          waterChannelDetail.ExternalID,
					Name:        waterChannelDetail.Name,
					ActualDebit: waterSurfaceElevationAndDebit.ActualDebit,
				},
				Geometry: Geometry{
					Type: "Point",
					Coordinates: []float64{
						longitude,
						latitude,
					},
				},
			}
			features = append(features, feature)
		}

		geoJSON := GeoJSONFeatureWaterDoorDebitListCollection{
			Type:     "FeatureCollection",
			Features: features,
		}

		return &ListWaterChannelDoorWithActualDebitResp{
			GeoJSON: geoJSON,
		}, nil
	}
}
