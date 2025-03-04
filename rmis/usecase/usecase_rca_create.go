package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/helper"
	sharedModel "shared/model"
)

type RcaCreateUseCaseReq struct {
	PemilikRisiko                      string   `json:"pemilik_risiko"`
	TahunPenilaian                     string   `json:"tahun_penilaian"`
	Why                                []string `json:"why"`
	KegiatanPengendalian               string   `json:"kegiatan_pengendalian"`
	PenyebabRisikoID                   string   `json:"penyebab_risiko_id"`
	IdentifikasiRisikoStrategisPemdaId string   `json:"identifikasi_risiko_strategis_pemda_id"`
}

type RcaCreateUseCaseRes struct {
	ID string `json:"id"`
}

type RcaCreateUseCase = core.ActionHandler[RcaCreateUseCaseReq, RcaCreateUseCaseRes]

func ImplRcaCreateUseCase(
	generateId gateway.GenerateId,
	createRca gateway.RcaSave,
	IdentifikasiRisikoStrategisPemdaGetByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
	PenyebabRisikoGetByID gateway.PenyebabRisikoGetByID,
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

		// Penyebab Risiko
		_, err = PenyebabRisikoGetByID(ctx, gateway.PenyebabRisikoGetByIDReq{ID: req.PenyebabRisikoID})
		if err != nil {
			return nil, fmt.Errorf("error getting penyebab risiko table: %v", err)
		}

		// Identifikasi Risiko Strategis Pemda
		_, err = IdentifikasiRisikoStrategisPemdaGetByID(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{ID: req.IdentifikasiRisikoStrategisPemdaId})
		if err != nil {
			return nil, fmt.Errorf("error getting identifikasi risiko strategis pemda table: %v", err)
		}

		whyJSON := helper.ToDataTypeJSONPtr(req.Why...)
		obj := model.Rca{
			ID:                                 &genObj.RandomId,
			IdentifikasiRisikoStrategisPemdaID: &req.IdentifikasiRisikoStrategisPemdaId,
			PemilikRisiko:                      &req.PemilikRisiko,
			TahunPenilaian:                     &tahunPenilaian,
			Why:                                whyJSON,
			PenyebabRisikoID:                   &req.PenyebabRisikoID,
			KegiatanPengendalian:               &req.KegiatanPengendalian,
			Status:                             sharedModel.StatusMenungguVerifikasi,
		}
		obj.SetAkarPenyebab()

		// Save the RCA entry
		if _, err = createRca(ctx, gateway.RcaSaveReq{Rca: obj}); err != nil {
			return nil, err
		}

		return &RcaCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
