package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) ProjectDeleteHandler(u usecase.ProjectDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodDelete,
		Url:     "/dashboard/projects/{id}",
		Access:  iammodel.MASTER_DATA_DAFTAR_PROYEK_DELETE,
		Summary: "Delete an Project",
		Tag:     "Master Data",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.ProjectDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
