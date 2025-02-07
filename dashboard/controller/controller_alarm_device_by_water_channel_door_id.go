package controller

import (
	"dashboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) DeviceByWaterChannelDoorIdHandler(u usecase.DeviceByWaterChannelDoorId) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/doors",
		Access:  model.DEFAULT_OPERATION,
		Summary: "Get all pintu",
		Tag:     "Dashboard",
		QueryParams: []helper.QueryParam{
			{Name: "water_channel_door_id", Type: "number", Description: "water channer door id", Required: true},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := controller.GetQueryInt(r, "water_channel_door_id", 0)
		req := usecase.DeviceByWaterChannelDoorIdReq{WaterChannelDoorId: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
