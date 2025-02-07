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

type AlarmConfigUpdateReq struct {
	AlarmConfigCreateReq
	ID model.AlarmConfigID `json:"-"`
}

type AlarmConfigUpdateRes struct {
	ID model.AlarmConfigID `json:"id"`
}

type AlarmConfigUpdate = core.ActionHandler[AlarmConfigUpdateReq, AlarmConfigUpdateRes]

func ImplAlarmConfigUpdate(
	getOne sg.AlarmConfigGetOne,
	save sg.AlarmConfigSave,
	callGrafana gateway.GrafanaUpsert,
	getChannelName gateway.WaterChannelDoorByID,
	getDoorName gateway.DeviceByWaterChannelDoorId,
) AlarmConfigUpdate {
	return func(ctx context.Context, req AlarmConfigUpdateReq) (*AlarmConfigUpdateRes, error) {

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

		res, err := getOne(ctx, sg.AlarmConfigGetOneReq{ID: req.ID})
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

		res.Item.ChannelId = req.ChannelId
		res.Item.ChannelName = channelName
		res.Item.Priority = req.Priority
		res.Item.Metric = req.Metric
		res.Item.ConditionOperator = req.ConditionOperator
		res.Item.ConditionValue = req.ConditionValue
		res.Item.ConditionDuration = req.ConditionDuration
		res.Item.ConditionUnit = req.ConditionUnit
		res.Item.ReceiverAdam = req.ReceiverAdam
		res.Item.ReceiverHawa = req.ReceiverHawa
		res.Item.ReceiverDashboard = req.ReceiverDashboard
		res.Item.ReceiverWhatsapps = helper.ToDataTypeJSON(req.ReceiverWhatsapps...)
		res.Item.ReceiverEmails = helper.ToDataTypeJSON(req.ReceiverEmails...)

		if _, err := save(ctx, sg.AlarmConfigSaveReq{AlarmConfig: res.Item}); err != nil {
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

		if _, err = callGrafana(ctx, gateway.GrafanaUpsertReq{
			//
			IsUpdate:  true,
			Title:     title,
			Uid:       string(res.Item.ID),
			Duration:  dur,
			RuleGroup: ruleGroup,
			Query:     getQuery(req.AlarmConfigCreateReq),
		}); err != nil {
			return nil, err
		}

		return &AlarmConfigUpdateRes{ID: res.Item.ID}, nil
	}
}
