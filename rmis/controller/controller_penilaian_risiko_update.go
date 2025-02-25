package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenilaianRisikoUpdateHandler handles the creation of a new PenilaianRisiko
func (c Controller) PenilaianRisikoUpdateHandler(u usecase.PenilaianRisikoUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/penilaian-risiko/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penilaian-risiko",
			Relation:  "update",
		},
		Body:    usecase.PenilaianRisikoUpdateUseCaseReq{},
		Summary: "Update a Penilaian Risiko",
		Tag:     "Penilaian Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.PenilaianRisikoUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
