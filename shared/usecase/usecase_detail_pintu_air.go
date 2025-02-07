package usecase

import (
	"context"
	"encoding/json"
	"shared/core"
	gateway2 "shared/gateway"
	"shared/model"
	"strconv"
)

type DetailPintuAirReq struct {
	WaterChannelDoorID int `json:"water_channel_door_id"`
}

type CCTV struct {
	ExternalID      int    `json:"external_id"`
	Name            string `json:"name"`       // get from WaterChannelDevice.Name where Category = "cctv" and WaterChannelDevice.WaterChannelDoorID = x
	IPAddress       string `json:"ip_address"` // get from WaterChannelDevice.IPAddress where Category = "cctv" and WaterChannelDevice.WaterChannelDoorID = x
	HumanDetected   bool   `json:"human_detected"`
	GarbageDetected bool   `json:"garbage_detected"`
}

type Officer struct {
	Name        string `json:"name"`         // get from WaterChannelOfficer.Name where WaterChannelOfficer.WaterChannelDoorID = x
	Photo       string `json:"photo"`        // get from WaterChannelOfficer.Photo where WaterChannelOfficer.WaterChannelDoorID = x
	PhoneNumber string `json:"phone_number"` // get from WaterChannelOfficer.PhoneNumber where WaterChannelOfficer.WaterChannelDoorID = x
	Task        string `json:"task"`         // get from WaterChannelOfficer.Task where WaterChannelOfficer.WaterChannelDoorID = x
}

type DetailPintuAirRes struct {
	WaterChannelDoorName string    `json:"water_channel_door_name"`
	WaterChannelDoorID   string    `json:"water_channel_door_id"` // get from WaterChannelDoor.ExternalID WaterChannelDoor.ExternalID = x
	PhotoUrls            []string  `json:"photo_urls"`            // get from WaterChannelDoor.Videos where WaterChannelDoor.ExternalID = x
	Officers             []Officer `json:"officers"`              // see `Officer` struct
	CCTVs                []CCTV    `json:"cctvs"`                 // see `CCTV` struct
	GarbageDetected      bool      `json:"garbage_detected"`      // check if any cctv has garbage
	HumanDetected        bool      `json:"human_detected"`        // check if any cctv has human
}

type DetailPintuAirUseCase = core.ActionHandler[DetailPintuAirReq, DetailPintuAirRes]

func ImplDetailPintuAir(
	getWaterChannelDoorByID gateway2.GetWaterChannelDoorByID,
	getWaterChannelDevicesByDoorID gateway2.GetWaterChannelDevicesByDoorID,
	getWaterChannelOfficersByDoorID gateway2.GetWaterChannelOfficersByDoorID,
) DetailPintuAirUseCase {
	return func(ctx context.Context, req DetailPintuAirReq) (*DetailPintuAirRes, error) {

		// Get WaterChannelDoor
		doorRes, err := getWaterChannelDoorByID(ctx, gateway2.GetWaterChannelDoorByIDReq{WaterChannelDoorID: req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}
		door := doorRes.WaterChannelDoor

		// Get WaterChannelDevices
		devicesRes, err := getWaterChannelDevicesByDoorID(ctx, gateway2.GetWaterChannelDevicesByDoorIDReq{WaterChannelDoorID: door.ExternalID})
		if err != nil {
			return nil, err
		}

		// Get WaterChannelOfficers
		officersRes, err := getWaterChannelOfficersByDoorID(ctx, gateway2.GetWaterChannelOfficersByDoorIDReq{WaterChannelDoorID: door.ExternalID})
		if err != nil {
			return nil, err
		}

		cctv := MapCCTVs(devicesRes.Devices)
		garbageDetected := false
		humanDetected := false
		//check if any cctv has garbage
		for _, c := range cctv {
			if c.GarbageDetected {
				garbageDetected = true
			}
			if c.HumanDetected {
				humanDetected = true
			}
		}

		// Prepare response
		response := &DetailPintuAirRes{
			WaterChannelDoorName: door.Name,
			WaterChannelDoorID:   strconv.Itoa(door.ExternalID),
			PhotoUrls:            getPhotoUrls(door.Photos),
			Officers:             mapOfficers(officersRes.Officers),
			CCTVs:                MapCCTVs(devicesRes.Devices),
			GarbageDetected:      garbageDetected,
			HumanDetected:        humanDetected,
		}

		return response, nil
	}
}

func getPhotoUrls(photosJSON []byte) []string {
	var photos []string
	json.Unmarshal(photosJSON, &photos)
	return photos
}

func ParseFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func mapOfficers(officers []model.WaterChannelOfficer) []Officer {
	result := make([]Officer, len(officers))
	for i, o := range officers {
		result[i] = Officer{
			Name:        o.Name,
			Photo:       o.Photo,
			PhoneNumber: o.PhoneNumber,
			Task:        o.Task,
		}
	}
	return result
}

func MapOfficersForAI(officers []model.WaterChannelOfficer) []Officer {
	result := make([]Officer, len(officers))
	for i, o := range officers {
		result[i] = Officer{
			Name:        o.Name,
			PhoneNumber: o.PhoneNumber,
			Task:        o.Task,
		}
	}
	return result
}

func MapCCTVs(devices []model.WaterChannelDevice) []CCTV {
	var cctvs []CCTV
	for _, d := range devices {
		if d.Category == "cctv" {

			garbageDetected := false
			humanDetected := false
			if d.DetectedObject == "garbage" {
				garbageDetected = true
			} else if d.DetectedObject == "human" {
				humanDetected = true
			}

			cctvs = append(cctvs, CCTV{
				ExternalID:      d.ExternalID,
				Name:            d.Name,
				IPAddress:       d.IPAddress,
				HumanDetected:   humanDetected,
				GarbageDetected: garbageDetected,
			})
		}
	}
	return cctvs
}
