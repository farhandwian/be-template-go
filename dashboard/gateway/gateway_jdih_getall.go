package gateway

import (
	"context"
	"dashboard/model"
	"fmt"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type GetListJDIHReq struct {
	Keyword string
	Page    int
	Size    int
}

type GetListJDIHResp struct {
	JDIH  []model.JDIH
	Count int64
}
type GetListJDIHGateway = core.ActionHandler[GetListJDIHReq, GetListJDIHResp]

func ImplGetListJDIH(db *gorm.DB) GetListJDIHGateway {
	return func(ctx context.Context, req GetListJDIHReq) (*GetListJDIHResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("title LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.JDIH{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.JDIH

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetListJDIHResp{
			JDIH:  objs,
			Count: count,
		}, nil
	}
}
