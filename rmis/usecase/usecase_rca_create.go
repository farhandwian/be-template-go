package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/helper"
)

type RcaCreateUseCaseReq struct {
	NamaUnitPemilikRisiko              string   `json:"nama_unit_pemilik_risiko"`
	TahunPenilaian                     string   `json:"tahun_penilaian"`
	IdentifikasiRisikoStrategisPemdaId *string  `json:"identifikasi_risiko_strategis_pemda_id"`
	Why                                []string `json:"why"`
	JenisPenyebab                      string   `json:"jenis_penyebab"`
	KegiatanPengendalian               string   `json:"kegiatan_pengendalian"`
}

type RcaCreateUseCaseRes struct {
	ID string `json:"id"`
}

type RcaCreateUseCase = core.ActionHandler[RcaCreateUseCaseReq, RcaCreateUseCaseRes]

func ImplRcaCreateUseCase(
	generateId gateway.GenerateId,
	createRca gateway.RcaSave,
	IdentifikasiRisikoStrategisPemdaGetByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
) RcaCreateUseCase {
	return func(ctx context.Context, req RcaCreateUseCaseReq) (*RcaCreateUseCaseRes, error) {

		// Generate a unique ID
		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		tahunPenilaian, err := extractYear(req.TahunPenilaian)
		if err != nil {
			return nil, fmt.Errorf("invalid TahunPenilaian format: %v", err)
		}

		// Identifikasi Risiko Strategis Pemda
		identifikasiRisikoStrategisPemdaRes, err := IdentifikasiRisikoStrategisPemdaGetByID(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{ID: *req.IdentifikasiRisikoStrategisPemdaId})
		if err != nil {
			return nil, fmt.Errorf("error getting identifikasi risiko strategis pemda table")
		}

		whyJSON := helper.ToDataTypeJSONPtr(req.Why...)
		obj := model.Rca{
			ID:                    &genObj.RandomId,
			NamaUnitPemilikRisiko: &req.NamaUnitPemilikRisiko,
			TahunPenilaian:        &tahunPenilaian,
			PernyataanRisiko:      identifikasiRisikoStrategisPemdaRes.IdentifikasiRisikoStrategisPemda.UraianRisiko,
			Why:                   whyJSON,
			JenisPenyebab:         &req.JenisPenyebab,
			KegiatanPengendalian:  &req.KegiatanPengendalian,
		}

		// Save the RCA entry
		if _, err = createRca(ctx, gateway.RcaSaveReq{Rca: obj}); err != nil {
			return nil, err
		}

		return &RcaCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
