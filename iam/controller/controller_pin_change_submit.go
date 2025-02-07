package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/core"
	"shared/helper"
	"time"
)

func (c Controller) PinChangeSubmitHandler(u usecase.PinChangeSubmit) helper.APIData {

	type Body struct {
		OTP    string `json:"otp"`
		NewPIN string `json:"new_pin"`
	}

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodPost,
		Url:     "/pin/change/verify",
		Body:    Body{},
		Summary: "Submit pin changes",
		Tag:     "IAM - Pin Management",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		userID := core.GetDataFromContext[model.UserID](r.Context(), UserIDContext)

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.PinChangeSubmitReq{
			UserID:   userID,
			OTPValue: body.OTP,
			NewPIN:   body.NewPIN,
			Now:      time.Now(),
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authenticatedHandler := Authentication(handler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}
