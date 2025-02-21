package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"iam/model"
	"net/http"
	"shared/core"
	"shared/helper"
	"strings"

	ketoHelper "shared/helper/ory/keto"

	"github.com/google/uuid"
	ory "github.com/ory/client-go"
)

const requestIDKey core.ContextKey = "REQUEST_ID"

func RequestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

const UserIDContext core.ContextKey = "userID"

const UserAccessContext core.ContextKey = "userAccess"

func GetBearerToken(w http.ResponseWriter, r *http.Request) (string, string, bool) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", "authorization header required", false
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		return "", "invalid Authorization header format", false
	}

	return bearerToken[1], "", true
}

func Authentication(next http.HandlerFunc, jwt helper.JWTTokenizer) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		bearerToken, errMessage, ok := GetBearerToken(w, r)
		if !ok {
			writeJSON(w, http.StatusUnauthorized, Response{Status: "failed", Error: &errMessage})
			return
		}

		content, err := jwt.VerifyToken(bearerToken)
		if err != nil {
			msg := "unverified token"
			writeJSON(w, http.StatusUnauthorized, Response{Status: "failed", Error: &msg})
			return
		}

		var userTokenPayload model.UserTokenPayload
		if err := json.Unmarshal(content, &userTokenPayload); err != nil {
			msg := "incorrect token payload"
			writeJSON(w, http.StatusUnauthorized, Response{Status: "failed", Error: &msg})
			return
		}

		ctx := core.AttachDataToContext(r.Context(), UserAccessContext, userTokenPayload.UserAccess)
		ctx = core.AttachDataToContext(ctx, UserIDContext, userTokenPayload.UserID)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	}

}

func AuthenticationKratos(next http.HandlerFunc, ory *ory.APIClient, keto *ketoHelper.KetoGRPCClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("=== Incoming Request Headers ===")
		// for key, values := range r.Header {
		// 	for _, value := range values {
		// 		fmt.Printf("%s: %s\n", key, value)
		// 	}
		// }
		// fmt.Println("=================================")

		cookies := r.Header.Get("Cookie")
		// fmt.Println("cookies: " + cookies)

		// // Look up session.
		session, _, err := ory.FrontendAPI.ToSession(r.Context()).Cookie(cookies).Execute()
		// Check if a session exists and if it is active.
		// You could add your own logic here to check if the session is valid for the specific endpoint, e.g. using the `session.AuthenticatedAt` field.

		if err != nil || !*session.Active {
			msg := "unverified session cookie"
			writeJSON(w, http.StatusUnauthorized, Response{Status: "failed", Error: &msg})
			return
		}

		sessionBytes, err := json.Marshal(session)
		if err != nil {
			fmt.Println("Error marshaling session to JSON:", err)
			return
		}
		var sessionNew model.Session
		err = json.Unmarshal(sessionBytes, &sessionNew)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}

		// Cetak hasil parsing
		// fmt.Println("Parsed Session Struct:")
		// fmt.Printf("%+v\n", sessionNew)

		ua := model.NewUserAccessKeto(string(sessionNew.Identity.ID), keto)

		ctx := core.AttachDataToContext(r.Context(), UserAccessContext, ua)
		ctx = core.AttachDataToContext(ctx, UserIDContext, sessionNew.Identity.ID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}

}

func AuthenticationKeto(next http.HandlerFunc, jwt helper.JWTTokenizer, keto *ketoHelper.KetoGRPCClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		bearerToken, errMessage, ok := GetBearerToken(w, r)
		if !ok {
			writeJSON(w, http.StatusUnauthorized, Response{Status: "failed", Error: &errMessage})
			return
		}

		content, err := jwt.VerifyToken(bearerToken)
		if err != nil {
			msg := "unverified token"
			writeJSON(w, http.StatusUnauthorized, Response{Status: "failed", Error: &msg})
			return
		}

		var userTokenPayload model.UserTokenPayload
		if err := json.Unmarshal(content, &userTokenPayload); err != nil {
			msg := "incorrect token payload"
			writeJSON(w, http.StatusUnauthorized, Response{Status: "failed", Error: &msg})
			return
		}

		ua := model.NewUserAccessKeto(string(userTokenPayload.UserID), keto)

		ctx := core.AttachDataToContext(r.Context(), UserAccessContext, ua)
		ctx = core.AttachDataToContext(ctx, UserIDContext, userTokenPayload.UserID)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	}

}

func AuthorizationKeto(next http.HandlerFunc, access model.AccessKeto) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// by pass
		if access.Object == "none" {
			next.ServeHTTP(w, r)
			return
		}

		userAccess := core.GetDataFromContext(r.Context(), UserAccessContext, model.UserAccess(""))
		fmt.Println(userAccess)
		ua := core.GetDataFromContext[*model.UserAccessKeto](r.Context(), UserAccessContext, nil)
		fmt.Println(ua)
		if ua == nil || ua.UserID == "" {
			msg := "user not authenticated"
			writeJSON(w, http.StatusUnauthorized, Response{Status: "failed", Error: &msg})
			return
		}

		if !ua.HasAccess(r.Context(), access.Namespace, access.Relation, access.Object) {
			msg := "unauthorized operation"
			writeJSON(w, http.StatusForbidden, Response{Status: "failed", Error: &msg})
			return
		}

		next.ServeHTTP(w, r)

	}
}

func Authorization(next http.HandlerFunc, access model.Access) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// by pass
		if access == "0" {
			next.ServeHTTP(w, r)
			return
		}

		userAccess := core.GetDataFromContext(r.Context(), UserAccessContext, model.UserAccess(""))
		fmt.Println(userAccess)
		if !userAccess.HasAccess(access) {
			msg := "unauthorized operation"
			writeJSON(w, http.StatusForbidden, Response{Status: "failed", Error: &msg})
			return
		}

		next.ServeHTTP(w, r)

	}
}

// func AuthMiddleware(jwt helper.JWTTokenizer, next http.HandlerFunc, access model.Access) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		authHeader := r.Header.Get("Authorization")
// 		if authHeader == "" {
// 			http.Error(w, "Authorization header required", http.StatusUnauthorized)
// 			return
// 		}

// 		bearerToken := strings.Split(authHeader, " ")
// 		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
// 			http.Error(w, "invalid Authorization header format", http.StatusUnauthorized)
// 			return
// 		}

// 		content, err := jwt.VerifyToken(bearerToken[1])
// 		if err != nil {
// 			http.Error(w, "unverified token", http.StatusUnauthorized)
// 			return
// 		}

// 		var userTokenPayload model.UserTokenPayload
// 		if err := json.Unmarshal(content, &userTokenPayload); err != nil {
// 			http.Error(w, "incorrect token payload", http.StatusUnauthorized)
// 			return
// 		}

// 		if !userTokenPayload.UserAccess.HasAccess(access) {
// 			http.Error(w, "unauthorized operation", http.StatusUnauthorized)
// 			return
// 		}

// 		// r = AttachUserID(r, string(userTokenPayload.UserID))

// 		ctx := core.AttachDataToContext(r.Context(), UserIDContext, userTokenPayload.UserID)

// 		r = r.WithContext(ctx)

// 		next.ServeHTTP(w, r)
// 	}
// }
