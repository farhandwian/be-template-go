package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func PlayAudioHandler(mux *http.ServeMux, u usecase.AiPlayAudioUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/play-audio",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Ask AI to play Audio",
		Tag:     "Bigboard AI",
		Body:    usecase.AiPlayAudioReq{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		request, ok := controller.ParseJSON[usecase.AiPlayAudioReq](w, r)
		if !ok {
			return
		}
		req := usecase.AiPlayAudioReq{FileName: request.FileName}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
