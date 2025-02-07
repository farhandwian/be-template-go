package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"time"
)

func (c Controller) WaterQualityPostGetAllHandler(u usecase.ListWaterQualityPostUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/sihka/water-quality-posts",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get water quality posts",
		Tag:     "Sihka",
		QueryParams: []helper.QueryParam{
			{Name: "date", Type: "string", Description: "Filter date", Required: true},
			{Name: "river", Type: "string", Description: "river", Required: false},
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
		river := controller.GetQueryString(r, "river", "")
		date := controller.GetQueryString(r, "date", time.Now().Format(time.DateOnly))
		keyword := controller.GetQueryString(r, "keyword", "")
		sortBy := controller.GetQueryString(r, "sort_by", "")
		sortOrder := controller.GetQueryString(r, "sort_order", "")

		req := usecase.ListWaterQualityReq{
			Date:      date,
			River:     river,
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
