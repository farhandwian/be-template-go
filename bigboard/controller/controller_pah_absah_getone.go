package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func GetPahAbsahDetailHandler(mux *http.ServeMux, u usecase.GetPahAbsahDetailUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/bigboard/pah-absah/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get detail pah absah by ID",
		Tag:     "Bigboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		req := usecase.GetPahAbsahDetailReq{ID: id}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
