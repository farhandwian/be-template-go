package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func (c Controller) GetListIntake(u usecase.GetListIntakeUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/intakes",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get all Intake",
		Tag:     "Infrastruktur",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetListIntakeReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
