package usecase

import (
	"context"
	"shared/core"
	sharedGateway "shared/gateway"
	sharedModel "shared/model"
	"shared/usecase"
	"sort"
	"strconv"
)

type ListMainWaterChannelDoorReq struct {
	SortBy    string
	SortOrder string
}

type ListMainWaterChannelDoorMainRes struct {
	GeoJSON  GeoJSONFeatureMainWaterDoorListCollection `json:"geo_json"`
	Metadata *usecase.Metadata                         `json:"metadata"`
}

type GeoJSONFeatureMainWaterDoorListCollection struct {
	Type     string                     `json:"type"`
	Features []ListMainWaterDoorFeature `json:"features"`
}

type ListMainWaterDoorFeature struct {
	Type       string                  `json:"type"`
	Properties MainWaterDoorProperties `json:"properties"`
}

type MainWaterDoorProperties struct {
	ID                    int     `json:"id"`
	WaterChannelDoorName  string  `json:"water_channel_door_name"`
	WaterSurfaceElevation float64 `json:"water_surface_elevation"`
	ActualDebit           float64 `json:"actual_debit"`
	RequiredDebit         float64 `json:"required_debit"`
	Status                string  `json:"status"`
}

type ListMainWaterChannelDoorUseCase = core.ActionHandler[ListMainWaterChannelDoorReq, ListMainWaterChannelDoorMainRes]

