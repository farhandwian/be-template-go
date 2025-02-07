package grafana

import "net/http"

type GrafanaClient struct {
	baseURL    string
	authToken  string
	httpClient *http.Client
}

func NewGrafanaClient(baseURL, authToken string) *GrafanaClient {
	return &GrafanaClient{
		baseURL:    baseURL,
		authToken:  authToken,
		httpClient: &http.Client{},
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
	TraceID string `json:"traceID"`
}

type AlertRulesResponse []AlertRuleStruct

type Item struct {
	FieldName  string `json:"fieldName"`
	FieldValue any    `json:"fieldValue"`
}

type WhereItem struct {
	FieldName  string `json:"fieldName"`
	Operator   string `json:"operator"`
	FieldValue any    `json:"fieldValue"`
}

type AlertRuleStruct struct {
	IsUpdate      bool   `json:"isUpdate"`  //
	Title         string `json:"title"`     //
	RuleGroup     string `json:"ruleGroup"` //
	OrgID         int    `json:"orgId"`     //
	Uid           string `json:"uid"`
	For           string `json:"for"`
	FolderUID     string `json:"folderUID"`
	Webhook       string `json:"webhook"`
	DatasourceUid string `json:"datasourceUid"`
	Query         string `json:"query"`
	// OrderBy       string      `json:"orderBy"`
	// Items         []WhereItem `json:"item"`
	// TableName     string      `json:"tableName"`
	// ConditionField string      `json:"conditionField"`
}

type ThresholdEvaluatorType string

const (
	GREATER_THAN  ThresholdEvaluatorType = "gt"
	LESS_THAN     ThresholdEvaluatorType = "lt"
	WITHIN_RANGE  ThresholdEvaluatorType = "within_range"
	OUTSIDE_RANGE ThresholdEvaluatorType = "outside_range"
)

// ====

type AlertRule struct {
	UID          string               `json:"uid"`
	OrgID        int                  `json:"orgId"`
	Title        string               `json:"title"`
	RuleGroup    string               `json:"ruleGroup"`
	FolderUID    string               `json:"folderUID"`
	Condition    string               `json:"condition"`
	Data         []AlertData          `json:"data"`
	NoDataState  string               `json:"noDataState"`
	ExecErrState string               `json:"execErrState"`
	For          string               `json:"for"`
	IsPaused     bool                 `json:"isPaused"`
	Notification NotificationSettings `json:"notification_settings"`
	// Annotations  map[string]string    `json:"annotations"`
	// Labels       map[string]string    `json:"labels"`
}

type AlertData struct {
	RefID             string            `json:"refId"`
	RelativeTimeRange RelativeTimeRange `json:"relativeTimeRange"`
	DatasourceUID     string            `json:"datasourceUid"`
	Model             AlertDataModel    `json:"model"`
}

type RelativeTimeRange struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type AlertDataModel struct {
	Datasource    DatasourceInfo `json:"datasource"`
	EditorMode    string         `json:"editorMode"`
	Format        string         `json:"format"`
	IntervalMs    int            `json:"intervalMs"`
	MaxDataPoints int            `json:"maxDataPoints"`
	RawQuery      bool           `json:"rawQuery"`
	RawSQL        string         `json:"rawSql"`
	RefID         string         `json:"refId"`
	SQL           SQLConfig      `json:"sql"`
}

type DatasourceInfo struct {
	Type string `json:"type"`
	UID  string `json:"uid"`
}

type SQLConfig struct {
	Columns []SQLColumn     `json:"columns"`
	GroupBy []GroupByConfig `json:"groupBy"`
	Limit   int             `json:"limit"`
}

type ColumnParameter struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type SQLColumn struct {
	Alias      string            `json:"alias"`
	Parameters []ColumnParameter `json:"parameters"`
	Type       string            `json:"type"`
}
type GroupByConfig struct {
	Property PropertyConfig `json:"property"`
	Type     string         `json:"type"`
}

type PropertyConfig struct {
	Type string `json:"type"`
}

type NotificationSettings struct {
	Receiver string `json:"receiver"`
}
