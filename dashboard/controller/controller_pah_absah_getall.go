package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func (c Controller) GetListPahAbsah(u usecase.GetListPahAbsahUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/pah-absah",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get all Pah Absah",
		Tag:     "Infrastruktur",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetListPahAbsahReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
