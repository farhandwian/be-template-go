package controller

import (
	"dashboard/model"
	"dashboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"time"
)

func (c Controller) CreateJDIH(u usecase.CreateJDIHUseCase) helper.APIData {

	type Body struct {
		Title       string           `json:"title"`
		PublishedAt string           `json:"published_at"`
		Status      model.JDIHStatus `json:"status"`
	}

	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/dashboard/jdih",
		Access:  iammodel.MASTER_DATA_DAFTAR_JDIH_CREATE,
		Summary: "Create JDIH",
		Tag:     "Master Data",
		Body:    Body{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		payload, ok := controller.ParseJSON[Body](w, r)
		if !ok {
			return
		}

		publishedAt, _ := time.Parse(time.DateOnly, payload.PublishedAt)
		req := usecase.CreateJDIHReq{
			Title:       payload.Title,
			PublishedAt: publishedAt,
			Status:      payload.Status,
			Now:         time.Now(),
		}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
