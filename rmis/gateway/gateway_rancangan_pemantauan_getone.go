package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type RancanganPemantauanGetByIDReq struct {
	ID string
}

type RancanganPemantauanGetByIDRes struct {
	RancanganPemantauan model.RancanganPemantauan
}

type RancanganPemantauanGetByID = core.ActionHandler[RancanganPemantauanGetByIDReq, RancanganPemantauanGetByIDRes]

func ImplRancanganPemantauanGetByID(db *gorm.DB) RancanganPemantauanGetByID {
	return func(ctx context.Context, req RancanganPemantauanGetByIDReq) (*RancanganPemantauanGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var RancanganPemantauan model.RancanganPemantauan
		if err := query.First(&RancanganPemantauan, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("RancanganPemantauan id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &RancanganPemantauanGetByIDRes{RancanganPemantauan: RancanganPemantauan}, nil
	}
}
