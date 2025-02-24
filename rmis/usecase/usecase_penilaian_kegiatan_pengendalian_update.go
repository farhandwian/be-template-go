package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"shared/core"
)

type PenilaianKegiatanPengendalianUpdateUseCaseReq struct {
	ID                            string `json:"id"`
	NamaPemda                     string `json:"nama_pemda"`
	TahunPenilaian                string `json:"tahun_penilaian"`
	SpipID                        string `json:"spip_id"`
	KondisiLingkunganPengendalian string `json:"kondisi_lingkungan_pengendalian"`
	RencanaTindakPerbaikan        string `json:"rencana_tindak_perbaikan"`
	PenanggungJawab               string `json:"penanggung_jawab"`
	TargetWaktuPenyelesaian       string `json:"target_waktu_penyelesaian"`
}

type PenilaianKegiatanPengendalianUpdateUseCaseRes struct{}

type PenilaianKegiatanPengendalianUpdateUseCase = core.ActionHandler[PenilaianKegiatanPengendalianUpdateUseCaseReq, PenilaianKegiatanPengendalianUpdateUseCaseRes]

func ImplPenilaianKegiatanPengendalianUpdateUseCase(
	getPenilaianKegiatanPengendalianById gateway.PenilaianKegiatanPengendalianGetByID,
	updatePenilaianKegiatanPengendalian gateway.PenilaianKegiatanPengendalianSave,
	SpipById gateway.SpipGetByID,
) PenilaianKegiatanPengendalianUpdateUseCase {
	return func(ctx context.Context, req PenilaianKegiatanPengendalianUpdateUseCaseReq) (*PenilaianKegiatanPengendalianUpdateUseCaseRes, error) {

		res, err := getPenilaianKegiatanPengendalianById(ctx, gateway.PenilaianKegiatanPengendalianGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		penilaian := res.PenilaianKegiatanPengendalian
		penilaian.NamaPemda = &req.NamaPemda
		penilaian.TahunPenilaian = &req.TahunPenilaian
		penilaian.KondisiLingkunganPengendalian = &req.KondisiLingkunganPengendalian
		penilaian.RencanaTindakPerbaikan = &req.RencanaTindakPerbaikan
		penilaian.PenanggungJawab = &req.PenanggungJawab
		penilaian.TargetWaktuPenyelesaian = &req.TargetWaktuPenyelesaian

		SpipRes, err := SpipById(ctx, gateway.SpipGetByIDReq{ID: req.SpipID})
		if err != nil {
			return nil, fmt.Errorf("error getting SPIP Table: %v", err)
		}
		penilaian.SPIPId = &req.SpipID
		penilaian.SPIPName = SpipRes.SPIP.Nama
		if _, err := updatePenilaianKegiatanPengendalian(ctx, gateway.PenilaianKegiatanPengendalianSaveReq{PenilaianKegiatanPengendalian: penilaian}); err != nil {
			return nil, err
		}

		return &PenilaianKegiatanPengendalianUpdateUseCaseRes{}, nil
	}
}
