// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// SimpulanKondisiKelemahanLingkunganDeleteHandler handles deleting a SimpulanKondisiKelemahanLingkungan
func (c Controller) SimpulanKondisiKelemahanLingkunganDeleteHandler(u usecase.SimpulanKondisiKelemahanLingkunganDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/simpulan-kondisi-kelemahan-lingkungan/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "simpulan-kondisi-kelemahan-lingkungan",
			Relation:  "delete",
		},
		Summary: "Delete a SimpulanKondisiKelemahanLingkungan",
		Tag:     "Simpulan Kondisi Kelemahan Lingkungan",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.SimpulanKondisiKelemahanLingkunganDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
