// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IdentifikasiRisikoOperasionalOPDDeleteHandler handles deleting a IdentifikasiRisikoOperasionalOPD
func (c Controller) IdentifikasiRisikoOperasionalOPDDeleteHandler(u usecase.IdentifikasiRisikoOperasionalOPDDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/identifikasi-risiko-operasional-opd/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "identifikasi-risiko-operasional-opd",
			Relation:  "delete",
		},
		Summary: "Delete a  Identifikasi Risiko Operasional OPD",
		Tag:     "Identifikasi Risiko Operasional OPD",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.IdentifikasiRisikoOperasionalOPDDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
