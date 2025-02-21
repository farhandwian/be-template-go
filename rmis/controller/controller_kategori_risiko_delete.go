// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// KategoriRisikoDeleteHandler handles deleting a KategoriRisiko
func (c Controller) KategoriRisikoDeleteHandler(u usecase.KategoriRisikoDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/kategori-risikos/{id}",
		AccessTest: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "kategori-risikos",
			Relation:  "delete",
		},
		Summary: "Delete a KategoriRisiko",
		Tag:     "Kategori Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.KategoriRisikoDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
