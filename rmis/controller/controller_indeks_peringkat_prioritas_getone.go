package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IndeksPeringkatPrioritasGetByIDHandler handles getting a IndeksPeringkatPrioritas by ID
func (c Controller) IndeksPeringkatPrioritasGetByIDHandler(u usecase.IndeksPeringkatPrioritasGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/indeks-peringkat-prioritas/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "indeks-peringkat-prioritas",
			Relation:  "read",
		},
		Summary: "Get a Indeks Peringkat Prioritas by ID",
		Tag:     "Indeks Peringkat Prioritas",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.IndeksPeringkatPrioritasGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
