package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func GetVideoPathHandler(mux *http.ServeMux, u usecase.GetVideoPathUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/video-path",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Ask AI to get video path",
		Tag:     "Bigboard AI",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		req := usecase.AiGetVideoPathReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
