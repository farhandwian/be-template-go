package controller

import (
	"bigboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
)

func GetListWaterChannelDoorDebit(mux *http.ServeMux, u usecase.ListWaterChannelDoorWithActualDebitUseCase) helper.APIData {
	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/bigboard/water-channel-doors/actual-debit",
		Summary: "Get list water channel door actual debit",
		Tag:     "Bigboard",
	}
	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.ListWaterChannelDoorWithActualDebitReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
