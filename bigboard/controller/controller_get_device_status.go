package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func GetDeviceStatusHandler(mux *http.ServeMux, u usecase.GetDeviceStatusUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/bigboard/device-status",
		Summary: "Get device status",
		Tag:     "Bigboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetDeviceStatusReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData

}
