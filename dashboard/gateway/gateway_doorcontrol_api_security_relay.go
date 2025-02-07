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

type DoorControlSecurityRelayReq struct {
	Status    int
	IPAddress string
	Now       time.Time
}

type DoorControlSecurityRelayRes struct {
	Response any
}

type DoorControlSecurityRelay = core.ActionHandler[DoorControlSecurityRelayReq, DoorControlSecurityRelayRes]

func ImplDoorControlSecurityRelay_(login LoginManganti) DoorControlSecurityRelay {
	return func(ctx context.Context, req DoorControlSecurityRelayReq) (*DoorControlSecurityRelayRes, error) {

		// TODO dummy response

		return &DoorControlSecurityRelayRes{Response: "this is dummy response"}, nil
	}
}

func ImplDoorControlSecurityRelay(login LoginManganti) DoorControlSecurityRelay {
	return func(ctx context.Context, req DoorControlSecurityRelayReq) (*DoorControlSecurityRelayRes, error) {

		// authObj, err := login(ctx, LoginMangantiReq{Now: req.Now})
		// if err != nil {
		// 	return nil, err
		// }

		// http://172.25.101.2:4101/api/perangkat/action/security
		// {"ip_address":"10.0.17.226","status":0}

		// Create the request body
		requestBody, err := json.Marshal(map[string]any{
			"ip_address": req.IPAddress,
			"status":     req.Status,
		})
		if err != nil {
			return nil, fmt.Errorf("error creating request body: %v", err)
		}

		// url := fmt.Sprintf("%s/api/perangkat/action/security", os.Getenv("MANGANTI_KONTROL_PINTU_URL"))
		url := fmt.Sprintf("http://opshi:poiuy09876@" + req.IPAddress + "/secrelay?key=" + os.Getenv("MANGANTI_CONTROLLER_API_KEY"))

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

		return &DoorControlSecurityRelayRes{Response: result}, nil
	}
}
