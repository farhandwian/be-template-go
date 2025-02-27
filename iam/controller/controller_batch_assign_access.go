package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) BatchAssignAccessHandler(u usecase.BatchAssignAccessUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:   model.ANONYMOUS,
		Method:   http.MethodPost,
		Url:      "/auth/batch-assign-access",
		Body:     usecase.BatchAssignAccessReq{},
		Summary:  "Batch Assign with ory keto",
		Tag:      "IAM - Authentication",
		Examples: []helper.ExampleResponse{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		body, ok := ParseJSON[usecase.BatchAssignAccessReq](w, r)
		if !ok {
			return
		}

		req := usecase.BatchAssignAccessReq{
			Namespace:  body.Namespace,
			SubjectID:  body.SubjectID,
			Objects:    body.Objects,
			Operations: body.Operations,
			SubjectSet: body.SubjectSet,
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
