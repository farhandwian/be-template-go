package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"

	"iam/model"

	"gorm.io/gorm"
)

type UserGetOneByIDReq struct {
	UserID model.UserID
}

type UserGetOneByIDRes struct {
	User model.User
}

type UserGetOneByID = core.ActionHandler[UserGetOneByIDReq, UserGetOneByIDRes]

func ImplUserGetOneByID() UserGetOneByID {
	return func(ctx context.Context, request UserGetOneByIDReq) (*UserGetOneByIDRes, error) {

		return &UserGetOneByIDRes{}, nil
	}
}

func ImplUserGetOneByIDWithDatabase(db *gorm.DB) UserGetOneByID {
	return func(ctx context.Context, request UserGetOneByIDReq) (*UserGetOneByIDRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var user model.User
		if err := query.First(&user, request.UserID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// return nil, fmt.Errorf("user id %v not found", request.UserID)
				return nil, nil
			}
			return nil, core.NewInternalServerError(err)
		}

		return &UserGetOneByIDRes{
			User: user,
		}, nil
	}
}
