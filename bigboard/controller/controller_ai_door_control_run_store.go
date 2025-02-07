package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"time"
)

func AiDoorControllRunStore(mux *http.ServeMux, u usecase.AiDoorControlRunStoreUseCase) helper.APIData {
	// ai payload body to /bigboard/ai/door-control-run-store
	// {
	// 	"water_channel_door_id" : 1,
	// 	"controller_index": 0,
	// 	"up_by": 10, //naikan sebanyak 10
	// 	"down_by": 10, //turunkan sebanyak 10
	// 	"set_to": 10 //ubah ke 10
	// }

	type Body struct {
		WaterChannelDoorID *int     `json:"water_channel_door_id"`
		ControllerIndex    *int     `json:"controller_index"`
		UpBy               *float32 `json:"up_by"`
		DownBy             *float32 `json:"down_by"`
		SetTo              *float32 `json:"set_to"`
	}

	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/door-control-run-store",
		Body:    Body{},
		Access:  iammodel.PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_PINTU_AIR_UPDATE,
		Summary: "Ask AI to control door run store",
		Tag:     "Bigboard AI",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		reqBody, ok := controller.ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.AiDoorControlRunStoreReq{
			WaterChannelDoorID: reqBody.WaterChannelDoorID,
			ControllerIndex:    reqBody.ControllerIndex,
			UpBy:               reqBody.UpBy,
			DownBy:             reqBody.DownBy,
			SetTo:              reqBody.SetTo,
			Now:                time.Now().In(time.Local),
			Access:             string(apiData.Access),
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
