package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// RancanganPemantauanUpdateHandler handles the creation of a new RancanganPemantauan
func (c Controller) RancanganPemantauanUpdateHandler(u usecase.RancanganPemantauanUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPut,
		Url:    "/api/rancangan-pemantauan/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "rancangan-pemantauan",
			Relation:  "update",
		},
		Body:    usecase.RancanganPemantauanUpdateUseCaseReq{},
		Summary: "Update a Rancangan Pemantauan",
		Tag:     "Rancangan Pemantauan",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.RancanganPemantauanUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
