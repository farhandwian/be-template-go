package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func GetBMKGLocationHandler(mux *http.ServeMux, u usecase.GetBMKGLocationUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/bigboard/weather/location",
		Summary: "Get weather bigboard by location",
		Tag:     "Bigboard",
		QueryParams: []helper.QueryParam{
			{Name: "name", Type: "string", Description: "location name", Required: true},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		location := controller.GetQueryString(r, "name", "")
		req := usecase.GetBMKGLocationReq{Location: location}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData

}
