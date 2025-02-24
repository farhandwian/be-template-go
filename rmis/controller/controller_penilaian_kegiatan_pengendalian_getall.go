package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenilaianKegiatanPengendalian Get All handler
func (c Controller) PenilaianKegiatanPengendalianGetAllHandler(u usecase.PenilaianKegiatanPengendalianGetAllUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/penilaian-kegiatan-pengendalians",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penilaian-kegiatan-pengendalians",
			Relation:  "read",
		},
		Summary: "Get all Penilaian Kegiatan Pengendalian",
		Tag:     "Penilaian Kegiatan Pengendalian",
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
		req := usecase.PenilaianKegiatanPengendalianGetAllUseCaseReq{Page: page, Size: size, Keyword: keyword}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
