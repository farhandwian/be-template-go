package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func GetWaterRainPostDetailAIHandler(mux *http.ServeMux, u usecase.AiGetRainFallDetailUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/water-rain-post/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Ask AI to detail from water rain post",
		Tag:     "Bigboard AI",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		req := usecase.AiGetRainFallDetailReq{
			ID: id,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
