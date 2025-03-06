package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetapanKonteksRisikoStrategisRenstraOPDGetByIDReq struct {
	ID string
}

type PenetapanKonteksRisikoStrategisRenstraOPDGetByIDRes struct {
	PenetapanKonteksRisikoStrategisRenstraOPD model.PenetapanKonteksRisikoStrategisRenstraOPD
}

type PenetapanKonteksRisikoStrategisRenstraOPDGetByID = core.ActionHandler[PenetapanKonteksRisikoStrategisRenstraOPDGetByIDReq, PenetapanKonteksRisikoStrategisRenstraOPDGetByIDRes]

func ImplPenetapanKonteksRisikoStrategisRenstraOPDGetByID(db *gorm.DB) PenetapanKonteksRisikoStrategisRenstraOPDGetByID {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisRenstraOPDGetByIDReq) (*PenetapanKonteksRisikoStrategisRenstraOPDGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var PenetapanKonteksRisikoStrategisRenstraOPD model.PenetapanKonteksRisikoStrategisRenstraOPD
		if err := query.
			Joins("LEFT JOIN opds ON penetapan_konteks_risiko_strategis_renstra_opds.opd_id = opds.id").
			Where("penetapan_konteks_risiko_strategis_renstra_opds.id =?", req.ID).
			First(&PenetapanKonteksRisikoStrategisRenstraOPD).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("PenetapanKonteksRisikoStrategisRenstraOPD id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &PenetapanKonteksRisikoStrategisRenstraOPDGetByIDRes{PenetapanKonteksRisikoStrategisRenstraOPD: PenetapanKonteksRisikoStrategisRenstraOPD}, nil
	}
}
