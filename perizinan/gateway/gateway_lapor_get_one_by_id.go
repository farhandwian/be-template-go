package gateway

import (
	"context"
	"perizinan/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type LaporanPerizinanGetOneByIDReq struct {
	LaporanPerizinanID model.LaporanPerizinanID
}

type LaporanPerizinanGetOneByIDRes struct {
	LaporanPerizinan *model.LaporanPerizinan `json:"laporan_perizinan"`
}

type LaporanPerizinanGetOneByID = core.ActionHandler[LaporanPerizinanGetOneByIDReq, LaporanPerizinanGetOneByIDRes]

func ImplLaporanPerizinanGetOneByID(db *gorm.DB) LaporanPerizinanGetOneByID {
	return func(ctx context.Context, req LaporanPerizinanGetOneByIDReq) (*LaporanPerizinanGetOneByIDRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var laporanPerizinan model.LaporanPerizinan

		err := query.First(&laporanPerizinan, "id = ?", req.LaporanPerizinanID).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return &LaporanPerizinanGetOneByIDRes{LaporanPerizinan: nil}, nil
			}
			return nil, core.NewInternalServerError(err)
		}

		return &LaporanPerizinanGetOneByIDRes{LaporanPerizinan: &laporanPerizinan}, nil
	}
}
