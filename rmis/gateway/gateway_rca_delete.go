package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type RcaDeleteReq struct {
	ID string
}

type RcaDeleteRes struct{}

type RcaDelete = core.ActionHandler[RcaDeleteReq, RcaDeleteRes]

func ImplRcaDelete(db *gorm.DB) RcaDelete {
	return func(ctx context.Context, req RcaDeleteReq) (*RcaDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.Rca{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &RcaDeleteRes{}, nil
	}
}
