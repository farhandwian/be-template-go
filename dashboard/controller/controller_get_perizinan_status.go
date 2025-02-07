package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func (c Controller) GetPerizinanStatusHandler(u usecase.GetPerizinanStatusUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DASHBOARD_DATA_SI_JAGACAI_READ,
		Method:  http.MethodGet,
		Url:     "/dashboard/perizinan-status",
		Summary: "Get perizinan status",
		Tag:     "Dashboard - Main Page",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetPerizinanStatusReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData

}
