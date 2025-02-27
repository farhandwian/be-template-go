package gateway

import (
	"context"
	"fmt"
	"iam/model"
	"io"
	"shared/core"

	ory "github.com/ory/client-go"
)

type UserUpdateKratosReq struct {
	User model.UserKratosUpdate
}

type UserUpdateKratosRes struct {
	ID string `json:"id"`
}

type UserUpdateKratos = core.ActionHandler[UserUpdateKratosReq, UserUpdateKratosRes]

func ImplUserUpdateKratos(oryClient *ory.APIClient) UserUpdateKratos {
	return func(ctx context.Context, request UserUpdateKratosReq) (*UserUpdateKratosRes, error) {
		// Prepare traits based on the schema
		traits := map[string]interface{}{
			"email":          request.User.Email,
			"nama":           request.User.Nama,
			"no_telepon":     request.User.NoTelepon,
			"jabatan":        request.User.Jabatan,
			"akses_pengguna": request.User.AksesPengguna,
			"jenis_kelamin":  request.User.JenisKelamin,
		}

		// Update Kratos identity object
		identity := ory.UpdateIdentityBody{
			SchemaId: "c8cceb228f4fe2591ecbb5cd6041f587a9048729f28f0b7f99d3db15938c91ef56144cf71b44a81714829b0adfd361f7f58cf8fb216ee2ea52f6f3c8063e2981", // Ensure this matches your schema ID
			Traits:   traits,
		}

		authCtx := context.WithValue(ctx, ory.ContextAccessToken, "ory_pat_fm6j3rn3rvoN3akofwijUNXjLjJ5PGAL")
		// Call Ory Kratos API to create identity
		createdIdentity, resp, err := oryClient.IdentityAPI.UpdateIdentity(authCtx, string(request.User.ID)).
			UpdateIdentityBody(identity).
			Execute()

		// Handle API errors
		// if err != nil {
		// 	fmt.Println("Error creating identity in Kratos:", err)
		// 	return nil, err
		// }

		if err != nil {
			// Convert error to Ory API response error
			if httpResp := resp; httpResp != nil {
				// Try reading the response body
				bodyBytes, readErr := io.ReadAll(httpResp.Body)
				if readErr == nil {
					fmt.Println("Error response from Ory Kratos:", string(bodyBytes))
				} else {
					fmt.Println("Failed to read error response body:", readErr)
				}
			}
			fmt.Println("Error creating identity in Kratos:", err)
			return nil, err
		}

		defer resp.Body.Close()
		return &UserUpdateKratosRes{
			ID: createdIdentity.Id,
		}, nil
	}
}
