package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetapanKonteksRisikoOperasionalGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type PenetapanKonteksRisikoOperasionalGetAllRes struct {
	PenetapanKonteksRisikoOperasional []model.PenetapanKonteksRisikoOperasional `json:"penetapan_konteks_risiko_operasional"`
	Count                             int64                                     `json:"count"`
}

type PenetapanKonteksRisikoOperasionalGetAll = core.ActionHandler[PenetapanKonteksRisikoOperasionalGetAllReq, PenetapanKonteksRisikoOperasionalGetAllRes]

func ImplPenetapanKonteksRisikoOperasionalGetAll(db *gorm.DB) PenetapanKonteksRisikoOperasionalGetAll {
	return func(ctx context.Context, req PenetapanKonteksRisikoOperasionalGetAllReq) (*PenetapanKonteksRisikoOperasionalGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.PenetapanKonteksRisikoOperasional{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.PenetapanKonteksRisikoOperasional

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenetapanKonteksRisikoOperasionalGetAllRes{
			PenetapanKonteksRisikoOperasional: objs,
			Count:                             count,
		}, nil
	}
}
