package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) AccessResetHandler(u usecase.AccessReset) helper.APIData {

	type Body struct {
		Accesses []model.MapAccessID `json:"accesses"`
	}

	apiData := helper.APIData{
		Access:  model.MANAJEMEN_PENGGUNA_HAK_AKSES_UPDATE,
		Method:  http.MethodPost,
		Url:     "/users/{id}/access",
		Body:    Body{},
		Summary: "Set access to user",
		Tag:     "IAM - User Access",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		userID := r.PathValue("id")

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.AccessResetReq{
			UserID:   model.UserID(userID),
			Accesses: body.Accesses,
		}

		HandleUsecase(r.Context(), w, u, req)

	}

	authorizationHandler := Authorization(handler, apiData.Access)
	authenticatedHandler := Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
