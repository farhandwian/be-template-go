package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) GetCitanduyArea(u usecase.GetCitanduyAreaUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/sihka/map/citanduy-area",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get all raw water",
		Tag:     "Sihka",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetCitanduyAreaReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
