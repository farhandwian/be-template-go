package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) CheckAccessHandler(u usecase.CheckAccess) helper.APIData {

	type Body struct {
		FunctionName string `json:"function_name"`
	}

	apiData := helper.APIData{
		Access:  model.SERVICE_OPERATION,
		Method:  http.MethodPost,
		Url:     "/access",
		Body:    Body{},
		Summary: "Simple Access Checking",
		Tag:     "IAM - Access",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		bearerToken, errMsg, ok := GetBearerToken(w, r)
		if !ok {
			writeJSON(w, http.StatusUnauthorized, Response{Status: "failed", Error: &errMsg})
			return
		}

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.CheckAccessReq{
			AccessToken:  bearerToken,
			FunctionName: body.FunctionName,
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
