package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IdentifikasiRisikoStrategisOPDCreateHandler handles the creation of a new IdentifikasiRisikoStrategisOPD
func (c Controller) IdentifikasiRisikoStrategisOPDCreateHandler(u usecase.IdentifikasiRisikoStrategisOPDCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/identifikasi-risiko-strategis-opd",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "identifikasi-risiko-strategis-opd",
			Relation:  "create",
		},
		Body:    usecase.IdentifikasiRisikoStrategisOPDCreateUseCaseReq{},
		Summary: "Create a new Identifikasi Risiko Strategis OPD",
		Tag:     "Identifikasi Risiko Strategis OPD",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.IdentifikasiRisikoStrategisOPDCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
