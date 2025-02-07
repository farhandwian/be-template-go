package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"strconv"
)

func GetCertainCCTVHandler(mux *http.ServeMux, u usecase.AIGetCertainCCTVFromWaterChannelUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/cctv/water-channel-door/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Ask AI to get certain cctv from water channel door",
		Tag:     "Bigboard AI",
		Body:    usecase.AIGetCertainCCTVFromWaterChannelReq{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		request, ok := controller.ParseJSON[usecase.AIGetCertainCCTVFromWaterChannelReq](w, r)
		if !ok {
			return
		}
		idStr := r.PathValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			controller.Fail(w, err)
			return
		}

		req := usecase.AIGetCertainCCTVFromWaterChannelReq{
			WaterChannelDoorID: id,
			CCTVIndex:          request.CCTVIndex,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
