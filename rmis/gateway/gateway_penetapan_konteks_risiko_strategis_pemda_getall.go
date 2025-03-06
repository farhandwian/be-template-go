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

type PenetapanKonteksRisikoStrategisPemdaGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
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
				Where("nama_pemda LIKE ?", keyword).
				Or("penetapan_konteks_resiko_strategis LIKE ?", keyword)
		}

		var count int64

		if req.Status != "" {
			query = query.Where("status =?", req.Status)
		}

		if err := query.
			Model(&model.PenetapanKonteksRisikoStrategisPemda{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		allowedSortBy := map[string]bool{
			"nama_pemda":        true,
			"periode":           true,
			"penetapan_konteks": true,
			"sumber_data":       true,
		}

		sortBy, sortOrder, err := helper.ValidateSortParams(allowedSortBy, req.SortBy, req.SortOrder, "nama_pemda")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.PenetapanKonteksRisikoStrategisPemda

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
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
