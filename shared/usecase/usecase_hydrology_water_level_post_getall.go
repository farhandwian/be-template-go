package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"sort"
	"time"
)

type ListWaterLevelPostReq struct {
	Date       string
	FilterDate string
	River      string
	Keyword    string
	Page       int
	PageSize   int
	SortBy     string
	SortOrder  string
}

type ListWaterLevelPostResp struct {
	Date            string           `json:"date"`
	WaterLevelPosts []WaterLevelPost `json:"water_level_posts"`
	Metadata        *Metadata        `json:"metadata"`
}

type WaterLevelPost struct {
	ID           string    `json:"id"`
	ExternalID   int64     `json:"external_id"`
	Name         string    `json:"name"`
	River        string    `json:"river"`
	Officer      string    `json:"officer"`
	Manual       float64   `json:"manual"`
	Telemetry    float64   `json:"telemetry"`
	LatestUpdate time.Time `json:"latest_update"`
}

type ListWaterLevelPostUseCase = core.ActionHandler[ListWaterLevelPostReq, ListWaterLevelPostResp]

func ImplListWaterLevelPost(
	getListWaterLevelPostGateway gateway.GetWaterLevelPostListGateway,
	getListLatestManualWaterPost gateway.GetListLatestManualWaterPostGateway,
	getListLatestTelemetryWaterPost gateway.GetListLatestTelemetryWaterPostGateway,
) ListWaterLevelPostUseCase {
	return func(ctx context.Context, request ListWaterLevelPostReq) (*ListWaterLevelPostResp, error) {

		if request.Date == "" {
			request.FilterDate = time.Now().Format(time.DateOnly)
		} else {
			request.FilterDate = request.Date
		}

		// get water level post
		waterLevelPosts, err := getListWaterLevelPostGateway(ctx, gateway.GetWaterLevelPostListReq{
			River: request.River,
			Name:  request.Keyword,
		})
		if err != nil {
			return nil, err
		}

		// officers := make([]string, 0)
		// waterLevelPostIDs := make([]int64, 0)

		// for _, wLevelPost := range waterLevelPosts.WaterLevelPost {
		// 	officers = append(officers, wLevelPost.Officer)
		// 	waterLevelPostIDs = append(waterLevelPostIDs, wLevelPost.ExternalID)
		// }

		waterLevelLatestManual, err := getListLatestManualWaterPost(ctx, gateway.GetListLatestManualWaterPostReq{
			Date: request.FilterDate})
		if err != nil {
			return nil, err
		}

		waterLevelLatestManualMap := make(map[int64]float64)
		for _, wlManual := range waterLevelLatestManual.LatestManualWaterPost {
			waterLevelLatestManualMap[wlManual.WaterLevelPostID] = wlManual.TMA
		}

		waterLevelLatestTelemetry, err := getListLatestTelemetryWaterPost(ctx, gateway.GetListLatestTelemetryWaterPostReq{
			Date: request.FilterDate})
		if err != nil {
			return nil, err
		}

		waterLevelLatestTelemetryMap := make(map[int64]float64)
		waterLevelLatestTelemetrySamplingMap := make(map[int64]time.Time)
		for _, wlManual := range waterLevelLatestTelemetry.LatestManualWaterPost {
			waterLevelLatestTelemetryMap[wlManual.WaterLevelPostID] = wlManual.WaterLevel
			waterLevelLatestTelemetrySamplingMap[wlManual.WaterLevelPostID] = wlManual.CreatedAt
		}

		waterLevelPostsData := make([]WaterLevelPost, 0)
		for _, waterLevelPost := range waterLevelPosts.WaterLevelPost {

			if request.Date != "" && waterLevelLatestTelemetrySamplingMap[waterLevelPost.ExternalID].Format(time.DateOnly) != request.FilterDate {
				continue
			}
			waterLevelPostsData = append(waterLevelPostsData, WaterLevelPost{
				ID:           waterLevelPost.ID,
				ExternalID:   waterLevelPost.ExternalID,
				Name:         waterLevelPost.Name,
				River:        waterLevelPost.Location,
				Officer:      waterLevelPost.Officer,
				Manual:       waterLevelLatestManualMap[waterLevelPost.ExternalID],
				Telemetry:    waterLevelLatestTelemetryMap[waterLevelPost.ExternalID],
				LatestUpdate: waterLevelLatestTelemetrySamplingMap[waterLevelPost.ExternalID],
			})
		}

		if request.SortBy != "" {
			sortWaterLevelPost(waterLevelPostsData, request)
		}

		paginatedData := paginateWaterLevelposts(waterLevelPostsData, request.Page, request.PageSize)

		totalItems := len(waterLevelPostsData)
		totalPages := (totalItems + request.PageSize - 1) / request.PageSize

		return &ListWaterLevelPostResp{
			Date:            request.Date,
			WaterLevelPosts: paginatedData,
			Metadata: &Metadata{
				Pagination{
					Page:       request.Page,
					Limit:      request.PageSize,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}

func paginateWaterLevelposts(features []WaterLevelPost, page, pageSize int) []WaterLevelPost {
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(features) {
		start = len(features)
	}
	if end > len(features) {
		end = len(features)
	}

	return features[start:end]
}

func sortWaterLevelPost(waterLevelPostsData []WaterLevelPost, request ListWaterLevelPostReq) {
	sort.Slice(waterLevelPostsData, func(i, j int) bool {
		switch request.SortBy {
		case "name":
			if request.SortOrder == "asc" {
				return waterLevelPostsData[i].Name < waterLevelPostsData[j].Name
			}
			return waterLevelPostsData[i].Name > waterLevelPostsData[j].Name
		case "river":
			if request.SortOrder == "asc" {
				return waterLevelPostsData[i].River < waterLevelPostsData[j].River
			}
			return waterLevelPostsData[i].River > waterLevelPostsData[j].River
		case "officer":
			if request.SortOrder == "asc" {
				return waterLevelPostsData[i].Officer < waterLevelPostsData[j].Officer
			}
			return waterLevelPostsData[i].Officer > waterLevelPostsData[j].Officer
		case "manual":
			if request.SortOrder == "asc" {
				return waterLevelPostsData[i].Manual < waterLevelPostsData[j].Manual
			}
			return waterLevelPostsData[i].Manual > waterLevelPostsData[j].Manual
		case "telemetry":
			if request.SortOrder == "asc" {
				return waterLevelPostsData[i].Telemetry < waterLevelPostsData[j].Telemetry
			}
			return waterLevelPostsData[i].Telemetry > waterLevelPostsData[j].Telemetry
		case "latest_update":
			if request.SortOrder == "asc" {
				return waterLevelPostsData[i].LatestUpdate.Before(waterLevelPostsData[j].LatestUpdate)
			}
			return waterLevelPostsData[i].LatestUpdate.After(waterLevelPostsData[j].LatestUpdate)
		}
		return false
	})
}
