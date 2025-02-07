package gateway

import (
	"bigboard/model"
	"context"
	"shared/core"

	"gorm.io/gorm"
)

type GetListSensitiveJobReq struct {
	ID string `json:"id_sensitive_job"`
}

type GetListSensitiveJobGateway = core.ActionHandler[GetListSensitiveJobReq, model.SensitiveJobs]

func ImplGetOneSensitiveJob(db *gorm.DB) GetListSensitiveJobGateway {
	return func(ctx context.Context, request GetListSensitiveJobReq) (*model.SensitiveJobs, error) {

		var sensitiveJobs model.SensitiveJobs

		err := db.Table("sensitive_jobs").
			Select("*").
			Where("id = ?", request.ID).
			Find(&sensitiveJobs).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &model.SensitiveJobs{
			ID:       sensitiveJobs.ID,
			FuncName: sensitiveJobs.FuncName,
			Status:   sensitiveJobs.Status,
			Payload:  sensitiveJobs.Payload,
		}, nil
	}
}
