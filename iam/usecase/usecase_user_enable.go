package usecase

import (
	"context"
	"fmt"
	"iam/gateway"
	"iam/model"
	"shared/core"
)

type UserEnableReq struct {
	UserID model.UserID
	Enable bool
}

type UserEnableRes struct{}

type UserEnable = core.ActionHandler[UserEnableReq, UserEnableRes]

func ImplUserEnable(
	userGetOneByID gateway.UserGetOneByID,
	userSave gateway.UserSave,
) UserEnable {
	return func(ctx context.Context, request UserEnableReq) (*UserEnableRes, error) {

		userObj, err := userGetOneByID(ctx, gateway.UserGetOneByIDReq{UserID: request.UserID})
		if err != nil {
			return nil, err
		}

		if userObj == nil {
			return nil, fmt.Errorf("user id %v not found", request.UserID)
		}

		userObj.User.Enabled = request.Enable

		if _, err := userSave(ctx, gateway.UserSaveReq{User: userObj.User}); err != nil {
			return nil, err
		}

		return &UserEnableRes{}, nil
	}
}
