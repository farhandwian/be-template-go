package controller

import (
	"bigboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
)

func GetListWaterChannelDoorCCTV(mux *http.ServeMux, u usecase.ListWaterChannelDoorCCTVUseCase) helper.APIData {
	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/bigboard/water-channel-doors/cctv",
		Summary: "Get list water channel door cctv",
		Tag:     "Bigboard",
	}
	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.ListWaterChannelDoorCCTVReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
