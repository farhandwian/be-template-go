package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoStrategisOPDGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type IdentifikasiRisikoStrategisOPDGetAllRes struct {
	IdentifikasiRisikoStrategisOPD []model.IdentifikasiRisikoStrategisOPD `json:"identifikasi_risiko_strategis_opd"`
	Count                          int64                                  `json:"count"`
}

type IdentifikasiRisikoStrategisOPDGetAll = core.ActionHandler[IdentifikasiRisikoStrategisOPDGetAllReq, IdentifikasiRisikoStrategisOPDGetAllRes]

func ImplIdentifikasiRisikoStrategisOPDGetAll(db *gorm.DB) IdentifikasiRisikoStrategisOPDGetAll {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDGetAllReq) (*IdentifikasiRisikoStrategisOPDGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword).
				Or("kode LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.IdentifikasiRisikoStrategisOPD{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.IdentifikasiRisikoStrategisOPD

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoStrategisOPDGetAllRes{
			IdentifikasiRisikoStrategisOPD: objs,
			Count:                          count,
		}, nil
	}
}
