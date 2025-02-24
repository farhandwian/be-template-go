package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// SimpulanKondisiKelemahanLingkunganHandler handles getting a Simpulan SimpulanKondisiKelemahanLingkungan by ID
func (c Controller) SimpulanKondisiKelemahanLingkunganGetByIDHandler(u usecase.SimpulanKondisiKelemahanLingkunganGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/simpulan-kondisi-kelemahan-lingkungan/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "simpulan-kondisi-kelemahan-lingkungan",
			Relation:  "read",
		},
		Summary: "Get a Sub Unsur Simpulan SimpulanKondisiKelemahanLingkungan by ID",
		Tag:     "Sub Unsur Simpulan SimpulanKondisiKelemahanLingkungan",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.SimpulanKondisiKelemahanLingkunganGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
