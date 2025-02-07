package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/model"
)

func (c Controller) AlarmConfigGetOneHandler(u usecase.AlarmConfigGetOne) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/alarmconfig/{id}",
		Access:  iammodel.SISTEM_PERINGATAN_KONFIGURASI_READ,
		Summary: "Get a alarm config by id",
		Tag:     "Master Data - Alarm",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.AlarmConfigGetOneReq{ID: model.AlarmConfigID(id)}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
