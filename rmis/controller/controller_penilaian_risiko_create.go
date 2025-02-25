package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenilaianRisikoCreateHandler handles the creation of a new PenilaianRisiko
func (c Controller) PenilaianRisikoCreateHandler(u usecase.PenilaianRisikoCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/penilaian-risiko",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penilaian-risiko",
			Relation:  "create",
		},
		Body:    usecase.PenilaianRisikoCreateUseCaseReq{},
		Summary: "Create a new Penilaian Risiko",
		Tag:     "Penilaian Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.PenilaianRisikoCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
