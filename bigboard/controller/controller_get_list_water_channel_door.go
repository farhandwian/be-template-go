package controller

import (
	"bigboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
)

func GetListWaterChannelDoor(mux *http.ServeMux, u usecase.ListWaterChannelUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/bigboard/water-channel-doors/",
		Summary: "Get list water channel door",
		Tag:     "Bigboard",
		QueryParams: []helper.QueryParam{
			{Name: "name", Type: "string", Description: "Water channel door name", Required: false},
			{Name: "min_water_elevation", Type: "number", Description: "Minimum water elevation", Required: false},
			{Name: "max_water_elevation", Type: "number", Description: "Maximum water elevation", Required: false},
			{Name: "min_actual_debit", Type: "number", Description: "Minimum actual debit", Required: false},
			{Name: "max_actual_debit", Type: "number", Description: "Maximum actual debit", Required: false},
			{Name: "min_required_debit", Type: "number", Description: "Minimum required debit", Required: false},
			{Name: "max_required_debit", Type: "number", Description: "Maximum required debit", Required: false},
			{Name: "water_channel_id", Type: "number", Description: "Water channel ID", Required: false},
			{Name: "status_debit", Type: "string", Description: "Status Debit", Required: false},
			{Name: "status_tma", Type: "string", Description: "Status Sensor TMA", Required: false},
			{Name: "has_garbage", Type: "boolean", Description: "Has Garbage", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		name := controller.GetQueryString(r, "name", "")
		minWaterElevation := controller.GetQueryFloat(r, "min_water_elevation", 0)
		maxWaterElevation := controller.GetQueryFloat(r, "max_water_elevation", 0)
		minActualDebit := controller.GetQueryFloat(r, "min_actual_debit", 0)
		maxActualDebit := controller.GetQueryFloat(r, "max_actual_debit", 0)
		minRequiredDebit := controller.GetQueryFloat(r, "min_required_debit", 0)
		maxRequiredDebit := controller.GetQueryFloat(r, "max_required_debit", 0)
		waterChannelID := controller.GetQueryInt(r, "water_channel_id", 0)
		status_debit := controller.GetQueryString(r, "status_debit", "")
		status_tma := controller.GetQueryString(r, "status_tma", "")
		hasGarbage := controller.GetQueryBoolean(r, "has_garbage", false)

		req := usecase.ListWaterChannelReq{
			Name:              name,
			MinWaterElevation: minWaterElevation,
			MaxWaterElevation: maxWaterElevation,
			MinActualDebit:    minActualDebit,
			MaxActualDebit:    maxActualDebit,
			MinRequiredDebit:  minRequiredDebit,
			MaxRequiredDebit:  maxRequiredDebit,
			WaterChannelID:    waterChannelID,
			StatusDebit:       status_debit,
			StatusTMA:         status_tma,
			HasGarbage:        hasGarbage,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData

}
