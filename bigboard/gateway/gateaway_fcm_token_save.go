package gateway

import (
	"bigboard/model"
	"context"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type FCMTokenSaveReq struct {
	Token string
}

type FCMTokenSaveRes struct {
}

type FCMTokenSave = core.ActionHandler[FCMTokenSaveReq, FCMTokenSaveRes]

func ImplFCMTokenSave(db *gorm.DB) FCMTokenSave {
	return func(ctx context.Context, req FCMTokenSaveReq) (*FCMTokenSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		token := model.FCMToken{
			ID:    1,
			Token: req.Token,
		}

		if err := query.Save(&token).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &FCMTokenSaveRes{}, nil
	}
}
