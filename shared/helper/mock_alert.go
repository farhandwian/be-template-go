package helper

import (
	"math/rand"
	"shared/model"
	"time"

	"github.com/google/uuid"
)

func GenerateRandomAlarmData(n int) []model.AlarmHistory {

	// Define possible values for random selection
	operators := []model.AlarmOperator{"greater_than", "less_than", "equals"}
	priorities := []model.AlarmConfigPriority{"critical", "warning"}
	metrics := []model.AlarmMetric{"tma", "debit"}
	units := []model.UnitAlarm{"hour", "minute", "second"}

	// Generate random channel names
	channelNames := []string{
		"Regulator Laksel",
		"Bendung Katulampa",
		"Pos Manggarai",
		"Pintu Air Pasar Ikan",
		"Waduk Pluit",
	}

	result := make([]model.AlarmHistory, n)

	for i := 0; i < n; i++ {
		result[i] = model.AlarmHistory{

			ID:                model.AlarmHistoryID(uuid.New().String()),
			ChannelId:         rand.Intn(100) + 1,
			ChannelName:       channelNames[rand.Intn(len(channelNames))],
			Priority:          priorities[rand.Intn(len(priorities))],
			Metric:            metrics[rand.Intn(len(metrics))],
			ConditionOperator: operators[rand.Intn(len(operators))],
			ConditionValue:    3.14,
			ConditionDuration: 1,
			ConditionUnit:     units[rand.Intn(len(units))],
			AlarmConfigUID:    model.AlarmConfigID(uuid.New().String()),
			CreatedAt:         time.Now(),
			Status:            "firing",
			Values:            200.0 + rand.Float64()*20.0,
			// DoorId:            0,
			// DoorName:          "",
		}
	}

	return result
}
