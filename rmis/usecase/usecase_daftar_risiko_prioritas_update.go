package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	sharedModel "shared/model"
)

type DaftarRisikoPrioritasUpdateUseCaseReq struct {
	ID               string `json:"id"`
	TipeIdentifikasi string `json:"tipe_identifikasi" ` // "strategis_pemda", "operasional_opd", or "strategis_renstra_opd"
	IdentifikasiID   string `json:"identifikasi_id" `

	TipePenetapanKonteks string `json:"tipe_penetapan_konteks" ` // "strategis_pemda", "operasional", or "strategis_renstra_opd"
	PenetapanKonteksID   string `json:"penetapan_konteks_id" gorm:"type:VARCHAR(255)"`
}

type DaftarRisikoPrioritasUpdateUseCaseRes struct{}

type DaftarRisikoPrioritasUpdateUseCase = core.ActionHandler[DaftarRisikoPrioritasUpdateUseCaseReq, DaftarRisikoPrioritasUpdateUseCaseRes]

func ImplDaftarRisikoPrioritasUpdateUseCase(
	getDaftarRisikoPrioritasById gateway.DaftarRisikoPrioritasGetByID,
	updateDaftarRisikoPrioritas gateway.DaftarRisikoPrioritasSave,
	IdentifikasiRisikoStrategisPemdaByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
	IdentifikasiRisikoOperasionalOPDByID gateway.IdentifikasiRisikoOperasionalOPDGetByID,
	IdentifikasiRisikoStrategisOPDByID gateway.IdentifikasiRisikoStrategisOPDGetByID,
	PenetapanKonteksRisikoStrategisPemdaByID gateway.PenetapanKonteksRisikoStrategisPemdaGetByID,
	PenetapanKonteksRisikoOperasionalByID gateway.PenetapanKonteksRisikoOperasionalGetByID,
	PenetapanKonteksRisikoStrategisRenstraOPDByID gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByID,
) DaftarRisikoPrioritasUpdateUseCase {
	return func(ctx context.Context, req DaftarRisikoPrioritasUpdateUseCaseReq) (*DaftarRisikoPrioritasUpdateUseCaseRes, error) {

		res, err := getDaftarRisikoPrioritasById(ctx, gateway.DaftarRisikoPrioritasGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		switch req.TipeIdentifikasi {
		case string(model.TipeIdentifikasiStrategisPemda):
			_, err = IdentifikasiRisikoStrategisPemdaByID(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{ID: req.IdentifikasiID})
		case string(model.TipeIdentifikasiOperasional):
			_, err = IdentifikasiRisikoOperasionalOPDByID(ctx, gateway.IdentifikasiRisikoOperasionalOPDGetByIDReq{ID: req.IdentifikasiID})
		case string(model.TipeIdentifikasiStrategisOPD):
			_, err = IdentifikasiRisikoStrategisOPDByID(ctx, gateway.IdentifikasiRisikoStrategisOPDGetByIDReq{ID: req.IdentifikasiID})
		default:
			return nil, fmt.Errorf("invalid TipeIdentifikasi: %v", req.TipeIdentifikasi)
		}
		if err != nil {
			return nil, fmt.Errorf("error getting Identifikasi Risiko table: %v", err)
		}

		switch req.TipePenetapanKonteks {
		case string(model.TipeIdentifikasiStrategisPemda):
			_, err = PenetapanKonteksRisikoStrategisPemdaByID(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetByIDReq{ID: req.PenetapanKonteksID})
		case string(model.TipePenetapanKonteksOperasional):
			_, err = PenetapanKonteksRisikoOperasionalByID(ctx, gateway.PenetapanKonteksRisikoOperasionalGetByIDReq{ID: req.PenetapanKonteksID})
		case string(model.TipePenetapanKonteksStrategisRenstraOPD):
			_, err = PenetapanKonteksRisikoStrategisRenstraOPDByID(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByIDReq{ID: req.PenetapanKonteksID})
		default:
			return nil, fmt.Errorf("invalid TipePenetapanKonteks: %v", req.TipePenetapanKonteks)
		}

		daftarRisikoPrioritas := res.DaftarRisikoPrioritas
		daftarRisikoPrioritas.Status = sharedModel.StatusMenungguVerifikasi
		daftarRisikoPrioritas.TipeIdentifikasi = &req.TipeIdentifikasi
		daftarRisikoPrioritas.IdentifikasiID = &req.IdentifikasiID
		daftarRisikoPrioritas.TipePenetapanKonteks = &req.TipePenetapanKonteks
		daftarRisikoPrioritas.PenetapanKonteksID = &req.PenetapanKonteksID

		if _, err := updateDaftarRisikoPrioritas(ctx, gateway.DaftarRisikoPrioritasSaveReq{DaftarRisikoPrioritas: daftarRisikoPrioritas}); err != nil {
			return nil, err
		}

		return &DaftarRisikoPrioritasUpdateUseCaseRes{}, nil
	}
}
