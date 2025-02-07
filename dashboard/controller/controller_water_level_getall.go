package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func (c Controller) GetListWaterLevel(u usecase.GetListWaterLevelUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/sihka/map/water-levels",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get all water levels",
		Tag:     "Sihka",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetListWaterLevelReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
