package controller

import (
	"iam/controller"
	"iam/model"
	"monitoring/usecase"
	"net/http"
	"shared/helper"
	usecase2 "shared/usecase"
)

func ActivityMonitorGetAllHandler(mux *http.ServeMux, u usecase2.ActivityMonitorGetAllUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/monitor/activity",
		Access:  model.DEFAULT_OPERATION,
		Body:    usecase.ActivityMonitorCreateReq{},
		Summary: "Get activity monitor",
		Tag:     "Activity Monitoring",
		QueryParams: []helper.QueryParam{
			{Name: "page", Type: "number", Description: "page", Required: false},
			{Name: "page_size", Type: "number", Description: "size", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		page := controller.GetQueryInt(r, "page", 1)
		size := controller.GetQueryInt(r, "page_size", 15)

		req := usecase2.ActivityMonitorGetAllUseCaseReq{
			Page: page,
			Size: size,
		}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
