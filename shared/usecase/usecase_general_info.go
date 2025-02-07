package usecase

import (
	"context"
	"fmt"
	"shared/core"
	"shared/gateway"
)

type GeneralInfoReq struct{}

type GeneralInfoRes struct {
	AgriculturalInfo   string `json:"agricultural_info"`    //
	PlantedArea        string `json:"planted_area"`         // LuasTanam
	PlantingSeason     string `json:"planting_season"`      // MusimTanam
	CroppingPattern    string `json:"cropping_pattern"`     // PolaTanam
	WaterRequirement   string `json:"water_requirement"`    // KebutuhanAir
	WaterAvailability  string `json:"water_availability"`   // KetersediaanAir
	HistoricalChart    string `json:"historical_chart"`     //
	TotalSensorOnline  int    `json:"total_sensor_online"`  //
	TotalSensorOffline int    `json:"total_sensor_offline"` //
	TotalCCTV          int    `json:"total_cctv"`           //
	TotalStaff         int    `json:"total_staff"`          // TotalPetugas
}

type GeneralInfoUseCase = core.ActionHandler[GeneralInfoReq, GeneralInfoRes]

func ImplGeneralInfo(generalInfoGateway gateway.GetGeneralInfoGateway, getTotalRequirementDebitMainWaterChannelDoorGateway gateway.GetTotalRequirementDebitMainWaterChannelDoorGateway) GeneralInfoUseCase {
	return func(ctx context.Context, request GeneralInfoReq) (*GeneralInfoRes, error) {

		generalInfoData, err := generalInfoGateway(ctx, gateway.GetGeneralInfoReq{})
		if err != nil {
			return nil, err
		}

		requirementDebit, err := getTotalRequirementDebitMainWaterChannelDoorGateway(ctx, gateway.GetTotalRequirementDebitMainWaterChannelDoorReq{})
		if err != nil {
			return nil, err
		}

		return &GeneralInfoRes{
			AgriculturalInfo:   "",
			PlantedArea:        fmt.Sprintf("%s Ha", generalInfoData.GeneralInfo.PlantingArea),
			PlantingSeason:     generalInfoData.GeneralInfo.PlantingSeason,
			CroppingPattern:    generalInfoData.GeneralInfo.PlantingPattern,
			WaterRequirement:   fmt.Sprintf("%.2f Lt", requirementDebit.TotalRequirementDebit),
			HistoricalChart:    "",
			TotalSensorOnline:  3,
			TotalSensorOffline: 20,
			TotalCCTV:          8,
			TotalStaff:         17,
		}, nil
	}
}
