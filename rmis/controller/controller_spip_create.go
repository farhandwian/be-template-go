package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// SpipCreateHandler handles the creation of a new Spip
func (c Controller) SpipCreateHandler(u usecase.SpipCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/spips",
		AccessTest: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "spips",
			Relation:  "create",
		},
		Body:    usecase.SpipCreateUseCaseReq{},
		Summary: "Create a new SPIP",
		Tag:     "Sub Unsur SPIP",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.SpipCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
