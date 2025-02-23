package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type RcaGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type RcaGetAllRes struct {
	Rca   []model.Rca `json:"Rcas"`
	Count int64       `json:"count"`
}

type RcaGetAll = core.ActionHandler[RcaGetAllReq, RcaGetAllRes]

func ImplRcaGetAll(db *gorm.DB) RcaGetAll {
	return func(ctx context.Context, req RcaGetAllReq) (*RcaGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.Rca{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.Rca

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &RcaGetAllRes{
			Rca:   objs,
			Count: count,
		}, nil
	}
}
