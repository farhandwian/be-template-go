package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenetapanKonteksRisikoStrategisRenstraOPDCreateHandler handles the creation of a new PenetapanKonteksRisikoStrategisRenstraOPD
func (c Controller) PenetapanKonteksRisikoStrategisRenstraOPDCreateHandler(u usecase.PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/penetapan-konteks-risiko-strategis-renstra-opd",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-strategis-renstra-opd",
			Relation:  "create",
		},
		Body:    usecase.PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseReq{},
		Summary: "Create a new Penetapan Konteks Risiko Strategis Renstra OPD",
		Tag:     "Penetapan Konteks Risiko Strategis Renstra OPD",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
