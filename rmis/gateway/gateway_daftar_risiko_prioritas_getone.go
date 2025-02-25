package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type DaftarRisikoPrioritasGetByIDReq struct {
	ID string
}

type DaftarRisikoPrioritasGetByIDRes struct {
	DaftarRisikoPrioritas model.DaftarRisikoPrioritas
}

type DaftarRisikoPrioritasGetByID = core.ActionHandler[DaftarRisikoPrioritasGetByIDReq, DaftarRisikoPrioritasGetByIDRes]

func ImplDaftarRisikoPrioritasGetByID(db *gorm.DB) DaftarRisikoPrioritasGetByID {
	return func(ctx context.Context, req DaftarRisikoPrioritasGetByIDReq) (*DaftarRisikoPrioritasGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var DaftarRisikoPrioritas model.DaftarRisikoPrioritas
		if err := query.First(&DaftarRisikoPrioritas, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("DaftarRisikoPrioritas id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &DaftarRisikoPrioritasGetByIDRes{DaftarRisikoPrioritas: DaftarRisikoPrioritas}, nil
	}
}
