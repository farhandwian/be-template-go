package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) GetListClimatology(u usecase.GetListClimatologyUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/sihka/hydrology/climatology",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get all climatology",
		Tag:     "Sihka",
		QueryParams: []helper.QueryParam{
			{Name: "keyword", Type: "string", Description: "Keyword", Required: false},
			{Name: "sort_by", Type: "string", Description: "sort key", Required: false},
			{Name: "sort_order", Type: "string", Description: "sort order", Required: false},
			{Name: "page", Type: "integer", Description: "page", Required: true},
			{Name: "page_size", Type: "integer", Description: "page size", Required: true},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		keyword := controller.GetQueryString(r, "keyword", "")
		sortBy := controller.GetQueryString(r, "sort_by", "")
		sortOrder := controller.GetQueryString(r, "sort_order", "")
		page := controller.GetQueryInt(r, "page", 1)
		pageSize := controller.GetQueryInt(r, "page_size", 10)

		req := usecase.GetListClimatologyReq{
			Keyword:   keyword,
			Page:      page,
			PageSize:  pageSize,
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
