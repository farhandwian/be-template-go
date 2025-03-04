package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type RcaGetByIDReq struct {
	ID string
}

type RcaGetByIDRes struct {
	Rca model.Rca
}

type RcaGetByID = core.ActionHandler[RcaGetByIDReq, RcaGetByIDRes]

func ImplRcaGetByID(db *gorm.DB) RcaGetByID {
	return func(ctx context.Context, req RcaGetByIDReq) (*RcaGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var rca model.Rca
		if err := query.
			Select(`
		rcas.*, 
		identifikasi_risiko_strategis_pemdas.uraian_risiko AS identifikasi_risiko_uraian_risiko,
		penyebab_risikos.nama AS penyebab_risiko_nama
		`).
			Joins("LEFT JOIN identifikasi_risiko_strategis_pemdas ON rcas.identifikasi_risiko_strategis_pemda_id = identifikasi_risiko_strategis_pemdas.id").
			Joins("LEFT JOIN penyebab_risikos ON rcas.penyebab_risiko_id = penyebab_risikos.id").
			Where("rcas.id = ?", req.ID).
			First(&rca).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("rca id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &RcaGetByIDRes{Rca: rca}, nil
	}
}
