package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"shared/core"
	"shared/model"
)

type GetDeviceStatusReq struct{}

type GetDeviceStatusRes struct {
	UpDeviceResponse   *model.PrometheusResponse
	DownDeviceResponse *model.PrometheusResponse
}

type GetDeviceStatusGateway = core.ActionHandler[GetDeviceStatusReq, GetDeviceStatusRes]

func ImplGetDeviceStatusGateway() GetDeviceStatusGateway {
	return func(ctx context.Context, request GetDeviceStatusReq) (*GetDeviceStatusRes, error) {

		url := os.Getenv("PROMETHEUS_URL")

		upDeviceUrl := fmt.Sprintf("%s/api/v1/query?query=%s", url, `count(avg_over_time(probe_success{job="blackbox"}[5m])>0)`)
		upDeviceRes, err := fetchPrometheusAPI(upDeviceUrl)
		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error fetching up devices: %v", err))
		}

		downDeviceUrl := fmt.Sprintf("%s/api/v1/query?query=%s", url, `count(avg_over_time(probe_success{job="blackbox"}[5m])==0)`)
		downDeviceRes, err := fetchPrometheusAPI(downDeviceUrl)
		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error fetching down devices: %v", err))
		}

		return &GetDeviceStatusRes{
			UpDeviceResponse:   upDeviceRes,
			DownDeviceResponse: downDeviceRes,
		}, nil

	}
}
func fetchPrometheusAPI(url string) (*model.PrometheusResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, core.NewInternalServerError(fmt.Errorf("error making request: %v", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, core.NewInternalServerError(fmt.Errorf("error reading response body: %v", err))
	}

	var promResp model.PrometheusResponse
	err = json.Unmarshal(body, &promResp)
	if err != nil {
		return nil, core.NewInternalServerError(fmt.Errorf("error unmarshaling JSON: %v", err))
	}

	return &promResp, nil
}
