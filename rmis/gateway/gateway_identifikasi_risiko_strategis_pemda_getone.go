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
	IdentifikasiRisikoStrategisPemda model.IdentifikasiRisikoStrategisPemerintahDaerah
}

type IdentifikasiRisikoStrategisPemdaGetByID = core.ActionHandler[IdentifikasiRisikoStrategisPemdaGetByIDReq, IdentifikasiRisikoStrategisPemdaGetByIDRes]

func ImplIdentifikasiRisikoStrategisPemdaGetByID(db *gorm.DB) IdentifikasiRisikoStrategisPemdaGetByID {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaGetByIDReq) (*IdentifikasiRisikoStrategisPemdaGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var IdentifikasiRisikoStrategisPemda model.IdentifikasiRisikoStrategisPemerintahDaerah
		if err := query.
			Joins("LEFT JOIN kategori_risiko ON kategori_risiko_id = identifikasi_risiko_strategis_pemerintah_daerah.kategori_risiko_id").
			First(&IdentifikasiRisikoStrategisPemda, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("IdentifikasiRisikoStrategisPemda id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoStrategisPemdaGetByIDRes{IdentifikasiRisikoStrategisPemda: IdentifikasiRisikoStrategisPemda}, nil
	}
}
