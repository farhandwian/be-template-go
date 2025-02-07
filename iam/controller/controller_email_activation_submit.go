package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
	"time"
)

func (c Controller) EmailActivationSubmitHandler(u usecase.EmailActivationSubmit) helper.APIData {

	type Body struct {
		ActivationToken string `json:"activation_token"`
		Password        string `json:"password"`
		Pin             string `json:"pin"`
	}

	apiData := helper.APIData{
		Access:  model.ANONYMOUS,
		Method:  http.MethodPost,
		Url:     "/account/activate/verify",
		Body:    Body{},
		Summary: "Submit email activation",
		Tag:     "IAM - Account Management",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.EmailActivationSubmitReq{
			ActivationToken: body.ActivationToken,
			Password:        body.Password,
			Pin:             body.Pin,
			Now:             time.Now(),
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
