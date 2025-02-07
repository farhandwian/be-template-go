package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func (c Controller) GetWaterAvailabilityForecastHandler(u usecase.WaterAvailabilityForecastUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/dashboard/water-availability",
		Summary: "Get water availability forecast",
		Tag:     "Dashboard - Main Page",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.WaterAvailabilityForecastReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData

}
