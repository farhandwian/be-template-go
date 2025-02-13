package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) GetAccessKetoHandler(u usecase.UserGetAccessKetoUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:   model.ANONYMOUS,
		Method:   http.MethodGet,
		Url:      "/auth/{id}/access-keto",
		Body:     usecase.UserGetAccessKetoReq{},
		Summary:  "Get Access list user",
		Tag:      "IAM - Authentication",
		Examples: []helper.ExampleResponse{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("id")
		namespace := "app"

		req := usecase.UserGetAccessKetoReq{
			Namespace: namespace,
			UserID:    userID,
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
