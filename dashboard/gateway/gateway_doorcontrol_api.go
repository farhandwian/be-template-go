package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"shared/core"
	"time"
)

type DoorControlAPIReq struct {
	PerangkatID int
	IPAddress   string
	GateID      int
	Instruction string
	Position    float32
	Value       float32
	FullTime    float64
	MaxHeight   float64
	Limit       bool
	Now         time.Time
}

type DoorControlAPIRes struct {
	Response any
}

type DoorControlAPI = core.ActionHandler[DoorControlAPIReq, DoorControlAPIRes]

// func ImplDoorControlAPI(login LoginManganti) DoorControlAPI {
// 	return func(ctx context.Context, req DoorControlAPIReq) (*DoorControlAPIRes, error) {

// 		// TODO dummy response

// 		return &DoorControlAPIRes{Response: "this is dummy response"}, nil
// 	}
// }

func ImplDoorControlAPI_(login LoginManganti) DoorControlAPI {
	return func(ctx context.Context, req DoorControlAPIReq) (*DoorControlAPIRes, error) {

		deviceApiKey := os.Getenv("MANGANTI_CONTROLLER_API_KEY")
		// authObj, err := login(ctx, LoginMangantiReq{Now: req.Now})
		// if err != nil {
		// 	return nil, err
		// }

		// BCH_1 http://172.25.101.2:4104/pintu/14/perangkat/1987
		// http://172.25.101.2:4101/api/perangkat/action/instruction

		// {
		// 	"perangkat_id": 1987,
		// 	"ip_address": "10.0.15.26",
		// 	"gate": 2,
		// 	"instruction": "up",
		// 	"position": 28,
		// 	"value": 19,
		// 	"full_time": 60000,
		// 	"max_heigh": 100,
		// 	"limit": true
		// }

		// {
		// 	"perangkat_id": 1987,
		// 	"ip_address": "10.0.15.26",
		// 	"gate": 2,
		// 	"instruction": "stop",
		// 	"position": 0,
		// 	"value": 0,
		// 	"full_time": 60000,
		// 	"max_heigh": 100,
		// 	"limit": true
		// }

		// BCH_3 http://172.25.101.2:4104/pintu/16/perangkat/1583

		// {
		//   "perangkat_id": 1583,
		//   "ip_address": "10.0.14.234",
		//   "gate": 1,
		//   "instruction": "up",
		//   "position": 40,
		//   "value": 22,
		//   "full_time": 60000,
		//   "max_heigh": 100,
		//   "limit": true
		// }

		//==========================

		// {
		// 	"perangkat_id": 1974,         --> water_channel_devices.external_id
		// 	"ip_address"  : "10.0.8.226", --> water_channel_devices.ip_address
		// 	"full_time"   : 60000,        --> sementara asumsi konstan
		// 	"max_heigh"   : 100,          --> sementara asumsi konstan
		// 	"limit"       : true,         --> sementara asumsi konstan
		// 	"value"       : 19,           --> naikkan ke 19
		// 	"instruction" : "up",         --> naikkan ?, kalau turunkan 'down'? perlu test
		// 	"gate"        : 1,            --> apa ini ?
		// 	"position"    : 81,           --> apa ini ?
		// }

		// {
		// 	"water_channel_devices": [
		// 		{
		// 			"id" : "e180606d",
		// 			"external_id" : 1974,
		// 			"category" : "controller",
		// 			"name" : "BCH 58 KI",
		// 			"ip_address" : "10.0.8.226",
		// 			"group_relay" : 1,
		// 			"type" : "",
		// 			"full_time" : 60000,
		// 			"max_height_sensor" : 100,
		// 			"upper_limit" : 50,
		// 			"lower_limit" : 0,
		// 			"measurement_scale" : 0,
		// 			"water_channel_door_id" : 262,     --> sambung ke water_channel_doors
		// 		}
		// 	]
		// }

		// {
		// 	"water_channel_doors": [
		// 		{
		// 			"id" : "6f0ece8e",
		// 			"external_id" : 262,
		// 			"name" : "BCH 58 KI",
		// 			"latitude" : "-7.617821",
		// 			"longitude" : "108.931502",
		// 			"address" : "",
		// 			"ip_gateway" : "",
		// 			"photos" : "",
		// 			"width" : 0.0,
		// 			"cc" : "",
		// 			"smopi_channel_id" : 0,
		// 			"forecast_building_id" : null,
		// 			"area_size" : "55",
		// 			"debit_requirement" : "0",
		// 			"debit_actual" : "0",
		// 			"debit_recommendation" : "0",
		// 			"water_channel_id" : 2,            --> sambung ke water_channels
		// 		}
		// 	]
		// }

		// {
		// 	"water_channels": [
		// 		{
		// 			"id" : "f3775655",
		// 			"external_id" : 2,
		// 			"name" : "SALURAN PRIMER CIHAUR",
		// 			"address" : "UPTD SIDAREJA : BCH. 0 SAMPAI BCH. 39B, UPTD JERUKLEGI : BCH. 39C SAMPAI BCH. 75 ",
		// 			"photos" : "",
		// 			"irrigation_area_id" : 1,
		// 			"channel_id" : null,
		// 			"smopi_channel_id" : 0,
		// 		}
		// 	]
		// }

		// Create the request body
		requestBody, err := json.Marshal(map[string]any{
			"perangkat_id": req.PerangkatID,
			"ip_address":   req.IPAddress,
			"gate":         req.GateID,
			"instruction":  req.Instruction,
			"value":        req.Value,
			"position":     req.Position,
			"full_time":    req.FullTime,
			"max_heigh":    req.MaxHeight,
			"limit":        req.Limit,
		})

		if err != nil {
			return nil, fmt.Errorf("error creating request body: %v", err)
		}

		// url := fmt.Sprintf("%s/api/perangkat/action/instruction", os.Getenv("MANGANTI_KONTROL_PINTU_URL"))
		url := fmt.Sprintf("http://opshi:poiuy09876@" + req.IPAddress + "/instruction?key=" + deviceApiKey)

		request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Api-key", deviceApiKey)
		if err != nil {
			return nil, fmt.Errorf("error creating request: %w", err)
		}
		// request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authObj.Token))

		resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(request)
		if err != nil {
			return nil, fmt.Errorf("error making POST request: %v", err)
		}

		defer resp.Body.Close()

		// Check if the response status is OK (200)
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		// Read and print the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %v", err)
		}

		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, err
		}

		// handle if security relay is off
		if message, ok := result["message"].(string); ok {
			if message == "Maaf security relay dalam kondisi off" {
				return nil, fmt.Errorf("security relay is off")
			}
		}

		return &DoorControlAPIRes{Response: result}, nil
	}
}
