package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	sharedModel "shared/model"
)

type PenilaianKegiatanPengendalianCreateUseCaseReq struct {
	NamaPemda                     string `json:"nama_pemda"`
	TahunPenilaian                string `json:"tahun_penilaian"`
	SpipID                        string `json:"spip_id"`
	KondisiLingkunganPengendalian string `json:"kondisi_lingkungan_pengendalian"`
	RencanaTindakPerbaikan        string `json:"rencana_tindak_perbaikan"`
	PenanggungJawab               string `json:"penanggung_jawab"`
	TargetWaktuPenyelesaian       string `json:"target_waktu_penyelesaian"`
}

type PenilaianKegiatanPengendalianCreateUseCaseRes struct {
	ID string `json:"id"`
}

type PenilaianKegiatanPengendalianCreateUseCase = core.ActionHandler[PenilaianKegiatanPengendalianCreateUseCaseReq, PenilaianKegiatanPengendalianCreateUseCaseRes]

func ImplPenilaianKegiatanPengendalianCreateUseCase(
	generateId gateway.GenerateId,
	createPenilaianKegiatanPengendalian gateway.PenilaianKegiatanPengendalianSave,
	SpipById gateway.SpipGetByID,
) PenilaianKegiatanPengendalianCreateUseCase {
	return func(ctx context.Context, req PenilaianKegiatanPengendalianCreateUseCaseReq) (*PenilaianKegiatanPengendalianCreateUseCaseRes, error) {

		// Generate a unique ID
		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		_, err = SpipById(ctx, gateway.SpipGetByIDReq{ID: req.SpipID})
		if err != nil {
			return nil, fmt.Errorf("error getting SPIP Table: %v", err)
		}
		obj := model.PenilaianKegiatanPengendalian{
			ID:                            &genObj.RandomId,
			NamaPemda:                     &req.NamaPemda,
			TahunPenilaian:                &req.TahunPenilaian,
			SpipId:                        &req.SpipID,
			KondisiLingkunganPengendalian: &req.KondisiLingkunganPengendalian,
			RencanaTindakPerbaikan:        &req.RencanaTindakPerbaikan,
			PenanggungJawab:               &req.PenanggungJawab,
			Status:                        sharedModel.StatusMenungguVerifikasi,
			TargetWaktuPenyelesaian:       &req.TargetWaktuPenyelesaian,
		}

		// Save the PenilaianKegiatanPengendalian entry
		if _, err = createPenilaianKegiatanPengendalian(ctx, gateway.PenilaianKegiatanPengendalianSaveReq{PenilaianKegiatanPengendalian: obj}); err != nil {
			return nil, err
		}

		return &PenilaianKegiatanPengendalianCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
