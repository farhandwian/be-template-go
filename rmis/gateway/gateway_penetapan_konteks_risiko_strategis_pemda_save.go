package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetapanKonteksRisikoStrategisPemdaSaveReq struct {
	PenetepanKonteksRisikoStrategisPemda model.PenetapanKonteksRisikoStrategisPemda
}

type PenetapanKonteksRisikoStrategisPemdaSaveRes struct {
	ID string
}

type PenetapanKonteksRisikoStrategisPemdaSave = core.ActionHandler[PenetapanKonteksRisikoStrategisPemdaSaveReq, PenetapanKonteksRisikoStrategisPemdaSaveRes]

func ImplPenetepanKonteksRisikoStrategisPemdaSave(db *gorm.DB) PenetapanKonteksRisikoStrategisPemdaSave {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisPemdaSaveReq) (*PenetapanKonteksRisikoStrategisPemdaSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.PenetepanKonteksRisikoStrategisPemda).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenetapanKonteksRisikoStrategisPemdaSaveRes{ID: *req.PenetepanKonteksRisikoStrategisPemda.ID}, nil
	}
}
