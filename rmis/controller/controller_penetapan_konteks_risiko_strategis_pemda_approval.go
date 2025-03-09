package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenetapanKonteksRisikoStrategisPemdaApprovalHandler handles the creation of a new PenetapanKonteksRisikoStrategisPemda
func (c Controller) PenetapanKonteksRisikoStrategisPemdaApprovalHandler(u usecase.PenetapanKonteksRisikoStrategisPemdaApprovalUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/penetapan-konteks-risiko-strategis-pemdas/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-strategis-pemdas",
			Relation:  "Approval",
		},
		Body:    usecase.PenetapanKonteksRisikoStrategisPemdaApprovalUseCaseReq{},
		Summary: "Approval a Penetapan Konteks Risiko Strategis Pemda",
		Tag:     "Penetapan Konteks Risiko Strategis Pemda",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.PenetapanKonteksRisikoStrategisPemdaApprovalUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
