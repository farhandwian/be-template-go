package controller

import (
	"bigboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
)

func WebhookSpeedTestHandler(mux *http.ServeMux, u usecase.WebhookSpeedTestUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/monitoring/webhook/speedtest",
		Access:  model.DEFAULT_OPERATION,
		Body:    usecase.WebhookSpeedTestReq{},
		Summary: "Webhook for speedtest",
		Tag:     "Webhook",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.WebhookSpeedTestReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
