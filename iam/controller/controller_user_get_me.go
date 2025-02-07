package controller

import (
	"iam/gateway"
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/core"
	"shared/helper"
)

func (c Controller) UserGetMeHandler(u usecase.UserGetOne) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/users/me",
		Summary: "Get current user detail",
		Tag:     "IAM - User Management",
		Examples: []helper.ExampleResponse{
			{
				StatusCode: 200,
				Content: map[string]interface{}{
					"user": map[string]interface{}{
						"id":             "c5ec2448-df2a-4c99-8436-1cd67be771a0",
						"name":           "admin",
						"phone_number":   "08123456789",
						"email":          "admin@mail.com",
						"email_verified": "2024-10-02T17:05:16.402+07:00",
						"enabled":        true,
						"user_access":    "3",
						"created_at":     "2024-10-02T17:05:16.406+07:00",
						"updated_at":     "2024-10-02T17:05:17.113+07:00",
					},
				},
			},
			// {
			// 	StatusCode: 404,
			// 	Content: map[string]interface{}{
			// 		"error": "User not found",
			// 	},
			// },
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		userID := core.GetDataFromContext[model.UserID](r.Context(), UserIDContext)

		req := usecase.UserGetOneReq{
			UserGetOneByIDReq: gateway.UserGetOneByIDReq{
				UserID: userID,
			},
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authenticatedHandler := Authentication(handler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
