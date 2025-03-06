package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IdentifikasiRisikoStrategisOPDApprovalHandler handles the creation of a new IdentifikasiRisikoStrategisOPD
func (c Controller) IdentifikasiRisikoStrategisOPDApprovalHandler(u usecase.IdentifikasiRisikoStrategisOPDApprovalUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/identifikasi-risiko-strategis-opd/{id}/approval",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "identifikasi-risiko-strategis-opd",
			Relation:  "approval",
		},
		Body:    usecase.IdentifikasiRisikoStrategisOPDApprovalUseCaseReq{},
		Summary: "Approval a Identifikasi Risiko Strategis OPD",
		Tag:     "Identifikasi Risiko Strategis OPD",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.IdentifikasiRisikoStrategisOPDApprovalUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
