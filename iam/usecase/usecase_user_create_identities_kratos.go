package usecase

import (
	"context"
	"iam/gateway"
	"shared/core"
)

type UserCreateKratosRes struct{}

type UserCreateKratos = core.ActionHandler[gateway.UserCreateKratosReq, UserCreateKratosRes]

func ImplUserCreateKratos(
	userSave gateway.UserCreateKratos,
) UserCreateKratos {
	return func(ctx context.Context, request gateway.UserCreateKratosReq) (*UserCreateKratosRes, error) {

		if _, err := userSave(ctx, gateway.UserCreateKratosReq{User: request.User}); err != nil {
			return nil, err
		}

		return &UserCreateKratosRes{}, nil
	}
}
