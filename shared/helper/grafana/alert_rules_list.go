package grafana

// type AlertRulesListResponse []AlertRule

// // type AlertRule struct {
// // 	ID            int             `json:"id"`
// // 	Uid           string          `json:"uid"`
// // 	OrgID         int             `json:"orgID"`
// // 	FolderUID     string          `json:"folderUID"`
// // 	RuleGroup     string          `json:"ruleGroup"`
// // 	Title         string          `json:"title"`
// // 	Condition     string          `json:"condition"`
// // 	Data          []AlertRuleData `json:"data"`
// // 	Updated       string          `json:"updated"`
// // 	NoDataState   string          `json:"noDataState"`
// // 	ExecErrState  string          `json:"execErrState"`
// // 	For           AlertRuleFor    `json:"for"`
// // 	Provenance    string          `json:"provenance"`
// // 	IsPaused      bool            `json:"isPaused"`
// // 	Notifications struct {
// // 		Receiver string `json:"receiver"`
// // 	} `json:"notification_settings"`
// // }

// type AlertRuleData struct {
// 	RefID             string             `json:"refId"`
// 	QueryType         string             `json:"queryType"`
// 	RelativeTimeRange RelativeTimeRange  `json:"relativeTimeRange"`
// 	DatasourceUid     string             `json:"datasourceUid"`
// 	Model             AlertRuleDataModel `json:"model"`
// }

// type AlertRuleDataModel struct {
// 	EditorMode    string      `json:"editorMode,omitempty"`
// 	Format        string      `json:"format,omitempty"`
// 	IntervalMs    int         `json:"intervalMs"`
// 	MaxDataPoints int         `json:"maxDataPoints"`
// 	RawSql        string      `json:"rawSql,omitempty"`
// 	RefID         string      `json:"refId"`
// 	SQL           *SQLModel   `json:"sql,omitempty"`
// 	Table         string      `json:"table,omitempty"`
// 	Conditions    []Condition `json:"conditions,omitempty"`
// 	Expression    string      `json:"expression,omitempty"`
// 	Reducer       string      `json:"reducer,omitempty"`
// 	Type          string      `json:"type,omitempty"`
// }

// type SQLModel struct {
// 	Columns          []SQLColumn `json:"columns"`
// 	Limit            int         `json:"limit"`
// 	OrderBy          OrderBy     `json:"orderBy"`
// 	OrderByDirection string      `json:"orderByDirection"`
// 	WhereJsonTree    WhereTree   `json:"whereJsonTree"`
// 	WhereString      string      `json:"whereString"`
// }

// type ColumnParameter struct {
// 	Name string `json:"name"`
// 	Type string `json:"type"`
// }

// type OrderBy struct {
// 	Property struct {
// 		Name []string `json:"name"`
// 		Type string   `json:"type"`
// 	} `json:"property"`
// 	Type string `json:"type"`
// }

// type WhereTree struct {
// 	Children1 []WhereTreeChild `json:"children1"`
// 	ID        string           `json:"id"`
// 	Type      string           `json:"type"`
// }

// type WhereTreeChild struct {
// 	ID         string     `json:"id"`
// 	Properties Properties `json:"properties"`
// 	Type       string     `json:"type"`
// }

// type Properties struct {
// 	Field      string   `json:"field"`
// 	FieldSrc   string   `json:"fieldSrc"`
// 	Operator   string   `json:"operator"`
// 	Value      []int    `json:"value"`
// 	ValueError []any    `json:"valueError"`
// 	ValueSrc   []string `json:"valueSrc"`
// 	ValueType  []string `json:"valueType"`
// }

// type Condition struct {
// 	Evaluator struct {
// 		Params []float64 `json:"params"`
// 		Type   string    `json:"type"`
// 	} `json:"evaluator"`
// 	Operator struct {
// 		Type string `json:"type"`
// 	} `json:"operator"`
// 	Query struct {
// 		Params []string `json:"params"`
// 	} `json:"query"`
// 	Reducer struct {
// 		Params []any  `json:"params"`
// 		Type   string `json:"type"`
// 	} `json:"reducer"`
// 	Type string `json:"type"`
// }

