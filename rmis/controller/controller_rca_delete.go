// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// RcaDeleteHandler handles deleting a Rca
func (c Controller) RcaDeleteHandler(u usecase.RcaDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/rcas/{id}",
		AccessTest: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "rcas",
			Relation:  "delete",
		},
		Summary: "Delete a Root Cause Analysis (RCA)",
		Tag:     "Root Cause Analysis (RCA)",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.RcaDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
