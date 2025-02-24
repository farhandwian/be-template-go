package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IKUGetByIDReq struct {
	ID string
}

type IKUGetByIDRes struct {
	IKU model.IKU
}

type IKUGetByID = core.ActionHandler[IKUGetByIDReq, IKUGetByIDRes]

func ImplIKUGetByID(db *gorm.DB) IKUGetByID {
	return func(ctx context.Context, req IKUGetByIDReq) (*IKUGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var iku model.IKU
		if err := query.First(&iku, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("iku id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &IKUGetByIDRes{IKU: iku}, nil
	}
}
