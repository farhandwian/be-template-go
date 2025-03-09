package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// DaftarRisikoPrioritasApprovalHandler handles the creation of a new DaftarRisikoPrioritas
func (c Controller) DaftarRisikoPrioritasApprovalHandler(u usecase.DaftarRisikoPrioritasApprovalUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/daftar-risiko-prioritas/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "daftar-risiko-prioritas",
			Relation:  "approval",
		},
		Body:    usecase.DaftarRisikoPrioritasApprovalUseCaseReq{},
		Summary: "Approval a Daftar Risiko Prioritas",
		Tag:     "Daftar Risiko Prioritas",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.DaftarRisikoPrioritasApprovalUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
