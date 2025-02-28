package gateway

import (
	"context"
	"fmt"
	"iam/model"
	"shared/core"
	"sort"
	"strings"

	ory "github.com/ory/client-go"
)

type UserGetAllIdentitiesKratosReq struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	SortOrder string `json:"sort_order"`
	SortBy    string `json:"sort_by"`
	Keyword   string `json:"keyword"`
}

type UserGetAllIdentitiesKratosRes struct {
	Count int64                 `json:"count"`
	Items []model.UserKratosGet `json:"items"`
}

// json yang disimpan di identities hasil dari ory.IdentityAPI.ListIdentities
// [
// 	{
// 	"created_at": "2019-08-24T14:15:22Z",
// 	"credentials": {},
// 	"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
// 	"metadata_admin": { },
// 	"metadata_public": { },
// 	"organization_id": "string",
// 	"recovery_addresses": [],
// 	"schema_id": "string",
// 	"schema_url": "string",
// 	"state": "active",
// 	"state_changed_at": "2019-08-24T14:15:22Z",
// 	"traits": null,
// 	"updated_at": "2019-08-24T14:15:22Z",
// 	"verifiable_addresses": []
// 	}
// ]

type UserGetAllIdentitiesKratos = core.ActionHandler[UserGetAllIdentitiesKratosReq, UserGetAllIdentitiesKratosRes]

func ImplUserGetAllIdentitiesKratos(oryClient *ory.APIClient) UserGetAllIdentitiesKratos {
	return func(ctx context.Context, request UserGetAllIdentitiesKratosReq) (*UserGetAllIdentitiesKratosRes, error) {

		// Prepare API request parameters
		perPage := int64(request.Size) // Convert int to int64
		page := int64(request.Page)

		// Call Ory Kratos Admin API to list identities
		// Add authorization header to context
		authCtx := context.WithValue(ctx, ory.ContextAccessToken, "ory_pat_fm6j3rn3rvoN3akofwijUNXjLjJ5PGAL")

		identities, resp, err := oryClient.IdentityAPI.ListIdentities(authCtx).
			PageSize(perPage).
			Page(page).
			Execute()

		// Handle API errors
		if err != nil {
			fmt.Println("error ImplUserGetAllIdentitiesKratos Gateway:", err)
			return nil, err
		}
		defer resp.Body.Close()

		// Convert Ory identities to model.UserKratosGet format
		users := make([]model.UserKratosGet, 0)
		for _, identity := range identities {
			traits := extractTraits(identity.Traits)

			user := model.UserKratosGet{
				ID:            model.UserKratosID(identity.Id),
				Email:         traits["email"],
				Nama:          traits["nama"],
				AksesPengguna: traits["akses_pengguna"],
				NoTelepon:     traits["no_telepon"],
				Jabatan:       traits["jabatan"],
				JenisKelamin:  traits["jenis_kelamin"],
				CreatedAt:     *identity.CreatedAt,
				UpdatedAt:     *identity.UpdatedAt,
			}

			// Apply filter
			if request.Keyword != "" && !matchesFilter(user, request.Keyword) {
				continue
			}

			users = append(users, user)
		}

		// Apply sorting
		if request.SortBy != "" {
			sortUsers(users, request.SortBy, request.SortOrder)
		}

		// Return the response
		return &UserGetAllIdentitiesKratosRes{
			Count: int64(len(users)),
			Items: users,
		}, nil
	}
}

func extractTraits(traits interface{}) map[string]string {
	result := make(map[string]string)

	if traitsMap, ok := traits.(map[string]interface{}); ok {
		for key, value := range traitsMap {
			if strVal, valid := value.(string); valid {
				result[key] = strVal
			}
		}
	}
	return result
}

func matchesFilter(user model.UserKratosGet, filter string) bool {
	// Implement your filter logic here
	// Example: return true if the user's name or email contains the filter string
	return contains(user.Nama, filter) || contains(string(user.Email), filter) || contains(user.NoTelepon, filter)
}

func contains(source, substr string) bool {
	return strings.Contains(strings.ToLower(source), strings.ToLower(substr))
}

func sortUsers(users []model.UserKratosGet, sortBy, sortOrder string) {
	// Implement your sorting logic here
	// Example: sort by name or email
	sort.Slice(users, func(i, j int) bool {
		switch sortBy {
		case "name":
			if sortOrder == "desc" {
				return users[i].Nama > users[j].Nama
			}
			return users[i].Nama < users[j].Nama
		case "email":
			if sortOrder == "desc" {
				return users[i].Email > users[j].Email
			}
			return users[i].Email < users[j].Email
		case "no_telepon":
			if sortOrder == "desc" {
				return users[i].NoTelepon > users[j].NoTelepon
			}
			return users[i].NoTelepon < users[j].NoTelepon
		case "updated_at":
			if sortOrder == "desc" {
				return users[i].UpdatedAt.After(users[j].UpdatedAt)
			}
			return users[i].UpdatedAt.Before(users[j].UpdatedAt)
		default:
			return false
		}
	})
}
