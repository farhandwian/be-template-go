package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"os"
	"shared/core"
	"shared/helper"
	"strconv"
	"time"
)

func (c Controller) PinChangeRequestHandler(u usecase.PinChangeRequest) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodPost,
		Url:     "/pin/change/initiate",
		Summary: "Initiate change pin",
		Tag:     "IAM - Pin Management",
	}

	pinChangePageUrl := os.Getenv("PIN_CHANGE_PAGE_URL")

	otpExpirationInSecond, err := strconv.Atoi(os.Getenv("OTP_EXPIRATION_IN_SECOND"))
	if err != nil {
		panic(err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		userID := core.GetDataFromContext[model.UserID](r.Context(), UserIDContext)

		req := usecase.PinChangeRequestReq{
			UserID:            userID,
			PinChangePageUrl:  pinChangePageUrl,
			PinChangeDuration: time.Duration(otpExpirationInSecond) * time.Second,
			Now:               time.Now(),
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authenticatedHandler := Authentication(handler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
