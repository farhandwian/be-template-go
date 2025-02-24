package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// HasilAnalisisRisikoCreateHandler handles the creation of a new HasilAnalisisRisiko
func (c Controller) HasilAnalisisRisikoCreateHandler(u usecase.HasilAnalisisRisikoCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/hasil-analisis-risikos",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "hasil-analisis-risikos",
			Relation:  "create",
		},
		Body:    usecase.HasilAnalisisRisikoCreateUseCaseReq{},
		Summary: "Create a new Hasil Analisis Risiko",
		Tag:     "Hasil Analisis Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.HasilAnalisisRisikoCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
