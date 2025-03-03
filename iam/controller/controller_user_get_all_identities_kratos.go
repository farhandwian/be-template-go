package controller

import (
	"iam/gateway"
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) UserGetAllIdentitiesKratosHandler(u usecase.UserGetAllIdentitiesKratos) helper.APIData {

	apiData := helper.APIData{
		Access:  model.MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_READ,
		Method:  http.MethodGet,
		Url:     "/api/kratos/users",
		Summary: "Get all identity kratos users",
		Tag:     "IAM - User Management",
		QueryParams: []helper.QueryParam{
			{Name: "page", Type: "integer", Description: "Page number", Required: false},
			{Name: "size", Type: "integer", Description: "Number of items per page", Required: false},
			{Name: "keyword", Type: "string", Description: "filter", Required: false},
			{Name: "sort_order", Type: "string", Description: "desc or asc", Required: false},
			{Name: "sort_by", Type: "string", Description: "name, phone_number, email", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		page := GetQueryInt(r, "page", 1)
		size := GetQueryInt(r, "size", 10)
		sortBy := GetQueryString(r, "sort_by", "")
		sortOrder := GetQueryString(r, "sort_order", "")
		keywords := GetQueryString(r, "keyword", "")

		req := usecase.UserGetAllIdentitiesKratosReq{
			UserGetAllIdentitiesKratosReq: gateway.UserGetAllIdentitiesKratosReq{
				Page:      page,
				Size:      size,
				SortOrder: sortOrder,
				SortBy:    sortBy,
				Keyword:   keywords,
			},
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	// authorizationHandler := Authorization(handler, apiData.Access)
	// authenticatedHandler := Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
