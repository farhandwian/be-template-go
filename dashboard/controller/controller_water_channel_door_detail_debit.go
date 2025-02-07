package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
	"strconv"
)

func (c Controller) GetStatusDebitHandler(u usecase.GetStatusDebitUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/dashboard/status-debit/{id}",
		Summary: "Get status debit",
		Tag:     "Dashboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			controller.Fail(w, err)
			return
		}

		req := usecase.GetStatusDebitReq{
			WaterChannelDoorID: id,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData

}
