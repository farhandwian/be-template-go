package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// RekapitulasiHasilKuesionerGetByIDHandler handles getting a RekapitulasiHasilKuesioner by ID
func (c Controller) RekapitulasiHasilKuesionerGetByIDHandler(u usecase.RekapitulasiHasilKuesionerGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/rekapitulasi-hasil-kuesioners/{id}",
		AccessTest: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "rekapitulasi-hasil-kuesioners",
			Relation:  "read",
		},
		Summary: "Get a Sub Unsur RekapitulasiHasilKuesioner by ID",
		Tag:     "Sub Unsur RekapitulasiHasilKuesioner",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.RekapitulasiHasilKuesionerGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
