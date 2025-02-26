package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// RancanganPemantauan Get All handler
func (c Controller) RancanganPemantauanGetAllHandler(u usecase.RancanganPemantauanGetAllUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/rancangan-pemantauan",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "rancangan-pemantauan",
			Relation:  "read",
		},
		Summary: "Get all Rancangan Pemantauan",
		Tag:     "Rancangan Pemantauan",
		QueryParams: []helper.QueryParam{
			{Name: "keyword", Type: "string", Description: "name, pic or location", Required: false},
			{Name: "page", Type: "number", Description: "page", Required: false},
			{Name: "size", Type: "number", Description: "size", Required: false},
			{Name: "sortBy", Type: "string", Description: "sort by", Required: false},
			{Name: "sortOrder", Type: "string", Description: "sort order", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		page := controller.GetQueryInt(r, "page", 1)
		size := controller.GetQueryInt(r, "size", 10)
		keyword := controller.GetQueryString(r, "keyword", "")
		sortBy := controller.GetQueryString(r, "sortBy", "")
		sortOrder := controller.GetQueryString(r, "sortOrder", "")

		req := usecase.RancanganPemantauanGetAllUseCaseReq{Page: page, Size: size, Keyword: keyword, SortBy: sortBy, SortOrder: sortOrder}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
