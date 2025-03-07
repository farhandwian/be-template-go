package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"os"
	"shared/helper"
	"strconv"
	"time"
)

func (c Controller) LoginOTPSubmitHandler(u usecase.LoginOTPSubmit) helper.APIData {

	type Body struct {
		Email model.Email `json:"email"`
		OTP   string      `json:"otp"`
	}

	apiData := helper.APIData{
		Access:  model.ANONYMOUS,
		Method:  http.MethodPost,
		Url:     "/auth/login/otp",
		Body:    Body{},
		Summary: "Submit OTP for login",
		Tag:     "IAM - Authentication",
		Examples: []helper.ExampleResponse{
			{
				StatusCode: 200,
				Content: map[string]interface{}{
					"access_token":  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb250ZW50IjoiZXlKemRXSnFaV04wSWpvaVFVTkRSVk5UWDFSUFMwVk9JaXdpZFhObGNsOXBaQ0k2SW1JME9HVXpZalpsTFdSbFkyUXRORFpqTWkxaE0yVTJMVEJtTjJKa1ltTmtNVEppTlNJc0luVnpaWEpmWVdOalpYTnpJam9pTXlKOSIsImV4cCI6MTcyNzg4MjY5OX0.fpx7n6dwXmCgjG7M5i3auNlj81O2s7o-tygQcdEjL04",
					"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb250ZW50IjoiZXlKemRXSnFaV04wSWpvaVVrVkdVa1ZUU0Y5VVQwdEZUaUlzSW5WelpYSmZhV1FpT2lKaU5EaGxNMkkyWlMxa1pXTmtMVFEyWXpJdFlUTmxOaTB3WmpkaVpHSmpaREV5WWpVaUxDSjBiMnRsYmw5cFpDSTZJbVptTldZM05tUXlMVGN4TURNdE5EVXdOUzFoTnpZNExUSTBaVGsxWm1NeE0yTXpPQ0o5IiwiZXhwIjoxNzI3OTUxMDk5fQ.dquFpcH39kRLy6y5r3nJGBEGjyw86ysfZrSXYDHPa0A",
				},
			},
		},
	}

	refreshTokenInSecond, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_IN_SECOND"))
	if err != nil {
		panic(err)
	}

	accessTokenInSecond, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_IN_SECOND"))
	if err != nil {
		panic(err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.LoginOTPSubmitReq{
			Email:                body.Email,
			OTPValue:             body.OTP,
			RefreshTokenDuration: time.Duration(refreshTokenInSecond) * time.Second,
			AccessTokenDuration:  time.Duration(accessTokenInSecond) * time.Second,
			Now:                  time.Now(),
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
