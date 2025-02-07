package usecase

import (
	"context"
	"iam/gateway"
	"shared/core"
)

type UserGetAllReq struct{ gateway.UserGetAllReq }

type UserGetAllRes struct{ gateway.UserGetAllRes }

type UserGetAll = core.ActionHandler[UserGetAllReq, UserGetAllRes]

func ImplUserGetAll(
	userGetAll gateway.UserGetAll,
) UserGetAll {
	return func(ctx context.Context, request UserGetAllReq) (*UserGetAllRes, error) {

		if err := request.Validate(); err != nil {
			return nil, err
		}

		response, err := userGetAll(ctx, request.UserGetAllReq)
		if err != nil {
			return nil, err
		}

		return &UserGetAllRes{
			UserGetAllRes: *response,
		}, nil
	}
}

func (r UserGetAllReq) Validate() error {
	return nil
}
