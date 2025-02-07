package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
	"time"
)

func (c Controller) RegisterUserHandler(u usecase.RegisterUser) helper.APIData {

	type Body struct {
		Name        string            `json:"name"`
		Email       model.Email       `json:"email"`
		PhoneNumber model.PhoneNumber `json:"phone_number"`
	}

	apiData := helper.APIData{
		Access:  model.MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_CREATE,
		Method:  http.MethodPost,
		Url:     "/account/register",
		Body:    Body{},
		Summary: "Register user",
		Tag:     "IAM - Account Management",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.RegisterUserReq{
			Now:         time.Now(),
			Name:        body.Name,
			Email:       body.Email,
			PhoneNumber: body.PhoneNumber,
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := Authorization(handler, apiData.Access)
	authenticatedHandler := Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