func ImplListMainWaterChannelDoorUseCase(
	getListWaterChannelDoorGateway sharedGateway.GetListWaterChannelDoorGateway,
	getListWaterChannelTMAGateway sharedGateway.GetListWaterChannelTMAGateway,
	getListWaterChannelActualDebitGateway sharedGateway.GetListWaterChannelActualDebitGateway,
) ListMainWaterChannelDoorUseCase {
	return func(ctx context.Context, request ListMainWaterChannelDoorReq) (*ListMainWaterChannelDoorMainRes, error) {

		defaultPage := 1
		defaultPerPage := 10

		// Get water channel doors
		waterChannelDoors, err := getListWaterChannelDoorGateway(ctx, sharedGateway.GetListWaterChannelDoorReq{})
		if err != nil {
			return nil, err
		}

		// Fetch water channel TMA using the new function
		waterChannelTMA, err := getListWaterChannelTMAGateway(ctx, sharedGateway.GetListWaterChannelTMAReq{})
		if err != nil {
			return nil, err
		}
		waterChannelTMAMap := make(map[int]sharedModel.WaterSurfaceElevationList)
		for _, waterChannelTMA := range waterChannelTMA.WaterChannelTMA {
			waterChannelTMAMap[waterChannelTMA.WaterChannelDoorID] = waterChannelTMA
		}

		waterChannelActualDebit, err := getListWaterChannelActualDebitGateway(ctx, sharedGateway.GetListWaterChannelActualDebitReq{})
		if err != nil {
			return nil, err
		}
		waterChannelActualDebitMap := make(map[int]sharedModel.ActualDebitData)
		for _, waterChannelActualDebit := range waterChannelActualDebit.WaterChannelActualDebit {
			waterChannelActualDebitMap[waterChannelActualDebit.WaterChannelDoorID] = waterChannelActualDebit
		}

		features := filterWaterDoorsMain(waterChannelDoors.WaterChannels, waterChannelTMAMap, waterChannelActualDebitMap)

		//Apply sorting logic
		if request.SortBy != "" {
			sortWaterDoorsMain(features, request)
		}

		// Pagination
		paginatedFeatures := paginateMainFeatures(features, defaultPage, defaultPerPage)

		// Create the response GeoJSON
		geoJSON := GeoJSONFeatureMainWaterDoorListCollection{
			Type:     "FeatureCollection",
			Features: paginatedFeatures,
		}

		// Calculate pagination details
		totalItems := len(features)
		totalPages := (totalItems + defaultPerPage - 1) / defaultPerPage

		return &ListMainWaterChannelDoorMainRes{
			GeoJSON: geoJSON,
			Metadata: &usecase.Metadata{
				Pagination: usecase.Pagination{
					Page:       defaultPage,
					Limit:      defaultPerPage,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}

// filter only water doors only certain ids, 2,4,and 5
func filterWaterDoorsMain(waterChannelDoors []sharedModel.WaterChannelDoor,
	waterChannelTMAMap map[int]sharedModel.WaterSurfaceElevationList,
	waterChannelActualDebitMap map[int]sharedModel.ActualDebitData) []ListMainWaterDoorFeature {

	filteredFeatures := make([]ListMainWaterDoorFeature, 0)
	for _, waterChannelDoor := range waterChannelDoors {

		if waterChannelDoor.ExternalID == 2 || waterChannelDoor.ExternalID == 4 || waterChannelDoor.ExternalID == 5 {
			requiredDebit, _ := strconv.ParseFloat(waterChannelDoor.DebitRequirement, 32)
			waterChannelTMA := waterChannelTMAMap[waterChannelDoor.ExternalID]
			waterChannelActualDebit := waterChannelActualDebitMap[waterChannelDoor.ExternalID]

			feature := ListMainWaterDoorFeature{
				Type: "Feature",
				Properties: MainWaterDoorProperties{
					ID:                    waterChannelDoor.ExternalID,
					WaterChannelDoorName:  waterChannelDoor.Name,
					RequiredDebit:         requiredDebit,
					ActualDebit:           waterChannelActualDebit.ActualDebit,
					WaterSurfaceElevation: waterChannelTMA.WaterSurfaceElevation,
					Status:                getWaterChannelDoorStatus(waterChannelTMA.WaterSurfaceElevation, requiredDebit),
				},
			}
			filteredFeatures = append(filteredFeatures, feature)
		}

	}
	return filteredFeatures
}

// Sorting function
func sortWaterDoorsMain(features []ListMainWaterDoorFeature, request ListMainWaterChannelDoorReq) {
	sort.Slice(features, func(i, j int) bool {
		switch request.SortBy {
		case "id":
			if request.SortOrder == "asc" {
				return features[i].Properties.ID < features[j].Properties.ID
			}
			return features[i].Properties.ID > features[j].Properties.ID
		case "water_channel_door_name":
			if request.SortOrder == "asc" {
				return features[i].Properties.WaterChannelDoorName < features[j].Properties.WaterChannelDoorName
			}
			return features[i].Properties.WaterChannelDoorName > features[j].Properties.WaterChannelDoorName
		case "water_surface_elevation":
			if request.SortOrder == "asc" {
				return features[i].Properties.WaterSurfaceElevation < features[j].Properties.WaterSurfaceElevation
			}
			return features[i].Properties.WaterSurfaceElevation > features[j].Properties.WaterSurfaceElevation
		case "actual_debit":
			if request.SortOrder == "asc" {
				return features[i].Properties.ActualDebit < features[j].Properties.ActualDebit
			}
			return features[i].Properties.ActualDebit > features[j].Properties.ActualDebit
		case "required_debit":
			if request.SortOrder == "asc" {
				return features[i].Properties.RequiredDebit < features[j].Properties.RequiredDebit
			}
			return features[i].Properties.RequiredDebit > features[j].Properties.RequiredDebit
		case "status":
			if request.SortOrder == "asc" {
				return features[i].Properties.Status < features[j].Properties.Status
			}
			return features[i].Properties.Status > features[j].Properties.Status
		}
		return false
	})
}

// Pagination function
func paginateMainFeatures(features []ListMainWaterDoorFeature, page, pageSize int) []ListMainWaterDoorFeature {
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(features) {
		start = len(features) // Ensure start does not exceed number of features
	}
	if end > len(features) {
		end = len(features) // Ensure end does not exceed number of features
	}

	return features[start:end]
}

// func getWaterChannelDoorMainStatus(actualDebit, requiredDebit float64) string {
// 	if actualDebit < requiredDebit {
// 		return "Tidak Terpenuhi"
// 	} else {
// 		return "Terpenuhi"
// 	}
// }
