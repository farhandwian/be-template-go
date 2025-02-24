package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenyebabRisikoUpdateHandler handles the creation of a new PenyebabRisiko
func (c Controller) PenyebabRisikoUpdateHandler(u usecase.PenyebabRisikoUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/penyebab-risikos/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penyebab-risikos",
			Relation:  "update",
		},
		Body:    usecase.PenyebabRisikoUpdateUseCaseReq{},
		Summary: "Update a Penyebab Risiko",
		Tag:     "Penyebab Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.PenyebabRisikoUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
