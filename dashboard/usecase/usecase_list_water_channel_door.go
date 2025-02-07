package usecase

import (
	"context"
	"dashboard/gateway"
	"shared/core"
	sharedGateway "shared/gateway"
	"shared/model"
	sharedModel "shared/model"
	"shared/usecase"
	"sort"
	"strconv"
	"strings"
)

type ListWaterChannelReq struct {
	Name              string
	MinWaterElevation float64
	MaxWaterElevation float64
	MinActualDebit    float64
	MaxActualDebit    float64
	MinRequiredDebit  float64
	MaxRequiredDebit  float64
	WaterChannelID    int
	Status            string
	Page              int
	PageSize          int
	SortBy            string
	SortOrder         string
}

type ListWaterChannelRes struct {
	GeoJSON  GeoJSONFeatureWaterDoorListCollection `json:"geo_json"`
	Metadata *usecase.Metadata                     `json:"metadata"`
}

type GeoJSONFeatureWaterDoorListCollection struct {
	Type     string                 `json:"type"`
	Features []ListWaterDoorFeature `json:"features"`
}

type ListWaterDoorFeature struct {
	Type       string              `json:"type"`
	Properties WaterDoorProperties `json:"properties"`
	Geometry   usecase.Geometry    `json:"geometry"`
}

type WaterDoorProperties struct {
	ID                    int     `json:"id"`
	WaterChannelID        int     `json:"water_channel_id"`
	WaterChannelName      string  `json:"water_channel_name"`
	WaterChannelDoorName  string  `json:"water_channel_door_name"`
	WaterSurfaceElevation float64 `json:"water_surface_elevation"`
	ActualDebit           float64 `json:"actual_debit"`
	RequiredDebit         float64 `json:"required_debit"`
	CCTVCount             int     `json:"cctv_count"`
	OfficerCount          int     `json:"officer_count"`
	StatusDebit           string  `json:"status_debit"`
	GarbageDetected       bool    `json:"garbage_detected"`
	HumanDetected         bool    `json:"human_detected"`
	StatusTma             bool    `json:"status_tma"`
}

type ListWaterChannelDoorUseCase = core.ActionHandler[ListWaterChannelReq, ListWaterChannelRes]

