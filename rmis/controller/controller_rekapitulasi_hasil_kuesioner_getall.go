package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// RekapitulasiHasilKuesioner Get All handler
func (c Controller) RekapitulasiHasilKuesionerGetAllHandler(u usecase.RekapitulasiHasilKuesionerGetAllUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/rekapitulasi-hasil-kuesioners",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "rekapitulasi-hasil-kuesioners",
			Relation:  "read",
		},
		Summary: "Get all Sub Unsur RekapitulasiHasilKuesioner",
		Tag:     "Sub Unsur RekapitulasiHasilKuesioner",
		QueryParams: []helper.QueryParam{
			{Name: "keyword", Type: "string", Description: "name, pic or location", Required: false},
			{Name: "page", Type: "number", Description: "page", Required: false},
			{Name: "size", Type: "number", Description: "size", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		page := controller.GetQueryInt(r, "page", 1)
		size := controller.GetQueryInt(r, "size", 10)
		keyword := controller.GetQueryString(r, "keyword", "")
		req := usecase.RekapitulasiHasilKuesionerGetAllUseCaseReq{Page: page, Size: size, Keyword: keyword}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
