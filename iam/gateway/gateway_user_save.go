package gateway

import (
	"context"
	"iam/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type UserSaveReq struct {
	User model.User
}

type UserSaveRes struct {
}

type UserSave = core.ActionHandler[UserSaveReq, UserSaveRes]

func ImplUserSave() UserSave {
	return func(ctx context.Context, request UserSaveReq) (*UserSaveRes, error) {

		return &UserSaveRes{}, nil
	}
}

func ImplUserSaveWithDatabse(db *gorm.DB) UserSave {
	return func(ctx context.Context, request UserSaveReq) (*UserSaveRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&request.User).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &UserSaveRes{}, nil
	}
}
