package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// HasilAnalisisRisikoApprovalHandler handles the creation of a new HasilAnalisisRisiko
func (c Controller) HasilAnalisisRisikoApprovalHandler(u usecase.HasilAnalisisRisikoApprovalUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/hasil-analisis-risiko/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "hasil-analisis-risiko",
			Relation:  "approval",
		},
		Body:    usecase.HasilAnalisisRisikoApprovalUseCaseReq{},
		Summary: "Approval a Hasil Analisis Risiko",
		Tag:     "Hasil Analisis Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.HasilAnalisisRisikoApprovalUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
