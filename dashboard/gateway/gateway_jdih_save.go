package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type JDIHSaveReq struct {
	JDIH model.JDIH
}

type JDIHSaveResp struct {
	JDIH model.JDIH
}
type JDIHSaveGateway = core.ActionHandler[JDIHSaveReq, JDIHSaveResp]

func ImplJDIHSave(db *gorm.DB) JDIHSaveGateway {
	return func(ctx context.Context, request JDIHSaveReq) (*JDIHSaveResp, error) {
		query := middleware.GetDBFromContext(ctx, db)
		if err := query.Save(&request.JDIH).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &JDIHSaveResp{
			JDIH: request.JDIH,
		}, nil
	}
}
