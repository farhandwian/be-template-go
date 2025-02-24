package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

func (c Controller) SimpulanKondisiKelemahanLingkunganCreateHandler(u usecase.SimpulanKondisiKelemahanLingkunganCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/simpulan-kondisi-kelemahan-lingkungan",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "simpulan-kondisi-kelemahan-lingkungan",
			Relation:  "create",
		},
		Body:    usecase.SimpulanKondisiKelemahanLingkunganCreateUseCaseReq{},
		Summary: "Create a new SimpulanKondisiKelemahanLingkungan",
		Tag:     "Sub Unsur SimpulanKondisiKelemahanLingkungan",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.SimpulanKondisiKelemahanLingkunganCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
