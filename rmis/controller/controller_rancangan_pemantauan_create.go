package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// RancanganPemantauanCreateHandler handles the creation of a new RancanganPemantauan
func (c Controller) RancanganPemantauanCreateHandler(u usecase.RancanganPemantauanCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/rancangan-pemantauan",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "rancangan-pemantauan",
			Relation:  "create",
		},
		Body:    usecase.RancanganPemantauanCreateUseCaseReq{},
		Summary: "Create a new Rancangan Pemantauan",
		Tag:     "Rancangan Pemantauan",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.RancanganPemantauanCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
