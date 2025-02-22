package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetepanKonteksRisikoStrategisPemdaSaveReq struct {
	PenetepanKonteksRisikoStrategisPemda model.PenetapanKonteksRisikoStrategisPemda
}

type PenetepanKonteksRisikoStrategisPemdaSaveRes struct {
	ID string
}

type PenetepanKonteksRisikoStrategisPemdaSave = core.ActionHandler[PenetepanKonteksRisikoStrategisPemdaSaveReq, PenetepanKonteksRisikoStrategisPemdaSaveRes]

func ImplPenetepanKonteksRisikoStrategisPemdaSave(db *gorm.DB) PenetepanKonteksRisikoStrategisPemdaSave {
	return func(ctx context.Context, req PenetepanKonteksRisikoStrategisPemdaSaveReq) (*PenetepanKonteksRisikoStrategisPemdaSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.PenetepanKonteksRisikoStrategisPemda).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenetepanKonteksRisikoStrategisPemdaSaveRes{ID: *req.PenetepanKonteksRisikoStrategisPemda.ID}, nil
	}
}
