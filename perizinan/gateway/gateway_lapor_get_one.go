package gateway

import (
	"context"
	"perizinan/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type LaporanPerizinanGetOneReq struct {
	NomorSK               model.NomorSK `json:"nomor_sk"`
	PeriodePengambilanSDA model.Periode `json:"periode_pengambilan_sda"`
}

type LaporanPerizinanGetOneRes struct {
	LaporanPerizinan *model.LaporanPerizinan `json:"laporan_perizinan"`
}

type LaporanPerizinanGetOne = core.ActionHandler[LaporanPerizinanGetOneReq, LaporanPerizinanGetOneRes]

func ImplLaporanPerizinanGetOne(db *gorm.DB) LaporanPerizinanGetOne {
	return func(ctx context.Context, req LaporanPerizinanGetOneReq) (*LaporanPerizinanGetOneRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var laporanPerizinan model.LaporanPerizinan

		err := query.First(&laporanPerizinan, "no_sk = ? AND periode_pengambilan_sda = ?", req.NomorSK, req.PeriodePengambilanSDA).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return &LaporanPerizinanGetOneRes{LaporanPerizinan: nil}, nil
			}
			return nil, core.NewInternalServerError(err)
		}

		return &LaporanPerizinanGetOneRes{LaporanPerizinan: &laporanPerizinan}, nil
	}
}
