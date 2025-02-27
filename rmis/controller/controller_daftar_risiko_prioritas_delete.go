package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// DaftarRisikoPrioritasDeleteHandler handles deleting a DaftarRisikoPrioritas
func (c Controller) DaftarRisikoPrioritasDeleteHandler(u usecase.DaftarRisikoPrioritasDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/daftar-risiko-prioritas/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "daftar-risiko-prioritas",
			Relation:  "delete",
		},
		Summary: "Delete a Daftar Risiko Prioritas",
		Tag:     "Daftar Risiko Prioritas",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.DaftarRisikoPrioritasDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
