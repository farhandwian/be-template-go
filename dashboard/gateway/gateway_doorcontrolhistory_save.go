package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type DoorControlHistorySaveReq struct {
	DoorControlHistory *model.DoorControlHistory
}

type DoorControlHistorySaveRes struct {
	ID model.DoorControlHistoryID
}

type DoorControlHistorySave = core.ActionHandler[DoorControlHistorySaveReq, DoorControlHistorySaveRes]

func ImplDoorControlHistorySave(db *gorm.DB) DoorControlHistorySave {
	return func(ctx context.Context, req DoorControlHistorySaveReq) (*DoorControlHistorySaveRes, error) {

		if err := middleware.
			GetDBFromContext(ctx, db).
			Save(req.DoorControlHistory).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &DoorControlHistorySaveRes{ID: req.DoorControlHistory.ID}, nil
	}
}
