package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type RcaSaveReq struct {
	Rca model.Rca
}

type RcaSaveRes struct {
	ID string
}

type RcaSave = core.ActionHandler[RcaSaveReq, RcaSaveRes]

func ImplRcaSave(db *gorm.DB) RcaSave {
	return func(ctx context.Context, req RcaSaveReq) (*RcaSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.Rca).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &RcaSaveRes{ID: *req.Rca.ID}, nil
	}
}
