package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func GetListClimatologyMap(mux *http.ServeMux, u usecase.GetClimatologyPostUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/bigboard/climatology",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get all climatologies",
		Tag:     "Bigboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetListClimatologyPostReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
