package gateway

import (
	"bigboard/model"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"shared/core"
)

type FetchBMKGWeatherForecastReq struct {
	ADM4 string
}

type FetchBMKGWeatherForecastRes struct {
	BMKGResponse model.BMKGResponse
}

type FetchBMKGWeatherForecast = core.ActionHandler[FetchBMKGWeatherForecastReq, FetchBMKGWeatherForecastRes]

func ImplFetchBMKGWeatherForecast() FetchBMKGWeatherForecast {
	return func(ctx context.Context, request FetchBMKGWeatherForecastReq) (*FetchBMKGWeatherForecastRes, error) {
		url := fmt.Sprintf("https://api.bmkg.go.id/publik/prakiraan-cuaca?adm4=%s", request.ADM4)

		resp, err := http.Get(url)
		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error making request: %v", err))
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error reading response body: %v", err))
		}

		var bmkgResp model.BMKGResponse
		err = json.Unmarshal(body, &bmkgResp)
		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error unmarshaling JSON: %v", err))
		}

		return &FetchBMKGWeatherForecastRes{
			BMKGResponse: bmkgResp,
		}, nil
	}
}
