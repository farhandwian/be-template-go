package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

// AssetUpdateHandler handles updating a Asset
func (c Controller) AssetUpdateHandler(u usecase.AssetUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPut,
		Url:     "/dashboard/assets/{id}",
		Access:  iammodel.MASTER_DATA_DAFTAR_ASET_UPDATE,
		Body:    usecase.AssetUpdateUseCaseReq{},
		Summary: "Update a Asset",
		Tag:     "Master Data",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		request, ok := controller.ParseJSON[usecase.AssetUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		request.ID = id
		controller.HandleUsecase(r.Context(), w, u, request)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
