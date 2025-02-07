package usecase

import (
	"context"
	"fmt"
	"iam/gateway"
	"shared/core"
)

type UserGetOneReq struct{ gateway.UserGetOneByIDReq }

type UserGetOneRes struct{ gateway.UserGetOneByIDRes }

type UserGetOne = core.ActionHandler[UserGetOneReq, UserGetOneRes]

func ImplUserGetOne(
	userGetOneByID gateway.UserGetOneByID,
) UserGetOne {
	return func(ctx context.Context, request UserGetOneReq) (*UserGetOneRes, error) {

		if err := request.Validate(); err != nil {
			return nil, err
		}

		userObj, err := userGetOneByID(ctx, gateway.UserGetOneByIDReq{
			UserID: request.UserID,
		})
		if err != nil {
			return nil, err
		}

		if userObj == nil {
			return nil, fmt.Errorf("user id %v not found", request.UserID)
		}

		return &UserGetOneRes{
			UserGetOneByIDRes: *userObj,
		}, nil
	}
}

func (r UserGetOneReq) Validate() error {
	return nil
}
