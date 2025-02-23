package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IdentifikasiRisikoStrategisPemdaUpdateHandler handles the creation of a new IdentifikasiRisikoStrategisPemda
func (c Controller) IdentifikasiRisikoStrategisPemdaUpdateHandler(u usecase.IdentifikasiRisikoStrategisPemdaUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/identifikasi-risiko-strategis-pemdas/{id}",
		AccessTest: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "identifikasi-risiko-strategis-pemdas",
			Relation:  "update",
		},
		Body:    usecase.IdentifikasiRisikoStrategisPemdaUpdateUseCaseReq{},
		Summary: "Update a Identifikasi Risiko Strategis Pemda",
		Tag:     "Identifikasi Risiko Strategis Pemda",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.IdentifikasiRisikoStrategisPemdaUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
