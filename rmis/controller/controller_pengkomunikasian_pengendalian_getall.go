package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PengkomunikasianPengendalian Get All handler
func (c Controller) PengkomunikasianPengendalianGetAllHandler(u usecase.PengkomunikasianPengendalianGetAllUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/pengkomunikasian-pengendalian",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "pengkomunikasian-pengendalian",
			Relation:  "read",
		},
		Summary: "Get all Pengkomunikasian Pengendalian",
		Tag:     "Pengkomunikasian Pengendalian",
		QueryParams: []helper.QueryParam{
			{Name: "keyword", Type: "string", Description: "name, pic or location", Required: false},
			{Name: "page", Type: "number", Description: "page", Required: false},
			{Name: "size", Type: "number", Description: "size", Required: false},
			{Name: "sortBy", Type: "string", Description: "sort by", Required: false},
			{Name: "sortOrder", Type: "string", Description: "sort order", Required: false},
			{Name: "status", Type: "string", Description: "status filter", Required: false},
			{Name: "media", Type: "string", Description: "media filter", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		page := controller.GetQueryInt(r, "page", 1)
		size := controller.GetQueryInt(r, "size", 10)
		keyword := controller.GetQueryString(r, "keyword", "")
		sortBy := controller.GetQueryString(r, "sortBy", "")
		sortOrder := controller.GetQueryString(r, "sortOrder", "")
		media := controller.GetQueryString(r, "media", "")
		status := controller.GetQueryString(r, "status", "")

		req := usecase.PengkomunikasianPengendalianGetAllUseCaseReq{
			Page: page, Size: size, Keyword: keyword, SortBy: sortBy, SortOrder: sortOrder, Status: status, Media: media,
		}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
