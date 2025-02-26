package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type RancanganPemantauanUpdateUseCaseReq struct {
	ID                   string                 `json:"id"`
	MetodePemantauan     model.MetodePemantauan `json:"metode_pemantuan"`
	PenanggungJawab      string                 `json:"penanggung_jawab"`
	RencanaPenyelesaian  string                 `json:"rencana_penyelesaian"`
	RealisasiPelaksanaan string                 `json:"realisasi_pelaksanaan"`
	Keterangan           string                 `json:"keterangan"`
}

type RancanganPemantauanUpdateUseCaseRes struct{}

type RancanganPemantauanUpdateUseCase = core.ActionHandler[RancanganPemantauanUpdateUseCaseReq, RancanganPemantauanUpdateUseCaseRes]

func ImplRancanganPemantauanUpdateUseCase(
	getRancanganPemantauanById gateway.RancanganPemantauanGetByID,
	updateRancanganPemantauan gateway.RancanganPemantauanSave,
) RancanganPemantauanUpdateUseCase {
	return func(ctx context.Context, req RancanganPemantauanUpdateUseCaseReq) (*RancanganPemantauanUpdateUseCaseRes, error) {

		res, err := getRancanganPemantauanById(ctx, gateway.RancanganPemantauanGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		rancanganPemantauan := res.RancanganPemantauan

		rancanganPemantauan.MetodePemantauan = &req.MetodePemantauan
		rancanganPemantauan.PenanggungJawab = &req.PenanggungJawab
		rancanganPemantauan.RencanaPenyelesaian = &req.RencanaPenyelesaian
		rancanganPemantauan.RealisasiPelaksanaan = &req.RealisasiPelaksanaan
		rancanganPemantauan.Keterangan = &req.Keterangan

		if _, err := updateRancanganPemantauan(ctx, gateway.RancanganPemantauanSaveReq{RancanganPemantauan: rancanganPemantauan}); err != nil {
			return nil, err
		}

		return &RancanganPemantauanUpdateUseCaseRes{}, nil
	}
}
