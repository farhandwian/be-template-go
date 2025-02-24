package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IKUGetAllReq struct {
	ExternalID string
	Keyword    string
	Page       int
	Size       int
}

type IKUGetAllRes struct {
	IKU   []model.IKU `json:"iku"`
	Count int64       `json:"count"`
}

type IKUGetAll = core.ActionHandler[IKUGetAllReq, IKUGetAllRes]

func ImplIKUGetAll(db *gorm.DB) IKUGetAll {
	return func(ctx context.Context, req IKUGetAllReq) (*IKUGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		if req.ExternalID != "" {
			query = query.
				Where("external_id = ?", req.ExternalID)
		}

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("nama LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.IKU{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.IKU

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &IKUGetAllRes{
			IKU:   objs,
			Count: count,
		}, nil
	}
}
