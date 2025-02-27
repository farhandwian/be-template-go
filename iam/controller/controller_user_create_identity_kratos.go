package controller

import (
	"iam/gateway"
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
)

func (c Controller) UserCreateKratosHandler(u usecase.UserCreateKratos) helper.APIData {

	type Body struct {
		Email         string `json:"email"`
		Password      string `json:"password"`
		Nama          string `json:"nama"`
		NoTelepon     string `json:"no_telepon"`
		Jabatan       string `json:"jabatan"`
		AksesPengguna string `json:"akses_pengguna"`
		JenisKelamin  string `json:"jenis_kelamin"`
	}

	apiData := helper.APIData{
		Access:  model.MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_UPDATE,
		Method:  http.MethodPost,
		Url:     "/api/kratos/user",
		Body:    Body{},
		Summary: "Create identity user Kratos",
		Tag:     "IAM - User Management",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := gateway.UserCreateKratosReq{
			User: model.UserKratosCreate{
				Nama:          body.Nama,
				NoTelepon:     body.NoTelepon,
				Email:         body.Email,
				Password:      body.Password,
				Jabatan:       body.Jabatan,
				AksesPengguna: body.AksesPengguna,
				JenisKelamin:  body.JenisKelamin,
			},
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	// authorizationHandler := Authorization(handler, apiData.Access)
	// authenticatedHandler := Authentication(authorizationHandler, c.JWT)
	// c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
