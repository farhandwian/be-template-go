package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func GetCitanduyDas(mux *http.ServeMux, u usecase.GetCitanduyDASUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/bigboard/citanduy-das",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get citanduy das",
		Tag:     "Bigboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetCitanduyDASReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
