// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IdentifikasiRisikoStrategisPemdaDeleteHandler handles deleting a IdentifikasiRisikoStrategisPemda
func (c Controller) IdentifikasiRisikoStrategisPemdaDeleteHandler(u usecase.IdentifikasiRisikoStrategisPemdaDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/identifikasi-risiko-strategis-pemdas/{id}",
		AccessTest: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "identifikasi-risiko-strategis-pemdas",
			Relation:  "delete",
		},
		Summary: "Delete a  Identifikasi Risiko Strategis Pemda",
		Tag:     "Identifikasi Risiko Strategis Pemda",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.IdentifikasiRisikoStrategisPemdaDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