// func (c *GrafanaClient) ListAlertRules_() (AlertRulesListResponse, error) {
// 	req, err := http.NewRequest("GET",
// 		fmt.Sprintf("%s/api/v1/provisioning/alert-rules", c.baseURL),
// 		nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating request: %w", err)
// 	}

// 	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authToken))

// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("error making request: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		var errResp ErrorResponse
// 		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
// 			return nil, fmt.Errorf("error decoding error response: %w", err)
// 		}
// 		return nil, fmt.Errorf("API error: %s", errResp.Message)
// 	}

// 	var rules AlertRulesListResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&rules); err != nil {
// 		return nil, fmt.Errorf("error decoding response: %w", err)
// 	}

// 	return rules, nil
// }

// func MapToAlertRuleResponse(rules AlertRulesListResponse) AlertRulesResponse {
// 	var response AlertRulesResponse

// 	for _, rule := range rules {
// 		ruleResp := AlertRuleStruct{
// 			Title:     rule.Title,
// 			RuleGroup: rule.RuleGroup,
// 			OrgID:     rule.OrgID,
// 			Uid:       rule.UID,
// 			For:       rule.For,
// 			FolderUID: rule.FolderUID,
// 			Webhook:   rule.Notification.Receiver,
// 			// Labels:    rule.Labels,
// 		}

// 		// Get data from first data entry (A)
// 		if len(rule.Data) > 0 {
// 			firstData := rule.Data[0]
// 			ruleResp.DatasourceUid = firstData.DatasourceUID

// 			if firstData.Model.SQL != nil {
// 				ruleResp.Limit = firstData.Model.SQL.Limit
// 				ruleResp.TableName = firstData.Model.Table
// 				// ruleResp.WhereString = firstData.Model.SQL.WhereString

// 				if len(firstData.Model.SQL.OrderBy.Property.Name) > 0 {
// 					ruleResp.OrderBy = firstData.Model.SQL.OrderBy.Property.Name[0]
// 				}

// 				// Get column aliases
// 				for _, col := range firstData.Model.SQL.Columns {
// 					if len(col.Parameters) > 0 {
// 						if col.Alias == "\"time\"" {
// 							ruleResp.TimeAlias = col.Parameters[0].Name
// 						} else if col.Alias == "\"value\"" {
// 							ruleResp.ValueAlias = col.Parameters[0].Name
// 						}
// 					}
// 				}

// 				ruleResp.Items = make([]Item, len(firstData.Model.SQL.WhereJsonTree.Children1))
// 				// Get where tree info
// 				if len(firstData.Model.SQL.WhereJsonTree.Children1) > 0 {
// 					for i, child := range firstData.Model.SQL.WhereJsonTree.Children1 {
// 						// ruleResp.Items[i].Id = child.ID
// 						ruleResp.Items[i].FieldName = child.Properties.Field
// 						if len(child.Properties.Value) > 0 {
// 							ruleResp.Items[i].FieldValue = child.Properties.Value[0]
// 						}
// 					}

// 					// child := firstData.Model.SQL.WhereJsonTree.Children1[0]
// 					// ruleResp.ItemId = child.ID
// 					// ruleResp.ItemFieldName = child.Properties.Field
// 					// if len(child.Properties.Value) > 0 {
// 					// 	ruleResp.ItemFieldValue = child.Properties.Value[0]
// 					// }
// 				}
// 				// ruleResp.JsonTreeId = firstData.Model.SQL.WhereJsonTree.ID
// 			}
// 		}

// 		// Get evaluator info from last data entry (C)
// 		if len(rule.Data) > 2 {
// 			lastData := rule.Data[2]
// 			if len(lastData.Model.Conditions) > 0 {
// 				cond := lastData.Model.Conditions[0]
// 				if len(cond.Evaluator.Params) > 0 {
// 					ruleResp.ThresholdEvaluatorValue = cond.Evaluator.Params
// 				}
// 				ruleResp.ThresholdEvaluatorType = ThresholdEvaluatorType(cond.Evaluator.Type)
// 			}
// 		}

// 		response = append(response, ruleResp)
// 	}

// 	return response
// }
