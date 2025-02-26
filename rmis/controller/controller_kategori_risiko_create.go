package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PencatatanKejadianRisikoCreateHandler handles the creation of a new PencatatanKejadianRisiko
func (c Controller) PencatatanKejadianRisikoCreateHandler(u usecase.PencatatanKejadianRisikoCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/pencatatan-kejadian-risiko",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "pencatatan-kejadian-risiko",
			Relation:  "create",
		},
		Body:    usecase.PencatatanKejadianRisikoCreateUseCaseReq{},
		Summary: "Create a new Pencatatan Kejadian Risiko",
		Tag:     "Pencatatan Kejadian Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.PencatatanKejadianRisikoCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
