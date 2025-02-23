package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"

	"gorm.io/datatypes"
)

type RcaCreateUseCaseReq struct {
	NamaUnitPemilikRisiko string   `json:"nama_unit_pemilik_risiko"`
	TahunPenilaian        string   `json:"tahun_penilaian"`
	PernyataanRisiko      string   `json:"pernyataan_risiko"`
	Why                   []string `json:"why"`
	JenisPenyebab         string   `json:"jenis_penyebab"`
	KegiatanPengendalian  string   `json:"kegiatan_pengendalian"`
}

type RcaCreateUseCaseRes struct {
	ID string `json:"id"`
}

type RcaCreateUseCase = core.ActionHandler[RcaCreateUseCaseReq, RcaCreateUseCaseRes]

func ImplRcaCreateUseCase(
	generateId gateway.GenerateId,
	createRca gateway.RcaSave,
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

		// Convert Why ([]string) to JSON
		whyJSON, err := json.Marshal(req.Why)
		if err != nil {
			return nil, err
		}
		why := datatypes.JSON(whyJSON)

		// Construct the Rca object (Fully Initialized)
		obj := model.Rca{
			ID:                    &genObj.RandomId,
			NamaUnitPemilikRisiko: &req.NamaUnitPemilikRisiko,
			TahunPenilaian:        &tahunPenilaian, // Only store year
			PenyebabRisiko:        &req.PernyataanRisiko,
			Why:                   &why, // JSON formatted data
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
