package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func (c Controller) ActivityMonitorGetAllHandler(u usecase.ActivityMonitorGetAllUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/monitor/activity",
		Access:  model.DASHBOARD_MONITORING_AKTIVITAS_READ,
		Summary: "Get activity monitor",
		Tag:     "Dashboard - Main Page",
		QueryParams: []helper.QueryParam{
			{Name: "page", Type: "number", Description: "page", Required: false},
			{Name: "page_size", Type: "number", Description: "size", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		page := controller.GetQueryInt(r, "page", 1)
		size := controller.GetQueryInt(r, "page_size", 10)

		req := usecase.ActivityMonitorGetAllUseCaseReq{
			Page: page,
			Size: size,
		}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
