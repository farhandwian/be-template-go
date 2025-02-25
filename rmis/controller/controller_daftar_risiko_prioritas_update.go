package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// DaftarRisikoPrioritasUpdateHandler handles the creation of a new DaftarRisikoPrioritas
func (c Controller) DaftarRisikoPrioritasUpdateHandler(u usecase.DaftarRisikoPrioritasUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/daftar-risiko-prioritas/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "daftar-risiko-prioritas",
			Relation:  "update",
		},
		Body:    usecase.DaftarRisikoPrioritasUpdateUseCaseReq{},
		Summary: "Update a Daftar Risiko Prioritas",
		Tag:     "Daftar Risiko Prioritas",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.DaftarRisikoPrioritasUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
