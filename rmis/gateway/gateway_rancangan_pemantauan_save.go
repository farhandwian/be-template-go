package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type RancanganPemantauanSaveReq struct {
	RancanganPemantauan model.RancanganPemantauan
}

type RancanganPemantauanSaveRes struct {
	ID string
}

type RancanganPemantauanSave = core.ActionHandler[RancanganPemantauanSaveReq, RancanganPemantauanSaveRes]

func ImplRancanganPemantauanSave(db *gorm.DB) RancanganPemantauanSave {
	return func(ctx context.Context, req RancanganPemantauanSaveReq) (*RancanganPemantauanSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Save(&req.RancanganPemantauan).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &RancanganPemantauanSaveRes{ID: *req.RancanganPemantauan.ID}, nil
	}
}
