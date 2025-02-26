package gateway

import (
	"context"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type RancanganPemantauanDeleteReq struct {
	ID string
}

type RancanganPemantauanDeleteRes struct{}

type RancanganPemantauanDelete = core.ActionHandler[RancanganPemantauanDeleteReq, RancanganPemantauanDeleteRes]

func ImplRancanganPemantauanDelete(db *gorm.DB) RancanganPemantauanDelete {
	return func(ctx context.Context, req RancanganPemantauanDeleteReq) (*RancanganPemantauanDeleteRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if err := query.Delete(&model.RancanganPemantauan{}, "id = ?", req.ID).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &RancanganPemantauanDeleteRes{}, nil
	}
}
