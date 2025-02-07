package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"perizinan/model"
	"perizinan/usecase"
	"shared/helper"
	"time"
)

func LaporanPerizinanSKPeriodeHandler(mux *http.ServeMux, u usecase.LaporanPerizinanSKPeriodeUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/api/perizinan/sk-periode",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get Laporan by SK and Periode",
		Tag:     "Laporan Perizinan",
		QueryParams: []helper.QueryParam{
			{Name: "nomor_sk", Type: "string", Description: "Nomor SK", Required: false},
			{Name: "periode_pengambilan_sda", Type: "string", Description: "Masa Berlaku SK", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		nomorSK := model.NomorSK(controller.GetQueryString(r, "nomor_sk", ""))
		periode := model.Periode(controller.GetQueryString(r, "periode_pengambilan_sda", ""))

		request := usecase.LaporanPerizinanSKPeriodeUseCaseReq{
			NomorSK:               nomorSK,
			PeriodePengambilanSDA: periode,
			Min:                   time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC),
			Now:                   time.Now(),
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
