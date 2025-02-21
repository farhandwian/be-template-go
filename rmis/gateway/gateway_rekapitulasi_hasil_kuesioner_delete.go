// File: gateway/gateway_asset.go

package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type RekapitulasiHasilKuesionerDeleteReq struct {
	ID string
}

type RekapitulasiHasilKuesionerDeleteRes struct{}

type RekapitulasiHasilKuesionerDelete = core.ActionHandler[RekapitulasiHasilKuesionerDeleteReq, RekapitulasiHasilKuesionerDeleteRes]

func ImplRekapitulasiHasilKuesionerDelete(db *gorm.DB) RekapitulasiHasilKuesionerDelete {
	return func(ctx context.Context, req RekapitulasiHasilKuesionerDeleteReq) (*RekapitulasiHasilKuesionerDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.RekapitulasiHasilKuesioner{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &RekapitulasiHasilKuesionerDeleteRes{}, nil
	}
}
