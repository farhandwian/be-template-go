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

func TestLoginIntegration(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	testCases := []struct {
		name           string
		email          string
		password       string
		expectedStatus int
		mux            *http.ServeMux
	}{
		{
			name:           "Valid Login",
			email:          "test@example.com",
			password:       "validpassword",
			expectedStatus: http.StatusOK,
			mux: LoginDependency(
				MockGateway(func(req gateway.UserGetAllReq) (*gateway.UserGetAllRes, error) {
					return &gateway.UserGetAllRes{
						Count: 1,
						Items: []model.User{
							{
								ID:              "test-user-id",
								Email:           "test@example.com",
								Password:        "$2a$10$NMukYgyPsghIe2VjhZcJx.0NY5CAn/YM2jjgF66GQdIgORq9ZyGgC",
								EmailVerifiedAt: time.Now().Add(-24 * time.Hour),
								Enabled:         true,
							},
						},
					}, nil
				}),
				MockGateway(func(req gateway.PasswordValidateReq) (*gateway.PasswordValidateRes, error) {
					return &gateway.PasswordValidateRes{}, nil
				}),
				MockGateway(func(req gateway.PasswordEncryptReq) (*gateway.PasswordEncryptRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.GenerateRandomReq) (*gateway.GenerateRandomRes, error) {
					return &gateway.GenerateRandomRes{Random: "123456"}, nil
				}),
				MockGateway(func(req gateway.UserSaveReq) (*gateway.UserSaveRes, error) {
					return &gateway.UserSaveRes{}, nil
				}),
				MockGateway(func(req gateway.SendOTPReq) (*gateway.SendOTPRes, error) {
					return &gateway.SendOTPRes{}, nil
				}),
			),
		},
		{
			name:           "Valid Login but disabled",
			email:          "disabled@example.com",
			password:       "validpassword",
			expectedStatus: http.StatusBadRequest,
			mux: LoginDependency(
				MockGateway(func(req gateway.UserGetAllReq) (*gateway.UserGetAllRes, error) {
					return &gateway.UserGetAllRes{
						Count: 1,
						Items: []model.User{
							{
								ID:              "disabled-user-id",
								Email:           "disabled@example.com",
								Password:        "$2a$10$NMukYgyPsghIe2VjhZcJx.0NY5CAn/YM2jjgF66GQdIgORq9ZyGgC",
								EmailVerifiedAt: time.Now().Add(-24 * time.Hour),
								Enabled:         false,
							},
						},
					}, nil
				}),
				MockGateway(func(req gateway.PasswordValidateReq) (*gateway.PasswordValidateRes, error) {
					return &gateway.PasswordValidateRes{}, nil
				}),
				MockGateway(func(req gateway.PasswordEncryptReq) (*gateway.PasswordEncryptRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.GenerateRandomReq) (*gateway.GenerateRandomRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.UserSaveReq) (*gateway.UserSaveRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.SendOTPReq) (*gateway.SendOTPRes, error) {
					return nil, nil
				}),
			),
		},
		{
			name:           "Valid Login but not activated yet",
			email:          "unverified@example.com",
			password:       "validpassword",
			expectedStatus: http.StatusBadRequest,
			mux: LoginDependency(
				MockGateway(func(req gateway.UserGetAllReq) (*gateway.UserGetAllRes, error) {
					return &gateway.UserGetAllRes{
						Count: 1,
						Items: []model.User{
							{
								ID:       "unverified-user-id",
								Email:    "unverified@example.com",
								Password: "$2a$10$NMukYgyPsghIe2VjhZcJx.0NY5CAn/YM2jjgF66GQdIgORq9ZyGgC",
								Enabled:  true,
							},
						},
					}, nil
				}),
				MockGateway(func(req gateway.PasswordValidateReq) (*gateway.PasswordValidateRes, error) {
					return &gateway.PasswordValidateRes{}, nil
				}),
				MockGateway(func(req gateway.PasswordEncryptReq) (*gateway.PasswordEncryptRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.GenerateRandomReq) (*gateway.GenerateRandomRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.UserSaveReq) (*gateway.UserSaveRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.SendOTPReq) (*gateway.SendOTPRes, error) {
					return nil, nil
				}),
			),
		},
		{
			name:           "User not found",
			email:          "nonexistent@example.com",
			password:       "validpassword",
			expectedStatus: http.StatusBadRequest,
			mux: LoginDependency(
				MockGateway(func(req gateway.UserGetAllReq) (*gateway.UserGetAllRes, error) {
					return &gateway.UserGetAllRes{
						Count: 0,
						Items: []model.User{},
					}, nil
				}),
				MockGateway(func(req gateway.PasswordValidateReq) (*gateway.PasswordValidateRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.PasswordEncryptReq) (*gateway.PasswordEncryptRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.GenerateRandomReq) (*gateway.GenerateRandomRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.UserSaveReq) (*gateway.UserSaveRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.SendOTPReq) (*gateway.SendOTPRes, error) {
					return nil, nil
				}),
			),
		},
		{
			name:           "Invalid Email",
			email:          "invalid-email",
			password:       "password",
			expectedStatus: http.StatusBadRequest,
			mux: LoginDependency(
				MockGateway(func(req gateway.UserGetAllReq) (*gateway.UserGetAllRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.PasswordValidateReq) (*gateway.PasswordValidateRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.PasswordEncryptReq) (*gateway.PasswordEncryptRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.GenerateRandomReq) (*gateway.GenerateRandomRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.UserSaveReq) (*gateway.UserSaveRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.SendOTPReq) (*gateway.SendOTPRes, error) {
					return nil, nil
				}),
			),
		},
		{
			name:           "Invalid Password",
			email:          "test@example.com",
			password:       "invalidpassword",
			expectedStatus: http.StatusBadRequest,
			mux: LoginDependency(
				MockGateway(func(req gateway.UserGetAllReq) (*gateway.UserGetAllRes, error) {
					return &gateway.UserGetAllRes{
						Count: 1,
						Items: []model.User{
							{
								ID:              "test-user-id",
								Email:           "test@example.com",
								Password:        "$2a$10$NMukYgyPsghIe2VjhZcJx.0NY5CAn/YM2jjgF66GQdIgORq9ZyGgC",
								EmailVerifiedAt: time.Now().Add(-24 * time.Hour),
								Enabled:         true,
							},
						},
					}, nil
				}),
				MockGateway(func(req gateway.PasswordValidateReq) (*gateway.PasswordValidateRes, error) {
					return nil, errors.New("invalid password")
				}),
				MockGateway(func(req gateway.PasswordEncryptReq) (*gateway.PasswordEncryptRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.GenerateRandomReq) (*gateway.GenerateRandomRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.UserSaveReq) (*gateway.UserSaveRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.SendOTPReq) (*gateway.SendOTPRes, error) {
					return nil, nil
				}),
			),
		},
		{
			name:           "Empty Password",
			email:          "test@example.com",
			password:       "",
			expectedStatus: http.StatusBadRequest,
			mux: LoginDependency(
				MockGateway(func(req gateway.UserGetAllReq) (*gateway.UserGetAllRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.PasswordValidateReq) (*gateway.PasswordValidateRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.PasswordEncryptReq) (*gateway.PasswordEncryptRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.GenerateRandomReq) (*gateway.GenerateRandomRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.UserSaveReq) (*gateway.UserSaveRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.SendOTPReq) (*gateway.SendOTPRes, error) {
					return nil, nil
				}),
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body := struct {
				Email    model.Email `json:"email"`
				Password string      `json:"password"`
			}{
				Email:    model.Email(tc.email),
				Password: tc.password,
			}

			bodyBytes, err := json.Marshal(body)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			tc.mux.ServeHTTP(rr, req)

			if rr.Code != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, rr.Code)
			}

			if rr.Code == http.StatusOK {
				var loginResp struct{}
				if err := json.NewDecoder(rr.Body).Decode(&loginResp); err != nil {
					t.Errorf("Failed to decode response body: %v", err)
				}
			}
		})
	}
}

