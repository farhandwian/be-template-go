package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoStrategisOPDGetByIDReq struct {
	ID string
}

type IdentifikasiRisikoStrategisOPDGetByIDRes struct {
	IdentifikasiRisikoStrategisOPD model.IdentifikasiRisikoStrategisOPD
}

type IdentifikasiRisikoStrategisOPDGetByID = core.ActionHandler[IdentifikasiRisikoStrategisOPDGetByIDReq, IdentifikasiRisikoStrategisOPDGetByIDRes]

func ImplIdentifikasiRisikoStrategisOPDGetByID(db *gorm.DB) IdentifikasiRisikoStrategisOPDGetByID {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDGetByIDReq) (*IdentifikasiRisikoStrategisOPDGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var IdentifikasiRisikoStrategisOPD model.IdentifikasiRisikoStrategisOPD
		if err := query.
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_renstra_opds ON identifikasi_risiko_strategis_opds.penetapan_konteks_risiko_strategis_renstra_id = penetapan_konteks_risiko_strategis_renstra_opds.id").
			Joins("LEFT JOIN kategori_risikos ON identifikasi_risiko_strategis_opds.kategori_risiko_id = kategori_risikos.id").
			Where("identifikasi_risiko_strategis_opds.id =?", req.ID).
			First(&IdentifikasiRisikoStrategisOPD).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("IdentifikasiRisikoStrategisOPD id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoStrategisOPDGetByIDRes{IdentifikasiRisikoStrategisOPD: IdentifikasiRisikoStrategisOPD}, nil
	}
}
