package controller

import (
	"dashboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) GetListMainWaterChannelDoor(u usecase.ListMainWaterChannelDoorUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DASHBOARD_TABEL_DATA_KONDISI_PEMENUHAN_AIR_IRIGASI_READ,
		Method:  http.MethodGet,
		Url:     "/dashboard/main-water-channel-doors",
		Summary: "Get list main water channel door",
		Tag:     "Dashboard - Main Page",
		QueryParams: []helper.QueryParam{
			{Name: "sort_by", Type: "string", Description: "sort key", Required: false},
			{Name: "sort_order", Type: "string", Description: "sort order", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		sortBy := controller.GetQueryString(r, "sort_by", "")
		sortOrder := controller.GetQueryString(r, "sort_order", "")

		req := usecase.ListMainWaterChannelDoorReq{
			SortBy:    sortBy,
			SortOrder: sortOrder,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData

}
