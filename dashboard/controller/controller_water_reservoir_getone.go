package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func (c Controller) GetWaterReservoirDetailHandler(u usecase.GetWaterReservoirDetailUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/water-reservoirs/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get detail water reservoir by ID",
		Tag:     "Infrastruktur",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		req := usecase.GetWaterReservoirDetailReq{ID: id}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
