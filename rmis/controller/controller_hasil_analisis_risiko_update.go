package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// HasilAnalisisRisikoUpdateHandler handles the creation of a new HasilAnalisisRisiko
func (c Controller) HasilAnalisisRisikoUpdateHandler(u usecase.HasilAnalisisRisikoUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/hasil-analisis-risikos/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "hasil-analisis-risikos",
			Relation:  "update",
		},
		Body:    usecase.HasilAnalisisRisikoUpdateUseCaseReq{},
		Summary: "Update a Hasil Analisis Risiko",
		Tag:     "Hasil Analisis Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.HasilAnalisisRisikoUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
