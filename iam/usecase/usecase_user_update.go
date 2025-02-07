package usecase

import (
	"context"
	"fmt"
	"iam/gateway"
	"iam/model"
	"shared/core"
)

type UserUpdateReq struct {
	UserID      model.UserID
	Name        string
	PhoneNumber model.PhoneNumber
	Email       model.Email
}

type UserUpdateRes struct{}

type UserUpdate = core.ActionHandler[UserUpdateReq, UserUpdateRes]

func ImplUserUpdate(
	userGetOneByID gateway.UserGetOneByID,
	userSave gateway.UserSave,
) UserUpdate {
	return func(ctx context.Context, request UserUpdateReq) (*UserUpdateRes, error) {

		userObj, err := userGetOneByID(ctx, gateway.UserGetOneByIDReq{
			UserID: request.UserID,
		})
		if err != nil {
			return nil, err
		}

		if userObj == nil {
			return nil, fmt.Errorf("user id %v not found", request.UserID)
		}

		userObj.User.Email = request.Email
		userObj.User.PhoneNumber = request.PhoneNumber
		userObj.User.Name = request.Name

		if _, err := userSave(ctx, gateway.UserSaveReq{User: userObj.User}); err != nil {
			return nil, err
		}

		return &UserUpdateRes{}, nil
	}
}
