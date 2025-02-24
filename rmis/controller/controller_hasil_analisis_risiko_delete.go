// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// HasilAnalisisRisikoDeleteHandler handles deleting a HasilAnalisisRisiko
func (c Controller) HasilAnalisisRisikoDeleteHandler(u usecase.HasilAnalisisRisikoDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/hasil-analisis-risikos/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "hasil-analisis-risikos",
			Relation:  "delete",
		},
		Summary: "Delete a Hasil Analisis Risiko",
		Tag:     "Hasil Analisis Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.HasilAnalisisRisikoDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
