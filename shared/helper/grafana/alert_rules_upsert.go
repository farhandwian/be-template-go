package grafana

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strings"
// )

// type AlertRulesCreateRequest = AlertRuleStruct

// func (c *GrafanaClient) UpsertAlertRule(req *AlertRulesCreateRequest) error {

// 	mapStringAny := make([]map[string]any, 0)

// 	for _, item := range req.Items {

// 		mapStringAny = append(mapStringAny, map[string]any{
// 			// "id": item.Id,
// 			"properties": map[string]any{
// 				"field":      item.FieldName,
// 				"fieldSrc":   "field",
// 				"operator":   "equal",
// 				"value":      []any{item.FieldValue},
// 				"valueError": []any{nil},
// 				"valueSrc":   []string{"value"},
// 				"valueType":  []string{"number"},
// 			},
// 			"type": "rule",
// 		})

// 	}

// 	whereString := formatItems(req.Items)

// 	payload := map[string]any{
// 		"title":        req.Title,     // DONE
// 		"ruleGroup":    req.RuleGroup, // DONE
// 		"orgID":        req.OrgID,     // DONE
// 		"uid":          req.Uid,       // DONE
// 		"for":          req.For,       // DONE
// 		"folderUID":    req.FolderUID, // DONE
// 		"noDataState":  "NoData",      // DONE
// 		"isPaused":     false,         // DONE
// 		"execErrState": "Error",       // DONE
// 		"condition":    "C",           // DONE
// 		"notification_settings": map[string]any{
// 			"receiver": req.Webhook,
// 		},
// 		"data": []map[string]any{
// 			{ // First data element (A)
// 				"refId":         "A",               // DONE
// 				"datasourceUid": req.DatasourceUid, // DONE
// 				"relativeTimeRange": map[string]int{
// 					"from": 600,
// 					"to":   0,
// 				},
// 				"model": map[string]any{
// 					"editorMode":    "builder",
// 					"format":        "time_series",
// 					"intervalMs":    1000,
// 					"maxDataPoints": 43200,
// 					"rawSql":        generateSQLQuery(*req),
// 					"refId":         "A",
// 					"table":         req.TableName,
// 					"sql": map[string]any{
// 						"columns": []map[string]any{
// 							{
// 								"alias": "\"time\"",
// 								"parameters": []map[string]any{
// 									{"name": req.TimeAlias, "type": "functionParameter"},
// 								},
// 								"type": "function",
// 							},
// 							{
// 								"alias": "\"value\"",
// 								"parameters": []map[string]any{
// 									{"name": req.ValueAlias, "type": "functionParameter"},
// 								},
// 								"type": "function",
// 							},
// 						},
// 						"limit": req.Limit,
// 						"orderBy": map[string]any{
// 							"property": map[string]any{
// 								"name": []string{req.OrderBy},
// 								"type": "string",
// 							},
// 							"type": "property",
// 						},
// 						"orderByDirection": "DESC",
// 						"whereJsonTree": map[string]any{
// 							// "id":        req.JsonTreeId,
// 							"children1": mapStringAny,
// 							"type":      "group",
// 						},
// 						"whereString": whereString,
// 					},
// 				}, // DONE
// 			},
// 			{ // Second data element (B)
// 				"refId":         "B",        // DONE
// 				"queryType":     "",         // DONE
// 				"datasourceUid": "__expr__", // DONE
// 				"relativeTimeRange": map[string]int{
// 					"from": 0,
// 					"to":   0,
// 				}, // DONE
// 				"model": map[string]any{
// 					"conditions": []map[string]any{
// 						{
// 							"evaluator": map[string]any{
// 								"params": []any{},
// 								"type":   "gt",
// 							},
// 							"operator": map[string]any{
// 								"type": "and",
// 							},
// 							"query": map[string]any{
// 								"params": []string{"B"},
// 							},
// 							"reducer": map[string]any{
// 								"params": []any{},
// 								"type":   "last",
// 							},
// 							"type": "query",
// 						},
// 					},
// 					"expression":    "A",
// 					"intervalMs":    1000,
// 					"maxDataPoints": 43200,
// 					"reducer":       "mean",
// 					"refId":         "B",
// 					"type":          "reduce",
// 				},
// 			},
// 			{ // Third data element (C)
// 				"refId":         "C",        // DONE
// 				"queryType":     "",         // DONE
// 				"datasourceUid": "__expr__", // DONE
// 				"relativeTimeRange": map[string]int{
// 					"from": 0,
// 					"to":   0,
// 				}, // DONE
// 				"model": map[string]any{
// 					"conditions": []map[string]any{
// 						{
// 							"evaluator": map[string]any{
// 								"params": req.ThresholdEvaluatorValue,
// 								"type":   req.ThresholdEvaluatorType,
// 							},
// 							"operator": map[string]any{
// 								"type": "and",
// 							},
// 							"query": map[string]any{
// 								"params": []string{"C"},
// 							},
// 							"reducer": map[string]any{
// 								"params": []any{},
// 								"type":   "last",
// 							},
// 							"type": "query",
// 						},
// 					},
// 					"expression":    "B",
// 					"intervalMs":    1000,
// 					"maxDataPoints": 43200,
// 					"refId":         "C",
// 					"type":          "threshold",
// 				},
// 			},
// 		}, // DONE
// 	}

