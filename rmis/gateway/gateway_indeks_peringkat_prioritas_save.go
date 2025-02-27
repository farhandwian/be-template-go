package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IndeksPeringkatPrioritasSaveReq struct {
	IndeksPeringkatPrioritas model.IndeksPeringkatPrioritas
}

type IndeksPeringkatPrioritasSaveRes struct {
	ID string
}

type IndeksPeringkatPrioritasSave = core.ActionHandler[IndeksPeringkatPrioritasSaveReq, IndeksPeringkatPrioritasSaveRes]

func ImplIndeksPeringkatPrioritasSave(db *gorm.DB) IndeksPeringkatPrioritasSave {
	return func(ctx context.Context, req IndeksPeringkatPrioritasSaveReq) (*IndeksPeringkatPrioritasSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.IndeksPeringkatPrioritas).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IndeksPeringkatPrioritasSaveRes{ID: *req.IndeksPeringkatPrioritas.ID}, nil
	}
}
