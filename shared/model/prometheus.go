package model

type PrometheusResponse struct {
	Status         string         `json:"status"`
	PrometheusData PrometheusData `json:"data"`
}

type PrometheusData struct {
	ResultType string             `json:"resultType"`
	Result     []PrometheusResult `json:"result"`
}
type PrometheusResult struct {
	Metric PrometheusMetric `json:"metric"`
	Value  []interface{}    `json:"value"`
}
type PrometheusMetric struct {
	Name string `json:"name"`
}
