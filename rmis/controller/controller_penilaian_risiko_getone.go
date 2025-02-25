package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenilaianRisikoGetByIDHandler handles getting a PenilaianRisiko by ID
func (c Controller) PenilaianRisikoGetByIDHandler(u usecase.PenilaianRisikoGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/penilaian-risiko/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penilaian-risiko",
			Relation:  "read",
		},
		Summary: "Get a Penilaian Risiko by ID",
		Tag:     "Penilaian Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.PenilaianRisikoGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
