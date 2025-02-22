package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetapanKonteksRisikoStrategisPemdaGetByIDReq struct {
	ID string
}

type PenetapanKonteksRisikoStrategisPemdaGetByIDRes struct {
	PenetapanKonteksRisikoStrategisPemda model.PenetapanKonteksRisikoStrategisPemda
}

type PenetapanKonteksRisikoStrategisPemdaGetByID = core.ActionHandler[PenetapanKonteksRisikoStrategisPemdaGetByIDReq, PenetapanKonteksRisikoStrategisPemdaGetByIDRes]

func ImplPenetapanKonteksRisikoStrategisPemdaGetByID(db *gorm.DB) PenetapanKonteksRisikoStrategisPemdaGetByID {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisPemdaGetByIDReq) (*PenetapanKonteksRisikoStrategisPemdaGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var PenetapanKonteksRisikoStrategisPemda model.PenetapanKonteksRisikoStrategisPemda
		if err := query.First(&PenetapanKonteksRisikoStrategisPemda, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("PenetapanKonteksRisikoStrategisPemda id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &PenetapanKonteksRisikoStrategisPemdaGetByIDRes{PenetapanKonteksRisikoStrategisPemda: PenetapanKonteksRisikoStrategisPemda}, nil
	}
}
