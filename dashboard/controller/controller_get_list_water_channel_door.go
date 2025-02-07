package controller

import (
	"dashboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) GetListWaterChannelDoor(u usecase.ListWaterChannelDoorUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.PINTU_AIR_TABEL_DAFTAR_PINTU_READ,
		Method:  http.MethodGet,
		Url:     "/dashboard/water-channel-doors/",
		Summary: "Get list water channel door",
		Tag:     "Dashboard",
		QueryParams: []helper.QueryParam{
			{Name: "name", Type: "string", Description: "Water channel door name", Required: false},
			{Name: "min_water_elevation", Type: "number", Description: "Minimum water elevation", Required: false},
			{Name: "max_water_elevation", Type: "number", Description: "Maximum water elevation", Required: false},
			{Name: "min_actual_debit", Type: "number", Description: "Minimum actual debit", Required: false},
			{Name: "max_actual_debit", Type: "number", Description: "Maximum actual debit", Required: false},
			{Name: "min_required_debit", Type: "number", Description: "Minimum required debit", Required: false},
			{Name: "max_required_debit", Type: "number", Description: "Maximum required debit", Required: false},
			{Name: "water_channel_id", Type: "number", Description: "Water channel ID", Required: false},
			{Name: "status", Type: "string", Description: "Status", Required: false},
			{Name: "page", Type: "number", Description: "Page", Required: true},
			{Name: "page_size", Type: "number", Description: "page size", Required: true},
			{Name: "sort_by", Type: "string", Description: "sort key", Required: false},
			{Name: "sort_order", Type: "string", Description: "sort order", Required: false},
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
		status := controller.GetQueryString(r, "status", "")
		page := controller.GetQueryInt(r, "page", 1)
		pageSize := controller.GetQueryInt(r, "page_size", 10)
		sortBy := controller.GetQueryString(r, "sort_by", "")
		sortOrder := controller.GetQueryString(r, "sort_order", "")

		req := usecase.ListWaterChannelReq{
			Name:              name,
			MinWaterElevation: minWaterElevation,
			MaxWaterElevation: maxWaterElevation,
			MinActualDebit:    minActualDebit,
			MaxActualDebit:    maxActualDebit,
			MinRequiredDebit:  minRequiredDebit,
			MaxRequiredDebit:  maxRequiredDebit,
			WaterChannelID:    waterChannelID,
			Status:            status,
			Page:              page,
			PageSize:          pageSize,
			SortBy:            sortBy,
			SortOrder:         sortOrder,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData

}
