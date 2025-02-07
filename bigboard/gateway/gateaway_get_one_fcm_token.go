package gateway

import (
	"bigboard/model"
	"context"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type GetOneFCMTokenReq struct {
	ID int `json:"id"`
}

type GetOneFCMToken = core.ActionHandler[GetOneFCMTokenReq, model.FCMToken]

func ImplGetOneFCMToken(db *gorm.DB) GetOneFCMToken {
	return func(ctx context.Context, req GetOneFCMTokenReq) (*model.FCMToken, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var token model.FCMToken

		if err := query.First(&token, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &token, nil
	}
}
