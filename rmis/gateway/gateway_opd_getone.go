package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type OPDGetByIDReq struct {
	ID string
}

type OPDGetByIDRes struct {
	OPD model.OPD
}

type OPDGetByID = core.ActionHandler[OPDGetByIDReq, OPDGetByIDRes]

func ImplOPDGetByID(db *gorm.DB) OPDGetByID {
	return func(ctx context.Context, req OPDGetByIDReq) (*OPDGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var opd model.OPD
		if err := query.First(&opd, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("opd id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &OPDGetByIDRes{OPD: opd}, nil
	}
}
