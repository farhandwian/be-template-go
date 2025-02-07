package gateway

import (
	"context"
	"errors"
	"fmt"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type IntakeGetDetailReq struct {
	ID string
}
type IntakeGetDetailResp struct {
	Intake model.Intake
}

type IntakeGetDetailGateway = core.ActionHandler[IntakeGetDetailReq, IntakeGetDetailResp]

func ImplIntakeGetDetailGateway(db *gorm.DB) IntakeGetDetailGateway {
	return func(ctx context.Context, request IntakeGetDetailReq) (*IntakeGetDetailResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var intake model.Intake

		err := query.Where("id=?", request.ID).First(&intake).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("intake id %v is not found", request.ID)
			}
			return nil, core.NewInternalServerError(err)
		}
		return &IntakeGetDetailResp{
			Intake: intake,
		}, err
	}
}
