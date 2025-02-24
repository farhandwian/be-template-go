package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenilaianKegiatanPengendalianGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type PenilaianKegiatanPengendalianGetAllRes struct {
	PenilaianKegiatanPengendalian []model.PenilaianKegiatanPengendalian `json:"penilai_kegiatan_pengendalians"`
	Count                         int64                                 `json:"count"`
}

type PenilaianKegiatanPengendalianGetAll = core.ActionHandler[PenilaianKegiatanPengendalianGetAllReq, PenilaianKegiatanPengendalianGetAllRes]

func ImplPenilaianKegiatanPengendalianGetAll(db *gorm.DB) PenilaianKegiatanPengendalianGetAll {
	return func(ctx context.Context, req PenilaianKegiatanPengendalianGetAllReq) (*PenilaianKegiatanPengendalianGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.PenilaianKegiatanPengendalian{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.PenilaianKegiatanPengendalian

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &PenilaianKegiatanPengendalianGetAllRes{
			PenilaianKegiatanPengendalian: objs,
			Count:                         count,
		}, nil
	}
}
