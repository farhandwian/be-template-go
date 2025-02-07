package gateway

import (
	"context"
	"fmt"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type DoorControlGetOneReq struct {
	ID model.DoorControlID
}

type DoorControlGetOneRes struct {
	Item *model.DoorControl
}

type DoorControlGetOne = core.ActionHandler[DoorControlGetOneReq, DoorControlGetOneRes]

func ImplDoorControlGetOne(db *gorm.DB) DoorControlGetOne {
	return func(ctx context.Context, req DoorControlGetOneReq) (*DoorControlGetOneRes, error) {

		var obj model.DoorControl

		if err := middleware.
			GetDBFromContext(ctx, db).
			First(&obj, "id = ?", req.ID).
			Error; err != nil {

			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("id %v is not found", req.ID)
			}

			return nil, core.NewInternalServerError(err)
		}

		return &DoorControlGetOneRes{Item: &obj}, nil
	}
}
