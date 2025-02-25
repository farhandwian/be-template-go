package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// KriteriaDampakGetByIDHandler handles getting a KriteriaDampak by ID
func (c Controller) KriteriaDampakGetByIDHandler(u usecase.KriteriaDampakGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/kriteria-dampak/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "kriteria-dampak",
			Relation:  "read",
		},
		Summary: "Get a Kriteria Dampak by ID",
		Tag:     "Kriteria Dampak",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.KriteriaDampakGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
