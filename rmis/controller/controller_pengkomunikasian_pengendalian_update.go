package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PengkomunikasianPengendalianUpdateHandler handles the creation of a new PengkomunikasianPengendalian
func (c Controller) PengkomunikasianPengendalianUpdateHandler(u usecase.PengkomunikasianPengendalianUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/pengkomunikasian-pengendalian/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "pengkomunikasian-pengendalian",
			Relation:  "update",
		},
		Body:    usecase.PengkomunikasianPengendalianUpdateUseCaseReq{},
		Summary: "Update a Pengkomunikasian Pengendalian",
		Tag:     "Pengkomunikasian Pengendalian",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.PengkomunikasianPengendalianUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
