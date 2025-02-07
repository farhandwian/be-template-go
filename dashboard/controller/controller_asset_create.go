package controller

import (
	"dashboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
)

// AssetCreateHandler handles the creation of a new Asset
func (c Controller) AssetCreateHandler(u usecase.AssetCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/dashboard/assets",
		Access:  model.MASTER_DATA_DAFTAR_ASET_CREATE,
		Body:    usecase.AssetCreateUseCaseReq{},
		Summary: "Create a new Asset",
		Tag:     "Master Data",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.AssetCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
