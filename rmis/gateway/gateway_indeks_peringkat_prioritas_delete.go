package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IndeksPeringkatPrioritasDeleteReq struct {
	ID string
}

type IndeksPeringkatPrioritasDeleteRes struct{}

type IndeksPeringkatPrioritasDelete = core.ActionHandler[IndeksPeringkatPrioritasDeleteReq, IndeksPeringkatPrioritasDeleteRes]

func ImplIndeksPeringkatPrioritasDelete(db *gorm.DB) IndeksPeringkatPrioritasDelete {
	return func(ctx context.Context, req IndeksPeringkatPrioritasDeleteReq) (*IndeksPeringkatPrioritasDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.IndeksPeringkatPrioritas{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IndeksPeringkatPrioritasDeleteRes{}, nil
	}
}
