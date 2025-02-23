package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoStrategisPemdaGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type IdentifikasiRisikoStrategisPemdaGetAllRes struct {
	IdentifikasiRisikoStrategisPemda []model.IdentifikasiRisikoStrategisPemerintahDaerah `json:"identifikasi_risiko_strategis_pemda"`
	Count                            int64                                               `json:"count"`
}

type IdentifikasiRisikoStrategisPemdaGetAll = core.ActionHandler[IdentifikasiRisikoStrategisPemdaGetAllReq, IdentifikasiRisikoStrategisPemdaGetAllRes]

func ImplIdentifikasiRisikoStrategisPemdaGetAll(db *gorm.DB) IdentifikasiRisikoStrategisPemdaGetAll {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaGetAllReq) (*IdentifikasiRisikoStrategisPemdaGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword).
				Or("kode LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.IdentifikasiRisikoStrategisPemerintahDaerah{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.IdentifikasiRisikoStrategisPemerintahDaerah

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoStrategisPemdaGetAllRes{
			IdentifikasiRisikoStrategisPemda: objs,
			Count:                            count,
		}, nil
	}
}
