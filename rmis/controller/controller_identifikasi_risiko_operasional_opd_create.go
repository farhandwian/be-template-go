package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IdentifikasiRisikoOperasionalOPDCreateHandler handles the creation of a new IdentifikasiRisikoOperasionalOPD
func (c Controller) IdentifikasiRisikoOperasionalOPDCreateHandler(u usecase.IdentifikasiRisikoOperasionalOPDCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/identifikasi-risiko-operasional-opd",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "identifikasi-risiko-operasional-opd",
			Relation:  "create",
		},
		Body:    usecase.IdentifikasiRisikoOperasionalOPDCreateUseCaseReq{},
		Summary: "Create a new Identifikasi Risiko Operasional OPD",
		Tag:     "Identifikasi Risiko Operasional OPD",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.IdentifikasiRisikoOperasionalOPDCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
