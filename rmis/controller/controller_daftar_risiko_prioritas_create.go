package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// DaftarRisikoPrioritasCreateHandler handles the creation of a new DaftarRisikoPrioritas
func (c Controller) DaftarRisikoPrioritasCreateHandler(u usecase.DaftarRisikoPrioritasCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/daftar-risiko-prioritas",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "daftar-risiko-prioritas",
			Relation:  "create",
		},
		Body:    usecase.DaftarRisikoPrioritasCreateUseCaseReq{},
		Summary: "Create a new Daftar Risiko Prioritas",
		Tag:     "Daftar Risiko Prioritas",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.DaftarRisikoPrioritasCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