// 	jsonData, err := json.Marshal(payload)
// 	if err != nil {
// 		return fmt.Errorf("error marshaling request: %w", err)
// 	}

// 	method := http.MethodPost
// 	url := fmt.Sprintf("%s/api/v1/provisioning/alert-rules", c.baseURL)
// 	if req.IsUpdate {
// 		method = http.MethodPut
// 		url = fmt.Sprintf("%s/api/v1/provisioning/alert-rules/%s", c.baseURL, req.Uid)
// 	}

// 	request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return fmt.Errorf("error creating request: %w", err)
// 	}

// 	request.Header.Set("Content-Type", "application/json")
// 	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authToken))

// 	resp, err := c.httpClient.Do(request)
// 	if err != nil {
// 		return fmt.Errorf("error making request: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	if req.IsUpdate {
// 		if resp.StatusCode != http.StatusOK {

// 			var errResp ErrorResponse
// 			if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
// 				return fmt.Errorf("error decoding error response: %w", err)
// 			}
// 			return fmt.Errorf("API error: %s", errResp.Message)
// 		}
// 	} else {
// 		if resp.StatusCode != http.StatusCreated {

// 			var errResp ErrorResponse
// 			if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
// 				return fmt.Errorf("error decoding error response: %w", err)
// 			}
// 			return fmt.Errorf("API error: %s", errResp.Message)
// 		}
// 	}

// 	return nil
// }

// func formatItems(items []Item) string {
// 	if len(items) == 0 {
// 		return ""
// 	}

// 	var result strings.Builder
// 	result.WriteString("(")

// 	for i, item := range items {
// 		if i > 0 {
// 			result.WriteString(" AND ")
// 		}
// 		result.WriteString(fmt.Sprintf("%s = %v", item.FieldName, item.FieldValue))
// 	}

// 	result.WriteString(")")
// 	return result.String()
// }

// func generateSQLQuery(alert AlertRulesCreateRequest) string {
// 	// Base query structure
// 	query := fmt.Sprintf(
// 		"SELECT %s AS \"time\", %s AS \"value\" FROM %s",
// 		alert.TimeAlias,
// 		alert.ValueAlias,
// 		alert.TableName,
// 	)

// 	// Add WHERE clause if Items exist
// 	whereClause := formatItems(alert.Items)
// 	if whereClause != "" {
// 		query += " WHERE " + strings.Trim(whereClause, "()")
// 	}

// 	// Add ORDER BY if specified
// 	if alert.OrderBy != "" {
// 		query += fmt.Sprintf(" ORDER BY %s DESC", alert.OrderBy)
// 	}

// 	// Add LIMIT if specified
// 	if alert.Limit > 0 {
// 		query += fmt.Sprintf(" LIMIT %d", alert.Limit)
// 	}

// 	return query
// }
