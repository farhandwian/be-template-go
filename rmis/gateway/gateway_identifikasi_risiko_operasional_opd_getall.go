package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoOperasionalOPDGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type IdentifikasiRisikoOperasionalOPDGetAllRes struct {
	IdentifikasiRisikoOperasionalOPD []model.IdentifikasiRisikoOperasionalOPD `json:"identifikasi_risiko_operasional_opd"`
	Count                            int64                                    `json:"count"`
}

type IdentifikasiRisikoOperasionalOPDGetAll = core.ActionHandler[IdentifikasiRisikoOperasionalOPDGetAllReq, IdentifikasiRisikoOperasionalOPDGetAllRes]

func ImplIdentifikasiRisikoOperasionalOPDGetAll(db *gorm.DB) IdentifikasiRisikoOperasionalOPDGetAll {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDGetAllReq) (*IdentifikasiRisikoOperasionalOPDGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword).
				Or("kode LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.IdentifikasiRisikoOperasionalOPD{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.IdentifikasiRisikoOperasionalOPD

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoOperasionalOPDGetAllRes{
			IdentifikasiRisikoOperasionalOPD: objs,
			Count:                            count,
		}, nil
	}
}
