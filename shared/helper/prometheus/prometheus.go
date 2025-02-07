package prometheus

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"log"
	"os"
	"time"
)

func InitPrometheusClient() v1.API {
	client, err := api.NewClient(api.Config{
		Address: os.Getenv("PROMETHEUS_URL"),
	})
	if err != nil {
		log.Fatalf("Error creating Prometheus client: %v", err)
	}

	apiClient := v1.NewAPI(client)

	return apiClient
}

func RunQueryRange(apiClient v1.API, query string) int {
	end := time.Now()
	start := end.Add(-6 * time.Hour)

	rangeQuery := v1.Range{
		Start: start,
		End:   end,
		Step:  time.Minute,
	}

	value, _, err := apiClient.QueryRange(context.Background(), query, rangeQuery)
	if err != nil {
		log.Fatalf("Error running range query '%s': %v", query, err)
	}

	switch v := value.(type) {
	case model.Matrix:
		if len(v) > 0 {
			return int(v[0].Values[len(v[0].Values)-1].Value)
		}
	default:
		log.Printf("Unexpected value type for query '%s': %v", query, value)
	}
	return 0
}

func RunQuery(apiClient v1.API, query string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	value, _, err := apiClient.Query(ctx, query, time.Now())
	if err != nil {
		return "", fmt.Errorf("error running query '%s': %w", query, err)
	}

	return value.String(), nil
}
