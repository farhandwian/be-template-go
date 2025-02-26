package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PencatatanKejadianRisikoUpdateHandler handles the creation of a new PencatatanKejadianRisiko
func (c Controller) PencatatanKejadianRisikoUpdateHandler(u usecase.PencatatanKejadianRisikoUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/pencatatan-kejadian-risiko/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "pencatatan-kejadian-risiko",
			Relation:  "update",
		},
		Body:    usecase.PencatatanKejadianRisikoUpdateUseCaseReq{},
		Summary: "Update a Pencatatan Kejadian Risiko",
		Tag:     "Pencatatan Kejadian Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.PencatatanKejadianRisikoUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
