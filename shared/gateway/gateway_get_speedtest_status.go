package gateway

import (
	"bigboard/model"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"shared/core"
)

type GetSpeedtestStatusReq struct {
}

type GetSpeedtestStatusRes struct {
	SpeedtestResponse model.SpeedtestResponse
}

type GetSpeedtestStatusGateway = core.ActionHandler[GetSpeedtestStatusReq, GetSpeedtestStatusRes]

func ImplGetSpeedtestStatusGateway() GetSpeedtestStatusGateway {
	return func(ctx context.Context, request GetSpeedtestStatusReq) (*GetSpeedtestStatusRes, error) {
		url := os.Getenv("SPEEDTEST_URL")

		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error sending request:", err)
			return nil, nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error reading response body: %v", err))
		}

		var speedtestResp model.SpeedtestResponse
		err = json.Unmarshal(body, &speedtestResp)
		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error unmarshaling JSON: %v", err))
		}

		return &GetSpeedtestStatusRes{
			SpeedtestResponse: speedtestResp,
		}, nil
	}
}