func ImplListWaterChannel(
	getListWaterChannelDoorGateway sharedGateway.GetListWaterChannelDoorGateway,
	getListWaterChannelTMAGateway sharedGateway.GetListWaterChannelTMAGateway,
	getListWaterChannelActualDebitGateway sharedGateway.GetListWaterChannelActualDebitGateway,
	getListCCTVCountGateway gateway.GetListCCTVCountGateway,
	getListOfficerCountGateway gateway.GetListWaterChannelDoorOfficerGateway,
	getCCTVDevices sharedGateway.GetCCTVDevices,
) ListWaterChannelDoorUseCase {
	return func(ctx context.Context, request ListWaterChannelReq) (*ListWaterChannelRes, error) {
		// Get water channel doors
		waterChannelDoors, err := getListWaterChannelDoorGateway(ctx, sharedGateway.GetListWaterChannelDoorReq{Name: request.Name})
		if err != nil {
			return nil, err
		}

		// Fetch CCTV devices
		cctvDevices, err := getCCTVDevices(ctx, sharedGateway.GetCCTVDevicesReq{})
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

		cctvCount, err := getListCCTVCountGateway(ctx, gateway.GetListCCTVCountReq{})
		if err != nil {
			return nil, err
		}
		waterChannelCCTVCountMap := make(map[int]int)
		for _, waterChannelCCTVCount := range cctvCount.CCTVDevices {
			waterChannelCCTVCountMap[waterChannelCCTVCount.WaterChannelDoorID] = waterChannelCCTVCount.CCTVCount
		}

		officerCount, err := getListOfficerCountGateway(ctx, gateway.GetListWaterChannelDoorOfficerReq{})
		if err != nil {
			return nil, err
		}
		waterChannelOfficerCountMap := make(map[int]int)
		for _, waterChannelOfficerCount := range officerCount.WaterChannels {
			waterChannelOfficerCountMap[waterChannelOfficerCount.WaterChannelDoorID] = waterChannelOfficerCount.OfficerCount
		}

		// Apply filtering logic
		features := filterWaterDoors(waterChannelDoors.WaterChannels,
			waterChannelTMAMap,
			waterChannelActualDebitMap,
			waterChannelCCTVCountMap,
			waterChannelOfficerCountMap,
			cctvDevices.Devices,
			request)

		//Apply sorting logic
		if request.SortBy != "" {
			sortWaterDoors(features, request)
		}

		// Pagination
		paginatedFeatures := paginateFeatures(features, request.Page, request.PageSize)

		// Create the response GeoJSON
		geoJSON := GeoJSONFeatureWaterDoorListCollection{
			Type:     "FeatureCollection",
			Features: paginatedFeatures,
		}

		// Calculate pagination details
		totalItems := len(features)
		totalPages := (totalItems + request.PageSize - 1) / request.PageSize

		return &ListWaterChannelRes{
			GeoJSON: geoJSON,
			Metadata: &usecase.Metadata{
				Pagination: usecase.Pagination{
					Page:       request.Page,
					Limit:      request.PageSize,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}

// Filtering function
func filterWaterDoors(
	waterChannelDoors []sharedModel.WaterChannelDoor,
	waterChannelTMAMap map[int]sharedModel.WaterSurfaceElevationList,
	waterChannelActualDebitMap map[int]sharedModel.ActualDebitData,
	waterChannelCCTVCountMap map[int]int,
	waterChannelOfficerCountMap map[int]int,
	cctvDevices []model.WaterChannelDevice,
	request ListWaterChannelReq,
) []ListWaterDoorFeature {
	var features []ListWaterDoorFeature
	for _, waterChannelDoor := range waterChannelDoors {
		waterChannelTMA := waterChannelTMAMap[waterChannelDoor.ExternalID]
		waterChannelActualDebit := waterChannelActualDebitMap[waterChannelDoor.ExternalID]
		waterChannelCCTVCount := waterChannelCCTVCountMap[waterChannelDoor.ExternalID]
		waterChannelOfficerCount := waterChannelOfficerCountMap[waterChannelDoor.ExternalID]

		//filter criteria by water channel ID
		if request.WaterChannelID != 0 && waterChannelDoor.WaterChannelID != request.WaterChannelID {
			continue // Skip entry if it doesn't match the WaterChannelID filter
		}

		// filter criteria by water elevation
		if (request.MinWaterElevation != 0 && waterChannelTMA.WaterSurfaceElevation < request.MinWaterElevation) ||
			(request.MaxWaterElevation != 0 && waterChannelTMA.WaterSurfaceElevation > request.MaxWaterElevation) {
			continue // Skip entry if it doesn't match the filters
		}

		if (request.MinActualDebit != 0 && waterChannelActualDebit.ActualDebit < request.MinActualDebit) ||
			(request.MaxActualDebit != 0 && waterChannelActualDebit.ActualDebit > request.MaxActualDebit) {
			continue // Skip entry if it doesn't match the filters
		}

		currentRequiredDebit, _ := strconv.ParseFloat(waterChannelDoor.DebitRequirement, 32)
		if (request.MinRequiredDebit != 0 && currentRequiredDebit < request.MinRequiredDebit) ||
			(request.MaxRequiredDebit != 0 && currentRequiredDebit > request.MaxRequiredDebit) {
			continue
		}

		// if request.Status != "" && strings.ToLower(request.Status) != strings.ToLower(getWaterChannelDoorStatus(waterChannelActualDebit.ActualDebit, currentRequiredDebit)) {
		// 	continue
		// }

		if request.Status != "" && !strings.EqualFold(request.Status, getWaterChannelDoorStatus(waterChannelActualDebit.ActualDebit, currentRequiredDebit)) {
			continue
		}

		garbageDetected := false
		humanDetected := false
		for _, cctvDevice := range cctvDevices {
			if cctvDevice.WaterChannelDoorID == waterChannelDoor.ExternalID {
				if cctvDevice.DetectedObject == "garbage" {
					garbageDetected = true
				} else if cctvDevice.DetectedObject == "human" {
					humanDetected = true
				}
			}
		}

		longitude, _ := strconv.ParseFloat(strings.TrimSpace(waterChannelDoor.Longitude), 64)
		latitude, _ := strconv.ParseFloat(strings.TrimSpace(waterChannelDoor.Latitude), 64)

		requiredDebit, _ := strconv.ParseFloat(waterChannelDoor.DebitRequirement, 32)
		feature := ListWaterDoorFeature{
			Type: "Feature",
			Properties: WaterDoorProperties{
				ID:                    waterChannelDoor.ExternalID,
				WaterChannelName:      waterChannelDoor.WaterChannelName,
				WaterChannelID:        waterChannelDoor.WaterChannelID,
				WaterChannelDoorName:  waterChannelDoor.Name,
				WaterSurfaceElevation: waterChannelTMA.WaterSurfaceElevation,
				ActualDebit:           waterChannelActualDebit.ActualDebit,
				CCTVCount:             waterChannelCCTVCount,
				OfficerCount:          waterChannelOfficerCount,
				RequiredDebit:         requiredDebit,
				StatusDebit:           getWaterChannelDoorStatus(waterChannelActualDebit.ActualDebit, requiredDebit),
				GarbageDetected:       garbageDetected,
				HumanDetected:         humanDetected,
				StatusTma:             waterChannelTMA.Status,
			},
			Geometry: usecase.Geometry{
				Type:        "Point",
				Coordinates: []float64{longitude, latitude},
			},
		}
		features = append(features, feature)
	}
	return features
}

// Sorting function
func sortWaterDoors(features []ListWaterDoorFeature, request ListWaterChannelReq) {
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
		case "water_channel_name":
			if request.SortOrder == "asc" {
				return strings.TrimSpace(features[i].Properties.WaterChannelName) < strings.TrimSpace(features[j].Properties.WaterChannelName)
			}
			return strings.TrimSpace(features[i].Properties.WaterChannelName) > strings.TrimSpace(features[j].Properties.WaterChannelName)
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
		case "status_debit":
			if request.SortOrder == "asc" {
				return features[i].Properties.StatusDebit < features[j].Properties.StatusDebit
			}
			return features[i].Properties.StatusDebit > features[j].Properties.StatusDebit
		case "status_tma":
			if request.SortOrder == "asc" {
				return !features[i].Properties.StatusTma && features[j].Properties.StatusTma
			}
			return features[i].Properties.StatusTma && !features[j].Properties.StatusTma
		}
		return false
	})
}

// Pagination function
func paginateFeatures(features []ListWaterDoorFeature, page, pageSize int) []ListWaterDoorFeature {
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

func getWaterChannelDoorStatus(actualDebit, requiredDebit float64) string {
	if actualDebit < requiredDebit {
		return "Tidak Terpenuhi"
	} else {
		return "Terpenuhi"
	}
}
