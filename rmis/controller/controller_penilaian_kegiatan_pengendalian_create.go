package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenilaianKegiatanPengendalianCreateHandler handles the creation of a new PenilaianKegiatanPengendalian
func (c Controller) PenilaianKegiatanPengendalianCreateHandler(u usecase.PenilaianKegiatanPengendalianCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/penilaian-kegiatan-pengendalians",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penilaian-kegiatan-pengendalians",
			Relation:  "create",
		},
		Body:    usecase.PenilaianKegiatanPengendalianCreateUseCaseReq{},
		Summary: "Create a new Penilaian Kegiatan Pengendalian",
		Tag:     "Penilaian Kegiatan Pengendalian",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.PenilaianKegiatanPengendalianCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
