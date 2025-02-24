// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// RekapitulasiHasilKuesionerDeleteHandler handles deleting a RekapitulasiHasilKuesioner
func (c Controller) RekapitulasiHasilKuesionerDeleteHandler(u usecase.RekapitulasiHasilKuesionerDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/rekapitulasi-hasil-kuesioners/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "rekapitulasi-hasil-kuesioners",
			Relation:  "delete",
		},
		Summary: "Delete a Rekapitulasi hasil kuesioner",
		Tag:     "Rekapitulasi hasil kuesioner",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.RekapitulasiHasilKuesionerDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
