package usecase

import (
	"context"
	"fmt"
	"perizinan/gateway"
	"perizinan/model"
	"shared/core"
	"time"
)

type LaporanPerizinanExistUseCaseReq struct {
	NomorSK               model.NomorSK `json:"nomor_sk"`
	PeriodePengambilanSDA model.Periode `json:"periode_pengambilan_sda"`
	Min                   time.Time
	Now                   time.Time
}

type LaporanPerizinanExistUseCaseRes struct{}

type LaporanPerizinanExistUseCase = core.ActionHandler[LaporanPerizinanExistUseCaseReq, LaporanPerizinanExistUseCaseRes]

func ImplLaporanPerizinanExistUseCase(
	getOneSK gateway.SKPerizinanGetOne,
	getOneLapor gateway.LaporanPerizinanGetOne,
) LaporanPerizinanExistUseCase {
	return func(ctx context.Context, req LaporanPerizinanExistUseCaseReq) (*LaporanPerizinanExistUseCaseRes, error) {

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

		resLapor, err := getOneLapor(ctx, gateway.LaporanPerizinanGetOneReq{
			NomorSK:               req.NomorSK,
			PeriodePengambilanSDA: req.PeriodePengambilanSDA,
		})
		if err != nil {
			return nil, err
		}

		_ = resLapor

		// if resLapor.LaporanPerizinan != nil && resLapor.LaporanPerizinan.Status == model.StatusLaporSubmitted {
		// 	return nil, fmt.Errorf("laporan perizinan dengan SK nomor '%s' dalam periode %s sudah disubmit", req.NomorSK, req.PeriodePengambilanSDA)
		// }

		return &LaporanPerizinanExistUseCaseRes{}, nil
	}
}
