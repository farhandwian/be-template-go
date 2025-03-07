package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	sharedModel "shared/model"
)

type DaftarRisikoPrioritasCreateUseCaseReq struct {
	TipeIdentifikasi string `json:"tipe_identifikasi" ` // "strategis_pemda", "operasional_opd", or "strategis_renstra_opd"
	IdentifikasiID   string `json:"identifikasi_id" `

	TipePenetapanKonteks string `json:"tipe_penetapan_konteks" ` // "strategis_pemda", "operasional", or "strategis_renstra_opd"
	PenetapanKonteksID   string `json:"penetapan_konteks_id" gorm:"type:VARCHAR(255)"`
}

type DaftarRisikoPrioritasCreateUseCaseRes struct {
	ID string `json:"id"`
}

type DaftarRisikoPrioritasCreateUseCase = core.ActionHandler[DaftarRisikoPrioritasCreateUseCaseReq, DaftarRisikoPrioritasCreateUseCaseRes]

func ImplDaftarRisikoPrioritasCreateUseCase(
	generateId gateway.GenerateId,
	createDaftarRisikoPrioritas gateway.DaftarRisikoPrioritasSave,
	IdentifikasiRisikoStrategisPemdaByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
	IdentifikasiRisikoOperasionalOPDByID gateway.IdentifikasiRisikoOperasionalOPDGetByID,
	IdentifikasiRisikoStrategisOPDByID gateway.IdentifikasiRisikoStrategisOPDGetByID,
	PenetapanKonteksRisikoStrategisPemdaByID gateway.PenetapanKonteksRisikoStrategisPemdaGetByID,
	PenetapanKonteksRisikoOperasionalByID gateway.PenetapanKonteksRisikoOperasionalGetByID,
	PenetapanKonteksRisikoStrategisRenstraOPDByID gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByID,
) DaftarRisikoPrioritasCreateUseCase {
	return func(ctx context.Context, req DaftarRisikoPrioritasCreateUseCaseReq) (*DaftarRisikoPrioritasCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
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
			return nil, err
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

		if err != nil {
			return nil, err
		}
		obj := model.DaftarRisikoPrioritas{
			ID:                   &genObj.RandomId,
			TipeIdentifikasi:     &req.TipeIdentifikasi,
			TipePenetapanKonteks: &req.TipePenetapanKonteks,
			IdentifikasiID:       &req.IdentifikasiID,
			PenetapanKonteksID:   &req.PenetapanKonteksID,

			Status: sharedModel.StatusMenungguVerifikasi,
		}

		if _, err = createDaftarRisikoPrioritas(ctx, gateway.DaftarRisikoPrioritasSaveReq{DaftarRisikoPrioritas: obj}); err != nil {
			return nil, err
		}

		return &DaftarRisikoPrioritasCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
