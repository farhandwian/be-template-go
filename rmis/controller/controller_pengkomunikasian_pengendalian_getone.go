package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PengkomunikasianPengendalianGetByIDHandler handles getting a PengkomunikasianPengendalian by ID
func (c Controller) PengkomunikasianPengendalianGetByIDHandler(u usecase.PengkomunikasianPengendalianGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/pengkomunikasian-pengendalians/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "pengkomunikasian-pengendalians",
			Relation:  "read",
		},
		Summary: "Get a Pengkomunikasian Pengendalian by ID",
		Tag:     "Pengkomunikasian Pengendalian",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.PengkomunikasianPengendalianGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
