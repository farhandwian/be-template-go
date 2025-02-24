package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IKUSaveReq struct {
	IKU model.IKU
}

type IKUSaveRes struct {
	ID string
}

type IKUSave = core.ActionHandler[IKUSaveReq, IKUSaveRes]

func ImplIKUSave(db *gorm.DB) IKUSave {
	return func(ctx context.Context, req IKUSaveReq) (*IKUSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.IKU).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IKUSaveRes{ID: *req.IKU.ID}, nil
	}
}
