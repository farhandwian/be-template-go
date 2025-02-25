package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenetapanKonteksRisikoOperasionalUpdateHandler handles the creation of a new PenetapanKonteksRisikoOperasional
func (c Controller) PenetapanKonteksRisikoOperasionalUpdateHandler(u usecase.PenetapanKonteksRisikoOperasionalUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/penetapan-konteks-risiko-operasional/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-operasional",
			Relation:  "update",
		},
		Body:    usecase.PenetapanKonteksRisikoOperasionalUpdateUseCaseReq{},
		Summary: "Update a Penetapan Konteks Risiko Operasional",
		Tag:     "Penetapan Konteks Risiko Operasional",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.PenetapanKonteksRisikoOperasionalUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
