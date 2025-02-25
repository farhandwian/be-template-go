package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type OPDSaveReq struct {
	OPD model.OPD
}

type OPDSaveRes struct {
	ID string
}

type OPDSave = core.ActionHandler[OPDSaveReq, OPDSaveRes]

func ImplOPDSave(db *gorm.DB) OPDSave {
	return func(ctx context.Context, req OPDSaveReq) (*OPDSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.OPD).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &OPDSaveRes{ID: *req.OPD.ID}, nil
	}
}
