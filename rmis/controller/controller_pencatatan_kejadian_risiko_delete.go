// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PencatatanKejadianRisikoDeleteHandler handles deleting a PencatatanKejadianRisiko
func (c Controller) PencatatanKejadianRisikoDeleteHandler(u usecase.PencatatanKejadianRisikoDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/pencatatan-kejadian-risiko/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "pencatatan-kejadian-risiko",
			Relation:  "delete",
		},
		Summary: "Delete a Pencatatan Kejadian Risiko",
		Tag:     "Pencatatan Kejadian Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.PencatatanKejadianRisikoDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
