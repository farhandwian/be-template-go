package gateway

import (
	"context"
	"dashboard/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"shared/core"
)

type FetchBMKGWeatherLocationReq struct {
	Location string
}

type FetchBMKGWeatherLocationRes struct {
	Location []model.BmkgLocation
}

type FetchBMKGWeatherLocationGateway = core.ActionHandler[FetchBMKGWeatherLocationReq, FetchBMKGWeatherLocationRes]

func ImplFetchBMKGWeatherLocationGateway() FetchBMKGWeatherLocationGateway {
	return func(ctx context.Context, req FetchBMKGWeatherLocationReq) (*FetchBMKGWeatherLocationRes, error) {
		url := fmt.Sprintf("https://www.bmkg.go.id/api/cuaca/search?q=%s", req.Location)

		// Create a new HTTP request
		httpReq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error creating request: %v", err))
		}

		// Add the X-API-Key header
		httpReq.Header.Add("X-API-Key", "7Byoez9uECzCqtlqjH89fVtioyyySUIT")

		// Send the request
		client := &http.Client{}
		resp, err := client.Do(httpReq)
		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error making request: %v", err))
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error reading response body: %v", err))
		}

		var bmkgResp []model.BmkgLocation
		err = json.Unmarshal(body, &bmkgResp)
		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error unmarshaling JSON: %v\nRaw Response: %s", err, string(body)))
		}

		return &FetchBMKGWeatherLocationRes{
			Location: bmkgResp,
		}, nil
	}
}
