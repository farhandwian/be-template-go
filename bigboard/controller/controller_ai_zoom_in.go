package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"strconv"
)

func AiZoomInHandler(mux *http.ServeMux, u usecase.AiZoomInUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/zoom-in/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Ask AI to zoom in to water channel door",
		Tag:     "Bigboard AI",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			controller.Fail(w, err)
			return
		}

		req := usecase.AiZoomInReq{
			WaterChannelDoorID: id,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
