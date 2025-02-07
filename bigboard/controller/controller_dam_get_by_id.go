package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func GetDamDetailHandler(mux *http.ServeMux, u usecase.GetDamDetailUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/bigboard/dams/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get detail dam by ID",
		Tag:     "Bigboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		req := usecase.GetDamDetailReq{ID: id}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
