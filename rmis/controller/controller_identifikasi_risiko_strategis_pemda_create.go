package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IdentifikasiRisikoStrategisPemdaCreateHandler handles the creation of a new IdentifikasiRisikoStrategisPemda
func (c Controller) IdentifikasiRisikoStrategisPemdaCreateHandler(u usecase.IdentifikasiRisikoStrategisPemdaCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/identifikasi-risiko-strategis-pemdas",
		AccessTest: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "identifikasi-risiko-strategis-pemdas",
			Relation:  "create",
		},
		Body:    usecase.IdentifikasiRisikoStrategisPemdaCreateUseCaseReq{},
		Summary: "Create a new Identifikasi Risiko Strategis Pemda",
		Tag:     "Identifikasi Risiko Strategis Pemda",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.IdentifikasiRisikoStrategisPemdaCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
