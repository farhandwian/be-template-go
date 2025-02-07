package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"strconv"
)

func GetWaterChannelDetailAIHandler(mux *http.ServeMux, u usecase.AiGetDetailWaterChannelUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/water-channel-door/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Ask AI to detail from water channel door",
		Tag:     "Bigboard AI",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			controller.Fail(w, err)
			return
		}

		req := usecase.AiGetDetailWaterChannelReq{
			WaterChannelID: id,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
