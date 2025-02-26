package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// RancanganPemantauanGetByIDHandler handles getting a RancanganPemantauan by ID
func (c Controller) RancanganPemantauanGetByIDHandler(u usecase.RancanganPemantauanGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/penyebab-risikos/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penyebab-risikos",
			Relation:  "read",
		},
		Summary: "Get a Penyebab Risiko by ID",
		Tag:     "Penyebab Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.RancanganPemantauanGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
