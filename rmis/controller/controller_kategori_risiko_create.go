package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// KategoriRisikoCreateHandler handles the creation of a new KategoriRisiko
func (c Controller) KategoriRisikoCreateHandler(u usecase.KategoriRisikoCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/kategori-risikos",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "kategori-risikos",
			Relation:  "create",
		},
		Body:    usecase.KategoriRisikoCreateUseCaseReq{},
		Summary: "Create a new KategoriRisiko",
		Tag:     "Sub Unsur KategoriRisiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.KategoriRisikoCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
