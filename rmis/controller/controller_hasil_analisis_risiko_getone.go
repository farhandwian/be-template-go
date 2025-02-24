package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// HasilAnalisisRisikoGetByIDHandler handles getting a HasilAnalisisRisiko by ID
func (c Controller) HasilAnalisisRisikoGetByIDHandler(u usecase.HasilAnalisisRisikoGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/hasil-analisis-risikos/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "hasil-analisis-risikos",
			Relation:  "read",
		},
		Summary: "Get a Hasil Analisis Risiko by ID",
		Tag:     "Hasil Analisis Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.HasilAnalisisRisikoGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
