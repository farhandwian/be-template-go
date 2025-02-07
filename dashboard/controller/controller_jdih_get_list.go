package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) GetListJDIH(u usecase.GetListDocumentAndLawUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/jdih",
		Access:  iammodel.MASTER_DATA_DAFTAR_JDIH_READ,
		Summary: "Get all JDIH",
		Tag:     "Master Data",
		QueryParams: []helper.QueryParam{
			{Name: "keyword", Type: "string", Description: "document name", Required: false},
			{Name: "page", Type: "number", Description: "Page", Required: false},
			{Name: "size", Type: "number", Description: "Size", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		page := controller.GetQueryInt(r, "page", 1)
		size := controller.GetQueryInt(r, "size", 10)
		keyword := controller.GetQueryString(r, "keyword", "")
		req := usecase.GetListJDIHReq{Keyword: keyword, Page: page, Size: size}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
