package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IdentifikasiRisikoStrategisPemdaGetByIDHandler handles getting a IdentifikasiRisikoStrategisPemda by ID
func (c Controller) IdentifikasiRisikoStrategisPemdaGetByIDHandler(u usecase.IdentifikasiRisikoStrategisPemdaGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/identifikasi-risiko-strategis-pemdas/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "identifikasi-risiko-strategis-pemdas",
			Relation:  "read",
		},
		Summary: "Get a  Identifikasi Risiko Strategis Pemda by ID",
		Tag:     " Identifikasi Risiko Strategis Pemda",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.IdentifikasiRisikoStrategisPemdaGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
