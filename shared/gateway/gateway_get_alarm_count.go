package gateway

import (
	"context"
	"gorm.io/gorm"
	"shared/core"
)

type GetAlarmCountReq struct {
}

type GetAlarmCountRes struct {
	WarningCount  int
	CriticalCount int
}

type GetAlarmCountGateway = core.ActionHandler[GetAlarmCountReq, GetAlarmCountRes]

func ImplGetAlarmCountGateway(db *gorm.DB) GetAlarmCountGateway {
	return func(ctx context.Context, req GetAlarmCountReq) (*GetAlarmCountRes, error) {
		var AlarmCount GetAlarmCountRes
		err := db.Raw(`
						SELECT 
			SUM(CASE WHEN priority = 'warning' THEN 1 ELSE 0 END) AS warning_count,
			SUM(CASE WHEN priority = 'critical' THEN 1 ELSE 0 END) AS critical_count
		FROM 
			alarm_histories
		WHERE 
			created_at >= CONVERT_TZ(NOW(), '+00:00', 'Asia/Jakarta') - INTERVAL 24 HOUR;
		`).Scan(&AlarmCount).Error
		if err != nil {
			return nil, err
		}
		res := GetAlarmCountRes{
			WarningCount:  AlarmCount.WarningCount,
			CriticalCount: AlarmCount.CriticalCount,
		}

		return &res, nil
	}
}
