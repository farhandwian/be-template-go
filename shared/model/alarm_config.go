package model

import (
	"fmt"

	"gorm.io/datatypes"
)

type AlarmConfigID string

type AlarmConfig struct {
	ID          AlarmConfigID `json:"id"`
	ChannelId   int           `json:"channel_id"`
	ChannelName string        `json:"channel_name"`
	// DoorId            int                 `json:"door_id"`
	// DoorName          string              `json:"door_name"`
	Priority          AlarmConfigPriority `json:"priority"`
	Metric            AlarmMetric         `json:"metric"`
	ConditionOperator AlarmOperator       `json:"condition_operator"`
	ConditionValue    float64             `json:"condition_value"`
	ConditionDuration int                 `json:"condition_duration"`
	ConditionUnit     UnitAlarm           `json:"condition_unit"`
	ReceiverAdam      bool                `json:"receiver_adam"`
	ReceiverHawa      bool                `json:"receiver_hawa"`
	ReceiverBigboard  bool                `json:"receiver_bigboard"`
	ReceiverDashboard bool                `json:"receiver_dashboard"`
	ReceiverWhatsapps datatypes.JSON      `json:"receiver_whatsapps"`
	ReceiverEmails    datatypes.JSON      `json:"receiver_emails"`
}

type AlarmConfigPriority string

const (
	AlarmConfigPriorityWarning  = "warning"
	AlarmConfigPriorityCritical = "critical"
)

func (a AlarmConfigPriority) Validate() error {
	if a != AlarmConfigPriorityWarning && a != AlarmConfigPriorityCritical {
		return fmt.Errorf("priority must be one of warning or critical")
	}
	return nil
}

type UnitAlarm string

const (
	UnitSecond UnitAlarm = "seconds"
	UnitMinute UnitAlarm = "minutes"
	UnitHour   UnitAlarm = "hours"
)

func (a UnitAlarm) Validate() error {
	if a != UnitSecond && a != UnitMinute && a != UnitHour {
		return fmt.Errorf("unit must be one of seconds, minutes or hours")
	}
	return nil
}

type AlarmOperator string

const (
	OperatorLessThan    AlarmOperator = "less_than"
	OperatorGreaterThan AlarmOperator = "greater_than"
	OperatorEqualTo     AlarmOperator = "equals"
)

func (a AlarmOperator) Validate() error {
	if a != OperatorLessThan && a != OperatorGreaterThan && a != OperatorEqualTo {
		return fmt.Errorf("operator must be one of less_than, greater_than or equals")
	}
	return nil
}

type AlarmMetric string

const (
	Debit         AlarmMetric = "debit_threshold"
	CurahHujan    AlarmMetric = "pch_threshold"
	DugaAir       AlarmMetric = "pda_threshold"
	TmaThreshold  AlarmMetric = "tma_threshold"
	TmaPrediction AlarmMetric = "tma_prediction"
)

func (a AlarmMetric) Validate() error {
	if a != Debit && a != TmaPrediction && a != TmaThreshold && a != CurahHujan && a != DugaAir {
		return fmt.Errorf("metric must be one of debit_threshold, pch_threshold, pda_threshold, tma_threshold or tma_prediction")
	}
	return nil
}
