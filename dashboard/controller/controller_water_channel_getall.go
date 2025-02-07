package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) GetListWaterChannel(u usecase.ListWaterChannelUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/water-channels",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get all water channels",
		Tag:     "Dashboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetAllWaterChannelUseCaseReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
