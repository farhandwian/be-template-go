package usecase

import (
	"context"
	"fmt"
	"iam/gateway"
	"iam/model"
	"shared/core"
)

type LogoutReq struct {
	UserID model.UserID
}

type LogoutRes struct{}

type Logout = core.ActionHandler[LogoutReq, LogoutRes]

func ImplLogout(

	userGetOneByID gateway.UserGetOneByID,
	saveUser gateway.UserSave,

) Logout {
	return func(ctx context.Context, request LogoutReq) (*LogoutRes, error) {

		if err := request.Validate(); err != nil {
			return nil, err
		}

		userObj, err := userGetOneByID(ctx, gateway.UserGetOneByIDReq{UserID: request.UserID})
		if err != nil {
			return nil, err
		}

		if userObj == nil {
			return nil, fmt.Errorf("user id %v not found", request.UserID)
		}

		userObj.User.RefreshTokenID = ""

		if _, err := saveUser(ctx, gateway.UserSaveReq{User: userObj.User}); err != nil {
			return nil, err
		}

		return &LogoutRes{}, nil
	}
}

func (r LogoutReq) Validate() error {

	if r.UserID == "" {
		return fmt.Errorf("user if must not empty")
	}

	return nil
}
