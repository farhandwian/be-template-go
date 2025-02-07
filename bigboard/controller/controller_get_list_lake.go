package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func GetListLake(mux *http.ServeMux, u usecase.GetListLakeUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/bigboard/lakes",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get all lakes",
		Tag:     "Bigboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetListLakeReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
