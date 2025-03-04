package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoStrategisPemdaSaveReq struct {
	IdentifikasiRisikoStrategisPemda model.IdentifikasiRisikoStrategisPemda
}

type IdentifikasiRisikoStrategisPemdaSaveRes struct {
	ID string
}

type IdentifikasiRisikoStrategisPemdaSave = core.ActionHandler[IdentifikasiRisikoStrategisPemdaSaveReq, IdentifikasiRisikoStrategisPemdaSaveRes]

func ImplIdentifikasiRisikoStrategisPemdaSave(db *gorm.DB) IdentifikasiRisikoStrategisPemdaSave {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaSaveReq) (*IdentifikasiRisikoStrategisPemdaSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.IdentifikasiRisikoStrategisPemda).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoStrategisPemdaSaveRes{ID: *req.IdentifikasiRisikoStrategisPemda.ID}, nil
	}
}
