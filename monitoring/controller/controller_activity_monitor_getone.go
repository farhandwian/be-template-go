package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"monitoring/usecase"
	"net/http"
	"shared/helper"
)

func ActivityMonitorDetailHandler(mux *http.ServeMux, u usecase.ActivityMonitorGetOneUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/monitor/activity/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get one activity monitor",
		Tag:     "Activity Monitoring",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		req := usecase.ActivityMonitorGetOneReq{ID: id}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
