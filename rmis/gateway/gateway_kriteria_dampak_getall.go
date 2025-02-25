package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type KriteriaDampakGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type KriteriaDampakGetAllRes struct {
	KriteriaDampak []model.KriteriaDampak `json:"KriteriaDampaks"`
	Count          int64                  `json:"count"`
}

type KriteriaDampakGetAll = core.ActionHandler[KriteriaDampakGetAllReq, KriteriaDampakGetAllRes]

func ImplKriteriaDampakGetAll(db *gorm.DB) KriteriaDampakGetAll {
	return func(ctx context.Context, req KriteriaDampakGetAllReq) (*KriteriaDampakGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.KriteriaDampak{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.KriteriaDampak

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &KriteriaDampakGetAllRes{
			KriteriaDampak: objs,
			Count:          count,
		}, nil
	}
}
