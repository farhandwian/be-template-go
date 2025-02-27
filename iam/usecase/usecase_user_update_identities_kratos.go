package usecase

import (
	"context"
	"iam/gateway"
	"shared/core"
)

type UserUpdateKratosRes struct{}

type UserUpdateKratos = core.ActionHandler[gateway.UserUpdateKratosReq, UserUpdateKratosRes]

func ImplUserUpdateKratos(
	userSave gateway.UserUpdateKratos,
) UserUpdateKratos {
	return func(ctx context.Context, request gateway.UserUpdateKratosReq) (*UserUpdateKratosRes, error) {

		if _, err := userSave(ctx, gateway.UserUpdateKratosReq{User: request.User}); err != nil {
			return nil, err
		}

		return &UserUpdateKratosRes{}, nil
	}
}
