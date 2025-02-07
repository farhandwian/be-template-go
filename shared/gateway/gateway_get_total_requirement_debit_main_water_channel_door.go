package gateway

import (
	"context"
	"gorm.io/gorm"
	"shared/core"
)

type GetTotalRequirementDebitMainWaterChannelDoorReq struct {
}

type GetTotalRequirementDebitMainWaterChannelDoorRes struct {
	TotalRequirementDebit float32
}

type GetTotalRequirementDebitMainWaterChannelDoorGateway = core.ActionHandler[GetTotalRequirementDebitMainWaterChannelDoorReq, GetTotalRequirementDebitMainWaterChannelDoorRes]

func ImplGetTotalRequirementDebitMainWaterChannelDoorGateway(db *gorm.DB) GetTotalRequirementDebitMainWaterChannelDoorGateway {
	return func(ctx context.Context, req GetTotalRequirementDebitMainWaterChannelDoorReq) (*GetTotalRequirementDebitMainWaterChannelDoorRes, error) {
		var result GetTotalRequirementDebitMainWaterChannelDoorRes
		err := db.Raw(`
          SELECT SUM(water_requirement) AS total_requirement_debit 
          FROM water_channels
          WHERE external_id IN (2,3,4)
       `).Scan(&result).Error
		if err != nil {
			return nil, err
		}
		return &GetTotalRequirementDebitMainWaterChannelDoorRes{
			TotalRequirementDebit: result.TotalRequirementDebit,
		}, nil
	}
}
