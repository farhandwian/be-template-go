package gateway

import (
	"context"
	"fmt"
	"io"
	"shared/core"

	ory "github.com/ory/client-go"
)

type UserDeleteKratosReq struct {
	ID string `json:"id"`
}

type UserDeleteKratosRes struct {
	ID string `json:"id"`
}

type UserDeleteKratos = core.ActionHandler[UserDeleteKratosReq, UserDeleteKratosRes]

func ImplUserDeleteKratos(oryClient *ory.APIClient) UserDeleteKratos {
	return func(ctx context.Context, request UserDeleteKratosReq) (*UserDeleteKratosRes, error) {
		authCtx := context.WithValue(ctx, ory.ContextAccessToken, "ory_pat_fm6j3rn3rvoN3akofwijUNXjLjJ5PGAL")

		// Call Ory Kratos API to delete identity
		resp, err := oryClient.IdentityAPI.DeleteIdentity(authCtx, request.ID).Execute()

		// Handle API errors
		if err != nil {
			// Convert error to Ory API response error
			if resp != nil {
				bodyBytes, readErr := io.ReadAll(resp.Body)
				if readErr == nil {
					fmt.Println("Error response from Ory Kratos:", string(bodyBytes))
				} else {
					fmt.Println("Failed to read error response body:", readErr)
				}
			}
			fmt.Println("Error deleting identity in Kratos:", err)
			return nil, err
		}
		defer resp.Body.Close()

		// Return deleted identity ID
		return &UserDeleteKratosRes{
			ID: request.ID,
		}, nil
	}
}
