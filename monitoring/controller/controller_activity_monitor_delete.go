package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"monitoring/usecase"
	"net/http"
	"shared/helper"
)

func ActivityMonitorDeleteHandler(mux *http.ServeMux, u usecase.ActivityMonitorDelete) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodDelete,
		Url:     "/monitor/activity/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Delete activity monitor",
		Tag:     "Activity Monitoring",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		req := usecase.ActivityMonitorDeleteReq{ID: id}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
