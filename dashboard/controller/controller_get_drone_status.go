package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func (c Controller) GetDroneStatusHandler(u usecase.GetDroneStatusUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DASHBOARD_DATA_AUTONOMOUS_DRONE_PATROLLING_SYSTEM_READ,
		Method:  http.MethodGet,
		Url:     "/dashboard/drone-status",
		Summary: "Get drone status",
		Tag:     "Dashboard - Main Page",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetDroneStatusReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData

}
