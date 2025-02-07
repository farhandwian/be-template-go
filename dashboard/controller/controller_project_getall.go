package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) ProjectGetAllHandler(u usecase.ProjectGetAllUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/projects",
		Access:  iammodel.MASTER_DATA_DAFTAR_PROYEK_READ,
		Summary: "Get all Projects",
		Tag:     "Master Data",
		QueryParams: []helper.QueryParam{
			{Name: "keyword", Type: "string", Description: "document name", Required: false},
			{Name: "page", Type: "integer", Description: "page", Required: false},
			{Name: "size", Type: "integer", Description: "size", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		page := controller.GetQueryInt(r, "page", 1)
		size := controller.GetQueryInt(r, "size", 10)
		keyword := controller.GetQueryString(r, "keyword", "")
		req := usecase.ProjectGetAllUseCaseReq{Page: page, Size: size, Keyword: keyword}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
