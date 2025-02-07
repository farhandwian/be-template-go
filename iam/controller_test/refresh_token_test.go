package controllertest

import (
	"bytes"
	"encoding/json"
	"errors"
	"iam/controller"
	"iam/gateway"
	"iam/model"
	"iam/usecase"
	"net/http"
	"net/http/httptest"
	"shared/helper"
	"shared/middleware"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestRefreshTokenIntegration(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	testCases := []struct {
		name           string
		refreshToken   string
		expectedStatus int
		expectedToken  bool
		mux            *http.ServeMux
	}{
		{
			name:           "Valid Refresh Token",
			refreshToken:   "valid-refresh-token",
			expectedStatus: http.StatusOK,
			expectedToken:  true,
			mux: RefreshTokenDependency(
				MockGateway(func(req gateway.UserGetOneByIDReq) (*gateway.UserGetOneByIDRes, error) {
					return &gateway.UserGetOneByIDRes{
						User: model.User{
							ID:              "test-user-id",
							Email:           "test@example.com",
							EmailVerifiedAt: time.Now().Add(-24 * time.Hour),
							Enabled:         true,
							RefreshTokenID:  "valid-refresh-token-id",
							UserAccess:      model.NewUserAccess(),
						},
					}, nil
				}),
				MockGateway(func(req gateway.GenerateJWTReq) (*gateway.GenerateJWTRes, error) {
					return &gateway.GenerateJWTRes{
						JWTToken: "new-access-token",
					}, nil
				}),
				MockGateway(func(req gateway.ValidateJWTReq) (*gateway.ValidateJWTRes, error) {
					return &gateway.ValidateJWTRes{
						Payload: []byte(`{"subject":"REFRESH_TOKEN","user_id":"test-user-id","token_id":"valid-refresh-token-id"}`),
					}, nil
				}),
			),
		},
		{
			name:           "Invalid Refresh Token",
			refreshToken:   "invalid-refresh-token",
			expectedStatus: http.StatusBadRequest,
			expectedToken:  false,
			mux: RefreshTokenDependency(
				MockGateway(func(req gateway.UserGetOneByIDReq) (*gateway.UserGetOneByIDRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.GenerateJWTReq) (*gateway.GenerateJWTRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.ValidateJWTReq) (*gateway.ValidateJWTRes, error) {
					return nil, errors.New("invalid token")
				}),
			),
		},
		{
			name:           "User Not Found",
			refreshToken:   "valid-refresh-token",
			expectedStatus: http.StatusBadRequest,
			expectedToken:  false,
			mux: RefreshTokenDependency(
				MockGateway(func(req gateway.UserGetOneByIDReq) (*gateway.UserGetOneByIDRes, error) {
					return nil, errors.New("user not found")
				}),
				MockGateway(func(req gateway.GenerateJWTReq) (*gateway.GenerateJWTRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.ValidateJWTReq) (*gateway.ValidateJWTRes, error) {
					return &gateway.ValidateJWTRes{
						Payload: []byte(`{"subject":"REFRESH_TOKEN","user_id":"non-existent-user-id","token_id":"valid-refresh-token-id"}`),
					}, nil
				}),
			),
		},
		{
			name:           "User Disabled",
			refreshToken:   "valid-refresh-token",
			expectedStatus: http.StatusBadRequest,
			expectedToken:  false,
			mux: RefreshTokenDependency(
				MockGateway(func(req gateway.UserGetOneByIDReq) (*gateway.UserGetOneByIDRes, error) {
					return &gateway.UserGetOneByIDRes{
						User: model.User{
							ID:              "test-user-id",
							Email:           "test@example.com",
							EmailVerifiedAt: time.Now().Add(-24 * time.Hour),
							Enabled:         false,
							RefreshTokenID:  "valid-refresh-token-id",
						},
					}, nil
				}),
				MockGateway(func(req gateway.GenerateJWTReq) (*gateway.GenerateJWTRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.ValidateJWTReq) (*gateway.ValidateJWTRes, error) {
					return &gateway.ValidateJWTRes{
						Payload: []byte(`{"subject":"REFRESH_TOKEN","user_id":"test-user-id","token_id":"valid-refresh-token-id"}`),
					}, nil
				}),
			),
		},
		{
			name:           "Invalid Refresh Token ID",
			refreshToken:   "valid-refresh-token",
			expectedStatus: http.StatusBadRequest,
			expectedToken:  false,
			mux: RefreshTokenDependency(
				MockGateway(func(req gateway.UserGetOneByIDReq) (*gateway.UserGetOneByIDRes, error) {
					return &gateway.UserGetOneByIDRes{
						User: model.User{
							ID:              "test-user-id",
							Email:           "test@example.com",
							EmailVerifiedAt: time.Now().Add(-24 * time.Hour),
							Enabled:         true,
							RefreshTokenID:  "different-refresh-token-id",
						},
					}, nil
				}),
				MockGateway(func(req gateway.GenerateJWTReq) (*gateway.GenerateJWTRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.ValidateJWTReq) (*gateway.ValidateJWTRes, error) {
					return &gateway.ValidateJWTRes{
						Payload: []byte(`{"subject":"REFRESH_TOKEN","user_id":"test-user-id","token_id":"valid-refresh-token-id"}`),
					}, nil
				}),
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body := struct {
				RefreshToken string `json:"refresh_token"`
			}{
				RefreshToken: tc.refreshToken,
			}

			bodyBytes, err := json.Marshal(body)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			req := httptest.NewRequest("POST", "/token/refresh", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			tc.mux.ServeHTTP(rr, req)

			if rr.Code != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, rr.Code)
			}

			if tc.expectedStatus == http.StatusOK {
				var response struct {
					AccessToken string `json:"AccessToken"`
				}

				if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
					t.Errorf("Failed to decode response body: %v", err)
				}

				if tc.expectedToken {
					if response.AccessToken == "" {
						t.Errorf("Expected non-empty AccessToken, got empty string")
					}
				} else {
					if response.AccessToken != "" {
						t.Errorf("Expected empty AccessToken, got %s", response.AccessToken)
					}
				}
			}
		})
	}
}

func RefreshTokenDependency(
	userGetOneByID gateway.UserGetOneByID,
	generateJWT gateway.GenerateJWT,
	validateJWT gateway.ValidateJWT,
) *http.ServeMux {
	mux := http.NewServeMux()

	userGetOneByIDWithLogging := middleware.Logging(userGetOneByID, 4)
	generateJWTWithLogging := middleware.Logging(generateJWT, 4)
	validateJWTWithLogging := middleware.Logging(validateJWT, 4)

	refreshTokenUseCase := usecase.ImplRefreshToken(
		userGetOneByIDWithLogging,
		generateJWTWithLogging,
		validateJWTWithLogging,
	)

	refreshTokenUseCaseWithLogging := middleware.Logging(refreshTokenUseCase, 0)

	mockJWTToken, _ := helper.NewJWTTokenizer("mock-secret-key")

	c := controller.Controller{
		Mux: mux,
		JWT: mockJWTToken,
	}

	helper.NewApiPrinter().
		Add(c.RefreshTokenHandler(refreshTokenUseCaseWithLogging)).
		Print()

	return mux
}
