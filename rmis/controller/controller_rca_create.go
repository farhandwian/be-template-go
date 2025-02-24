package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// RcaCreateHandler handles the creation of a new Rca
func (c Controller) RcaCreateHandler(u usecase.RcaCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/rcas",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "rcas",
			Relation:  "create",
		},
		Body:    usecase.RcaCreateUseCaseReq{},
		Summary: "Create a new Rca",
		Tag:     "Rca",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.RcaCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
