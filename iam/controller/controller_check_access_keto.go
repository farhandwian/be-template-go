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

func (c Controller) CheckAccessKetoHandler(u usecase.Login) helper.APIData {

	type Body struct {
		Email    model.Email `json:"email"`
		Password string      `json:"password"`
	}

	apiData := helper.APIData{
		Access:   model.ANONYMOUS,
		Method:   http.MethodPost,
		Url:      "/auth/login",
		Body:     Body{},
		Summary:  "Initiate user login",
		Tag:      "IAM - Authentication",
		Examples: []helper.ExampleResponse{},
	}

	otpExpirationInSecond, err := strconv.Atoi(os.Getenv("OTP_EXPIRATION_IN_SECOND"))
	if err != nil {
		panic(err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.LoginReq{
			Email:       body.Email,
			Password:    body.Password,
			OTPDuration: time.Duration(otpExpirationInSecond) * time.Second,
			Now:         time.Now(),
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
