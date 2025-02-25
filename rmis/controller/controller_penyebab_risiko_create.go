package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenyebabRisikoCreateHandler handles the creation of a new PenyebabRisiko
func (c Controller) PenyebabRisikoCreateHandler(u usecase.PenyebabRisikoCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/penyebab-risikos",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penyebab-risikos",
			Relation:  "create",
		},
		Body:    usecase.PenyebabRisikoCreateUseCaseReq{},
		Summary: "Create a new Penyebab Risiko",
		Tag:     "Penyebab Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.PenyebabRisikoCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
