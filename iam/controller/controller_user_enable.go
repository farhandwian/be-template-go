package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) UserEnableHandler(u usecase.UserEnable) helper.APIData {

	type Body struct {
		Status bool `json:"status"`
	}

	apiData := helper.APIData{
		Access:  model.MANAJEMEN_PENGGUNA_HAK_AKSES_UPDATE,
		Method:  http.MethodPost,
		Url:     "/users/{id}/enable",
		Summary: "Enable or disable user",
		Tag:     "IAM - User Management",
		Body:    Body{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		userID := r.PathValue("id")

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.UserEnableReq{
			UserID: model.UserID(userID),
			Enable: body.Status,
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := Authorization(handler, apiData.Access)
	authenticatedHandler := Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
