package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IdentifikasiRisikoOperasionalOPDApprovalHandler handles the creation of a new IdentifikasiRisikoOperasionalOPD
func (c Controller) IdentifikasiRisikoOperasionalOPDApprovalHandler(u usecase.IdentifikasiRisikoOperasionalOPDApprovalUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/identifikasi-risiko-operasional-opd/{id}/approval",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "identifikasi-risiko-operasional-opd",
			Relation:  "approval",
		},
		Body:    usecase.IdentifikasiRisikoOperasionalOPDApprovalUseCaseReq{},
		Summary: "Approval a Identifikasi Risiko Operasional OPD",
		Tag:     "Identifikasi Risiko Operasional OPD",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.IdentifikasiRisikoOperasionalOPDApprovalUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
