package controller

import (
	"dashboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/core"
	"shared/helper"
	"time"
)

func (c Controller) DoorControlRunSecurityRelayHandler(u usecase.DoorControlRunSecurityRelay) helper.APIData {

	type Body struct {
		WaterChannelDoorID int    `json:"water_channel_door_id"`
		DeviceID           int    `json:"device_id"`
		Pin                string `json:"pin"`
		Status             int    `json:"status"`
	}

	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/dashboard/doorcontrol-run-security-relay",
		Access:  model.PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_SECURITY_RELAY_UPDATE,
		Body:    Body{},
		Summary: "run door control security relay",
		Tag:     "Master Data - Door Control",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		reqBody, ok := controller.ParseJSON[Body](w, r)
		if !ok {
			return
		}

		request := usecase.DoorControlRunSecurityRelayReq{
			Pin:                model.PIN(reqBody.Pin),
			WaterChannelDoorID: reqBody.WaterChannelDoorID,
			DeviceID:           reqBody.DeviceID,
			OfficerId:          core.GetDataFromContext[model.UserID](r.Context(), controller.UserIDContext),
			Now:                time.Now().In(time.Local),
			Status:             reqBody.Status,
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
