package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenilaianKegiatanPengendalianUpdateHandler handles the creation of a new PenilaianKegiatanPengendalian
func (c Controller) PenilaianKegiatanPengendalianUpdateHandler(u usecase.PenilaianKegiatanPengendalianUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/penilaian-kegiatan-pengendalians/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penilaian-kegiatan-pengendalians",
			Relation:  "update",
		},
		Body:    usecase.PenilaianKegiatanPengendalianUpdateUseCaseReq{},
		Summary: "Update a Penilaian Kegiatan Pengendalian",
		Tag:     "Penilaian Kegiatan Pengendalian",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.PenilaianKegiatanPengendalianUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
