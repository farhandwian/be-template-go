// File: controller/controller_Asset.go

package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

// AssetDeleteHandler handles deleting a Asset
func (c Controller) AssetDeleteHandler(u usecase.AssetDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodDelete,
		Url:     "/dashboard/assets/{id}",
		Access:  iammodel.MASTER_DATA_DAFTAR_ASET_DELETE,
		Summary: "Delete a Asset",
		Tag:     "Master Data",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.AssetDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
