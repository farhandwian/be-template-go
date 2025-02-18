package controller

import (
	"iam/model"
	"net/http"
	"shared/helper"
	ory "shared/helper/ory"
)

func (c Controller) LoginHandler(s ory.ORYServer) helper.APIData {

	type Body struct {
		Email    model.Email `json:"email"`
		Password string      `json:"password"`
	}

	apiData := helper.APIData{
		Access:   model.ANONYMOUS,
		Method:   http.MethodPost,
		Url:      "/auth/login",
		Body:     Body{},
		Summary:  "Initiate user login",
		Tag:      "IAM - Authentication",
		Examples: []helper.ExampleResponse{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}

// package controller

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	hydra "github.com/ory/hydra-client-go/v2"
// 	kratos "github.com/ory/kratos-client-go"
// )

// var (
// 	hydraAdminURL = "http://localhost:4445"
// 	client        = hydra.NewAPIClient(hydra.NewConfiguration())
// )

// // Login Handler
// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	loginChallenge := r.URL.Query().Get("login_challenge")
// 	if loginChallenge == "" {
// 		http.Error(w, "Missing login_challenge", http.StatusBadRequest)
// 		return
// 	}

// 	// Check if user has an active Kratos session
// 	session, err := checkKratosSession(r)
// 	if err != nil || !session.Active {
// 		http.Redirect(w, r, fmt.Sprintf("%s/self-service/login/browser", kratosPublicURL), http.StatusFound)
// 		return
// 	}

// 	// Accept login challenge in Hydra
// 	adminAPI := client.AdminApi
// 	response, _, err := adminAPI.AcceptLoginRequest(r.Context()).
// 		LoginChallenge(loginChallenge).
// 		AcceptLoginRequest(hydra.AcceptLoginRequest{Subject: session.Identity.Id}).
// 		Execute()

// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error accepting login: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(w, r, response.RedirectTo, http.StatusFound)
// }

// // Check Kratos session
// func checkKratosSession(r *http.Request) (*kratos.Session, error) {
// 	client := kratos.NewAPIClient(kratos.NewConfiguration())
// 	api := client.V0alpha2Api
// 	sessionCookie := r.Header.Get("Cookie")

// 	session, _, err := api.ToSession(r.Context()).Cookie(sessionCookie).Execute()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return session, nil
// }
