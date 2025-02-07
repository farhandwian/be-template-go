package controller

import (
	"dashboard/usecase"
	"fmt"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"net/url"
	"regexp"
	"shared/helper"
	sm "shared/model"
	"strings"
	"time"
)

func (c Controller) AlarmConfigWebhookHandler(u usecase.AlarmConfigWebhook) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/alert",
		Access:  iammodel.WEBHOOK_OPERATION,
		Summary: "Webhook called by Grafana",
		Tag:     "Webhook",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[AlertNotification](w, r)
		if !ok {
			return
		}

		var alerts []sm.Alert
		for _, a := range request.Alerts {

			uid, err := extractIdentifierFromURL(a.GeneratorURL)
			if err != nil {
				controller.Fail(w, err)
				return
			}

			alerts = append(alerts, sm.Alert{
				UID:    uid,
				Status: a.Status,
				Values: a.Values.B,
			})
		}

		req := usecase.AlarmConfigWebhookReq{
			Alerts: alerts,
			Now:    time.Now(),
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}

type AlertValues struct {
	B float64 `json:"B,omitempty"`
	C float64 `json:"C,omitempty"`
}

type Alert struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
	SilenceURL   string            `json:"silenceURL"`
	DashboardURL string            `json:"dashboardURL"`
	PanelURL     string            `json:"panelURL"`
	Values       AlertValues       `json:"values"`
	ValueString  string            `json:"valueString"`
}

type AlertNotification struct {
	Receiver          string            `json:"receiver"`
	Status            string            `json:"status"`
	Alerts            []Alert           `json:"alerts"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
	OrgId             int               `json:"orgId"`
	Title             string            `json:"title"`
	State             string            `json:"state"`
	Message           string            `json:"message"`
}

func extractIdentifierFromURL(urlStr string) (string, error) {
	// Jika URL kosong, return error
	if urlStr == "" {
		return "", fmt.Errorf("empty URL provided")
	}

	// Parse URL ke struktur URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %v", err)
	}

	// Split path menjadi segments
	// URL: "http://localhost:3000/alerting/grafana/hello123/view"
	// akan menjadi: ["", "alerting", "grafana", "hello123", "view"]
	segments := strings.Split(parsedURL.Path, "/")

	// Cari identifier yang berada setelah "grafana" dalam path
	for i, segment := range segments {
		if segment == "grafana" && i+1 < len(segments) {
			// Tambahan validasi untuk identifier
			identifier := segments[i+1]

			// Pastikan identifier tidak kosong
			if identifier == "" {
				return "", fmt.Errorf("empty identifier found in URL")
			}

			// Pastikan identifier bukan "view" atau string lain yang tidak diinginkan
			if identifier == "view" {
				return "", fmt.Errorf("invalid identifier found in URL")
			}

			// Validasi format identifier (misalnya hanya alphanumeric)
			match, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", identifier)
			if !match {
				return "", fmt.Errorf("identifier contains invalid characters")
			}

			return identifier, nil
		}
	}

	return "", fmt.Errorf("identifier not found in URL")
}
