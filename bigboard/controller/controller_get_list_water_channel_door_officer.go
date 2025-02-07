package controller

import (
	"bigboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
)

func GetListWaterChannelDoorOfficer(mux *http.ServeMux, u usecase.ListWaterChannelDoorOfficerUseCase) helper.APIData {
	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/bigboard/water-channel-doors/officer",
		Summary: "Get list water channel door officer",
		Tag:     "Bigboard",
	}
	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.ListWaterChannelDoorOfficerReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
