package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"perizinan/usecase"
	"shared/helper"
)

func SKGetAllHandler(mux *http.ServeMux, u usecase.SKGetAllUseCase, jwt helper.JWTTokenizer) helper.APIData {

	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/api/sk",
		Access:  iammodel.SI_JAGACAI_TABLE_REKOMTEK_READ,
		Summary: "Get All SK",
		Tag:     "Laporan Perizinan",
		QueryParams: []helper.QueryParam{
			{Name: "keyword", Type: "string", Description: "nosk", Required: false},
			{Name: "page", Type: "number", Description: "page", Required: false},
			{Name: "size", Type: "number", Description: "size", Required: false},
			{Name: "sort_by", Type: "string", Description: "sort key", Required: false},
			{Name: "sort_order", Type: "string", Description: "sort order", Required: false},
		},
	}

	// checkAccess := ImplCheckAccess()

	handler := func(w http.ResponseWriter, r *http.Request) {

		// // Simple check access right
		// {
		// 	bearerToken, _, ok := GetBearerToken(w, r)
		// 	if !ok {
		// 		http.Error(w, "unauthorized operation", http.StatusUnauthorized)
		// 		return
		// 	}

		// 	checkAccessRes, err := checkAccess(context.TODO(), CheckAccessReq{
		// 		FunctionName: "sk_read",
		// 		Token:        bearerToken,
		// 	})

		// 	if err != nil {
		// 		http.Error(w, "unauthorized operation", http.StatusUnauthorized)
		// 		return
		// 	}

		// 	if !checkAccessRes.Data.CanAccess {
		// 		http.Error(w, "unauthorized operation", http.StatusUnauthorized)
		// 		return
		// 	}
		// }

		page := controller.GetQueryInt(r, "page", 1)
		size := controller.GetQueryInt(r, "size", 10)
		keyword := controller.GetQueryString(r, "keyword", "")
		sortBy := controller.GetQueryString(r, "sort_by", "")
		sortOrder := controller.GetQueryString(r, "sort_order", "")

		request := usecase.SKGetAllUseCaseReq{Page: page, Size: size, Keyword: keyword, SortBy: sortBy, SortOrder: sortOrder}
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	// mux.HandleFunc(apiData.GetMethodUrl(), handler)

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, jwt)
	mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
