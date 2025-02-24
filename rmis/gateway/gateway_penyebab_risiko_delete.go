package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenyebabRisikoDeleteReq struct {
	ID string
}

type PenyebabRisikoDeleteRes struct{}

type PenyebabRisikoDelete = core.ActionHandler[PenyebabRisikoDeleteReq, PenyebabRisikoDeleteRes]

func ImplPenyebabRisikoDelete(db *gorm.DB) PenyebabRisikoDelete {
	return func(ctx context.Context, req PenyebabRisikoDeleteReq) (*PenyebabRisikoDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.PenyebabRisiko{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenyebabRisikoDeleteRes{}, nil
	}
}
