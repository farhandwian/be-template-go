package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func (c Controller) DoorControlHistoryGetAllHandler(u usecase.DoorControlHistoryGetAll) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/doorcontrolhistory",
		Access:  iammodel.PINTU_AIR_DETAIL_PINTU_AIR_TABLE_DATA_RIWAYAT_PENGOPERASIAN_PINTU_AIR_READ,
		Summary: "Get all door control history",
		Tag:     "Master Data - Door Control",
		QueryParams: []helper.QueryParam{
			{Name: "water_channel_door_id", Type: "number", Description: "water_channel_door_id", Required: false},
			{Name: "device_id", Type: "number", Description: "device_id", Required: false},
			{Name: "page", Type: "number", Description: "page", Required: false},
			{Name: "size", Type: "number", Description: "size", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		page := controller.GetQueryInt(r, "page", 1)
		size := controller.GetQueryInt(r, "size", 10)
		wcdid := controller.GetQueryInt(r, "water_channel_door_id", 0)
		did := controller.GetQueryInt(r, "device_id", 0)
		req := usecase.DoorControlHistoryGetAllReq{
			Page:               page,
			Size:               size,
			WaterChannelDoorID: wcdid,
			DeviceID:           did,
		}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
