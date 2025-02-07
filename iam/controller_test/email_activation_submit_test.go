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

	"github.com/joho/godotenv"
)

func TestEmailActivationSubmitIntegration(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	testCases := []struct {
		name            string
		activationToken string
		password        string
		pin             string
		expectedStatus  int
		mux             *http.ServeMux
	}{
		{
			name:            "Valid Activation Submit",
			activationToken: "valid-token",
			password:        "newpassword",
			pin:             "1234",
			expectedStatus:  http.StatusOK,
			mux: EmailActivationSubmitDependency(
				MockGateway(func(req gateway.ValidateJWTReq) (*gateway.ValidateJWTRes, error) {
					return &gateway.ValidateJWTRes{
						Payload: []byte(`{"subject":"EMAIL_ACTIVATION","user_id":"valid-user-id"}`),
					}, nil
				}),
				MockGateway(func(req gateway.UserGetOneByIDReq) (*gateway.UserGetOneByIDRes, error) {
					return &gateway.UserGetOneByIDRes{
						User: model.User{
							ID:    "valid-user-id",
							Email: "test@example.com",
						},
					}, nil
				}),
				MockGateway(func(req gateway.UserSaveReq) (*gateway.UserSaveRes, error) {

					return &gateway.UserSaveRes{}, nil
				}),
				MockGateway(func(req gateway.PasswordEncryptReq) (*gateway.PasswordEncryptRes, error) {
					return &gateway.PasswordEncryptRes{
						PasswordEncrypted: "encrypted_password",
					}, nil
				}),
			),
		},
		{
			name:            "Invalid Token",
			activationToken: "invalid-token",
			password:        "newpassword",
			pin:             "1234",
			expectedStatus:  http.StatusBadRequest,
			mux: EmailActivationSubmitDependency(
				MockGateway(func(req gateway.ValidateJWTReq) (*gateway.ValidateJWTRes, error) {
					return nil, errors.New("invalid token")
				}),
				MockGateway(func(req gateway.UserGetOneByIDReq) (*gateway.UserGetOneByIDRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.UserSaveReq) (*gateway.UserSaveRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.PasswordEncryptReq) (*gateway.PasswordEncryptRes, error) {
					return nil, nil
				}),
			),
		},
		{
			name:            "User Not Found",
			activationToken: "valid-token",
			password:        "newpassword",
			pin:             "1234",
			expectedStatus:  http.StatusBadRequest,
			mux: EmailActivationSubmitDependency(
				MockGateway(func(req gateway.ValidateJWTReq) (*gateway.ValidateJWTRes, error) {
					return &gateway.ValidateJWTRes{
						Payload: []byte(`{"subject":"EMAIL_ACTIVATION","user_id":"non-existent-user-id"}`),
					}, nil
				}),
				MockGateway(func(req gateway.UserGetOneByIDReq) (*gateway.UserGetOneByIDRes, error) {
					return nil, errors.New("user not found")
				}),
				MockGateway(func(req gateway.UserSaveReq) (*gateway.UserSaveRes, error) {
					return nil, nil
				}),
				MockGateway(func(req gateway.PasswordEncryptReq) (*gateway.PasswordEncryptRes, error) {
					return nil, nil
				}),
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body := struct {
				ActivationToken string `json:"activation_token"`
				Password        string `json:"password"`
				Pin             string `json:"pin"`
			}{
				ActivationToken: tc.activationToken,
				Password:        tc.password,
				Pin:             tc.pin,
			}

			bodyBytes, err := json.Marshal(body)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			req := httptest.NewRequest("POST", "/activation/submit", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			tc.mux.ServeHTTP(rr, req)

			if rr.Code != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, rr.Code)
			}

			if rr.Code == http.StatusOK {
				var response struct{}
				if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
					t.Errorf("Failed to decode response body: %v", err)
				}
			}
		})
	}
}

func EmailActivationSubmitDependency(
	validateJWT gateway.ValidateJWT,
	userGetOneByID gateway.UserGetOneByID,
	userSave gateway.UserSave,
	passwordEncrypt gateway.PasswordEncrypt,
) *http.ServeMux {
	mux := http.NewServeMux()

	validateJWTWithLogging := middleware.Logging(validateJWT, 4)
	userGetOneByIDWithLogging := middleware.Logging(userGetOneByID, 4)
	userSaveWithLogging := middleware.Logging(userSave, 4)
	passwordEncryptWithLogging := middleware.Logging(passwordEncrypt, 4)

	emailActivationSubmitUseCase := usecase.ImplEmailActivationSubmit(
		validateJWTWithLogging,
		userGetOneByIDWithLogging,
		userSaveWithLogging,
		passwordEncryptWithLogging,
	)

	emailActivationSubmitUseCaseWithLogging := middleware.Logging(emailActivationSubmitUseCase, 0)

	c := controller.Controller{
		Mux: mux,
	}

	helper.NewApiPrinter().
		Add(c.EmailActivationSubmitHandler(emailActivationSubmitUseCaseWithLogging)).
		Print()

	return mux
}
