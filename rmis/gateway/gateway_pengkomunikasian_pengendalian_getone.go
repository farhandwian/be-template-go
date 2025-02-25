package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PengkomunikasianPengendalianGetByIDReq struct {
	ID string
}

type PengkomunikasianPengendalianGetByIDRes struct {
	PengkomunikasianPengendalian model.PengkomunikasianPengendalian
}

type PengkomunikasianPengendalianGetByID = core.ActionHandler[PengkomunikasianPengendalianGetByIDReq, PengkomunikasianPengendalianGetByIDRes]

func ImplPengkomunikasianPengendalianGetByID(db *gorm.DB) PengkomunikasianPengendalianGetByID {
	return func(ctx context.Context, req PengkomunikasianPengendalianGetByIDReq) (*PengkomunikasianPengendalianGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var PengkomunikasianPengendalian model.PengkomunikasianPengendalian
		if err := query.First(&PengkomunikasianPengendalian, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("PengkomunikasianPengendalian id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &PengkomunikasianPengendalianGetByIDRes{PengkomunikasianPengendalian: PengkomunikasianPengendalian}, nil
	}
}
