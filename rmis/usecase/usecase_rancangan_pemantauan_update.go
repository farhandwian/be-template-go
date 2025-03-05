package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	sharedModel "shared/model"
)

type RancanganPemantauanUpdateUseCaseReq struct {
	ID                       string                 `json:"id"`
	PenilaianRisikoID        string                 `json:"penilaian_risiko_id"`
	MetodePemantauan         model.MetodePemantauan `json:"metode_pemantuan"`
	PenanggungJawab          string                 `json:"penanggung_jawab"`
	RencanaWaktuPemantauan   string                 `json:"rencana_waktu_pemantauan"`
	RealisasiWaktuPemantauan string                 `json:"realisasi_waktu_pemantauan"`
	Keterangan               string                 `json:"keterangan"`
}

type RancanganPemantauanUpdateUseCaseRes struct{}

type RancanganPemantauanUpdateUseCase = core.ActionHandler[RancanganPemantauanUpdateUseCaseReq, RancanganPemantauanUpdateUseCaseRes]

func ImplRancanganPemantauanUpdateUseCase(
	getRancanganPemantauanById gateway.RancanganPemantauanGetByID,
	updateRancanganPemantauan gateway.RancanganPemantauanSave,
	PenilaianRisikoByID gateway.PenilaianRisikoGetByID,
) RancanganPemantauanUpdateUseCase {
	return func(ctx context.Context, req RancanganPemantauanUpdateUseCaseReq) (*RancanganPemantauanUpdateUseCaseRes, error) {

		res, err := getRancanganPemantauanById(ctx, gateway.RancanganPemantauanGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		_, err = PenilaianRisikoByID(ctx, gateway.PenilaianRisikoGetByIDReq{ID: req.PenilaianRisikoID})
		if err != nil {
			return nil, err
		}

		rancanganPemantauan := res.RancanganPemantauan

		rancanganPemantauan.PenilaianRisikoID = &req.PenilaianRisikoID
		rancanganPemantauan.MetodePemantauan = &req.MetodePemantauan
		rancanganPemantauan.PenanggungJawab = &req.PenanggungJawab
		rancanganPemantauan.RencanaWaktuPemantauan = &req.RencanaWaktuPemantauan
		rancanganPemantauan.RealisasiWaktuPemantauan = &req.RealisasiWaktuPemantauan
		rancanganPemantauan.Keterangan = &req.Keterangan
		rancanganPemantauan.Status = sharedModel.StatusMenungguVerifikasi

		if _, err := updateRancanganPemantauan(ctx, gateway.RancanganPemantauanSaveReq{RancanganPemantauan: rancanganPemantauan}); err != nil {
			return nil, err
		}

		return &RancanganPemantauanUpdateUseCaseRes{}, nil
	}
}
