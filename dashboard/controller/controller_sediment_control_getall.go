package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func (c Controller) GetSedimentControl(u usecase.GetListSedimentControlUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/sediment-controls",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get all sediment control",
		Tag:     "Infrastruktur",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetListSedimentControlReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
