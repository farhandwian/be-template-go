package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PencatatanKejadianRisikoGetByIDHandler handles getting a PencatatanKejadianRisiko by ID
func (c Controller) PencatatanKejadianRisikoGetByIDHandler(u usecase.PencatatanKejadianRisikoGetByIDUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodGet,
		Url:    "/api/pencatatan-kejadian-resiko/{id}",
		AccessKeto: iammodel.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "pencatatan-kejadian-resiko",
			Relation:  "read",
		},
		Summary: "Get a Pencatatan Kejadian Risiko by ID",
		Tag:     "Pencatatan Kejadian Risiko",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.PencatatanKejadianRisikoGetByIDUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
