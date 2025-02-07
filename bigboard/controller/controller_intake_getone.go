package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func GetIntakeDetailHandler(mux *http.ServeMux, u usecase.GetIntakeDetailUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/bigboard/intakes/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get detail intake by ID",
		Tag:     "Bigboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		req := usecase.GetIntakeDetailReq{ID: id}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
