package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"iam/model"
	"io/ioutil"
	"net/http"
	"shared/helper"
	oryConfig "shared/helper/ory"

	ory "github.com/ory/client-go"
)

func (c Controller) RegisterUserHandler(s oryConfig.ORYServer) helper.APIData {

	type Body struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
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

		// Start Kratos Registration Flow
		flowURL := fmt.Sprintf("%s/self-service/registration/flows", s.KratosPublicEndpoint)
		resp, err := http.Get(flowURL)
		if err != nil || resp.StatusCode != http.StatusOK {
			http.Error(w, "Failed to initiate registration flow", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var flow ory.RegistrationFlow
		json.NewDecoder(resp.Body).Decode(&flow)

		// Prepare registration payload
		registerData := map[string]string{
			"traits.email": body.Email,
			"password":     body.Password,
			"method":       "password",
			"csrf_token":   flow.Ui.Nodes[0].Attributes.UiNodeInputAttributes.Value.(string), // CSRF Token
		}
		payload, _ := json.Marshal(registerData)

		// Submit registration request
		registerURL := fmt.Sprintf("%s/self-service/registration?flow=%s", s.KratosPublicEndpoint, flow.Id)
		reqRegister, _ := http.NewRequest(http.MethodPost, registerURL, ioutil.NopCloser(bytes.NewReader(payload)))
		reqRegister.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		respRegister, err := client.Do(reqRegister)
		if err != nil || respRegister.StatusCode != http.StatusOK {
			http.Error(w, "Registration failed", http.StatusBadRequest)
			return
		}
		defer respRegister.Body.Close()

		// Parse response & return success
		var registerResponse ory.SuccessfulNativeRegistration
		json.NewDecoder(respRegister.Body).Decode(&registerResponse)

		Success(w, registerResponse)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

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
