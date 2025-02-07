package usecase

import (
	"context"
	"dashboard/gateway"
	"fmt"
	"shared/core"
	sg "shared/gateway"
	"shared/helper"
	"shared/model"
)

type AlarmConfigCreateReq struct {
	ChannelId          int                       `json:"channel_id"`
	Priority           model.AlarmConfigPriority `json:"priority"`
	Metric             model.AlarmMetric         `json:"metric"`
	ConditionOperator  model.AlarmOperator       `json:"condition_operator"`
	ConditionValue     float64                   `json:"condition_value"`
	ConditionDuration  int                       `json:"condition_duration"`
	ConditionUnit      model.UnitAlarm           `json:"condition_unit"`
	PredictionInterval int                       `json:"prediction_interval"`
	PredictionUnit     model.UnitAlarm           `json:"prediction_unit"`
	ReceiverAdam       bool                      `json:"receiver_adam"`
	ReceiverHawa       bool                      `json:"receiver_hawa"`
	ReceiverDashboard  bool                      `json:"receiver_dashboard"`
	ReceiverWhatsapps  []string                  `json:"receiver_whatsapps"`
	ReceiverEmails     []string                  `json:"receiver_emails"`
}

type AlarmConfigCreateRes struct {
	ID model.AlarmConfigID `json:"id"`
}

type AlarmConfigCreate = core.ActionHandler[AlarmConfigCreateReq, AlarmConfigCreateRes]

func ImplAlarmConfigCreate(
	generateId gateway.GenerateId,
	save sg.AlarmConfigSave,
	callGrafana gateway.GrafanaUpsert,
	getChannelName gateway.WaterChannelDoorByID,
	getDoorName gateway.DeviceByWaterChannelDoorId,
) AlarmConfigCreate {
	return func(ctx context.Context, req AlarmConfigCreateReq) (*AlarmConfigCreateRes, error) {

		if err := req.Metric.Validate(); err != nil {
			return nil, err
		}

		if err := req.ConditionOperator.Validate(); err != nil {
			return nil, err
		}

		if err := req.ConditionUnit.Validate(); err != nil {
			return nil, err
		}

		if err := req.Priority.Validate(); err != nil {
			return nil, err
		}

		if req.Metric == model.TmaPrediction && req.PredictionInterval == 0 {
			return nil, fmt.Errorf("prediction_interval must be greater than 0")
		}

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		channel, err := getChannelName(ctx, gateway.WaterChannelDoorByIDReq{ID: req.ChannelId})
		if err != nil {
			return nil, err
		}

		if len(channel.Items) == 0 {
			return nil, fmt.Errorf("channel with id %v not found", req.ChannelId)
		}

		channelName := channel.Items[0].Name

		obj := model.AlarmConfig{
			ID:                model.AlarmConfigID(genObj.RandomId),
			ChannelId:         req.ChannelId,
			ChannelName:       channelName,
			Priority:          req.Priority,
			Metric:            req.Metric,
			ConditionOperator: req.ConditionOperator,
			ConditionValue:    req.ConditionValue,
			ConditionDuration: req.ConditionDuration,
			ConditionUnit:     req.ConditionUnit,
			ReceiverAdam:      req.ReceiverAdam,
			ReceiverHawa:      req.ReceiverHawa,
			ReceiverDashboard: req.ReceiverDashboard,
			ReceiverWhatsapps: helper.ToDataTypeJSON(req.ReceiverWhatsapps...),
			ReceiverEmails:    helper.ToDataTypeJSON(req.ReceiverEmails...),
		}

		if _, err = save(ctx, sg.AlarmConfigSaveReq{AlarmConfig: &obj}); err != nil {
			return nil, err
		}

		dur, err := getDuration(req.ConditionUnit, req.ConditionDuration)
		if err != nil {
			return nil, err
		}

		title, err := getTitle(req.Metric, req.ConditionOperator, req.ChannelId, req.ConditionValue, dur)
		if err != nil {
			return nil, err
		}

		ruleGroup, err := getRuleGroup(req.Metric)
		if err != nil {
			return nil, err
		}

		if _, err := callGrafana(ctx, gateway.GrafanaUpsertReq{
			//
			IsUpdate:  false,
			Title:     title,
			Uid:       string(obj.ID),
			Duration:  dur,
			RuleGroup: ruleGroup,
			Query:     getQuery(req),
		}); err != nil {
			return nil, err
		}

		return &AlarmConfigCreateRes{ID: obj.ID}, nil
	}
}

