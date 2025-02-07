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

func (c Controller) PasswordResetRequestHandler(u usecase.PasswordResetRequest) helper.APIData {

	type Body struct {
		UserID model.UserID `json:"user_id"`
	}

	apiData := helper.APIData{
		Access:  model.MANAJEMEN_PENGGUNA_RESET_KATA_SANDI_UPDATE,
		Method:  http.MethodPost,
		Url:     "/password/reset/initiate",
		Body:    Body{},
		Summary: "Initiate password reset",
		Tag:     "IAM - Password Management",
	}

	passwordResetPageUrl := os.Getenv("PASSWORD_RESET_PAGE_URL")

	emailExpirationInSecond, err := strconv.Atoi(os.Getenv("EMAIL_EXPIRATION_IN_SECOND"))
	if err != nil {
		panic(err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.PasswordResetRequestReq{
			UserID:                body.UserID,
			PasswordResetPageUrl:  passwordResetPageUrl,
			PasswordResetDuration: time.Duration(emailExpirationInSecond) * time.Second,
			Now:                   time.Now(),
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := Authorization(handler, apiData.Access)
	authenticatedHandler := Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
