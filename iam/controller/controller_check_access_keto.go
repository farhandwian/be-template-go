package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) CheckAccessKetoHandler(u usecase.CheckAccessKetoUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:   model.ANONYMOUS,
		Method:   http.MethodGet,
		Url:      "/auth/check-access-keto",
		Body:     usecase.CheckAccessKetoReq{},
		Summary:  "Check access with ory keto",
		Tag:      "IAM - Authentication",
		Examples: []helper.ExampleResponse{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		// namespace := r.URL.Query().Get("namespace")
		// subjectID := r.URL.Query().Get("subject_id")
		// object := r.URL.Query().Get("object")
		// relation := r.URL.Query().Get("relation")

		namespace := "app"
		object := "testing"
		relation := "owner"
		subjectID := "user:123"
		req := usecase.CheckAccessKetoReq{
			Namespace: namespace,
			SubjectID: subjectID,
			Object:    object,
			Relation:  relation,
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
