package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// KategoriRisikoUpdateHandler handles the creation of a new KategoriRisiko
func (c Controller) KategoriRisikoUpdateHandler(u usecase.KategoriRisikoUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/kategori-risikos/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "kategori-risikos",
			Relation:  "update",
		},
		Body:    usecase.KategoriRisikoUpdateUseCaseReq{},
		Summary: "Update a Kategori Risiko",
		Tag:     "Kategori Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.KategoriRisikoUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
