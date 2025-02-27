package usecase

import (
	"context"
	"iam/gateway"
	"shared/core"
)

type UserDeleteKratosReq struct {
	ID string `json:"id"`
}

type UserDeleteKratosRes struct {
}

type UserDeleteKratos = core.ActionHandler[UserDeleteKratosReq, UserDeleteKratosRes]

func ImplUserDeleteKratos(
	userDelete gateway.UserDeleteKratos,
) UserDeleteKratos {
	return func(ctx context.Context, request UserDeleteKratosReq) (*UserDeleteKratosRes, error) {

		if _, err := userDelete(ctx, gateway.UserDeleteKratosReq{ID: request.ID}); err != nil {
			return nil, err
		}

		return &UserDeleteKratosRes{}, nil
	}
}
