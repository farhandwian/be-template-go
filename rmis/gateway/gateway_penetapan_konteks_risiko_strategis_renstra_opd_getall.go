package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetapanKonteksRisikoStrategisRenstraOPDGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type PenetapanKonteksRisikoStrategisRenstraOPDGetAllRes struct {
	PenetapanKonteksRisikoStrategisRenstraOPD []model.PenetapanKonteksRisikoStrategisRenstraOPD `json:"penetapan_konteks_risiko_strategis_renstra_opd"`
	Count                                     int64                                             `json:"count"`
}

type PenetapanKonteksRisikoStrategisRenstraOPDGetAll = core.ActionHandler[PenetapanKonteksRisikoStrategisRenstraOPDGetAllReq, PenetapanKonteksRisikoStrategisRenstraOPDGetAllRes]

func ImplPenetapanKonteksRisikoStrategisRenstraOPDGetAll(db *gorm.DB) PenetapanKonteksRisikoStrategisRenstraOPDGetAll {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisRenstraOPDGetAllReq) (*PenetapanKonteksRisikoStrategisRenstraOPDGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.PenetapanKonteksRisikoStrategisRenstraOPD{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.PenetapanKonteksRisikoStrategisRenstraOPD

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenetapanKonteksRisikoStrategisRenstraOPDGetAllRes{
			PenetapanKonteksRisikoStrategisRenstraOPD: objs,
			Count: count,
		}, nil
	}
}
