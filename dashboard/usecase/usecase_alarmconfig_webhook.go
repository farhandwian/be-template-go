package usecase

import (
	"context"
	"dashboard/gateway"
	"encoding/json"
	"fmt"
	"shared/constant"
	"shared/core"
	sg "shared/gateway"
	"shared/model"
	sharedModel "shared/model"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AlarmConfigWebhookReq struct {
	Alerts []model.Alert
	Now    time.Time
}

type AlarmConfigWebhookRes struct {
	ErrorMessage []error
}

type CompleteAlarmConfig struct {
	model.AlarmConfig
	model.Alert
}

type AlarmHistory struct {
	ChannelId         string
	ChannelName       string
	Priority          int
	Metric            string
	ConditionOperator string
	ConditionValue    float64
	ConditionDuration int
	ConditionUnit     string
	AlarmConfigUID    string
	CreatedAt         string
	Status            string
	Values            float64
	ID                string
	DoorName          string
}

type AlarmConfigWebhook = core.ActionHandler[AlarmConfigWebhookReq, AlarmConfigWebhookRes]

func ImplAlarmConfigWebhook(
	getOne sg.AlarmConfigGetOne,
	sseBigboard sg.SendSSEMessage,
	sseDashboard sg.SendSSEMessage,
	email gateway.SendEmail,
	whatsapp gateway.SendWhatsApp,
	saveHistory gateway.AlarmHistorySave,
	createActivityMonitoring sg.CreateActivityMonitoringGateway,
	generateId gateway.GenerateId,

) AlarmConfigWebhook {
	return func(ctx context.Context, req AlarmConfigWebhookReq) (*AlarmConfigWebhookRes, error) {

		// buat nampung multiple error
		var errs []error

		// seluruh informasi Alert
		alarmConfiMap := map[string]CompleteAlarmConfig{}

		alarmConfigArray := make([]model.AlarmHistory, 0)

		// terima multiple alert dari webhook grafana
		for _, a := range req.Alerts {

			res, err := getOne(ctx, sg.AlarmConfigGetOneReq{ID: model.AlarmConfigID(a.UID)})
			if err != nil {
				return nil, err
			}

			alarmConfiMap[a.UID] = CompleteAlarmConfig{
				AlarmConfig: *res.Item,
				Alert:       a,
			}

			id, _ := generateId(ctx, gateway.GenerateIdReq{})

			alarmConfigArray = append(alarmConfigArray, model.AlarmHistory{
				ChannelId:         res.Item.ChannelId,
				ChannelName:       res.Item.ChannelName,
				Priority:          res.Item.Priority,
				Metric:            res.Item.Metric,
				ConditionOperator: res.Item.ConditionOperator,
				ConditionValue:    res.Item.ConditionValue,
				ConditionDuration: res.Item.ConditionDuration,
				ConditionUnit:     res.Item.ConditionUnit,
				AlarmConfigUID:    model.AlarmConfigID(a.UID),
				CreatedAt:         req.Now,
				Status:            model.AlarmHistoryStatus(a.Status),
				Values:            a.Values,
				ID:                model.AlarmHistoryID(id.RandomId),
			})

		}

		if _, err := saveHistory(ctx, gateway.AlarmHistorySaveReq{AlarmHistory: alarmConfigArray}); err != nil {
			return nil, err
		}

		for _, m := range alarmConfiMap {

			if m.ReceiverAdam {

				if _, err := sseBigboard(ctx, sg.SendSSEMessageReq{
					Subject:      "alert",
					FunctionName: "alert",
					Data:         alarmConfigArray,
				}); err != nil {
					errs = append(errs, err)
				}
				if err := processAlerts(ctx, alarmConfigArray, createActivityMonitoring); err != nil {
					return nil, err
				}
			}

			if m.ReceiverDashboard {

				if _, err := sseDashboard(ctx, sg.SendSSEMessageReq{
					Subject:      "alert",
					FunctionName: "alert",
					Data:         alarmConfigArray,
				}); err != nil {
					errs = append(errs, err)
				}
			}

			emails, err := jsonToStringSlice(m.ReceiverEmails)
			if err != nil {
				errs = append(errs, err)
			} else {

				if _, err := email(ctx, gateway.SendEmailReq{
					EmailRecipients: emails,
					Subject:         "Alert !!!", // TODO lengkapi judul
					Body:            model.GenerateGroupedAlertMessagesForEmail(alarmConfigArray),
				}); err != nil {
					errs = append(errs, err)
				}
			}

			whatsapps, err := jsonToStringSlice(m.ReceiverWhatsapps)
			if err != nil {
				errs = append(errs, err)
			} else {

				if _, err := whatsapp(ctx, gateway.SendWhatsAppReq{
					PhoneNumbers: whatsapps,
					Message:      model.GenerateGroupedAlertMessagesForWhatsApp(alarmConfigArray),
				}); err != nil {
					errs = append(errs, err)
				}

			}

		}

		if len(errs) > 0 {
			return nil, mergeErrors(errs)
		}

		return &AlarmConfigWebhookRes{}, nil
	}
}

// MergeErrors menggabungkan slice of errors menjadi satu error.
// Jika slice kosong, return nil.
// Jika hanya ada satu error, return error tersebut.
// Jika ada multiple errors, gabungkan dengan newline separator.
func mergeErrors(errs []error) error {
	// Filter out nil errors
	var validErrs []error
	for _, err := range errs {
		if err != nil {
			validErrs = append(validErrs, err)
		}
	}

	// Handle special cases
	switch len(validErrs) {
	case 0:
		return nil
	case 1:
		return validErrs[0]
	}

	// Build error messages
	var errMsgs []string
	for _, err := range validErrs {
		errMsgs = append(errMsgs, err.Error())
	}

	// Join all error messages with newline
	return fmt.Errorf("%s", strings.Join(errMsgs, "\n"))
}

func jsonToStringSlice(data datatypes.JSON) ([]string, error) {
	if len(data) == 0 {
		return []string{}, nil
	}

	var result []string
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

const (
	template                = "Kondisi %s pada saluran %s pintu %s sudah mencapai %s %.2f %s."
	templateOnlyChannelName = "Kondisi %s pada saluran %s sudah mencapai %s %.2f %s."
	templateOnlyDoorName    = "Kondisi %s pada pintu %s sudah mencapai %s %.2f %s."
)

func getCondition(condition string) string {
	switch condition {
	case "less_than":
		return "kurang dari"
	case "less_than_equal":
		return "kurang dari atau sama dengan"
	case "equal":
		return "sama dengan"
	case "greater_than":
		return "lebih dari"
	case "greater_than_equal":
		return "lebih dari atau sama dengan"
	default:
		return ""
	}
}

func generateMetricName(metric string) string {
	switch metric {
	case "tma_threshold":
		metric = "Tinggi Muka Air"
	case "pda_threshold":
		metric = "Duga Air"
	case "debit_threshold":
		metric = "Debit"
	case "pch_threshold":
		metric = "Curah Hujan"
	case "tma_prediction":
		metric = "Tinggi Muka Air"
	default:
		// No change, keep the original metric name
	}
	return metric
}

func getUnit(metric string) string {
	switch metric {
	case "Tinggi Muka Air":
		metric = "cm"
	case "Duga air":
		metric = "cm"
	case "Debit":
		metric = "lt/dt"
	case "Curah Hujan":
		metric = ""
	}
	return metric
}

// generateAlertMessage generates the alert message based on the data.
func generateAlertMessage(alerts []sharedModel.AlarmHistory) string {
	var messages []string
	for _, alert := range alerts {
		metric := generateMetricName(string(alert.Metric))
		unit := getUnit(metric)
		condition := getCondition(string(alert.ConditionOperator))

		var message string
		if alert.ChannelName != "" {
			message = fmt.Sprintf(
				"Kondisi %s pada saluran %s sudah mencapai %s %.2f %s.",
				metric,
				alert.ChannelName,
				condition,
				alert.ConditionValue,
				unit,
			)
		} else {
			message = fmt.Sprintf(
				"Kondisi %s sudah mencapai %s %.2f %s.",
				metric,
				condition,
				alert.ConditionValue,
				unit,
			)
		}

		messages = append(messages, message)
	}
	return strings.Join(messages, "\n")
}

func processAlerts(ctx context.Context, alarmConfigArray []sharedModel.AlarmHistory, createActivityMonitoring sg.CreateActivityMonitoringGateway) error {
	for _, alert := range alarmConfigArray {

		description := generateAlertMessage([]sharedModel.AlarmHistory{alert})
		_, err := createActivityMonitoring(ctx, sg.CreateActivityMonitoringReq{
			ActivityMonitor: sharedModel.ActivityMonitor{
				ID:           uuid.NewString(),
				UserName:     "",
				Category:     constant.MONITORING_TYPE_ALARM,
				ActivityTime: time.Now(),
				Description:  description,
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