func getQuery(req AlarmConfigCreateReq) string {

	var mapOperator = map[model.AlarmOperator]string{
		model.OperatorGreaterThan: ">=",
		model.OperatorLessThan:    "<=",
		model.OperatorEqualTo:     "=",
	}

	if req.Metric == model.CurahHujan {
		return fmt.Sprintf(`
			SELECT rain FROM (
				SELECT rain FROM hydrology_rain_hourlies 
				WHERE 1 = 1
					AND rain_post_id = %d 
					AND count > 0 
					AND updated_at > current_date
				ORDER BY hour DESC LIMIT 1
			)
			WHERE rain %s %f;
		`, req.ChannelId, mapOperator[req.ConditionOperator], req.ConditionValue)
	}

	if req.Metric == model.DugaAir {
		return fmt.Sprintf(`
			SELECT water_level FROM (
				SELECT water_level FROM hydrology_water_level_telemetries 
				WHERE 1 = 1
					AND water_level_post_id = %d
				ORDER BY timestamp DESC LIMIT 1
			) 
			WHERE water_level %s %f;
		`, req.ChannelId, mapOperator[req.ConditionOperator], req.ConditionValue)
	}

	if req.Metric == model.Debit {
		return fmt.Sprintf(`
			SELECT actual_debit FROM (
				SELECT actual_debit FROM actual_debit_data 
				WHERE 1 = 1
					AND water_channel_door_id = %d
				ORDER BY timestamp DESC LIMIT 1
			)
			WHERE actual_debit %s %f;
		`, req.ChannelId, mapOperator[req.ConditionOperator], req.ConditionValue)
	}

	if req.Metric == model.TmaThreshold {
		return fmt.Sprintf(`
			SELECT water_level FROM (
				SELECT water_level FROM water_surface_elevation_data
				WHERE 1 = 1
					AND water_channel_door_id = %d
					AND status = true
				ORDER BY timestamp DESC LIMIT 1
			)
			WHERE water_level %s %f;
		`, req.ChannelId, mapOperator[req.ConditionOperator], req.ConditionValue)
	}

	if req.Metric == model.TmaPrediction {

		interval := fmt.Sprintf("%d %s", req.PredictionInterval, req.PredictionUnit)

		return fmt.Sprintf(`
			SELECT water_level FROM (
				SELECT water_level FROM forecast_water_surface_elevations
				WHERE 1 = 1
					AND water_channel_door_id = %d
					AND timestamp > NOW() + INTERVAL '%s'
				ORDER BY timestamp DESC LIMIT 1
			)
			WHERE water_level %s %f;
		`, req.ChannelId, interval, mapOperator[req.ConditionOperator], req.ConditionValue)
	}

	return ""
}

func getRuleGroup(ma model.AlarmMetric) (string, error) {

	var mapRuleGroup = map[model.AlarmMetric]string{
		model.Debit:         "alarm_debit_rulegroup",
		model.CurahHujan:    "alarm_pch_rulegroup",
		model.DugaAir:       "alarm_pda_rulegroup",
		model.TmaPrediction: "alarm_tma_prediction_rulegroup",
		model.TmaThreshold:  "alarm_tma_threshold_rulegroup",
	}

	a, ok := mapRuleGroup[ma]
	if !ok {
		return "", fmt.Errorf("metric %s not found", ma)
	}

	return a, nil
}

func getDuration(runit model.UnitAlarm, dur int) (string, error) {
	var mapUnit = map[model.UnitAlarm]string{
		model.UnitHour:   "h",
		model.UnitMinute: "m",
		model.UnitSecond: "s",
	}

	a, ok := mapUnit[runit]
	if !ok {
		return "", fmt.Errorf("unit %s not found", runit)
	}
	return fmt.Sprintf("%d%s", dur, a), nil
}

func getTitle(metric model.AlarmMetric, alarmOperator model.AlarmOperator, channelId int, val float64, dur string) (string, error) {

	var mapOperator = map[model.AlarmOperator]string{
		model.OperatorGreaterThan: "gt",
		model.OperatorLessThan:    "lt",
		model.OperatorEqualTo:     "eq",
	}

	a, ok := mapOperator[alarmOperator]
	if !ok {
		return "", fmt.Errorf("operator %s not found", alarmOperator)
	}

	return fmt.Sprintf("%s_%d_%s_%.0f_%s", metric, channelId, a, val, dur), nil
}
