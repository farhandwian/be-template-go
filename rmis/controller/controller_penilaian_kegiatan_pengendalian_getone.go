package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenilaianKegiatanPengendalianGetByIDHandler handles getting a PenilaianKegiatanPengendalian by ID
func (c Controller) PenilaianKegiatanPengendalianGetByIDHandler(u usecase.PenilaianKegiatanPengendalianGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/penilaian-kegiatan-pengendalians/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penilaian-kegiatan-pengendalians",
			Relation:  "read",
		},
		Summary: "Get a Penilaian Kegiatan Pengendalian by ID",
		Tag:     "Penilaian Kegiatan Pengendalian",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.PenilaianKegiatanPengendalianGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
