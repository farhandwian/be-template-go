package usecase

import (
	"context"
	"fmt"
	"perizinan/gateway"
	"perizinan/model"
	"shared/core"
	"time"
)

type LaporanPerizinanSKPeriodeUseCaseReq struct {
	NomorSK               model.NomorSK `json:"nomor_sk"`
	PeriodePengambilanSDA model.Periode `json:"periode_pengambilan_sda"`
	Min                   time.Time
	Now                   time.Time
}

type LaporanPerizinanSKPeriodeUseCaseRes struct {
	LaporanPerizinan *model.LaporanPerizinan `json:"laporan_perizinan"`
	SKPerizinan      *model.SKPerizinan      `json:"sk_perizinan"`
}

type LaporanPerizinanSKPeriodeUseCase = core.ActionHandler[LaporanPerizinanSKPeriodeUseCaseReq, LaporanPerizinanSKPeriodeUseCaseRes]

func ImplLaporanPerizinanSKPeriodeUseCase(
	getOneSK gateway.SKPerizinanGetOne,
	getOneLaporan gateway.LaporanPerizinanGetOne,
) LaporanPerizinanSKPeriodeUseCase {
	return func(ctx context.Context, req LaporanPerizinanSKPeriodeUseCaseReq) (*LaporanPerizinanSKPeriodeUseCaseRes, error) {

		if err := req.PeriodePengambilanSDA.Validate(req.Min, req.Now); err != nil {
			return nil, err
		}

		if err := req.NomorSK.Validate(); err != nil {
			return nil, err
		}

		resSK, err := getOneSK(ctx, gateway.SKPerizinanGetOneReq{
			NomorSK: req.NomorSK,
		})
		if err != nil {
			return nil, err
		}

		if resSK.SKPerizinan == nil {
			return nil, fmt.Errorf("laporan perizinan dengan SK no '%s' tidak ditemukan", req.NomorSK)
		}

		resLapor, err := getOneLaporan(ctx, gateway.LaporanPerizinanGetOneReq{
			NomorSK:               req.NomorSK,
			PeriodePengambilanSDA: req.PeriodePengambilanSDA,
		})
		if err != nil {
			return nil, err
		}

		if resLapor.LaporanPerizinan == nil {

			return &LaporanPerizinanSKPeriodeUseCaseRes{
				LaporanPerizinan: model.LaporanPerizinanEmpty(req.NomorSK, req.PeriodePengambilanSDA),
				SKPerizinan:      resSK.SKPerizinan,
			}, nil

		}

		return &LaporanPerizinanSKPeriodeUseCaseRes{
			LaporanPerizinan: resLapor.LaporanPerizinan,
			SKPerizinan:      resSK.SKPerizinan,
		}, nil
	}
}
