package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetapanKonteksRisikoStrategisPemdaGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type PenetapanKonteksRisikoStrategisPemdaGetAllRes struct {
	PenetapanKonteksRisikoStrategisPemda []model.PenetapanKonteksRisikoStrategisPemda `json:"penetapan_konteks_risiko_strategis_pemda"`
	Count                                int64                                        `json:"count"`
}

type PenetapanKonteksRisikoStrategisPemdaGetAll = core.ActionHandler[PenetapanKonteksRisikoStrategisPemdaGetAllReq, PenetapanKonteksRisikoStrategisPemdaGetAllRes]

func ImplPenetapanKonteksRisikoStrategisPemdaGetAll(db *gorm.DB) PenetapanKonteksRisikoStrategisPemdaGetAll {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisPemdaGetAllReq) (*PenetapanKonteksRisikoStrategisPemdaGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.PenetapanKonteksRisikoStrategisPemda{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.PenetapanKonteksRisikoStrategisPemda

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenetapanKonteksRisikoStrategisPemdaGetAllRes{
			PenetapanKonteksRisikoStrategisPemda: objs,
			Count:                                count,
		}, nil
	}
}
