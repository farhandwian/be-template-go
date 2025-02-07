package controller

import (
	"iam/gateway"
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) UserGetAccessHandler(u usecase.UserGetAccess) helper.APIData {

	apiData := helper.APIData{
		Access:  model.MANAJEMEN_PENGGUNA_HAK_AKSES_UPDATE,
		Method:  http.MethodGet,
		Url:     "/users/{id}/access",
		Summary: "Get user access by id",
		Tag:     "IAM - User Access",
		Examples: []helper.ExampleResponse{
			{
				StatusCode: 200,
				Content: map[string]interface{}{
					"accesses": []map[string]interface{}{
						{
							"Description": "function 1",
							"Enabled":     false,
							"ID":          1,
						},
						{
							"Description": "function 2",
							"Enabled":     true,
							"ID":          2,
						},
						{
							"Description": "function 3",
							"Enabled":     false,
							"ID":          3,
						},
						{
							"Description": "function 4",
							"Enabled":     false,
							"ID":          4,
						},
						{
							"Description": "function 5",
							"Enabled":     false,
							"ID":          5,
						},
					},
				},
			},
			// {
			// 	StatusCode: 403,
			// 	Content: map[string]interface{}{
			// 		"error": "Insufficient permissions to view access functions",
			// 	},
			// },
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		userID := r.PathValue("id")

		req := usecase.UserGetAccessReq{
			UserGetOneByIDReq: gateway.UserGetOneByIDReq{
				UserID: model.UserID(userID),
			},
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := Authorization(handler, apiData.Access)
	authenticatedHandler := Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
