package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func GetRainfallDetailHandler(mux *http.ServeMux, u usecase.GetRainFallDetailUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/bigboard/rainfalls/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get detail rainfall",
		Tag:     "Bigboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		req := usecase.GetRainFallDetailReq{ID: id}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
