package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"
	"shared/middleware"
	sharedModel "shared/model"

	"gorm.io/gorm"
)

type CctvImageProcessingSaveReq struct {
	CctvImageProcessing model.CctvImageProcessing
}

type CctvImageProcessingSaveRes struct {
	CctvImageProcessing model.CctvImageProcessing
}

type CctvImageProcessingSave = core.ActionHandler[CctvImageProcessingSaveReq, CctvImageProcessingSaveRes]

func ImplCctvImageProcessingSave(mariaDB *gorm.DB, tsdb *gorm.DB) CctvImageProcessingSave {
	return func(ctx context.Context, req CctvImageProcessingSaveReq) (*CctvImageProcessingSaveRes, error) {
		//save to WaterChannelDevice
		query := middleware.GetDBFromContext(ctx, mariaDB)

		if err := query.Model(&sharedModel.WaterChannelDevice{}).
			Where("ip_address = ?", req.CctvImageProcessing.IP).
			Updates(map[string]interface{}{
				"detected_object": req.CctvImageProcessing.DetectedObject,
			}).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		//save to CctvImageProcessing for history if detection exist
		if req.CctvImageProcessing.DetectedObject != "" {
			query = middleware.GetDBFromContext(ctx, tsdb)
			if err := query.Create(&req.CctvImageProcessing).Error; err != nil {
				return nil, core.NewInternalServerError(err)
			}
		}

		return &CctvImageProcessingSaveRes{CctvImageProcessing: req.CctvImageProcessing}, nil
	}
}
