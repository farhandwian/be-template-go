package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func OpenLayerHandler(mux *http.ServeMux, u usecase.AiOpenLayerUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/open-layer",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Ask AI to open layer",
		Tag:     "Bigboard AI",
		Body:    usecase.AiOpenLayerReq{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		request, ok := controller.ParseJSON[usecase.AiOpenLayerReq](w, r)
		if !ok {
			return
		}
		req := usecase.AiOpenLayerReq{LayersName: request.LayersName}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
