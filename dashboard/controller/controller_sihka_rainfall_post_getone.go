package controller

import (
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) RainfallPostCalculationGetDetailHandler(u usecase.RainfallDetailWithCalculationUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/dashboard/sihka/hydrology/rainfall-post/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Get detail rain fall posts",
		Tag:     "Sihka",
		QueryParams: []helper.QueryParam{
			{Name: "start_date", Type: "string", Description: "Filter date", Required: true},
			{Name: "end_date", Type: "string", Description: "Filter date", Required: true},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		startDate := controller.GetQueryString(r, "start_date", "")
		endDate := controller.GetQueryString(r, "end_date", "")

		req := usecase.DetailRainfallPostCalculationReq{
			StartDate: startDate,
			EndDate:   endDate,
			ID:        id,
		}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