func LoginDependency(
	userGetAll gateway.UserGetAll,
	passwordValidate gateway.PasswordValidate,
	passwordEncrypt gateway.PasswordEncrypt,
	generateRandom gateway.GenerateRandom,
	userSave gateway.UserSave,
	sendOTP gateway.SendOTP,
) *http.ServeMux {
	mux := http.NewServeMux()

	userGetAllWithLogging := middleware.Logging(userGetAll, 4)
	passwordValidateWithLogging := middleware.Logging(passwordValidate, 4)
	passwordEncryptWithLogging := middleware.Logging(passwordEncrypt, 4)
	generateRandomWithLogging := middleware.Logging(generateRandom, 4)
	userSaveWithLogging := middleware.Logging(userSave, 4)
	sendOTPWithLogging := middleware.Logging(sendOTP, 4)

	loginUseCase := usecase.ImplLogin(
		userGetAllWithLogging,
		sendOTPWithLogging,
		generateRandomWithLogging,
		passwordValidateWithLogging,
		passwordEncryptWithLogging,
		userSaveWithLogging,
	)

	loginUseCaseWithLogging := middleware.Logging(loginUseCase, 0)

	c := controller.Controller{
		Mux: mux,
	}

	helper.NewApiPrinter("", "").
		Add(c.LoginHandler(loginUseCaseWithLogging)).
		Print()

	return mux
}
