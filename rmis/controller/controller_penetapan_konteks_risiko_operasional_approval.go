package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenetapanKonteksRisikoOperasionalApprovalHandler handles the creation of a new PenetapanKonteksRisikoOperasional
func (c Controller) PenetapanKonteksRisikoOperasionalApprovalHandler(u usecase.PenetapanKonteksRisikoOperasionalApprovalUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/penetapan-konteks-risiko-operasional/{id}/approval",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-operasional",
			Relation:  "approval",
		},
		Body:    usecase.PenetapanKonteksRisikoOperasionalApprovalUseCaseReq{},
		Summary: "Approval a new Penetapan Konteks Risiko Operasional",
		Tag:     "Penetapan Konteks Risiko Operasional",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		request, ok := controller.ParseJSON[usecase.PenetapanKonteksRisikoOperasionalApprovalUseCaseReq](w, r)
		if !ok {
			return
		}

		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
