package gateway

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"shared/core"
	"shared/model"
	"strconv"
	"strings"
)

type GetServiceStatusReq struct{}

type GetServiceStatusRes struct {
	Service model.ServiceStatus `json:"service"`
}

type GetServiceStatusGateway = core.ActionHandler[GetServiceStatusReq, GetServiceStatusRes]

func ImplGetServiceStatusGateway() GetServiceStatusGateway {
	return func(ctx context.Context, request GetServiceStatusReq) (*GetServiceStatusRes, error) {
		url := os.Getenv("PROMETHEUS_URL")

		serviceStatusUrl := fmt.Sprintf("%s/api/v1/query?query=%s", url, "probe_success{job=\"blackbox-healthcheck\"}")
		serviceStatusRes, err := fetchPrometheusAPI(serviceStatusUrl)
		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error fetching up devices: %v", err))
		}

		serviceStatus, err := mapToServiceStatus(serviceStatusRes)
		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error mapping service status: %v", err))
		}

		return &GetServiceStatusRes{
			Service: *serviceStatus,
		}, nil
	}
}

func mapToServiceStatus(prometheusResponse *model.PrometheusResponse) (*model.ServiceStatus, error) {

	status := &model.ServiceStatus{}
	for _, result := range prometheusResponse.PrometheusData.Result {
		// Parse the success value
		success := result.Value[1].(string) == "1"

		// Map to the appropriate field in ServiceStatus
		switch strings.ToLower(result.Metric.Name) {
		case "dashboard":
			status.Dashboard = success
		case "adam - hawa":
			status.AdamHawa = success
		case "adapter manganti":
			status.AdapterManganti = success
		case "adapter hidrologi":
			status.AdapterHydrology = success
		case "sihka":
			status.Sihka = success
		case "jagacai":
			status.SiJagaCai = success
		}
	}

	return status, nil
}

//
//func ImplGetServiceStatusGateway() GetServiceStatusGateway {
//	return func(ctx context.Context, request GetServiceStatusReq) (*GetServiceStatusRes, error) {
//		prometheusClient := prometheus.InitPrometheusClient()
//
//		query := `probe_success{job="blackbox-healthcheck"}`
//		result, err := prometheus.RunQuery(prometheusClient, query)
//		if err != nil {
//			log.Fatalf("Error querying PrometheusResponse: %v", err)
//		}
//
//		status, err := parseServiceStatus(result)
//		if err != nil {
//			log.Fatalf("Error parsing result: %v", err)
//		}
//
//		return &GetServiceStatusRes{
//			Service: status,
//		}, nil
//
//	}
//}

func parseServiceStatus(input string) (model.ServiceStatus, error) {
	regex := regexp.MustCompile(`probe_success\{[^}]*name="([^"]+)"[^}]*\}\s*=>\s*(\d+)(?:\s*@\[[^\]]+\])?`)

	// More robust input cleaning
	var cleanedLines []string
	for _, line := range strings.Split(input, "\n") {
		trimmedLine := strings.TrimSpace(line)
		if strings.Contains(trimmedLine, "probe_success") &&
			!strings.HasPrefix(trimmedLine, "#") &&
			trimmedLine != "" {
			cleanedLines = append(cleanedLines, trimmedLine)
		}
	}
	cleanedInput := strings.Join(cleanedLines, "\n")

	matches := regex.FindAllStringSubmatch(cleanedInput, -1)

	if len(matches) == 0 {
		return model.ServiceStatus{}, fmt.Errorf("no matches found in input")
	}

	status := model.ServiceStatus{}

	for _, match := range matches {
		if len(match) >= 3 {
			name := match[1]
			value := mustAtoi(match[2]) == 1

			switch name {
			case "Dashboard":
				status.Dashboard = value
			case "Adam - Hawa":
				status.AdamHawa = value
			case "Adapter Manganti":
				status.AdapterManganti = value
			case "Adapter Hidrologi":
				status.AdapterHydrology = value
			case "SIHKA":
				status.Sihka = value
			case "JagaCai":
				status.SiJagaCai = value
			default:
				fmt.Printf("Unhandled service name: %s\n", name)
			}
		}
	}

	return status, nil
}

func mustAtoi(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		panic(err) // In real application, handle the error properly
	}
	return value
}
