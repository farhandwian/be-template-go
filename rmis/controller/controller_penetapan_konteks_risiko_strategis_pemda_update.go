package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenetapanKonteksRisikoStrategisPemdaUpdateHandler handles the creation of a new PenetapanKonteksRisikoStrategisPemda
func (c Controller) PenetapanKonteksRisikoStrategisPemdaUpdateHandler(u usecase.PenetapanKonteksRisikoStrategisPemdaUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/penetapan-konteks-risiko-strategis-pemdas/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-strategis-pemdas",
			Relation:  "update",
		},
		Body:    usecase.PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseReq{},
		Summary: "Update a PenetapanKonteksRisikoStrategisPemda",
		Tag:     "Sub Unsur PenetapanKonteksRisikoStrategisPemda",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
