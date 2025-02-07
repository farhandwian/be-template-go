package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
	"time"
)

func (c Controller) PasswordResetSubmitHandler(u usecase.PasswordResetSubmit) helper.APIData {

	type Body struct {
		PasswordResetToken string `json:"password_reset_token"`
		NewPassword        string `json:"new_password"`
	}

	apiData := helper.APIData{
		Access:  model.ANONYMOUS,
		Method:  http.MethodPost,
		Url:     "/password/reset/verify",
		Body:    Body{},
		Summary: "Submit password reset",
		Tag:     "IAM - Password Management",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.PasswordResetSubmitReq{
			PasswordResetToken: body.PasswordResetToken,
			NewPassword:        body.NewPassword,
			Now:                time.Now(),
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
