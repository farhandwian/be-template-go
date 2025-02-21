package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type KategoriRisikoGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type KategoriRisikoGetAllRes struct {
	KategoriRisiko []model.KategoriRisiko `json:"kategori_risikos"`
	Count          int64                  `json:"count"`
}

type KategoriRisikoGetAll = core.ActionHandler[KategoriRisikoGetAllReq, KategoriRisikoGetAllRes]

func ImplKategoriRisikoGetAll(db *gorm.DB) KategoriRisikoGetAll {
	return func(ctx context.Context, req KategoriRisikoGetAllReq) (*KategoriRisikoGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword).
				Or("kode LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.KategoriRisiko{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.KategoriRisiko

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &KategoriRisikoGetAllRes{
			KategoriRisiko: objs,
			Count:          count,
		}, nil
	}
}
