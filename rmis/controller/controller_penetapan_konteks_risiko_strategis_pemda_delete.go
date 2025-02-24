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
		Url:    "/api/penetapan-konteks-risiko-strategis-pemdas/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-strategis-pemdas",
			Relation:  "delete",
		},
		Summary: "Delete a Rekapitulasi hasil kuesioner",
		Tag:     "Rekapitulasi hasil kuesioner",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.PenetapanKonteksRisikoDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
