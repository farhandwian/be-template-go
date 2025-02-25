package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// DaftarRisikoPrioritasGetByIDHandler handles getting a DaftarRisikoPrioritas by ID
func (c Controller) DaftarRisikoPrioritasGetByIDHandler(u usecase.DaftarRisikoPrioritasGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/daftar-risiko-prioritas/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "daftar-risiko-prioritas",
			Relation:  "read",
		},
		Summary: "Get a Daftar Risiko Prioritas by ID",
		Tag:     "Daftar Risiko Prioritas",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.DaftarRisikoPrioritasGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
