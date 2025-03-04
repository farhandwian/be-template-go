package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// HasilAnalisisRisiko Get All handler
func (c Controller) HasilAnalisisRisikoGetAllHandler(u usecase.HasilAnalisisRisikoGetAllUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/hasil-analisis-risiko",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "hasil-analisis-risiko",
			Relation:  "read",
		},
		Summary: "Get all Hasil Analisis Risiko",
		Tag:     "Hasil Analisis Risiko",
		QueryParams: []helper.QueryParam{
			{Name: "keyword", Type: "string", Description: "name, pic or location", Required: false},
			{Name: "page", Type: "number", Description: "page", Required: false},
			{Name: "size", Type: "number", Description: "size", Required: false},
			{Name: "sortBy", Type: "string", Description: "sort by", Required: false},
			{Name: "sortOrder", Type: "string", Description: "sort order", Required: false},
			{Name: "status", Type: "string", Description: "status filter", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		page := controller.GetQueryInt(r, "page", 1)
		size := controller.GetQueryInt(r, "size", 10)
		keyword := controller.GetQueryString(r, "keyword", "")
		sortBy := controller.GetQueryString(r, "sortBy", "")
		sortOrder := controller.GetQueryString(r, "sortOrder", "")
		status := controller.GetQueryString(r, "status", "")

		req := usecase.HasilAnalisisRisikoGetAllUseCaseReq{
			Page: page, Size: size, Keyword: keyword, SortBy: sortBy, SortOrder: sortOrder, Status: status,
		}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
