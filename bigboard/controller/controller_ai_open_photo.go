package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func OpenPhotoHandler(mux *http.ServeMux, u usecase.AiOpenPhotoUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/open-photo",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Ask AI to open photo",
		Tag:     "Bigboard AI",
		Body:    usecase.AiOpenPhotoReq{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		request, ok := controller.ParseJSON[usecase.AiOpenPhotoReq](w, r)

		if !ok {
			return
		}

		req := usecase.AiOpenPhotoReq{
			PhotoPath: request.PhotoPath,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
