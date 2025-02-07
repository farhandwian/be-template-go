package gateway

import (
	"context"
	"os"
	"shared/core"
	"shared/helper/grafana"
	"strconv"
)

type GrafanaUpsertReq struct {
	IsUpdate  bool
	Title     string
	Uid       string
	Duration  string
	RuleGroup string
	Query     string
	// TableName string
	// Items     []grafana.WhereItem
	// OrderBy   string
}

type GrafanaUpsertRes struct {
	UID string
}

type GrafanaUpsert = core.ActionHandler[GrafanaUpsertReq, GrafanaUpsertRes]

func ImplGrafanaUpsert(clientx *grafana.GrafanaClient) GrafanaUpsert {
	return func(ctx context.Context, req GrafanaUpsertReq) (*GrafanaUpsertRes, error) {

		client := grafana.NewGrafanaClient(os.Getenv("GRAFANA_URL"), os.Getenv("GRAFANA_API_KEY"))

		orgId, err := strconv.Atoi(os.Getenv("GRAFANA_ORG_ID"))
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		if err := client.UpsertAlertRule(&grafana.AlertRulesCreateRequest{
			IsUpdate:      req.IsUpdate,
			OrgID:         orgId,
			FolderUID:     os.Getenv("GRAFANA_FOLDER_UID"),
			Webhook:       os.Getenv("GRAFANA_WEBHOOK"),
			DatasourceUid: os.Getenv("GRAFANA_DATASOURCE_ID"),
			RuleGroup:     req.RuleGroup,
			Title:         req.Title,
			Uid:           req.Uid,
			For:           req.Duration,
			Query:         req.Query,
			// OrderBy:       req.OrderBy,
			// TableName:     req.TableName,
			// Items:         req.Items,
		}); err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GrafanaUpsertRes{
			UID: req.Uid,
		}, nil
	}
}
