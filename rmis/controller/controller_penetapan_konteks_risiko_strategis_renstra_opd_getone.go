package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenetapanKonteksRisikoStrategisRenstraOPDHandler handles getting a RekapitulasiHasilKuesioner by ID
func (c Controller) PenetapanKonteksRisikoStrategisRenstraOPDGetOneHandler(u usecase.PenetapanKonteksRisikoRenstraOPDGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/penetapan-konteks-risiko-strategis-renstra-opd/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-strategis-renstra-opd",
			Relation:  "read",
		},
		Summary: "Get a Penetapan Konteks Risiko Strategis Renstra OPD by ID",
		Tag:     "Penetapan Konteks Risiko Strategis Renstra OPD",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.PenetapanKonteksRisikoRenstraOPDGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
