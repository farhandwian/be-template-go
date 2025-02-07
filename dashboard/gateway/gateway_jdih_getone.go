package gateway

import (
	"context"
	"dashboard/model"
	"errors"
	"fmt"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type GetJDIHByIDReq struct {
	ID string
}

type GetJDIHByIDResp struct {
	JDIH model.JDIH
}
type GetJDIHByIDGateway = core.ActionHandler[GetJDIHByIDReq, GetJDIHByIDResp]

func ImplGetJDIHByID(db *gorm.DB) GetJDIHByIDGateway {
	return func(ctx context.Context, request GetJDIHByIDReq) (*GetJDIHByIDResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var JDIH model.JDIH
		if err := query.First(&JDIH, "id = ?", request.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("document and law information id %v is not found", request.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &GetJDIHByIDResp{JDIH: JDIH}, nil
	}
}
