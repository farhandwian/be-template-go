package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// SpipGetByIDHandler handles getting a Spip by ID
func (c Controller) SpipGetByIDHandler(u usecase.SpipGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/spips/{id}",
		AccessTest: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "spips",
			Relation:  "read",
		},
		Summary: "Get a Sub Unsur Spip by ID",
		Tag:     "Sub Unsur SPIP",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.SpipGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
