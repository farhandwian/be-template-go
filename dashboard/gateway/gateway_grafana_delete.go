package gateway

import (
	"context"
	"os"
	"shared/core"
	"shared/helper/grafana"
	"shared/model"
)

type AlarmConfigXDeleteReq struct {
	UID model.AlarmConfigID
}

type AlarmConfigXDeleteRes struct{}

type AlarmConfigXDelete = core.ActionHandler[AlarmConfigXDeleteReq, AlarmConfigXDeleteRes]

func ImplAlarmConfigXDelete(clientx *grafana.GrafanaClient) AlarmConfigXDelete {
	return func(ctx context.Context, req AlarmConfigXDeleteReq) (*AlarmConfigXDeleteRes, error) {

		client := grafana.NewGrafanaClient(os.Getenv("GRAFANA_URL"), os.Getenv("GRAFANA_API_KEY"))

		if err := client.DeleteAlertRule(string(req.UID)); err != nil {
			return nil, err
		}

		return &AlarmConfigXDeleteRes{}, nil
	}
}
