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
		if err := query.First(&rca, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("rca id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &RcaGetByIDRes{Rca: rca}, nil
	}
}
