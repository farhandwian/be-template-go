package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
	"strconv"
	"time"
)

// GetChartPintuAir

func (c Controller) GetChartPintuAirHandler(u usecase.ChartPintuAirUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/dashboard/chart-pintu-air/{id}",
		Summary: "Get Chart Pintu Air",
		Tag:     "Dashboard - Chart",
		QueryParams: []helper.QueryParam{
			{Name: "min", Type: "string", Description: "format 2024-11-25 02:00:00", Required: false},
			{Name: "max", Type: "string", Description: "format 2024-11-25 02:10:10", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		idInStr := r.PathValue("id")

		id, err := strconv.Atoi(idInStr)
		if err != nil {
			controller.Fail(w, err)
			return
		}

		min := controller.GetQueryString(r, "min", "0")
		max := controller.GetQueryString(r, "max", "0")

		minTime, err := time.Parse("2006-01-02 15:04:05", min)
		if err != nil {
			minTime = time.Now().Add(-12 * time.Hour)
		}

		maxTime, err := time.Parse("2006-01-02 15:04:05", max)
		if err != nil {
			maxTime = time.Now()
		}

		req := usecase.ChartPintuAirReq{
			WaterChannelDoorID: id,
			// MinTime:            time.Now().Add(-24 * time.Hour),
			// MaxTime:            time.Now(),
			MinTime: minTime,
			MaxTime: maxTime,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData

}
