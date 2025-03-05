package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	sharedModel "shared/model"
)

type RancanganPemantauanCreateUseCaseReq struct {
	PenilaianRisikoID        string                 `json:"penilaian_risiko_id"`
	MetodePemantauan         model.MetodePemantauan `json:"metode_pemantuan"`
	PenanggungJawab          string                 `json:"penanggung_jawab"`
	RencanaWaktuPemantauan   string                 `json:"rencana_waktu_pemantauan"`
	RealisasiWaktuPemantauan string                 `json:"realisasi_waktu_pemantauan"`
	Keterangan               string                 `json:"keterangan"`
}

type RancanganPemantauanCreateUseCaseRes struct {
	ID string `json:"id"`
}

type RancanganPemantauanCreateUseCase = core.ActionHandler[RancanganPemantauanCreateUseCaseReq, RancanganPemantauanCreateUseCaseRes]

func ImplRancanganPemantauanCreateUseCase(
	generateId gateway.GenerateId,
	createRancanganPemantauan gateway.RancanganPemantauanSave,
	PenilaianRisikoByID gateway.PenilaianRisikoGetByID,
) RancanganPemantauanCreateUseCase {
	return func(ctx context.Context, req RancanganPemantauanCreateUseCaseReq) (*RancanganPemantauanCreateUseCaseRes, error) {

		// Generate a unique ID
		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		_, err = PenilaianRisikoByID(ctx, gateway.PenilaianRisikoGetByIDReq{ID: req.PenilaianRisikoID})
		if err != nil {
			return nil, err
		}

		obj := model.RancanganPemantauan{
			ID:                       &genObj.RandomId,
			PenilaianRisikoID:        &req.PenilaianRisikoID,
			MetodePemantauan:         &req.MetodePemantauan,
			PenanggungJawab:          &req.PenanggungJawab,
			RencanaWaktuPemantauan:   &req.RencanaWaktuPemantauan,
			RealisasiWaktuPemantauan: &req.RealisasiWaktuPemantauan,
			Keterangan:               &req.Keterangan,
			Status:                   sharedModel.StatusMenungguVerifikasi,
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
