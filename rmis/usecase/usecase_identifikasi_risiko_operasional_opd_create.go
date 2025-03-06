package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type IdentifikasiRisikoOperasionalOPDCreateUseCaseReq struct {
	PenetapanKonteksRisikoOperasionalOpdID string `json:"penetapan_konteks_risiko_operasional_opd_id"`
	KategoriRisikoID                       string `json:"kategori_risiko_id"`
	RcaID                                  string `json:"rca_id"`
	UraianSebab                            string `json:"uraian_sebab"`
	SumberSebab                            string `json:"sumber_sebab"`
	Controllable                           string `json:"controllable"`
	UraianDampak                           string `json:"uraian_dampak"`
	PihakDampak                            string `json:"pihak_dampak"`
	Kegiatan                               string `json:"kegiatan"`
	Indikatorkeluaran                      string `json:"indikator_keluaran"`
	PemilikRisiko                          string `json:"pemilik_resiko"`
	UraianRisiko                           string `json:"uraian_resiko"`
	TahapRisiko                            string `json:"tahap_risiko"`
}

type IdentifikasiRisikoOperasionalOPDCreateUseCaseRes struct {
	ID string `json:"id"`
}

type IdentifikasiRisikoOperasionalOPDCreateUseCase = core.ActionHandler[
	IdentifikasiRisikoOperasionalOPDCreateUseCaseReq,
	IdentifikasiRisikoOperasionalOPDCreateUseCaseRes,
]

func ImplIdentifikasiRisikoOperasionalOPDCreateUseCase(
	generateId gateway.GenerateId,
	createIdentifikasiRisikoOperasionalOPD gateway.IdentifikasiRisikoOperasionalOPDSave,
	kodeRisikoByID gateway.KategoriRisikoGetByID,
	getOneOPD gateway.OPDGetByID,
	penetapanKonteksRisikoOperasionalOpdID gateway.PenetapanKonteksRisikoOperasionalGetByID,
) IdentifikasiRisikoOperasionalOPDCreateUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDCreateUseCaseReq) (*IdentifikasiRisikoOperasionalOPDCreateUseCaseRes, error) {
		// Generate a unique ID
		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		kategoriRisikoRes, err := kodeRisikoByID(ctx, gateway.KategoriRisikoGetByIDReq{ID: req.KategoriRisikoID})
		if err != nil {
			return nil, err
		}

		penetapanKonteksRisikoOperasionalOpdID, err := penetapanKonteksRisikoOperasionalOpdID(ctx, gateway.PenetapanKonteksRisikoOperasionalGetByIDReq{ID: req.PenetapanKonteksRisikoOperasionalOpdID})
		if err != nil {
			return nil, err
		}

		opdRes, err := getOneOPD(ctx, gateway.OPDGetByIDReq{ID: *penetapanKonteksRisikoOperasionalOpdID.PenetapanKonteksRisikoOperasional.OpdID})
		if err != nil {
			return nil, err
		}

		obj := model.IdentifikasiRisikoOperasionalOPD{
			ID:                &genObj.RandomId,
			KategoriRisikoID:  &req.KategoriRisikoID,
			UraianRisiko:      &req.UraianRisiko,
			TahapRisiko:       &req.TahapRisiko,
			PemilikRisiko:     &req.PemilikRisiko,
			Controllable:      &req.Controllable,
			UraianDampak:      &req.UraianDampak,
			PihakDampak:       &req.PihakDampak,
			UraianSebab:       &req.UraianSebab,
			SumberSebab:       &req.SumberSebab,
			Kegiatan:          &req.Kegiatan,
			Indikatorkeluaran: &req.Indikatorkeluaran,
		}

		obj.GenerateKodeRisiko(
			*penetapanKonteksRisikoOperasionalOpdID.PenetapanKonteksRisikoOperasional.TahunPenilaian, *kategoriRisikoRes.KategoriRisiko.Kode, *opdRes.OPD.Kode,
		)
		// Save the new entry
		if _, err = createIdentifikasiRisikoOperasionalOPD(ctx, gateway.IdentifikasiRisikoOperasionalOPDSaveReq{
			IdentifikasiRisikoOperasionalOPD: obj,
		}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoOperasionalOPDCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
