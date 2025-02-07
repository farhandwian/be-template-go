package controller

import (
	"iam/controller"
	"iam/model"
	"monitoring/usecase"
	"net/http"
	"shared/helper"
)

func ActivityMonitorCreateHandler(mux *http.ServeMux, u usecase.ActivityMonitorCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/monitor/activity",
		Access:  model.DEFAULT_OPERATION,
		Body:    usecase.ActivityMonitorCreateReq{},
		Summary: "Create activity monitor",
		Tag:     "Activity Monitoring",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.ActivityMonitorCreateReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
