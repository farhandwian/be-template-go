package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) EmployeeUpdateHandler(u usecase.EmployeeUpdateUseCase) helper.APIData {

	apiData := helper.APIData{
		Method:  http.MethodPut,
		Url:     "/dashboard/employees/{id}",
		Access:  iammodel.MASTER_DATA_DAFTAR_KEPEGAWAIAN_UPDATE,
		Body:    usecase.EmployeeUpdateUseCaseReq{},
		Summary: "Update an Employee",
		Tag:     "Master Data",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		req, ok := controller.ParseJSON[usecase.EmployeeUpdateUseCaseReq](w, r)
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
