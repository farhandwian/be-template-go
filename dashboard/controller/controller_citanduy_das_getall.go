package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) GetCitanduyDas(u usecase.GetCitanduyDASUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/sihka/map/citanduy-das",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get citanduy das",
		Tag:     "Sihka",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetCitanduyDASReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
