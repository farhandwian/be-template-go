package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"perizinan/usecase"
	"shared/helper"
)

func LaporanPerizinanGetAllHandler(mux *http.ServeMux, u usecase.LaporGetAllUseCase, jwt helper.JWTTokenizer) helper.APIData {

	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/api/perizinan",
		Access:  iammodel.SI_JAGACAI_SISTEM_PEMANTAUAN_TABEL_DAFTAR_LAPORAN_READ,
		Summary: "Get All Laporan",
		Tag:     "Laporan Perizinan",
		QueryParams: []helper.QueryParam{
			{Name: "keyword", Type: "string", Description: "no sk", Required: false},
			{Name: "periode_pengambilan_sda", Type: "string", Description: "periode pengambilan sda", Required: false},
			{Name: "page", Type: "integer", Description: "page", Required: false},
			{Name: "size", Type: "integer", Description: "size", Required: false},
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
		// 		FunctionName: "perizinan_read",
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

		keyword := controller.GetQueryString(r, "keyword", "")
		periodePengambilanSda := controller.GetQueryString(r, "periode_pengambilan_sda", "")
		page := controller.GetQueryInt(r, "page", 1)
		size := controller.GetQueryInt(r, "size", 10)
		sortBy := controller.GetQueryString(r, "sort_by", "")
		sortOrder := controller.GetQueryString(r, "sort_order", "")

		request := usecase.LaporGetAllUseCaseReq{Page: page, Size: size, Keyword: keyword, PeriodePengambilanSda: periodePengambilanSda, SortBy: sortBy, SortOrder: sortOrder}
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	// mux.HandleFunc(apiData.GetMethodUrl(), handler)

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, jwt)
	mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
