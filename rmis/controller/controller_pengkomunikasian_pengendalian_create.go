package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PengkomunikasianPengendalianCreateHandler handles the creation of a new PengkomunikasianPengendalian
func (c Controller) PengkomunikasianPengendalianCreateHandler(u usecase.PengkomunikasianPengendalianCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/pengkomunikasian-pengendalian",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "pengkomunikasian-pengendalian",
			Relation:  "create",
		},
		Body:    usecase.PengkomunikasianPengendalianCreateUseCaseReq{},
		Summary: "Create a new Pengkomunikasian Pengendalian",
		Tag:     "Pengkomunikasian Pengendalian",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.PengkomunikasianPengendalianCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
