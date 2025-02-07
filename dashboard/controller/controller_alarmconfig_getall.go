package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/model"
)

func (c Controller) AlarmConfigGetAllHandler(u usecase.AlarmConfigGetAll) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/alarmconfig",
		Access:  iammodel.SISTEM_PERINGATAN_KONFIGURASI_READ,
		Summary: "Get all alarm config",
		Tag:     "Master Data - Alarm",
		QueryParams: []helper.QueryParam{
			{Name: "page", Type: "number", Description: "page", Required: false},
			{Name: "size", Type: "number", Description: "size", Required: false},
			{Name: "water_channel_door_id", Type: "string", Description: "water_channel_door_id", Required: false},
			{Name: "priority", Type: "string", Description: "priority critical or warning", Required: false},
			{Name: "metric", Type: "string", Description: "metric tma, gate_level or debit", Required: false},
			{Name: "sort_order", Type: "string", Description: "desc or asc", Required: false},
			{Name: "sort_by", Type: "string", Description: "channel_name, door_name, priority, metric, created_at", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		page := controller.GetQueryInt(r, "page", 1)
		size := controller.GetQueryInt(r, "size", 10)

		wcdi := controller.GetQueryInt(r, "water_channel_door_id", 0)
		priority := controller.GetQueryString(r, "priority", "")
		metric := controller.GetQueryString(r, "metric", "")

		sortBy := controller.GetQueryString(r, "sort_by", "")
		sortOrder := controller.GetQueryString(r, "sort_order", "")

		req := usecase.AlarmConfigGetAllReq{
			Page:               page,
			Size:               size,
			WaterChannelDoorID: wcdi,
			Priority:           model.AlarmConfigPriority(priority),
			Metric:             model.AlarmMetric(metric),
			SortOrder:          sortOrder,
			SortBy:             sortBy,
		}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
