package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type DoorControlSaveReq struct {
	DoorControl *model.DoorControl
}

type DoorControlSaveRes struct {
	ID model.DoorControlID
}

type DoorControlSave = core.ActionHandler[DoorControlSaveReq, DoorControlSaveRes]

func ImplDoorControlSave(db *gorm.DB) DoorControlSave {
	return func(ctx context.Context, req DoorControlSaveReq) (*DoorControlSaveRes, error) {

		if err := middleware.
			GetDBFromContext(ctx, db).
			Save(req.DoorControl).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &DoorControlSaveRes{ID: req.DoorControl.ID}, nil
	}
}
