package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func (c Controller) GetDeviceStatusHandler(u usecase.GetDeviceStatusUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DASHBOARD_DATA_PERANGKAT_READ,
		Method:  http.MethodGet,
		Url:     "/dashboard/device-status",
		Summary: "Get device status",
		Tag:     "Dashboard - Main Page",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetDeviceStatusReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData

}
