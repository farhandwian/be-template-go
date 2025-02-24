package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// RekapitulasiHasilKuesionerCreateHandler handles the creation of a new RekapitulasiHasilKuesioner
func (c Controller) RekapitulasiHasilKuesionerCreateHandler(u usecase.RekapitulasiHasilKuesionerCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/rekapitulasi-hasil-kuesioners",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "rekapitulasi-hasil-kuesioners",
			Relation:  "create",
		},
		Body:    usecase.RekapitulasiHasilKuesionerCreateUseCaseReq{},
		Summary: "Create a new RekapitulasiHasilKuesioner",
		Tag:     "Sub Unsur RekapitulasiHasilKuesioner",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.RekapitulasiHasilKuesionerCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
