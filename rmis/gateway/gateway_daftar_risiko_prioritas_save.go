package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type DaftarRisikoPrioritasSaveReq struct {
	DaftarRisikoPrioritas model.DaftarRisikoPrioritas
}

type DaftarRisikoPrioritasSaveRes struct {
	ID string
}

type DaftarRisikoPrioritasSave = core.ActionHandler[DaftarRisikoPrioritasSaveReq, DaftarRisikoPrioritasSaveRes]

func ImplDaftarRisikoPrioritasSave(db *gorm.DB) DaftarRisikoPrioritasSave {
	return func(ctx context.Context, req DaftarRisikoPrioritasSaveReq) (*DaftarRisikoPrioritasSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.DaftarRisikoPrioritas).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &DaftarRisikoPrioritasSaveRes{ID: *req.DaftarRisikoPrioritas.ID}, nil
	}
}
