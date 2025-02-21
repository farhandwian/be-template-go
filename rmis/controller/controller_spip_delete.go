// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// SpipDeleteHandler handles deleting a Spip
func (c Controller) SpipDeleteHandler(u usecase.SpipDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/spips/{id}",
		AccessTest: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "spips",
			Relation:  "delete",
		},
		Summary: "Delete a Spip",
		Tag:     "Sub Unsur SPIP",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.SpipDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
