package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	sharedModel "shared/model"
)

type PencatatanKejadianRisikoCreateUseCaseReq struct {
	PenetapanKonteksRisikoStrategisPemdaID string `json:"penetapan_konteks_risiko_strategis_pemda_id"`
	IdentifikasiRisikoStrategisPemdaID     string `json:"identifikasi_risiko_strategis_pemda_id"`

	TanggalTerjadiRisiko string `json:"tanggal_terjadi_risiko"`
	SebabRisiko          string `json:"sebab_risiko"`
	DampakRisiko         string `json:"dampak_risiko"`
	KeteranganRisiko     string `json:"keterangan_risiko"`
}

type PencatatanKejadianRisikoCreateUseCaseRes struct {
	ID string `json:"id"`
}

type PencatatanKejadianRisikoCreateUseCase = core.ActionHandler[PencatatanKejadianRisikoCreateUseCaseReq, PencatatanKejadianRisikoCreateUseCaseRes]

func ImplPencatatanKejadianRisikoCreateUseCase(
	generateId gateway.GenerateId,
	createPencatatanKejadianRisiko gateway.PencatatanKejadianRisikoSave,
	PenetapanKonteksRisikoStrategisPemdaByID gateway.PenetapanKonteksRisikoStrategisPemdaGetByID,
	IdentifikasiRisikoStrategisPemdaByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
) PencatatanKejadianRisikoCreateUseCase {
	return func(ctx context.Context, req PencatatanKejadianRisikoCreateUseCaseReq) (*PencatatanKejadianRisikoCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		_, err = PenetapanKonteksRisikoStrategisPemdaByID(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetByIDReq{ID: req.PenetapanKonteksRisikoStrategisPemdaID})
		if err != nil {
			return nil, err
		}

		_, err = IdentifikasiRisikoStrategisPemdaByID(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{ID: req.IdentifikasiRisikoStrategisPemdaID})
		if err != nil {
			return nil, err
		}

		obj := model.PencatatanKejadianRisiko{
			ID:                                     &genObj.RandomId,
			PenetapanKonteksRisikoStrategisPemdaID: &req.PenetapanKonteksRisikoStrategisPemdaID,
			IdentifikasiRisikoStrategisPemdaID:     &req.IdentifikasiRisikoStrategisPemdaID,
			TanggalTerjadiRisiko:                   &req.TanggalTerjadiRisiko,
			SebabRisiko:                            &req.SebabRisiko,
			DampakRisiko:                           &req.DampakRisiko,
			KeteranganRisiko:                       &req.KeteranganRisiko,
			Status:                                 sharedModel.StatusMenungguVerifikasi,
		}

		if _, err = createPencatatanKejadianRisiko(ctx, gateway.PencatatanKejadianRisikoSaveReq{PencatatanKejadianRisiko: obj}); err != nil {
			return nil, err
		}

		return &PencatatanKejadianRisikoCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
