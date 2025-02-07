package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func GetPhotoPathHandler(mux *http.ServeMux, u usecase.GetPhotoPathUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/photo-path",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Ask AI to get photo path",
		Tag:     "Bigboard AI",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		req := usecase.AiGetPhotoPathReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
