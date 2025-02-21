package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type SpipSaveReq struct {
	Spip model.SPIP
}

type SpipSaveRes struct {
	ID string
}

type SpipSave = core.ActionHandler[SpipSaveReq, SpipSaveRes]

func ImplSpipSave(db *gorm.DB) SpipSave {
	return func(ctx context.Context, req SpipSaveReq) (*SpipSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.Spip).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &SpipSaveRes{ID: *req.Spip.ID}, nil
	}
}
