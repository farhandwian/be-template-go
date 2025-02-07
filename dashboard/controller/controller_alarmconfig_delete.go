// File: controller/controller_Asset.go

package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/model"
)

func (c Controller) AlarmConfigDeleteHandler(u usecase.AlarmConfigDelete) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodDelete,
		Url:     "/dashboard/alarmconfig/{id}",
		Access:  iammodel.SISTEM_PERINGATAN_KONFIGURASI_DELETE,
		Summary: "Delete a alarm config",
		Tag:     "Master Data - Alarm",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.AlarmConfigDeleteReq{ID: model.AlarmConfigID(id)}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
