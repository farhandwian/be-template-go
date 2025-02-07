package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func OpenVideoHandler(mux *http.ServeMux, u usecase.AiOpenVideoUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/open-video",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Ask AI to open video",
		Tag:     "Bigboard AI",
		Body:    usecase.AiOpenVideoReq{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		request, ok := controller.ParseJSON[usecase.AiOpenVideoReq](w, r)

		if !ok {
			return
		}

		req := usecase.AiOpenVideoReq{
			VideoPath: request.VideoPath,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
