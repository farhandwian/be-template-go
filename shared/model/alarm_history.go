package model

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type AlarmHistoryID string

type AlarmHistory struct {
	ID                AlarmHistoryID      `json:"id"`
	ChannelId         int                 `json:"channel_id"`
	ChannelName       string              `json:"channel_name"`
	Priority          AlarmConfigPriority `json:"priority"`
	Metric            AlarmMetric         `json:"metric"`
	ConditionOperator AlarmOperator       `json:"condition_operator"`
	ConditionValue    float64             `json:"condition_value"`
	ConditionDuration int                 `json:"condition_duration"`
	ConditionUnit     UnitAlarm           `json:"condition_unit"`
	AlarmConfigUID    AlarmConfigID       `json:"alarm_config_uid"`
	CreatedAt         time.Time           `json:"created_at"`
	Status            AlarmHistoryStatus  `json:"status"`
	Values            float64             `json:"values"`
	// DoorId            int                 `json:"door_id"`
	// DoorName          string              `json:"door_name"`
}

type AlarmHistoryStatus string

const (
	Firing   AlarmHistoryStatus = "firing"
	Resolved AlarmHistoryStatus = "resolved"
)

type alertGroup struct {
	priority AlarmConfigPriority
	alerts   []string
}

func (ah AlarmHistory) generateAlertForWhatsApp() string {
	operator := ""
	switch ah.ConditionOperator {
	case OperatorLessThan:
		operator = "Kurang Dari"
	case OperatorGreaterThan:
		operator = "Lebih Dari"
	case OperatorEqualTo:
		operator = "Sama Dengan"
	}

	// doorInfo := ""
	// if ah.DoorId != 0 {
	// 	doorInfo = fmt.Sprintf("pintu *%s* ", ah.DoorName)
	// }

	// return fmt.Sprintf("%s\nKondisi *%s* pada saluran *%s* %ssudah mencapai *%s* *%.0f cm*",
	return fmt.Sprintf("%s\nKondisi *%s* pada saluran *%s* sudah mencapai *%s* *%.0f cm*",
		ah.CreatedAt.Format("2006-01-02 15:04:05"),
		// ah.Metric, ah.ChannelName, doorInfo, operator, ah.ConditionValue)
		ah.Metric, ah.ChannelName, operator, ah.ConditionValue)
}

func (ah AlarmHistory) generateAlertForEmail() string {
	operator := ""
	switch ah.ConditionOperator {
	case OperatorLessThan:
		operator = "Kurang Dari"
	case OperatorGreaterThan:
		operator = "Lebih Dari"
	case OperatorEqualTo:
		operator = "Sama Dengan"
	}

	// doorInfo := ""
	// if ah.DoorId != 0 {
	// 	doorInfo = fmt.Sprintf(" pintu <strong>%s</strong>", ah.DoorName)
	// }

	// return fmt.Sprintf("%s<br>Kondisi <strong>%s</strong> pada saluran <strong>%s</strong>%s sudah mencapai <strong>%s</strong> <strong>%.0f cm</strong>",
	return fmt.Sprintf("%s<br>Kondisi <strong>%s</strong> pada saluran <strong>%s</strong> sudah mencapai <strong>%s</strong> <strong>%.0f cm</strong>",
		ah.CreatedAt.Format("2006-01-02 15:04:05"),
		// ah.Metric, ah.ChannelName, doorInfo, operator, ah.ConditionValue)
		ah.Metric, ah.ChannelName, operator, ah.ConditionValue)
}

func GenerateGroupedAlertMessagesForWhatsApp(alarms []AlarmHistory) string {
	groups := []alertGroup{
		{priority: AlarmConfigPriorityCritical, alerts: []string{}},
		{priority: AlarmConfigPriorityWarning, alerts: []string{}},
	}

	// Sort alarms
	sort.Slice(alarms, func(i, j int) bool {
		if alarms[i].Priority != alarms[j].Priority {
			return alarms[i].Priority == AlarmConfigPriorityCritical
		}
		return alarms[i].CreatedAt.After(alarms[j].CreatedAt)
	})

	for _, alarm := range alarms {
		if alarm.Status != Firing {
			continue
		}

		for i := range groups {
			if groups[i].priority == alarm.Priority {
				groups[i].alerts = append(groups[i].alerts, alarm.generateAlertForWhatsApp())
				break
			}
		}
	}

	var result strings.Builder
	firstGroup := true

	for _, group := range groups {
		if len(group.alerts) == 0 {
			continue
		}

		if !firstGroup {
			result.WriteString("\n\n")
		}

		switch group.priority {
		case AlarmConfigPriorityCritical:
			result.WriteString("*🚨 Critical Alert! 🚨 ------*\n\n")
		case AlarmConfigPriorityWarning:
			result.WriteString("*⚠️ Warning Alert! ⚠️ ------*\n\n")
		}

		result.WriteString(strings.Join(group.alerts, "\n\n"))
		firstGroup = false
	}

	return result.String()
}

func GenerateGroupedAlertMessagesForEmail(alarms []AlarmHistory) string {
	groups := []alertGroup{
		{priority: AlarmConfigPriorityCritical, alerts: []string{}},
		{priority: AlarmConfigPriorityWarning, alerts: []string{}},
	}

	// Sort alarms
	sort.Slice(alarms, func(i, j int) bool {
		if alarms[i].Priority != alarms[j].Priority {
			return alarms[i].Priority == AlarmConfigPriorityCritical
		}
		return alarms[i].CreatedAt.After(alarms[j].CreatedAt)
	})

	for _, alarm := range alarms {
		if alarm.Status != Firing {
			continue
		}

		for i := range groups {
			if groups[i].priority == alarm.Priority {
				groups[i].alerts = append(groups[i].alerts, alarm.generateAlertForEmail())
				break
			}
		}
	}

	var result strings.Builder
	firstGroup := true

	for _, group := range groups {
		if len(group.alerts) == 0 {
			continue
		}

		if !firstGroup {
			result.WriteString("<br><br>")
		}

		switch group.priority {
		case AlarmConfigPriorityCritical:
			result.WriteString("<strong>🚨 Critical Alert! 🚨</strong><br><br>")
		case AlarmConfigPriorityWarning:
			result.WriteString("<strong>⚠️Warning Alert! ⚠️</strong><br><br>")
		}

		result.WriteString(strings.Join(group.alerts, "<br><br>"))
		firstGroup = false
	}

	return result.String()
}
