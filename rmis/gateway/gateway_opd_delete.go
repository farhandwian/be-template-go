// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type OPDDeleteReq struct {
	ID string
}

type OPDDeleteRes struct{}

type OPDDelete = core.ActionHandler[OPDDeleteReq, OPDDeleteRes]

func ImplOPDDelete(db *gorm.DB) OPDDelete {
	return func(ctx context.Context, req OPDDeleteReq) (*OPDDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.OPD{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &OPDDeleteRes{}, nil
	}
}
