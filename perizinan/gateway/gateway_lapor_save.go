package gateway

import (
	"context"
	"perizinan/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type LaporanPerizinanSaveReq struct {
	LaporanPerizinan *model.LaporanPerizinan
}

type LaporanPerizinanSaveRes struct {
	ID model.LaporanPerizinanID
}

type LaporanPerizinanSave = core.ActionHandler[LaporanPerizinanSaveReq, LaporanPerizinanSaveRes]

func ImplLaporanPerizinanSave(db *gorm.DB) LaporanPerizinanSave {
	return func(ctx context.Context, req LaporanPerizinanSaveReq) (*LaporanPerizinanSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(req.LaporanPerizinan).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &LaporanPerizinanSaveRes{ID: req.LaporanPerizinan.ID}, nil
	}
}
