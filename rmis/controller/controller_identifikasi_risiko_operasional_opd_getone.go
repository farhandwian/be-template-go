package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IdentifikasiRisikoOperasionalOPDGetByIDHandler handles getting a IdentifikasiRisikoOperasionalOPD by ID
func (c Controller) IdentifikasiRisikoOperasionalOPDGetByIDHandler(u usecase.IdentifikasiRisikoOperasionalOPDGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/identifikasi-risiko-operasional-opd/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "identifikasi-risiko-operasional-opd",
			Relation:  "read",
		},
		Summary: "Get a Identifikasi Risiko Operasional OPD by ID",
		Tag:     "Identifikasi Risiko Operasional OPD",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.IdentifikasiRisikoOperasionalOPDGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
