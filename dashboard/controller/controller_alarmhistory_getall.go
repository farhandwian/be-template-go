package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
	"time"
)

func (c Controller) AlarmHistoryGetAllHandler(u usecase.AlarmHistoryGetAll) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/alarmhistory",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get all alarm history",
		Tag:     "Master Data - Alarm",
		QueryParams: []helper.QueryParam{
			{Name: "water_channel_door_id", Type: "number", Description: "water_channel_door_id", Required: false},
			{Name: "device_id", Type: "number", Description: "device_id", Required: false},
			{Name: "page", Type: "number", Description: "page", Required: false},
			{Name: "size", Type: "number", Description: "size", Required: false},
			{Name: "min", Type: "string", Description: "format 2024-11-25 02:00:00", Required: false},
			{Name: "max", Type: "string", Description: "format 2024-11-25 02:10:10", Required: false},
			{Name: "metric", Type: "string", Description: "debit, tma or gate_level", Required: false},
			{Name: "priority", Type: "string", Description: "critical or warning", Required: false},
			{Name: "sort_order", Type: "string", Description: "desc or asc", Required: false},
			{Name: "sort_by", Type: "string", Description: "channel_name, door_name, priority, metric, created_at", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		page := controller.GetQueryInt(r, "page", 1)
		size := controller.GetQueryInt(r, "size", 10)
		wcdid := controller.GetQueryInt(r, "water_channel_door_id", 0)
		did := controller.GetQueryInt(r, "device_id", 0)

		min := controller.GetQueryString(r, "min", "0")
		max := controller.GetQueryString(r, "max", "0")

		priority := controller.GetQueryString(r, "priority", "")
		metric := controller.GetQueryString(r, "metric", "")

		sortBy := controller.GetQueryString(r, "sort_by", "")
		sortOrder := controller.GetQueryString(r, "sort_order", "")

		minTime, err := time.Parse("2006-01-02 15:04:05", min)
		if err != nil {
			minTime = time.Time{}
		}

		maxTime, err := time.Parse("2006-01-02 15:04:05", max)
		if err != nil {
			maxTime = time.Time{}
		}

		req := usecase.AlarmHistoryGetAllReq{
			Page:               page,
			Size:               size,
			WaterChannelDoorID: wcdid,
			DeviceID:           did,
			MinTime:            minTime,
			MaxTime:            maxTime,
			Priority:           priority,
			Metric:             metric,
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
