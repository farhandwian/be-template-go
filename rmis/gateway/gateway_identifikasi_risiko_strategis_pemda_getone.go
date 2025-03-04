package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoStrategisPemdaGetByIDReq struct {
	ID string
}

type IdentifikasiRisikoStrategisPemdaGetByIDRes struct {
	IdentifikasiRisikoStrategisPemda model.IdentifikasiRisikoStrategisPemda
}

type IdentifikasiRisikoStrategisPemdaGetByID = core.ActionHandler[IdentifikasiRisikoStrategisPemdaGetByIDReq, IdentifikasiRisikoStrategisPemdaGetByIDRes]

func ImplIdentifikasiRisikoStrategisPemdaGetByID(db *gorm.DB) IdentifikasiRisikoStrategisPemdaGetByID {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaGetByIDReq) (*IdentifikasiRisikoStrategisPemdaGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var result model.IdentifikasiRisikoStrategisPemda
		// Use Preload to load the KategoriRisiko association
		if err := query.
			Preload("KategoriRisiko").
			Preload("Rca").
			Preload("PenetapanKonteksRisikoStrategisPemda").
			First(&result, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("IdentifikasiRisikoStrategisPemda id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoStrategisPemdaGetByIDRes{
			IdentifikasiRisikoStrategisPemda: result,
		}, nil
	}
}
