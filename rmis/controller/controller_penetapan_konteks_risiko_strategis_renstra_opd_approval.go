package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenetapanKonteksRisikoStrategisRenstraOPDApprovalHandler handles the creation of a new PenetapanKonteksRisikoStrategisRenstraOPD
func (c Controller) PenetapanKonteksRisikoStrategisRenstraOPDApprovalHandler(u usecase.PenetapanKonteksRisikoStrategisRenstraOPDApprovalUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/penetapan-konteks-risiko-strategis-renstra-opd/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-strategis-renstra-opd",
			Relation:  "Approval",
		},
		Body:    usecase.PenetapanKonteksRisikoStrategisRenstraOPDApprovalUseCaseReq{},
		Summary: "Approval a Penetapan Konteks Risiko Strategis Renstra OPD",
		Tag:     "Penetapan Konteks Risiko Strategis Renstra OPD",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.PenetapanKonteksRisikoStrategisRenstraOPDApprovalUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
