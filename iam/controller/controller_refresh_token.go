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

func (c Controller) RefreshTokenHandler(u usecase.RefreshToken) helper.APIData {

	type Body struct {
		RefreshToken string `json:"refresh_token"`
	}

	apiData := helper.APIData{
		Access:  model.ANONYMOUS,
		Method:  http.MethodPost,
		Url:     "/auth/refresh-token",
		Body:    Body{},
		Summary: "Get the new access token",
		Tag:     "IAM - Authentication",
		Examples: []helper.ExampleResponse{
			{
				StatusCode: 200,
				Content: map[string]interface{}{
					"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb250ZW50IjoiZXlKemRXSnFaV04wSWpvaVFVTkRSVk5UWDFSUFMwVk9JaXdpZFhObGNsOXBaQ0k2SW1JME9HVXpZalpsTFdSbFkyUXRORFpqTWkxaE0yVTJMVEJtTjJKa1ltTmtNVEppTlNJc0luVnpaWEpmWVdOalpYTnpJam9pTXlKOSIsImV4cCI6MTcyNzg4MjY5OX0.fpx7n6dwXmCgjG7M5i3auNlj81O2s7o-tygQcdEjL04",
				},
			},
		},
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

		req := usecase.RefreshTokenReq{
			RefreshToken:        body.RefreshToken,
			AccessTokenDuration: time.Duration(accessTokenInSecond) * time.Second,
			Now:                 time.Now(),
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
