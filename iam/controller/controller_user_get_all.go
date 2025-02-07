package controller

import (
	"iam/gateway"
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) UserGetAllHandler(u usecase.UserGetAll) helper.APIData {

	apiData := helper.APIData{
		Access:  model.MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_READ,
		Method:  http.MethodGet,
		Url:     "/users",
		Summary: "Get all users",
		Tag:     "IAM - User Management",
		QueryParams: []helper.QueryParam{
			{Name: "page", Type: "integer", Description: "Page number", Required: false},
			{Name: "size", Type: "integer", Description: "Number of items per page", Required: false},
			{Name: "user_id", Type: "string", Description: "Filter by user ID", Required: false},
			{Name: "email", Type: "string", Description: "Filter by email", Required: false},
			{Name: "phone_number", Type: "string", Description: "Filter by phone number", Required: false},
			{Name: "name_like", Type: "string", Description: "Filter by name (partial match)", Required: false},
			{Name: "sort_order", Type: "string", Description: "desc or asc", Required: false},
			{Name: "sort_by", Type: "string", Description: "name, phone_number, email", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		page := GetQueryInt(r, "page", 1)
		size := GetQueryInt(r, "size", 10)
		userID := GetQueryString(r, "user_id", "")
		email := GetQueryString(r, "email", "")
		phoneNumber := GetQueryString(r, "phone_number", "")
		nameLike := GetQueryString(r, "name_like", "")

		sortBy := GetQueryString(r, "sort_by", "")
		sortOrder := GetQueryString(r, "sort_order", "")

		req := usecase.UserGetAllReq{
			UserGetAllReq: gateway.UserGetAllReq{
				Page:        page,
				Size:        size,
				UserID:      model.UserID(userID),
				Email:       model.Email(email),
				PhoneNumber: model.PhoneNumber(phoneNumber),
				NameLike:    nameLike,
				SortOrder:   sortOrder,
				SortBy:      sortBy,
			},
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := Authorization(handler, apiData.Access)
	authenticatedHandler := Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
