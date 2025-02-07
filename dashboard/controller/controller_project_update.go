package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) ProjectUpdateHandler(u usecase.ProjectUpdateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPut,
		Url:     "/dashboard/projects/{id}",
		Access:  iammodel.MASTER_DATA_DAFTAR_PROYEK_UPDATE,
		Summary: "Update an Project",
		Tag:     "Master Data",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		req, ok := controller.ParseJSON[usecase.ProjectUpdateUseCaseReq](w, r)
		if !ok {
			return
		}
		req.ID = id

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
