// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PengkomunikasianPengendalianDeleteHandler handles deleting a PengkomunikasianPengendalian
func (c Controller) PengkomunikasianPengendalianDeleteHandler(u usecase.PengkomunikasianPengendalianDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/pengkomunikasian-pengendalian/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "pengkomunikasian-pengendalian",
			Relation:  "delete",
		},
		Summary: "Delete a Pengkomunikasian Pengendalian",
		Tag:     "Pengkomunikasian Pengendalian",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.PengkomunikasianPengendalianDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
