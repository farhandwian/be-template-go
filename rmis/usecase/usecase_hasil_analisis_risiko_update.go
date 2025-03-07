package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	sharedModel "shared/model"
)

type HasilAnalisisRisikoUpdateUseCaseReq struct {
	ID               string `json:"-"`
	SkalaDampak      int    `json:"skala_dampak"`
	SkalaKemungkinan int    `json:"skala_kemungkinan"`

	TipeIdentifikasi model.TipeIdentifikasi `json:"tipe_identifikasi"` // Can be "strategis_pemda", "operasional_opd", "strategis_renstra_opd"
	IdentifikasiID   string                 `json:"identifikasi_id"`

	TipePenetapanKonteks model.TipePenetapanKonteks `json:"tipe_penetapan_konteks"` // Can be "strategis_pemda", "operasional", "strategis_renstra_opd"
	PenetapanKonteksID   string                     `json:"penetapan_konteks_id"`
}

type HasilAnalisisRisikoUpdateUseCaseRes struct{}

type HasilAnalisisRisikoUpdateUseCase = core.ActionHandler[HasilAnalisisRisikoUpdateUseCaseReq, HasilAnalisisRisikoUpdateUseCaseRes]

func ImplHasilAnalisisRisikoUpdateUseCase(
	getHasilAnalisisRisikoById gateway.HasilAnalisisRisikoGetByID,
	updateHasilAnalisisRisiko gateway.HasilAnalisisRisikoSave,
	indeksPeringkatPrioritasByID gateway.IndeksPeringkatPrioritasGetByID,
	IndeksPeringkatPrioritasCreate gateway.IndeksPeringkatPrioritasSave,
	IdentifikasiRisikoStrategisPemdaByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
	IdentifikasiRisikoOperasionalOPDByID gateway.IdentifikasiRisikoOperasionalOPDGetByID,
	IdentifikasiRisikoStrategisOPDByID gateway.IdentifikasiRisikoStrategisOPDGetByID,
	PenetapanKonteksRisikoStrategisPemdaByID gateway.PenetapanKonteksRisikoStrategisPemdaGetByID,
	PenetapanKonteksRisikoOperasionalByID gateway.PenetapanKonteksRisikoOperasionalGetByID,
	PenetapanKonteksRisikoStrategisRenstraOPDByID gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByID,
	KategoriRisikoByID gateway.KategoriRisikoGetByID,
) HasilAnalisisRisikoUpdateUseCase {
	return func(ctx context.Context, req HasilAnalisisRisikoUpdateUseCaseReq) (*HasilAnalisisRisikoUpdateUseCaseRes, error) {

		res, err := getHasilAnalisisRisikoById(ctx, gateway.HasilAnalisisRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		var kategoriRisikoID *string
		var nomorUraian *int

		// Fetch Identifikasi Risiko Data
		switch req.TipeIdentifikasi {
		case model.TipeIdentifikasiStrategisPemda:
			identifikasiRisikoRes, err := IdentifikasiRisikoStrategisPemdaByID(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{ID: req.IdentifikasiID})
			if err != nil {
				return nil, err
			}
			kategoriRisikoID = identifikasiRisikoRes.IdentifikasiRisikoStrategisPemda.KategoriRisikoID
			nomorUraian = identifikasiRisikoRes.IdentifikasiRisikoStrategisPemda.NomorUraian
		case model.TipeIdentifikasiOperasional:
			identifikasiRisikoRes, err := IdentifikasiRisikoOperasionalOPDByID(ctx, gateway.IdentifikasiRisikoOperasionalOPDGetByIDReq{ID: req.IdentifikasiID})
			if err != nil {
				return nil, err
			}
			kategoriRisikoID = identifikasiRisikoRes.IdentifikasiRisikoOperasionalOPD.KategoriRisikoID
			nomorUraian = identifikasiRisikoRes.IdentifikasiRisikoOperasionalOPD.NomorUraian
		case model.TipeIdentifikasiStrategisOPD:
			identifikasiRisikoRes, err := IdentifikasiRisikoStrategisOPDByID(ctx, gateway.IdentifikasiRisikoStrategisOPDGetByIDReq{ID: req.IdentifikasiID})
			if err != nil {
				return nil, err
			}
			kategoriRisikoID = identifikasiRisikoRes.IdentifikasiRisikoStrategisOPD.KategoriRisikoID
			nomorUraian = identifikasiRisikoRes.IdentifikasiRisikoStrategisOPD.NomorUraian
		default:
			return nil, fmt.Errorf("invalid tipe_identifikasi: %s", req.TipeIdentifikasi)
		}

		// Fetch Penetapan Konteks Data
		switch req.TipePenetapanKonteks {
		case model.TipePenetapanKonteksStrategisPemda:
			_, err := PenetapanKonteksRisikoStrategisPemdaByID(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetByIDReq{ID: req.PenetapanKonteksID})
			if err != nil {
				return nil, err
			}
		case model.TipePenetapanKonteksOperasional:
			_, err := PenetapanKonteksRisikoOperasionalByID(ctx, gateway.PenetapanKonteksRisikoOperasionalGetByIDReq{ID: req.PenetapanKonteksID})
			if err != nil {
				return nil, err
			}
		case model.TipePenetapanKonteksStrategisRenstraOPD:
			_, err := PenetapanKonteksRisikoStrategisRenstraOPDByID(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByIDReq{ID: req.PenetapanKonteksID})
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("invalid tipe_penetapan_konteks: %s", req.TipePenetapanKonteks)
		}
		hasilAnalisisRisiko := res.HasilAnalisisRisiko

		hasilAnalisisRisiko.SkalaDampak = &req.SkalaDampak
		hasilAnalisisRisiko.SkalaKemungkinan = &req.SkalaKemungkinan
		hasilAnalisisRisiko.IdentifikasiID = &req.IdentifikasiID
		hasilAnalisisRisiko.PenetapanKonteksID = &req.PenetapanKonteksID
		hasilAnalisisRisiko.Status = sharedModel.StatusMenungguVerifikasi
		hasilAnalisisRisiko.SetSkalaRisiko()

		if _, err := updateHasilAnalisisRisiko(ctx, gateway.HasilAnalisisRisikoSaveReq{HasilAnalisisRisiko: hasilAnalisisRisiko}); err != nil {
			return nil, err
		}

		kategoriRisikoRes, err := KategoriRisikoByID(ctx, gateway.KategoriRisikoGetByIDReq{ID: *kategoriRisikoID})
		if err != nil {
			return nil, err
		}

		indeksPeringkatPrioritasByIDRes, err := indeksPeringkatPrioritasByID(ctx, gateway.IndeksPeringkatPrioritasGetByIDReq{ID: *hasilAnalisisRisiko.ID})
		if err != nil {
			return nil, err
		}
		indeksPeringkatPrioritas := indeksPeringkatPrioritasByIDRes.IndeksPeringkatPrioritas

		indeksPeringkatPrioritas.SetToleransiRisiko(*kategoriRisikoRes.KategoriRisiko.Kode)
		indeksPeringkatPrioritas.SetMitigasi(*hasilAnalisisRisiko.SkalaRisiko)

		fmt.Sprintf("TEST", nomorUraian)
		indeksPeringkatPrioritas.SetIntermediateRank(*hasilAnalisisRisiko.SkalaRisiko, 1)
		// indeksPeringkatPrioritas.SetIntermediateRank(*hasilAnalisisRisiko.SkalaRisiko, *nomorUraian)
		if _, err = IndeksPeringkatPrioritasCreate(ctx, gateway.IndeksPeringkatPrioritasSaveReq{IndeksPeringkatPrioritas: indeksPeringkatPrioritas}); err != nil {
			return nil, err
		}

		return &HasilAnalisisRisikoUpdateUseCaseRes{}, nil
	}
}
