// File: controller/controller_Asset.go

package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenetapanKonteksRisikoStrategisPemdaDeleteHandler handles deleting a PenetapanKonteksRisikoStrategisPemda
func (c Controller) PenetapanKonteksRisikoStrategisPemdaDeleteHandler(u usecase.PenetapanKonteksRisikoDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/penetapan-konteks-risiko-strategis-pemda/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-strategis-pemda",
			Relation:  "delete",
		},
		Summary: "Delete a Penetapan Konteks Risiko Strategis Pemdar",
		Tag:     "Penetapan Konteks Risiko Strategis Pemda",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.PenetapanKonteksRisikoDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
