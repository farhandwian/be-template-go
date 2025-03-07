package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	sharedModel "shared/model"
)

type HasilAnalisisRisikoCreateUseCaseReq struct {
	SkalaDampak      int                    `json:"skala_dampak"`
	SkalaKemungkinan int                    `json:"skala_kemungkinan"`
	TipeIdentifikasi model.TipeIdentifikasi `json:"tipe_identifikasi"` // Can be "strategis_pemda", "operasional_opd", "strategis_renstra_opd"
	IdentifikasiID   string                 `json:"identifikasi_id"`

	TipePenetapanKonteks model.TipePenetapanKonteks `json:"tipe_penetapan_konteks"` // Can be "strategis_pemda", "operasional", "strategis_renstra_opd"
	PenetapanKonteksID   string                     `json:"penetapan_konteks_id"`
}

type HasilAnalisisRisikoCreateUseCaseRes struct {
	ID string `json:"id"`
}

type HasilAnalisisRisikoCreateUseCase = core.ActionHandler[HasilAnalisisRisikoCreateUseCaseReq, HasilAnalisisRisikoCreateUseCaseRes]

func ImplHasilAnalisisRisikoCreateUseCase(
	generateId gateway.GenerateId,
	createHasilAnalisisRisiko gateway.HasilAnalisisRisikoSave,
	IdentifikasiRisikoStrategisPemdaByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
	IdentifikasiRisikoOperasionalOPDByID gateway.IdentifikasiRisikoOperasionalOPDGetByID,
	IdentifikasiRisikoStrategisOPDByID gateway.IdentifikasiRisikoStrategisOPDGetByID,
	PenetapanKonteksRisikoStrategisPemdaByID gateway.PenetapanKonteksRisikoStrategisPemdaGetByID,
	PenetapanKonteksRisikoOperasionalByID gateway.PenetapanKonteksRisikoOperasionalGetByID,
	PenetapanKonteksRisikoStrategisRenstraOPDByID gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByID,
	IndeksPeringkatPrioritasCreate gateway.IndeksPeringkatPrioritasSave,
	KategoriRisikoByID gateway.KategoriRisikoGetByID,
) HasilAnalisisRisikoCreateUseCase {
	return func(ctx context.Context, req HasilAnalisisRisikoCreateUseCaseReq) (*HasilAnalisisRisikoCreateUseCaseRes, error) {

		// Generate ID
		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
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

		// Create Object
		obj := model.HasilAnalisisRisiko{
			ID:                   &genObj.RandomId,
			TipeIdentifikasi:     (*string)(&req.TipeIdentifikasi),
			IdentifikasiID:       &req.IdentifikasiID,
			TipePenetapanKonteks: (*string)(&req.TipePenetapanKonteks),
			PenetapanKonteksID:   &req.PenetapanKonteksID,
			SkalaDampak:          &req.SkalaDampak,
			SkalaKemungkinan:     &req.SkalaKemungkinan,
			Status:               sharedModel.StatusMenungguVerifikasi,
		}

		obj.SetSkalaRisiko()

		// Save to Database
		if _, err = createHasilAnalisisRisiko(ctx, gateway.HasilAnalisisRisikoSaveReq{HasilAnalisisRisiko: obj}); err != nil {
			return nil, err
		}

		kategoriRisikoRes, err := KategoriRisikoByID(ctx, gateway.KategoriRisikoGetByIDReq{ID: *kategoriRisikoID})
		if err != nil {
			return nil, err
		}

		// Generate ID for Indeks Peringkat Prioritas
		genObjIndeksPeringkat, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		objIndeksPeringkat := model.IndeksPeringkatPrioritas{
			ID:                    &genObjIndeksPeringkat.RandomId,
			HasilAnalisisRisikoID: obj.ID,
		}

		// Set Indeks Peringkat Prioritas
		objIndeksPeringkat.SetToleransiRisiko(*kategoriRisikoRes.KategoriRisiko.Nama)
		objIndeksPeringkat.SetMitigasi(*obj.SkalaRisiko)

		fmt.Sprintf("TEST", nomorUraian)
		objIndeksPeringkat.SetIntermediateRank(*obj.SkalaRisiko, 1)

		// LATER IF TEST PASS ACTIVE THIS
		// objIndeksPeringkat.SetIntermediateRank(*obj.SkalaRisiko, *nomorUraian)

		// Save Indeks Peringkat Prioritas
		if _, err = IndeksPeringkatPrioritasCreate(ctx, gateway.IndeksPeringkatPrioritasSaveReq{IndeksPeringkatPrioritas: objIndeksPeringkat}); err != nil {
			return nil, err
		}

		return &HasilAnalisisRisikoCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
