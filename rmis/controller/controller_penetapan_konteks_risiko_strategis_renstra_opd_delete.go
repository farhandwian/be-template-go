// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenetapanKonteksRisikoStrategisRenstraOPDDeleteHandler handles deleting a PenetapanKonteksRisikoStrategisRenstraOPD
func (c Controller) PenetapanKonteksRisikoStrategisRenstraOPDDeleteHandler(u usecase.PenetapanKonteksRisikoRenstraOPDDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/penetapan-konteks-risiko-strategis-renstra-opd/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-strategis-renstra-opd",
			Relation:  "delete",
		},
		Summary: "Delete a Penetapan Konteks Risiko Strategis Renstra OPD",
		Tag:     "Penetapan Konteks Risiko Strategis Renstra OPD",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.PenetapanKonteksRisikoRenstraOPDDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
