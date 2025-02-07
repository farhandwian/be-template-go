package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type DeleteJDIHReq struct {
	ID string
}

type DeleteJDIHResp struct {
}
type DeleteJDIHGateway = core.ActionHandler[DeleteJDIHReq, DeleteJDIHResp]

func ImplDeleteJDIH(db *gorm.DB) DeleteJDIHGateway {
	return func(ctx context.Context, request DeleteJDIHReq) (*DeleteJDIHResp, error) {
		query := middleware.GetDBFromContext(ctx, db)
		if err := query.Where("id = ?", request.ID).Delete(&model.JDIH{}).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &DeleteJDIHResp{}, nil
	}
}
