package controller

import (
	"dashboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) WaterChannelDoorByKeywordHandler(u usecase.WaterChannelDoorByKeyword) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/channels",
		Access:  model.DEFAULT_OPERATION,
		Summary: "Get all water channel doors name and id",
		Tag:     "Dashboard",
		QueryParams: []helper.QueryParam{
			{Name: "keyword", Type: "string", Description: "nama water channer doors", Required: true},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		keyword := controller.GetQueryString(r, "keyword", "")
		req := usecase.WaterChannelDoorByKeywordReq{Keyword: keyword}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
