package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) DeleteJDIH(u usecase.DeleteJDIHUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodDelete,
		Url:     "/dashboard/jdih/{id}",
		Access:  iammodel.MASTER_DATA_DAFTAR_JDIH_DELETE,
		Summary: "Delete JDIH",
		Tag:     "Master Data",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		req := usecase.DeleteJDIHReq{
			ID: id,
		}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
