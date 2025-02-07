package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"shared/core"
	"strings"
	"time"
)

type CheckAccessReq struct {
	FunctionName string
	Token        string
}

type Data struct {
	CanAccess bool `json:"can_access"`
}

type CheckAccessRes struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   Data    `json:"data"`
}

type CheckAccess = core.ActionHandler[CheckAccessReq, CheckAccessRes]

func ImplCheckAccess() CheckAccess {

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	return func(ctx context.Context, request CheckAccessReq) (*CheckAccessRes, error) {

		apiReq := map[string]any{
			"function_name": request.FunctionName,
		}

		// Marshal request body ke JSON
		jsonBody, err := json.Marshal(apiReq)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request: %w", err)
		}

		url := fmt.Sprintf("%s/access", os.Getenv("ACCESS_ENDPOINT"))

		// Membuat HTTP request
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonBody))
		if err != nil {
			return nil, fmt.Errorf("error creating request: %w", err)
		}

		// Set headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", request.Token))

		// Melakukan request
		resp, err := httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error making request: %w", err)
		}
		defer resp.Body.Close()

		// Mengecek status code
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		// Decode response
		var apiResp CheckAccessRes
		if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}

		return &apiResp, nil
	}
}

func GetBearerToken(w http.ResponseWriter, r *http.Request) (string, string, bool) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", "authorization header required", false
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		return "", "invalid Authorization header format", false
	}

	return bearerToken[1], "", true
}
