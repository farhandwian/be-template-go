package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type SpipGetAllReq struct {
	Keyword string
	Page    int
	Size    int
}

type SpipGetAllRes struct {
	SPIP  []model.Spip `json:"spips"`
	Count int64        `json:"count"`
}

type SpipGetAll = core.ActionHandler[SpipGetAllReq, SpipGetAllRes]

func ImplSpipGetAll(db *gorm.DB) SpipGetAll {
	return func(ctx context.Context, req SpipGetAllReq) (*SpipGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.Spip{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.Spip

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &SpipGetAllRes{
			SPIP:  objs,
			Count: count,
		}, nil
	}
}
