// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenetapanKonteksRisikoOperasionalDeleteHandler handles deleting a PenetapanKonteksRisikoOperasional
func (c Controller) PenetapanKonteksRisikoOperasionalDeleteHandler(u usecase.PenetapanKonteksRisikoOperasionalDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/penetapan-konteks-risiko-operasional/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-operasional",
			Relation:  "delete",
		},
		Summary: "Delete a Penetapan Konteks Risiko Operasional",
		Tag:     "Penetapan Konteks Risiko Operasional",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.PenetapanKonteksRisikoOperasionalDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
