package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// IKUCreateHandler handles the creation of a new IKU
func (c Controller) IKUCreateHandler(u usecase.IKUCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/iku",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "iku",
			Relation:  "create",
		},
		Body:    usecase.IKUCreateUseCaseReq{},
		Summary: "Create a new IKU",
		Tag:     "IKU",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.IKUCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
