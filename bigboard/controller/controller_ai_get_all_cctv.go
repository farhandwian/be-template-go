package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func GetAllCCTVHandler(mux *http.ServeMux, u usecase.AIGetAllCCTVFromWaterChannelUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/cctv",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Ask AI to get all cctv from water channel door",
		Tag:     "Bigboard AI",
		Body:    usecase.AIGetCCTVFromWaterChannelReq{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		request, ok := controller.ParseJSON[usecase.AIGetCCTVFromWaterChannelReq](w, r)
		if !ok {
			return
		}
		req := usecase.AIGetCCTVFromWaterChannelReq{WaterChannelDoorIDs: request.WaterChannelDoorIDs}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
