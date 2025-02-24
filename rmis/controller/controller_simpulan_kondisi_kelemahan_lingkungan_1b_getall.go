package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// SimpulanKondisiKelemahanLingkungan Get All handler
func (c Controller) SimpulanKondisiKelemahanLingkunganGetAllHandler(u usecase.SimpulanKondisiKelemahanLingkunganGetAllUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/simpulan-kondisi-kelemahan-lingkungans",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "simpulan-kondisi-kelemahan-lingkungans",
			Relation:  "read",
		},
		Summary: "Get all Sub Unsur SimpulanKondisiKelemahanLingkungan",
		Tag:     "Sub Unsur SimpulanKondisiKelemahanLingkungan",
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
		req := usecase.SimpulanKondisiKelemahanLingkunganGetAllUseCaseReq{Page: page, Size: size, Keyword: keyword}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
