package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type RancanganPemantauanCreateUseCaseReq struct {
	MetodePemantauan     model.MetodePemantauan `json:"metode_pemantuan"`
	PenanggungJawab      string                 `json:"penanggung_jawab"`
	RencanaPenyelesaian  string                 `json:"rencana_penyelesaian"`
	RealisasiPelaksanaan string                 `json:"realisasi_pelaksanaan"`
	Keterangan           string                 `json:"keterangan"`
}

type RancanganPemantauanCreateUseCaseRes struct {
	ID string `json:"id"`
}

type RancanganPemantauanCreateUseCase = core.ActionHandler[RancanganPemantauanCreateUseCaseReq, RancanganPemantauanCreateUseCaseRes]

func ImplRancanganPemantauanCreateUseCase(
	generateId gateway.GenerateId,
	createRancanganPemantauan gateway.RancanganPemantauanSave,
) RancanganPemantauanCreateUseCase {
	return func(ctx context.Context, req RancanganPemantauanCreateUseCaseReq) (*RancanganPemantauanCreateUseCaseRes, error) {

		// Generate a unique ID
		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		obj := model.RancanganPemantauan{
			ID:                   &genObj.RandomId,
			MetodePemantauan:     &req.MetodePemantauan,
			PenanggungJawab:      &req.PenanggungJawab,
			RencanaPenyelesaian:  &req.RencanaPenyelesaian,
			RealisasiPelaksanaan: &req.RealisasiPelaksanaan,
			Keterangan:           &req.Keterangan,
		}

		// Save the RancanganPemantauan entry
		if _, err = createRancanganPemantauan(ctx, gateway.RancanganPemantauanSaveReq{RancanganPemantauan: obj}); err != nil {
			return nil, err
		}

		return &RancanganPemantauanCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
