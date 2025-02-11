package controller

import (
	"iam/model"
	"iam/usecase"
	"net/http"
	"shared/helper"
	"time"
)

func (c Controller) RegisterUserHandler(u usecase.RegisterUser) helper.APIData {

	type Body struct {
		Name        string            `json:"name"`
		Email       model.Email       `json:"email"`
		PhoneNumber model.PhoneNumber `json:"phone_number"`
	}

	apiData := helper.APIData{
		Access:  model.MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_CREATE,
		Method:  http.MethodPost,
		Url:     "/account/register",
		Body:    Body{},
		Summary: "Register user",
		Tag:     "IAM - Account Management",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		body, ok := ParseJSON[Body](w, r)
		if !ok {
			return
		}

		req := usecase.RegisterUserReq{
			Now:         time.Now(),
			Name:        body.Name,
			Email:       body.Email,
			PhoneNumber: body.PhoneNumber,
		}

		HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := Authorization(handler, apiData.Access)
	authenticatedHandler := Authentication(authorizationHandler, c.JWT)
	c.Mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)

	return apiData
}

// package controller

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// var kratosPublicURL = "http://localhost:4433"

// // User registration request structure
// type RegisterRequest struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// // RegisterUserHandler - Handles direct API-based registration
// func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
// 	// Parse request body
// 	var req RegisterRequest
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		http.Error(w, "Invalid request format", http.StatusBadRequest)
// 		return
// 	}

// 	// Prepare registration payload for Kratos
// 	kratosRequestBody := map[string]interface{}{
// 		"method":   "password",
// 		"password": req.Password,
// 		"traits": map[string]string{
// 			"email": req.Email,
// 		},
// 	}
// 	jsonBody, _ := json.Marshal(kratosRequestBody)

// 	// Send registration request to Kratos API
// 	resp, err := http.Post(fmt.Sprintf("%s/self-service/registration?return_session_token=true", kratosPublicURL), "application/json", bytes.NewBuffer(jsonBody))
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error connecting to Kratos: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Read Kratos response
// 	body, _ := ioutil.ReadAll(resp.Body)

// 	// Forward Kratos response to client
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(resp.StatusCode)
// 	w.Write(body)
// }
