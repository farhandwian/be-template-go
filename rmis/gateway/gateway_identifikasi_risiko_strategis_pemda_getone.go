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

		if err := query.
			Select(`identifikasi_risiko_strategis_pemdas.*, 
				penetapan_konteks_risiko_strategis_pemdas.nama_pemda AS nama_pemda,
				penetapan_konteks_risiko_strategis_pemdas.tahun AS tahun,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS tujuan,
				penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintah AS urusan_pemerintah,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS penetapan_tujuan,
			`).
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON identifikasi_risiko_strategis_pemdas.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id").
			Joins("LEFT JOIN kategori_risikos ON identifikasi_risiko_strategis_pemdas.kategori_risiko_id = kategori_risikos.kategori_risikos").
			Joins("LEFT JOIN rcas ON identifikasi_risiko_strategis_pemdas.rca_id = rcas.id").
			Where("identifikasi_risiko_strategis_pemdas.id = ?", req.ID).
			First(&result).Error; err != nil {
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
