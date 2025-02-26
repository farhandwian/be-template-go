package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IdentifikasiRisikoStrategisOPDGetByIDHandler handles getting a IdentifikasiRisikoStrategisOPD by ID
func (c Controller) IdentifikasiRisikoStrategisOPDGetByIDHandler(u usecase.IdentifikasiRisikoStrategisOPDGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/identifikasi-risiko-strategis-opd/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "identifikasi-risiko-strategis-opd",
			Relation:  "read",
		},
		Summary: "Get a Identifikasi Risiko Strategis OPD by ID",
		Tag:     "Identifikasi Risiko Strategis OPD",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.IdentifikasiRisikoStrategisOPDGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
