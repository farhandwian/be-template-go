package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// KategoriRisikoGetByIDHandler handles getting a KategoriRisiko by ID
func (c Controller) KategoriRisikoGetByIDHandler(u usecase.KategoriRisikoGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/kategori-risikos/{id}",
		AccessTest: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "kategori-risikos",
			Relation:  "read",
		},
		Summary: "Get a Sub Unsur Kategori Risiko by ID",
		Tag:     "Sub Unsur Kategori Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.KategoriRisikoGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
