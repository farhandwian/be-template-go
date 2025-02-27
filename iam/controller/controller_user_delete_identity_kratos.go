package controller

import (
	"iam/usecase"
	"net/http"
	"shared/helper"
)

// UserDeleteKratosHandler handles deleting a DaftarRisikoPrioritas
func (c Controller) UserDeleteKratosHandler(u usecase.UserDeleteKratos) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodDelete,
		Url:     "/api/kratos/user/{id}",
		Summary: "Delete an identity user Kratos",
		Tag:     "IAM - User Management",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.UserDeleteKratosReq{ID: id}
		HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
