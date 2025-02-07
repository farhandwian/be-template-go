package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"time"
)

func (c Controller) EmployeeCreateHandler(u usecase.EmployeeCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/dashboard/employees",
		Access:  iammodel.MASTER_DATA_DAFTAR_KEPEGAWAIAN_CREATE,
		Body:    usecase.EmployeeCreateUseCaseReq{},
		Summary: "Create a new Employee",
		Tag:     "Master Data",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		request, ok := controller.ParseJSON[usecase.EmployeeCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		request.Now = time.Now()

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
