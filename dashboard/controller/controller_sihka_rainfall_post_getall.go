package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) RainfallPostGetAllHandler(u usecase.ListRainfallPostUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/sihka/hydrology/rainfall-post",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get rain fall posts",
		Tag:     "Sihka",
		QueryParams: []helper.QueryParam{
			{Name: "date", Type: "string", Description: "Filter date", Required: true},
			{Name: "city", Type: "string", Description: "City", Required: false},
			{Name: "telemetry", Type: "number", Description: "telemetry", Required: false},
			{Name: "vendor", Type: "string", Description: "vendor", Required: false},
			{Name: "keyword", Type: "string", Description: "keyword", Required: false},
			{Name: "sort_by", Type: "string", Description: "sort key", Required: false},
			{Name: "sort_order", Type: "string", Description: "sort order", Required: false},
			{Name: "page", Type: "integer", Description: "page", Required: true},
			{Name: "page_size", Type: "integer", Description: "page size", Required: true},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		page := controller.GetQueryInt(r, "page", 1)
		pageSize := controller.GetQueryInt(r, "page_size", 10)
		date := controller.GetQueryString(r, "date", "")
		city := controller.GetQueryString(r, "city", "")
		keyword := controller.GetQueryString(r, "keyword", "")
		sortBy := controller.GetQueryString(r, "sort_by", "")
		sortOrder := controller.GetQueryString(r, "sort_order", "")
		telemetry := controller.GetQueryFloat(r, "telemetry", 0)
		vendor := controller.GetQueryString(r, "vendor", "")

		req := usecase.ListRainfallPostReq{
			Date:      date,
			City:      city,
			Telemetry: telemetry,
			Vendor:    vendor,
			Page:      page,
			PageSize:  pageSize,
			Keyword:   keyword,
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
