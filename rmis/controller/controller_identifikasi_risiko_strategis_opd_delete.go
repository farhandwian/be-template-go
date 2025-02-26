// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IdentifikasiRisikoStrategisOPDDeleteHandler handles deleting a IdentifikasiRisikoStrategisOPD
func (c Controller) IdentifikasiRisikoStrategisOPDDeleteHandler(u usecase.IdentifikasiRisikoStrategisOPDDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/identifikasi-risiko-strategis-opd/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "identifikasi-risiko-strategis-opd",
			Relation:  "delete",
		},
		Summary: "Delete a  Identifikasi Risiko Strategis OPD",
		Tag:     "Identifikasi Risiko Strategis OPD",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.IdentifikasiRisikoStrategisOPDDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
