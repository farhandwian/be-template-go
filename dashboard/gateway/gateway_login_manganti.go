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
	"shared/helper"
	"time"
)

type LoginMangantiReq struct {
	Now time.Time
}

type LoginMangantiRes struct {
	Token string
}

type LoginManganti = core.ActionHandler[LoginMangantiReq, LoginMangantiRes]

func ImplLoginMangantiEmpty() LoginManganti {

	return func(ctx context.Context, req LoginMangantiReq) (*LoginMangantiRes, error) {

		return &LoginMangantiRes{Token: ""}, nil
	}

}

func ImplLoginManganti() LoginManganti {

	var token = ""

	return func(ctx context.Context, req LoginMangantiReq) (*LoginMangantiRes, error) {

		if err := helper.IsTokenValid(token, req.Now); err == nil {
			return &LoginMangantiRes{Token: token}, nil
		}

		requestBody, err := json.Marshal(map[string]any{
			"username": os.Getenv("MANGANTI_BEARER_USERNAME"),
			"password": os.Getenv("MANGANTI_BEARER_PASSWORD"),
		})
		if err != nil {
			return nil, fmt.Errorf("error creating request body: %v", err)
		}

		url := fmt.Sprintf("%s/api/auth/login", os.Getenv("MANGANTI_KONTROL_PINTU_URL"))

		request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
		if err != nil {
			return nil, fmt.Errorf("error creating request: %w", err)
		}

		request.Header.Set("Content-Type", "application/json")

		resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(request)
		if err != nil {
			return nil, fmt.Errorf("error making POST request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %v", err)
		}

		type AuthData struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		}

		type AuthResponse struct {
			Status  string   `json:"status"`
			Message string   `json:"message"`
			Data    AuthData `json:"data"`
		}

		var authResponse AuthResponse
		if err := json.Unmarshal(body, &authResponse); err != nil {
			return nil, err
		}

		token = authResponse.Data.AccessToken

		fmt.Printf(">> token %v\n", token)

		return &LoginMangantiRes{Token: token}, nil
	}
}
