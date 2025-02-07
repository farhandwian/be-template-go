package usecase

import (
	"context"
	"shared/core"
	sharedGateway "shared/gateway"
	sharedModel "shared/model"
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
	StatusTMA         string
	StatusDebit       string
	HasGarbage        bool
}

type ListWaterChannelRes struct {
	GeoJSON GeoJSONFeatureWaterDoorListCollection `json:"geo_json"`
}

type GeoJSONFeatureWaterDoorListCollection struct {
	Type     string                 `json:"type"`
	Features []ListWaterDoorFeature `json:"features"`
}

type ListWaterDoorFeature struct {
	Type       string              `json:"type"`
	Properties WaterDoorProperties `json:"properties"`
	Geometry   Geometry            `json:"geometry"`
}

type WaterDoorProperties struct {
	ID                    int     `json:"id"`
	Name                  string  `json:"name"`
	WaterSurfaceElevation float64 `json:"water_surface_elevation"`
	StatusTMA             bool    `json:"status_tma"`
	ActualDebit           float64 `json:"actual_debit"`
	// CCTVCount             int       `json:"cctv_count"`
	// OfficerCount          int       `json:"officer_count"`
	StatusDebit   string  `json:"status_debit"`
	RequiredDebit float64 `json:"required_debit"`
	// WaterChannelName      string    `json:"water_channel_door_name"`
	// WaterChannelAddress   string    `json:"water_channel_door_address"`
	// WaterGates            []float64 `json:"water_gates"`
	GarbageDetected bool `json:"garbage_detected"`
	HumanDetected   bool `json:"human_detected"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type ListWaterChannelUseCase = core.ActionHandler[ListWaterChannelReq, ListWaterChannelRes]

// TODO sasaran optimasi

func ImplListWaterChannel(
	getListWaterChannelDoorGateway sharedGateway.GetListWaterChannelDoorGateway,
	// getListWaterChannelCCTVCountGateway gateway.GetListWaterChannelDoorCCTVCountGateway,
	// getListWaterChannelOfficerCountGateway gateway.GetListWaterChannelDoorOfficerCountGateway,
	getListWaterChannelTMAGateway sharedGateway.GetListWaterChannelTMAGateway,
	getListWaterChannelActualDebitGateway sharedGateway.GetListWaterChannelActualDebitGateway,
	getCCTVDevices sharedGateway.GetCCTVDevices,
	getConfig sharedGateway.GetConfigGateway,

) ListWaterChannelUseCase {
	return func(ctx context.Context, request ListWaterChannelReq) (*ListWaterChannelRes, error) {

		// get config
		isGuestMode := false
		config, err := getConfig(ctx, sharedGateway.GetConfigReq{Name: "GUEST_MODE"})
		if err == nil {
			isGuestMode = config.Config.Value == "1"
		}

		// maria_db
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

		// Fetch water channel actual debit using the new function
		waterChannelActualDebit, err := getListWaterChannelActualDebitGateway(ctx, sharedGateway.GetListWaterChannelActualDebitReq{})
		if err != nil {
			return nil, err
		}
		waterChannelActualDebitMap := make(map[int]sharedModel.ActualDebitData)
		for _, waterChannelActualDebit := range waterChannelActualDebit.WaterChannelActualDebit {
			waterChannelActualDebitMap[waterChannelActualDebit.WaterChannelDoorID] = waterChannelActualDebit
		}

		features := filterWaterDoors(
			waterChannelDoors.WaterChannels,
			waterChannelTMAMap,
			waterChannelActualDebitMap,
			cctvDevices.Devices,
			request,
			isGuestMode)

		geoJSON := GeoJSONFeatureWaterDoorListCollection{
			Type:     "FeatureCollection",
			Features: features,
		}

		return &ListWaterChannelRes{
			GeoJSON: geoJSON,
		}, nil
	}
}

func getWaterChannelDoorStatus(actualDebit, requiredDebit float64) string {
	if actualDebit < requiredDebit {
		return "Tidak Terpenuhi"
	} else {
		return "Terpenuhi"
	}
}

// Filtering function
func filterWaterDoors(
	waterChannelDoors []sharedModel.WaterChannelDoor,
	waterChannelTMAMap map[int]sharedModel.WaterSurfaceElevationList,
	waterChannelActualDebitMap map[int]sharedModel.ActualDebitData,
	cctvDevices []sharedModel.WaterChannelDevice,
	request ListWaterChannelReq,
	isGuestMode bool,
) []ListWaterDoorFeature {
	var features []ListWaterDoorFeature
	for _, waterChannelDoor := range waterChannelDoors {
		waterChannelTMA := waterChannelTMAMap[waterChannelDoor.ExternalID]
		waterChannelActualDebit := waterChannelActualDebitMap[waterChannelDoor.ExternalID]

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

		if request.StatusDebit != "" && !strings.EqualFold(request.StatusDebit, getWaterChannelDoorStatus(waterChannelActualDebit.ActualDebit, currentRequiredDebit)) {
			continue
		}

		if request.StatusTMA != "" && strings.EqualFold(request.StatusTMA, "online") && !waterChannelTMA.Status {
			continue
		}

		if request.StatusTMA != "" && !strings.EqualFold(request.StatusTMA, "offline") && waterChannelTMA.Status {
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

		if request.HasGarbage && !garbageDetected {
			continue
		}

		statusTMA := waterChannelTMA.Status
		if isGuestMode {
			statusTMA = true
		}

		longitude, _ := strconv.ParseFloat(strings.TrimSpace(waterChannelDoor.Longitude), 64)
		latitude, _ := strconv.ParseFloat(strings.TrimSpace(waterChannelDoor.Latitude), 64)

		requiredDebit, _ := strconv.ParseFloat(waterChannelDoor.DebitRequirement, 32)
		feature := ListWaterDoorFeature{
			Type: "Feature",
			Properties: WaterDoorProperties{
				ID:   waterChannelDoor.ExternalID,
				Name: waterChannelDoor.Name,
				//WaterChannelAddress:     waterChannelDoor.Address,
				//WaterChannelName:        waterChannelDoor.WaterChannelName,
				WaterSurfaceElevation: waterChannelTMA.WaterSurfaceElevation,
				ActualDebit:           waterChannelActualDebit.ActualDebit,
				// CCTVCount:             cctvCount,
				// OfficerCount:          officerCount,
				RequiredDebit:   requiredDebit,
				StatusDebit:     getWaterChannelDoorStatus(waterChannelActualDebit.ActualDebit, requiredDebit),
				GarbageDetected: garbageDetected,
				HumanDetected:   humanDetected,
				StatusTMA:       statusTMA,
			},
			Geometry: Geometry{
				Type:        "Point",
				Coordinates: []float64{longitude, latitude},
			},
		}
		features = append(features, feature)
	}
	return features
}
