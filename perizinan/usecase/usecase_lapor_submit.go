package usecase

import (
	"context"
	"fmt"
	"perizinan/gateway"
	"perizinan/model"
	"shared/core"
)

type LaporanPerizinanSubmitUseCaseReq struct {
	LaporanPerizinanID model.LaporanPerizinanID `json:"laporan_perizinan_id"`
}

type LaporanPerizinanSubmitUseCaseRes struct {
	// Message string `json:"message"`
}

type LaporanPerizinanSubmitUseCase = core.ActionHandler[
	LaporanPerizinanSubmitUseCaseReq,
	LaporanPerizinanSubmitUseCaseRes,
]

func ImplLaporanPerizinanSubmitUseCase(
	//
	getOneLapor gateway.LaporanPerizinanGetOneByID,
	saveLaporanPerizinan gateway.LaporanPerizinanSave, //
) LaporanPerizinanSubmitUseCase {
	return func(ctx context.Context, req LaporanPerizinanSubmitUseCaseReq) (*LaporanPerizinanSubmitUseCaseRes, error) {

		resLapor, err := getOneLapor(ctx, gateway.LaporanPerizinanGetOneByIDReq{
			LaporanPerizinanID: req.LaporanPerizinanID,
		})
		if err != nil {
			return nil, err
		}

		if resLapor.LaporanPerizinan == nil {
			return nil, fmt.Errorf("laporan perizinan dengan id '%s' tidak ditemukan", req.LaporanPerizinanID)
		}

		if resLapor.LaporanPerizinan.Status == model.StatusLaporSubmitted {
			nomorSK := resLapor.LaporanPerizinan.NoSK
			periode := resLapor.LaporanPerizinan.PeriodePengambilanSDA
			return nil, fmt.Errorf("laporan perizinan dengan SK nomor '%s' dalam periode %s sudah pernah disubmit", nomorSK, periode)
		}

		resLapor.LaporanPerizinan.Status = model.StatusLaporSubmitted

		if _, err = saveLaporanPerizinan(ctx, gateway.LaporanPerizinanSaveReq{LaporanPerizinan: resLapor.LaporanPerizinan}); err != nil {
			return nil, err
		}

		return &LaporanPerizinanSubmitUseCaseRes{}, nil
	}
}
