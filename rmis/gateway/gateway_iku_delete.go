// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IKUDeleteReq struct {
	ID string
}

type IKUDeleteRes struct{}

type IKUDelete = core.ActionHandler[IKUDeleteReq, IKUDeleteRes]

func ImplIKUDelete(db *gorm.DB) IKUDelete {
	return func(ctx context.Context, req IKUDeleteReq) (*IKUDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.IKU{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IKUDeleteRes{}, nil
	}
}
