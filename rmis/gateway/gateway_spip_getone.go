package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type SpipGetByIDReq struct {
	ID string
}

type SpipGetByIDRes struct {
	SPIP model.SPIP
}

type SpipGetByID = core.ActionHandler[SpipGetByIDReq, SpipGetByIDRes]

func ImplSpipGetByID(db *gorm.DB) SpipGetByID {
	return func(ctx context.Context, req SpipGetByIDReq) (*SpipGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var spip model.SPIP
		if err := query.First(&spip, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("spip id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &SpipGetByIDRes{SPIP: spip}, nil
	}
}
