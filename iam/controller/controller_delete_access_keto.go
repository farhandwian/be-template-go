package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) DeleteAccessHandler(u usecase.DeleteAccessKetoUseCase) helper.APIData {

	type Body struct {
		Namespace string `json:"namespace"`
		SubjectID string `json:"subject_id"`
		Object    string `json:"object"`
		Relation  string `json:"relation"`
	}

	apiData := helper.APIData{
		Access:   model.ANONYMOUS,
		Method:   http.MethodDelete,
		Url:      "/auth/delete-access-keto",
		Body:     usecase.DeleteAccessKetoReq{},
		Summary:  "Check access with ory keto",
		Tag:      "IAM - Authentication",
		Examples: []helper.ExampleResponse{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.DeleteAccessKetoReq{
			Namespace: body.Namespace,
			SubjectID: body.SubjectID,
			Object:    body.Object,
			Relation:  body.Relation,
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
