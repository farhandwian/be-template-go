package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type IdentifikasiRisikoOperasionalOPDUpdateUseCaseReq struct {
	ID                                     string `json:"id"`
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

type IdentifikasiRisikoOperasionalOPDUpdateUseCaseRes struct{}

type IdentifikasiRisikoOperasionalOPDUpdateUseCase = core.ActionHandler[IdentifikasiRisikoOperasionalOPDUpdateUseCaseReq, IdentifikasiRisikoOperasionalOPDUpdateUseCaseRes]

func ImplIdentifikasiRisikoOperasionalOPDUpdateUseCase(
	getIdentifikasiRisikoOperasionalOPDById gateway.IdentifikasiRisikoOperasionalOPDGetByID,
	updateIdentifikasiRisikoOperasionalOPD gateway.IdentifikasiRisikoOperasionalOPDSave,
	kodeRisikoByID gateway.KategoriRisikoGetByID,
	RcaByID gateway.RcaGetByID,
	getOneOPD gateway.OPDGetByID,
	penetapanKonteksRisikoOperasionalOpdID gateway.PenetapanKonteksRisikoOperasionalGetByID,
) IdentifikasiRisikoOperasionalOPDUpdateUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDUpdateUseCaseReq) (*IdentifikasiRisikoOperasionalOPDUpdateUseCaseRes, error) {

		res, err := getIdentifikasiRisikoOperasionalOPDById(ctx, gateway.IdentifikasiRisikoOperasionalOPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		identifikasiOperasionalOpd := res.IdentifikasiRisikoOperasionalOPD
		identifikasiOperasionalOpd.TahapRisiko = &req.TahapRisiko
		identifikasiOperasionalOpd.KategoriRisikoID = &req.KategoriRisikoID
		identifikasiOperasionalOpd.UraianRisiko = &req.UraianRisiko
		identifikasiOperasionalOpd.PemilikRisiko = &req.PemilikRisiko
		identifikasiOperasionalOpd.Controllable = &req.Controllable
		identifikasiOperasionalOpd.UraianDampak = &req.UraianDampak
		identifikasiOperasionalOpd.PihakDampak = &req.PihakDampak
		identifikasiOperasionalOpd.Indikatorkeluaran = &req.Indikatorkeluaran
		identifikasiOperasionalOpd.UraianSebab = &req.UraianSebab
		identifikasiOperasionalOpd.SumberSebab = &req.SumberSebab
		identifikasiOperasionalOpd.Kegiatan = &req.Kegiatan
		identifikasiOperasionalOpd.KategoriRisikoID = &req.KategoriRisikoID
		identifikasiOperasionalOpd.RcaID = &req.RcaID
		identifikasiOperasionalOpd.PenetapanKonteksRisikoOperasionalOpdID = &req.PenetapanKonteksRisikoOperasionalOpdID
		identifikasiOperasionalOpd.Status = sharedModel.StatusMenungguVerifikasi

		if identifikasiOperasionalOpd.KategoriRisikoID == nil || *identifikasiOperasionalOpd.KategoriRisikoID == "" {
			return nil, fmt.Errorf("KategoriRisikoID is missing in the database record")
		}

		kategoriRisikoRes, err := kodeRisikoByID(ctx, gateway.KategoriRisikoGetByIDReq{
			ID: *identifikasiOperasionalOpd.KategoriRisikoID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get KategoriRisiko: %v", err)
		}

		penetapanKonteksRisikoOperasionalOpdID, err := penetapanKonteksRisikoOperasionalOpdID(ctx, gateway.PenetapanKonteksRisikoOperasionalGetByIDReq{ID: req.PenetapanKonteksRisikoOperasionalOpdID})
		if err != nil {
			return nil, err
		}

		opdRes, err := getOneOPD(ctx, gateway.OPDGetByIDReq{ID: *penetapanKonteksRisikoOperasionalOpdID.PenetapanKonteksRisikoOperasional.OpdID})
		if err != nil {
			return nil, fmt.Errorf("failed to get OPDName: %v", err)
		}

		if identifikasiOperasionalOpd.KodeRisiko == nil || *identifikasiOperasionalOpd.KodeRisiko == "" {
			identifikasiOperasionalOpd.GenerateKodeRisiko(*penetapanKonteksRisikoOperasionalOpdID.PenetapanKonteksRisikoOperasional.TahunPenilaian, *kategoriRisikoRes.KategoriRisiko.Kode, *opdRes.OPD.Kode)
		}

		// rcaRes, err := RcaByID(ctx, gateway.RcaGetByIDReq{ID: req.RcaID})
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to get RcaName: %v", err)
		// }

		if _, err := updateIdentifikasiRisikoOperasionalOPD(ctx, gateway.IdentifikasiRisikoOperasionalOPDSaveReq{IdentifikasiRisikoOperasionalOPD: identifikasiOperasionalOpd}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoOperasionalOPDUpdateUseCaseRes{}, nil
	}
}
