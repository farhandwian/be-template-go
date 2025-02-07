package grafana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AlertRulesCreateRequest = AlertRuleStruct

func GetPayloadAlertRule(req *AlertRulesCreateRequest) *AlertRule {
	return &AlertRule{
		UID:          req.Uid,                                     //
		Title:        req.Title,                                   //
		RuleGroup:    req.RuleGroup,                               //
		FolderUID:    req.FolderUID,                               //
		For:          req.For,                                     //
		OrgID:        req.OrgID,                                   //
		Notification: NotificationSettings{Receiver: req.Webhook}, //
		Data: []AlertData{
			{
				DatasourceUID: req.DatasourceUid, //
				Model: AlertDataModel{
					RawSQL:        req.Query, // generateSQLQuery(*req),
					RawQuery:      true,
					EditorMode:    "code",  //
					Format:        "table", //
					IntervalMs:    1000,    //
					MaxDataPoints: 43200,   //
					RefID:         "A",     //
					SQL: SQLConfig{
						Columns: []SQLColumn{{Parameters: []ColumnParameter{}, Type: "function"}},
						GroupBy: []GroupByConfig{{Property: PropertyConfig{Type: "string"}, Type: "groupBy"}},
						Limit:   50,
					},
				},
				RefID:             "A",                                 //
				RelativeTimeRange: RelativeTimeRange{From: 600, To: 0}, //
			},
		},
		Condition:    "A",     //
		NoDataState:  "OK",    //
		ExecErrState: "Error", //
		IsPaused:     false,   //
	}
}

func (c *GrafanaClient) UpsertAlertRule(req *AlertRulesCreateRequest) error {

	payload := GetPayloadAlertRule(req)

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling request: %w", err)
	}

	method := http.MethodPost
	url := fmt.Sprintf("%s/api/v1/provisioning/alert-rules", c.baseURL)
	if req.IsUpdate {
		method = http.MethodPut
		url = fmt.Sprintf("%s/api/v1/provisioning/alert-rules/%s", c.baseURL, req.Uid)
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authToken))

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if req.IsUpdate {
		if resp.StatusCode != http.StatusOK {

			var errResp ErrorResponse
			if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
				return fmt.Errorf("error decoding error response: %w", err)
			}
			return fmt.Errorf("API error: %s", errResp.Message)
		}
	} else {
		if resp.StatusCode != http.StatusCreated {

			var errResp ErrorResponse
			if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
				return fmt.Errorf("error decoding error response: %w", err)
			}
			return fmt.Errorf("API error: %s", errResp.Message)
		}
	}

	return nil
}

// func generateSQLQuery(alert AlertRulesCreateRequest) string {
// 	query := fmt.Sprintf("SELECT * FROM %s WHERE %s ORDER BY %s DESC LIMIT 1",
// 		alert.TableName,
// 		strings.Trim(formatItems(alert.Items), "()"),
// 		alert.OrderBy,
// 	)
// 	return query
// }

// func formatItems(items []WhereItem) string {
// 	if len(items) == 0 {
// 		return ""
// 	}

// 	var result strings.Builder
// 	result.WriteString("(")

// 	for i, item := range items {
// 		if i > 0 {
// 			result.WriteString(" AND ")
// 		}
// 		result.WriteString(fmt.Sprintf("%s %s %v", item.FieldName, item.Operator, item.FieldValue))
// 	}

// 	result.WriteString(")")
// 	return result.String()
// }
