package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"perizinan/usecase"
	"shared/helper"
)

// LaporanPerizinanSaveHandler handles the creation of a new LaporanPerizinan
func LaporanPerizinanSaveHandler(mux *http.ServeMux, u usecase.LaporanPerizinanSaveUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/api/perizinan",
		Access:  iammodel.DEFAULT_OPERATION,
		Body:    usecase.LaporanPerizinanSaveUseCaseReq{},
		Summary: "Save a new LaporanPerizinan",
		Tag:     "Laporan Perizinan",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.LaporanPerizinanSaveUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
