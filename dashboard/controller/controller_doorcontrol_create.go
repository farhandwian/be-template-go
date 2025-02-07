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

func (c Controller) DoorControlCreateHandler(u usecase.DoorControlCreate) helper.APIData {

	type Body struct {
		DateTime           string  `json:"date"`
		WaterChannelDoorID int     `json:"water_channel_door_id"`
		DeviceID           int     `json:"device_id"`
		OpenTarget         float32 `json:"open_target"`
		Reason             string  `json:"reason"`
	}

	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/dashboard/doorcontrol",
		Access:  model.DEFAULT_OPERATION,
		Body:    Body{},
		Summary: "Create door control",
		Tag:     "Master Data - Door Control",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		reqBody, ok := controller.ParseJSON[Body](w, r)
		if !ok {
			return
		}

		request := usecase.DoorControlCreateReq{
			DateTime:           reqBody.DateTime,
			WaterChannelDoorID: reqBody.WaterChannelDoorID,
			DeviceID:           reqBody.DeviceID,
			OpenTarget:         reqBody.OpenTarget,
			Reason:             reqBody.Reason,
			OfficerId:          core.GetDataFromContext[model.UserID](r.Context(), controller.UserIDContext),
			Now:                time.Now().In(time.Local),
			Tz:                 time.Local,
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
