package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type RancanganPemantauanGetByIDReq struct {
	ID string
}

type RancanganPemantauanGetByIDRes struct {
	RancanganPemantauan model.RancanganPemantauan
}

type RancanganPemantauanGetByID = core.ActionHandler[RancanganPemantauanGetByIDReq, RancanganPemantauanGetByIDRes]

func ImplRancanganPemantauanGetByID(db *gorm.DB) RancanganPemantauanGetByID {
	return func(ctx context.Context, req RancanganPemantauanGetByIDReq) (*RancanganPemantauanGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var RancanganPemantauan model.RancanganPemantauan
		if err := query.
			Joins("LEFT JOIN penilaian_risikos ON rancangan_pemantauans.penilaian_risiko_id = penilaian_risikos.id").
			Joins("LEFT JOIN daftar_risiko_prioritas ON penilaian_risikos.daftar_risiko_prioritas_id = daftar_risiko_prioritas.id").
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON daftar_risiko_prioritas.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id").
			Where("rancangan_pemantauans.id =?", req.ID).
			First(&RancanganPemantauan).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("RancanganPemantauan id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &RancanganPemantauanGetByIDRes{RancanganPemantauan: RancanganPemantauan}, nil
	}
}
