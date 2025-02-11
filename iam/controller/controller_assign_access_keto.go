package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) AssignAccessKetoHandler(u usecase.AssignAccessKetoUseCase) helper.APIData {

	type Body struct {
		Namespace string `json:"namespace"`
		SubjectID string `json:"subject_id"`
		Object    string `json:"object"`
		Relation  string `json:"relation"`
	}

	apiData := helper.APIData{
		Access:   model.ANONYMOUS,
		Method:   http.MethodPost,
		Url:      "/auth/assign-access-keto",
		Body:     usecase.AssignAccessKetoReq{},
		Summary:  "Check access with ory keto",
		Tag:      "IAM - Authentication",
		Examples: []helper.ExampleResponse{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.AssignAccessKetoReq{
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
