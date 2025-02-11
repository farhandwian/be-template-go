package controller

import (
	"example/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) ExampleGetAllHandler(u usecase.ExampleGetAllUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  iammodel.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/dashboard/example/",
		Summary: "Get list example",
		Tag:     "Example Tag",
		QueryParams: []helper.QueryParam{
			{Name: "test_string", Type: "string", Description: "Test parameter string", Required: false},
			{Name: "test_number", Type: "number", Description: "Test parameter number", Required: false},
			{Name: "test_boolean", Type: "boolean", Description: "Test parameter boolean", Required: false},
			{Name: "page", Type: "number", Description: "Page", Required: true},
			{Name: "page_size", Type: "number", Description: "page size", Required: true},
			{Name: "sort_by", Type: "string", Description: "sort key", Required: false},
			{Name: "sort_order", Type: "string", Description: "sort order", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		testString := controller.GetQueryString(r, "test_string", "")
		testNumber := controller.GetQueryInt(r, "test_number", 0)
		testBoolean := controller.GetQueryBoolean(r, "test_boolean", false)
		page := controller.GetQueryInt(r, "page", 1)
		pageSize := controller.GetQueryInt(r, "page_size", 10)
		sortBy := controller.GetQueryString(r, "sort_by", "")
		sortOrder := controller.GetQueryString(r, "sort_order", "")

		req := usecase.ExampleGetAllUseReq{
			TestString:  testString,
			TestNumber:  testNumber,
			TestBoolean: testBoolean,
			Page:        page,
			Size:        pageSize,
			SortBy:      sortBy,
			SortOrder:   sortOrder,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData

}
