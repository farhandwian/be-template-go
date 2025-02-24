package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// SimpulanKondisiKelemahanLingkunganUpdateHandler handles the creation of a new SimpulanKondisiKelemahanLingkungan
func (c Controller) SimpulanKondisiKelemahanLingkunganUpdateHandler(u usecase.SimpulanKondisiKelemahanLingkunganUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/simpulan-kondisi-kelemahan-lingkungans/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "simpulan-kondisi-kelemahan-lingkungans",
			Relation:  "update",
		},
		Body:    usecase.SimpulanKondisiKelemahanLingkunganUpdateUseCaseReq{},
		Summary: "Update a SimpulanKondisiKelemahanLingkungan",
		Tag:     "Sub Unsur SimpulanKondisiKelemahanLingkungan",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.SimpulanKondisiKelemahanLingkunganUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
