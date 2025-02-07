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

func (c Controller) PasswordChangeRequestHandler(u usecase.PasswordChangeRequest) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodPost,
		Url:     "/password/change/initiate",
		Summary: "Initiate change password",
		Tag:     "IAM - Password Management",
	}

	passwordChangePageUrl := os.Getenv("PASSWORD_CHANGE_PAGE_URL")

	otpExpirationInSecond, err := strconv.Atoi(os.Getenv("OTP_EXPIRATION_IN_SECOND"))
	if err != nil {
		panic(err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		userID := core.GetDataFromContext[model.UserID](r.Context(), UserIDContext)

		req := usecase.PasswordChangeRequestReq{
			UserID:                 userID,
			PasswordChangePageUrl:  passwordChangePageUrl,
			PasswordChangeDuration: time.Duration(otpExpirationInSecond) * time.Second,
			Now:                    time.Now(),
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authenticatedHandler := Authentication(handler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
