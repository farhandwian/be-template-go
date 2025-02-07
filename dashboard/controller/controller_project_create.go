package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) ProjectCreateHandler(u usecase.ProjectCreateUseCase) helper.APIData {

	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/dashboard/projects",
		Access:  iammodel.MASTER_DATA_DAFTAR_PROYEK_CREATE,
		Body:    usecase.ProjectCreateUseCaseReq{},
		Summary: "Create a new Project",
		Tag:     "Master Data",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.ProjectCreateUseCaseReq](w, r)
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
