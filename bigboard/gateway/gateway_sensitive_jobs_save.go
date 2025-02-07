package gateway

import (
	"bigboard/model"
	"context"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type SensitiveJobsSaveReq struct {
	SensitiveJobs model.SensitiveJobs
}

type SensitiveJobsSaveRes struct {
	ID model.SensitiveJobsID
}

type SensitiveJobsSave = core.ActionHandler[SensitiveJobsSaveReq, SensitiveJobsSaveRes]

func ImplSensitiveJobsSave(db *gorm.DB) SensitiveJobsSave {
	return func(ctx context.Context, req SensitiveJobsSaveReq) (*SensitiveJobsSaveRes, error) {
		query := middleware.GetDBFromContext(ctx, db)
		if err := query.Save(&req.SensitiveJobs).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &SensitiveJobsSaveRes{ID: req.SensitiveJobs.ID}, nil
	}
}
