package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenetapanKonteksRisikoOperasionalCreateHandler handles the creation of a new PenetapanKonteksRisikoOperasional
func (c Controller) PenetapanKonteksRisikoOperasionalCreateHandler(u usecase.PenetapanKonteksRisikoOperasionalCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/penetapan-konteks-risiko-operasional",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-operasional",
			Relation:  "create",
		},
		Body:    usecase.PenetapanKonteksRisikoOperasionalCreateUseCaseReq{},
		Summary: "Create a new Penetapan Konteks Risiko Operasional",
		Tag:     "Penetapan Konteks Risiko Operasional",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.PenetapanKonteksRisikoOperasionalCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
