package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func (c Controller) GetBMKGLocationHandler(u usecase.GetBMKGLocationUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/weather/location",
		Access:  model.DEFAULT_OPERATION,
		Summary: "Get all pintu",
		Tag:     "Dashboard",
		QueryParams: []helper.QueryParam{
			{Name: "name", Type: "string", Description: "location name", Required: true},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		location := controller.GetQueryString(r, "name", "")
		req := usecase.GetBMKGLocationReq{Location: location}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
