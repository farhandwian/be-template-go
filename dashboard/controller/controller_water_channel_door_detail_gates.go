package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
	"strconv"
)

func (c Controller) GetGateStatusHandler(u usecase.GetGateStatusUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/status-gates/{id}",
		Access:  iammodel.PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_PINTU_AIR_READ,
		Summary: "Gates Status",
		Tag:     "Dashboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			controller.Fail(w, err)
			return
		}

		req := usecase.GetGateStatusReq{WaterChannelDoorID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
