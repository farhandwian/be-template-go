package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
	"strconv"
)

func (c Controller) DetailWaterChannelDoorHandler(u usecase.DetailPintuAirUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/dashboard/water-channel-doors/{id}",
		Summary: "Get detail of water channel door",
		Tag:     "Dashboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			controller.Fail(w, err)
			return
		}

		req := usecase.DetailPintuAirReq{
			WaterChannelDoorID: id,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData

}
