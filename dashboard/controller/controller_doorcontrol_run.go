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

func (c Controller) DoorControlRunHandler(u usecase.DoorControlRunDirectly) helper.APIData {

	type Body struct {
		WaterChannelDoorID int     `json:"water_channel_door_id"`
		DeviceID           int     `json:"device_id"`
		OpenTarget         float32 `json:"open_target"`
		Reason             string  `json:"reason"`
		Pin                string  `json:"pin"`
	}

	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/dashboard/doorcontrol-run",
		Access:  model.PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_PINTU_AIR_UPDATE,
		Body:    Body{},
		Summary: "run door control directly",
		Tag:     "Master Data - Door Control",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		reqBody, ok := controller.ParseJSON[Body](w, r)
		if !ok {
			return
		}

		request := usecase.DoorControlRunDirectlyReq{
			Pin:                model.PIN(reqBody.Pin),
			WaterChannelDoorID: reqBody.WaterChannelDoorID,
			DeviceID:           reqBody.DeviceID,
			OpenTarget:         reqBody.OpenTarget,
			Reason:             reqBody.Reason,
			OfficerId:          core.GetDataFromContext[model.UserID](r.Context(), controller.UserIDContext),
			Now:                time.Now().In(time.Local),
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
