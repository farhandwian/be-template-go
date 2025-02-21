package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type RekapitulasiHasilKuesionerSaveReq struct {
	RekapitulasiHasilKuesioner model.RekapitulasiHasilKuesioner
}

type RekapitulasiHasilKuesionerSaveRes struct {
	ID string
}

type RekapitulasiHasilKuesionerSave = core.ActionHandler[RekapitulasiHasilKuesionerSaveReq, RekapitulasiHasilKuesionerSaveRes]

func ImplRekapitulasiHasilKuesionerSave(db *gorm.DB) RekapitulasiHasilKuesionerSave {
	return func(ctx context.Context, req RekapitulasiHasilKuesionerSaveReq) (*RekapitulasiHasilKuesionerSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.RekapitulasiHasilKuesioner).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &RekapitulasiHasilKuesionerSaveRes{ID: *req.RekapitulasiHasilKuesioner.ID}, nil
	}
}
