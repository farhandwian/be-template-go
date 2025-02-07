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

func (c Controller) EmailActivationRequestHandler(u usecase.EmailActivationRequest) helper.APIData {

	type Body struct {
		UserID model.UserID `json:"user_id"`
	}

	apiData := helper.APIData{
		Access:   model.MANAJEMEN_PENGGUNA_AKTIVASI_BUTTON_UPDATE,
		Method:   http.MethodPost,
		Url:      "/account/activate/initiate",
		Body:     Body{},
		Summary:  "Send email activation request",
		Tag:      "IAM - Account Management",
		Examples: []helper.ExampleResponse{},
	}

	emailActivationPageUrl := os.Getenv("EMAIL_ACTIVATION_PAGE_URL")

	emailExpirationInSecond, err := strconv.Atoi(os.Getenv("EMAIL_EXPIRATION_IN_SECOND"))
	if err != nil {
		panic(err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.EmailActivationRequestReq{
			UserID:                  body.UserID,
			Now:                     time.Now(),
			EmailActivationPageUrl:  emailActivationPageUrl,
			EmailActivationDuration: time.Duration(emailExpirationInSecond) * time.Second,
			ServerUrl:               os.Getenv("SERVER_URL"),
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := Authorization(handler, apiData.Access)
	authenticatedHandler := Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
