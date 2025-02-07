package gateway

import (
	"bigboard/model"
	"context"
	"gorm.io/gorm"
	"shared/core"
)

// SELECT  gate_level, "timestamp"
// FROM public.water_gates
// where water_channel_door_id = 2 and device_id = 4
// ORDER BY "timestamp" desc

type GetGeneralInfoReq struct{}

type GetGeneralInfoRes struct {
	GeneralInfo model.GeneralInfo
}

type GetGeneralInfoGateway = core.ActionHandler[GetGeneralInfoReq, GetGeneralInfoRes]

func ImplGetGeneralInfoGateway(db *gorm.DB) GetGeneralInfoGateway {
	return func(ctx context.Context, request GetGeneralInfoReq) (*GetGeneralInfoRes, error) {
		var generalInfo model.GeneralInfo

		err := db.
			Order("created_at DESC").
			Limit(1).
			Find(&generalInfo).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetGeneralInfoRes{
			GeneralInfo: generalInfo,
		}, nil
	}
}
