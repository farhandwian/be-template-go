package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) UserUpdateHandler(u usecase.UserUpdate) helper.APIData {

	type Body struct {
		Name        string
		PhoneNumber model.PhoneNumber
		Email       model.Email
	}

	apiData := helper.APIData{
		Access:  model.MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_UPDATE,
		Method:  http.MethodPut,
		Url:     "/users/{id}",
		Body:    Body{},
		Summary: "Update to user detail",
		Tag:     "IAM - User Management",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		userID := r.PathValue("id")

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.UserUpdateReq{
			UserID:      model.UserID(userID),
			Name:        body.Name,
			PhoneNumber: body.PhoneNumber,
			Email:       body.Email,
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := Authorization(handler, apiData.Access)
	authenticatedHandler := Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
