package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) GetAccessRoleKetoHandler(u usecase.GetAccessRoleKetoUseCase) helper.APIData {

	apiData := helper.APIData{
		AccessTest: model.AccessKeto{
			Namespace: "app",
			Object:    "none",
		},
		Method:   http.MethodGet,
		Url:      "/auth/access-role-keto",
		Body:     usecase.GetAccessRoleKetoReq{},
		Summary:  "Get Access list user",
		Tag:      "IAM - Authentication",
		Examples: []helper.ExampleResponse{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.GetAccessRoleKetoReq{
			Namespace: apiData.AccessTest.Namespace,
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
