package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenetapanKonteksRisikoOperasionalHandler handles getting a RekapitulasiHasilKuesioner by ID
func (c Controller) PenetapanKonteksRisikoOperasionalGetOneHandler(u usecase.PenetapanKonteksRisikoOperasionalGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/penetapan-konteks-risiko-operasional/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-operasional",
			Relation:  "read",
		},
		Summary: "Get a Penetapan Konteks Risiko Operasional by ID",
		Tag:     "Penetapan Konteks Risiko Operasional",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.PenetapanKonteksRisikoOperasionalGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
