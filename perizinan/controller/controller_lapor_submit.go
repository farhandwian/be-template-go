package controller

import (
	iammodel "iam/model"
	"net/http"
	"perizinan/model"
	"perizinan/usecase"
	"iam/controller"
	"shared/helper"
)

func LaporanPerizinanSubmitHandler(mux *http.ServeMux, u usecase.LaporanPerizinanSubmitUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/api/perizinan/{id}/submit",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Submit LaporanPerizinan",
		Tag:     "Laporan Perizinan",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request := usecase.LaporanPerizinanSubmitUseCaseReq{
			LaporanPerizinanID: model.LaporanPerizinanID(r.PathValue("id")),
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
