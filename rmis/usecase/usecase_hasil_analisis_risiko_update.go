package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"shared/core"
)

type HasilAnalisisRisikoUpdateUseCaseReq struct {
	ID                                     string `json:"-"`
	SkalaDampak                            int    `json:"skala_dampak"`
	SkalaKemungkinan                       int    `json:"skala_kemungkinan"`
	SkalaRisiko                            int    `json:"skala_risiko"`
	IdentifikasiRisikoStrategisPemdaID     string `json:"identifikasi_risiko_strategis_pemda_id"`
	PenetapanKonteksRisikoStrategisPemdaID string `json:"penetapan_konteks_risiko_strategis_pemda_id"`
}

type HasilAnalisisRisikoUpdateUseCaseRes struct{}

type HasilAnalisisRisikoUpdateUseCase = core.ActionHandler[HasilAnalisisRisikoUpdateUseCaseReq, HasilAnalisisRisikoUpdateUseCaseRes]

func ImplHasilAnalisisRisikoUpdateUseCase(
	getHasilAnalisisRisikoById gateway.HasilAnalisisRisikoGetByID,
	updateHasilAnalisisRisiko gateway.HasilAnalisisRisikoSave,
	indeksPeringkatPrioritasByID gateway.IndeksPeringkatPrioritasGetByID,
	IndeksPeringkatPrioritasCreate gateway.IndeksPeringkatPrioritasSave,
	IdentifikasiRisikoStrategisPemdaByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
	PenetapanKonteksRisikoStrategisPemdaByID gateway.PenetapanKonteksRisikoOperasionalGetByID,
) HasilAnalisisRisikoUpdateUseCase {
	return func(ctx context.Context, req HasilAnalisisRisikoUpdateUseCaseReq) (*HasilAnalisisRisikoUpdateUseCaseRes, error) {

		res, err := getHasilAnalisisRisikoById(ctx, gateway.HasilAnalisisRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		hasilAnalisisRisiko := res.HasilAnalisisRisiko
		identifikasiRisikoStrategisPemdaRes, err := IdentifikasiRisikoStrategisPemdaByID(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{ID: req.IdentifikasiRisikoStrategisPemdaID})
		if err != nil {
			return nil, fmt.Errorf("error getting Identifikasi Risiko Strategis Pemda table: %v", err)
		}

		_, err = PenetapanKonteksRisikoStrategisPemdaByID(ctx, gateway.PenetapanKonteksRisikoOperasionalGetByIDReq{ID: req.PenetapanKonteksRisikoStrategisPemdaID})
		if err != nil {
			return nil, fmt.Errorf("error getting Penetapan Konteks Risiko Strategis Pemda table: %v", err)
		}

		hasilAnalisisRisiko.SkalaDampak = &req.SkalaDampak
		hasilAnalisisRisiko.SkalaKemungkinan = &req.SkalaKemungkinan
		hasilAnalisisRisiko.IdentifikasiRisikoStrategisPemdaID = &req.IdentifikasiRisikoStrategisPemdaID
		hasilAnalisisRisiko.PenetapanKonteksRisikoStrategisPemdaID = &req.PenetapanKonteksRisikoStrategisPemdaID
		hasilAnalisisRisiko.SetSkalaRisiko()
		if _, err := updateHasilAnalisisRisiko(ctx, gateway.HasilAnalisisRisikoSaveReq{HasilAnalisisRisiko: hasilAnalisisRisiko}); err != nil {
			return nil, err
		}

		indeksPeringkatPrioritasByIDRes, err := indeksPeringkatPrioritasByID(ctx, gateway.IndeksPeringkatPrioritasGetByIDReq{ID: *hasilAnalisisRisiko.ID})
		if err != nil {
			return nil, err
		}
		indeksPeringkatPrioritas := indeksPeringkatPrioritasByIDRes.IndeksPeringkatPrioritas

		// indeksPeringkatPrioritas.SetToleransiRisiko(*identifikasiRisikoStrategisPemdaRes.IdentifikasiRisikoStrategisPemda.KategoriRisiko.Kode)
		indeksPeringkatPrioritas.SetMitigasi(*hasilAnalisisRisiko.SkalaRisiko)
		indeksPeringkatPrioritas.SetIntermediateRank(*hasilAnalisisRisiko.SkalaRisiko, *identifikasiRisikoStrategisPemdaRes.IdentifikasiRisikoStrategisPemda.NomorUraian)
		if _, err = IndeksPeringkatPrioritasCreate(ctx, gateway.IndeksPeringkatPrioritasSaveReq{IndeksPeringkatPrioritas: indeksPeringkatPrioritas}); err != nil {
			return nil, err
		}

		return &HasilAnalisisRisikoUpdateUseCaseRes{}, nil
	}
}
