package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/helper"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetapanKonteksRisikoOperasionalGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
}

type PenetapanKonteksRisikoOperasionalGetAllRes struct {
	PenetapanKonteksRisikoOperasional []model.PenetapanKonteksRisikoOperasionalResponse `json:"penetapan_konteks_risiko_operasional"`
	Count                             int64                                             `json:"count"`
}

type PenetapanKonteksRisikoOperasionalGetAll = core.ActionHandler[PenetapanKonteksRisikoOperasionalGetAllReq, PenetapanKonteksRisikoOperasionalGetAllRes]

func ImplPenetapanKonteksRisikoOperasionalGetAll(db *gorm.DB) PenetapanKonteksRisikoOperasionalGetAll {
	return func(ctx context.Context, req PenetapanKonteksRisikoOperasionalGetAllReq) (*PenetapanKonteksRisikoOperasionalGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		query = query.
			Joins("LEFT JOIN opds ON penetapan_konteks_risiko_operasionals.opd_id = opds.id")

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
		allowedSortBy := map[string]bool{
			"nama_pemda":  true,
			"periode":     true,
			"sumber_data": true,
			"status":      true,
		}

		allowerdForeignSortBy := map[string]string{
			"opd": "opds.nama",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.PenetapanKonteksRisikoOperasionalResponse

		if err := query.
			Select(`penetapan_konteks_risiko_operasionals.*, 
                    opds.nama AS opd_nama
					`).
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Order(orderClause).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenetapanKonteksRisikoOperasionalGetAllRes{
			PenetapanKonteksRisikoOperasional: objs,
			Count:                             count,
		}, nil
	}
}
