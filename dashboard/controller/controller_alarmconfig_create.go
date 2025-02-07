package controller

import (
	"dashboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) AlarmConfigCreateHandler(u usecase.AlarmConfigCreate) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/dashboard/alarmconfig",
		Access:  model.SISTEM_PERINGATAN_KONFIGURASI_CREATE,
		Body:    usecase.AlarmConfigCreateReq{},
		Summary: "Create alarm config",
		Tag:     "Master Data - Alarm",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.AlarmConfigCreateReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
