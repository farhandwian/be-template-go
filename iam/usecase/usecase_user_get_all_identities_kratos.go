package usecase

import (
	"context"
	"iam/gateway"
	"shared/core"
)

type UserGetAllIdentitiesKratosReq struct {
	gateway.UserGetAllIdentitiesKratosReq
}

type UserGetAllIdentitiesKratosRes struct {
	gateway.UserGetAllIdentitiesKratosRes
}

type UserGetAllIdentitiesKratos = core.ActionHandler[UserGetAllIdentitiesKratosReq, UserGetAllIdentitiesKratosRes]

func ImplUserGetAllIdentitiesKratos(
	userGetAllIdentitiesKratos gateway.UserGetAllIdentitiesKratos,
) UserGetAllIdentitiesKratos {
	return func(ctx context.Context, request UserGetAllIdentitiesKratosReq) (*UserGetAllIdentitiesKratosRes, error) {

		if err := request.Validate(); err != nil {
			return nil, err
		}

		response, err := userGetAllIdentitiesKratos(ctx, request.UserGetAllIdentitiesKratosReq)
		if err != nil {
			return nil, err
		}

		return &UserGetAllIdentitiesKratosRes{
			UserGetAllIdentitiesKratosRes: *response,
		}, nil
	}
}

func (r UserGetAllIdentitiesKratosReq) Validate() error {
	return nil
}
