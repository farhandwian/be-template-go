package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func GetServiceStatusHandler(mux *http.ServeMux, u usecase.GetServiceStatusUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/bigboard/service-status",
		Summary: "Get service status",
		Tag:     "Bigboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetServiceStatusReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData

}
