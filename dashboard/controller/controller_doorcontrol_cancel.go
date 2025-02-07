// File: controller/controller_Asset.go

package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/model"
	"time"
)

func (c Controller) DoorControlCancelHandler(u usecase.DoorControlCancel) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodDelete,
		Url:     "/dashboard/doorcontrol/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Cancel a door control",
		Tag:     "Master Data - Door Control",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.DoorControlCancelReq{
			ID:  model.DoorControlID(id),
			Now: time.Now().In(time.Local),
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
