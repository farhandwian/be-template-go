package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/core"
	"shared/helper"
	"time"
)

func (c Controller) PasswordChangeSubmitHandler(u usecase.PasswordChangeSubmit) helper.APIData {

	type Body struct {
		OTP         string `json:"otp"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodPost,
		Url:     "/password/change/verify",
		Body:    Body{},
		Summary: "Submit OTP for password changes",
		Tag:     "IAM - Password Management",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		userID := core.GetDataFromContext[model.UserID](r.Context(), UserIDContext)

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.PasswordChangeSubmitReq{
			UserID:      userID,
			OTPValue:    body.OTP,
			OldPassword: body.OldPassword,
			NewPassword: body.NewPassword,
			Now:         time.Now(),
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authenticatedHandler := Authentication(handler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
