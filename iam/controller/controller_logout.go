package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/core"
	"shared/helper"
)

func (c Controller) LogoutHandler(u usecase.Logout) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodPost,
		Url:     "/auth/logout",
		Summary: "Logout session",
		Tag:     "IAM - Authentication",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		userID := core.GetDataFromContext[model.UserID](r.Context(), UserIDContext)

		req := usecase.LogoutReq{
			UserID: userID,
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authenticatedHandler := Authentication(handler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
