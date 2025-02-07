package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"shared/core"
	"time"
)

type DoorControlSensorReq struct {
	Status    int
	IPAddress string
	Now       time.Time
}

type DoorControlSensorRes struct {
	Response any
}

type DoorControlSensor = core.ActionHandler[DoorControlSensorReq, DoorControlSensorRes]

func ImplDoorControlSensor_(login LoginManganti) DoorControlSensor {
	return func(ctx context.Context, req DoorControlSensorReq) (*DoorControlSensorRes, error) {

		// TODO dummy response

		return &DoorControlSensorRes{Response: "this is dummy response"}, nil
	}
}

func ImplDoorControlSensor(login LoginManganti) DoorControlSensor {
	return func(ctx context.Context, req DoorControlSensorReq) (*DoorControlSensorRes, error) {

		// authObj, err := login(ctx, LoginMangantiReq{Now: req.Now})
		// if err != nil {
		// 	return nil, err
		// }

		// http://172.25.101.2:4101/api/perangkat/action/seteeprom
		// {"ip_address":"10.0.17.226","index":3,"value":0}

		// Create the request body
		requestBody, err := json.Marshal(map[string]any{
			"ip_address": req.IPAddress,
			"value":      req.Status,
			"index":      3, // TODO darimana nilai index seteeprom ini?
		})
		if err != nil {
			return nil, fmt.Errorf("error creating request body: %v", err)
		}

		// url := fmt.Sprintf("%s/api/perangkat/action/seteeprom", os.Getenv("MANGANTI_KONTROL_PINTU_URL"))
		url := fmt.Sprintf("http://opshi:poiuy09876@" + req.IPAddress + "/seteeprom")

		request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
		if err != nil {
			return nil, fmt.Errorf("error creating request: %w", err)
		}

		request.Header.Set("Content-Type", "application/json")
		// request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authObj.Token))

		resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(request)
		if err != nil {
			return nil, fmt.Errorf("error making POST request: %v", err)
		}
		fmt.Println("resp", resp)
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

		var result any
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, err
		}

		return &DoorControlSensorRes{Response: result}, nil
	}
}
