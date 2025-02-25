package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type OPDGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type OPDGetAllRes struct {
	OPD   []model.OPD `json:"opds"`
	Count int64       `json:"count"`
}

type OPDGetAll = core.ActionHandler[OPDGetAllReq, OPDGetAllRes]

func ImplOPDGetAll(db *gorm.DB) OPDGetAll {
	return func(ctx context.Context, req OPDGetAllReq) (*OPDGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.OPD{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.OPD

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &OPDGetAllRes{
			OPD:   objs,
			Count: count,
		}, nil
	}
}
