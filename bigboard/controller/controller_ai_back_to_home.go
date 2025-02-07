package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func BackToHomeHandler(mux *http.ServeMux, u usecase.AiBackToHomeUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/back-to-home",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Ask AI to back to Home",
		Tag:     "Bigboard AI",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		req := usecase.AiBackToHomeReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
