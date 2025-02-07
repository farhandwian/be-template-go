package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/model"
)

func (c Controller) AlarmConfigUpdateHandler(u usecase.AlarmConfigUpdate) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPut,
		Url:     "/dashboard/alarmconfig/{id}",
		Access:  iammodel.SISTEM_PERINGATAN_KONFIGURASI_UPDATE,
		Body:    usecase.AlarmConfigCreateReq{},
		Summary: "Update a alarm config",
		Tag:     "Master Data - Alarm",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.AlarmConfigUpdateReq](w, r)
		if !ok {
			return
		}
		request.ID = model.AlarmConfigID(id)
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
