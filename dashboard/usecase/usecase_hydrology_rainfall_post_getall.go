package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"shared/model"
	"shared/usecase"
	"sort"
	"strings"
	"time"
)

type ListRainfallPostReq struct {
	Date       string
	FilterDate string
	City       string
	Telemetry  float64
	Vendor     string
	Keyword    string
	Page       int
	PageSize   int
	SortBy     string
	SortOrder  string
}

type ListRainfallPostRes struct {
	Date      string            `json:"date"`
	RainFalls []RainfallPost    `json:"rainfall_posts"`
	Metadata  *usecase.Metadata `json:"metadata"`
}

type RainfallPost struct {
	ID           string `json:"id"`
	ExternalID   int64  `json:"external_id"`
	LatestUpdate string `json:"latest_update"`

	Name    string `json:"name"`
	Officer string `json:"officer"`
	// WhatsappNumber string   `json:"whatsapp_number"`
	City           string   `json:"city"`
	Manual         *float64 `json:"manual"`
	Telemetry      float64  `json:"telemetry"`
	AttendanceData float64  `json:"attendance_data"`
	Vendor         string   `json:"vendor"`
	// Morning        float64  `json:"morning"`
	// Afternoon      float64  `json:"afternoon"`
	// Evening        float64  `json:"evening"`
	// Dawn           float64  `json:"dawn"`
}

type ListRainfallPostUseCase = core.ActionHandler[ListRainfallPostReq, ListRainfallPostRes]

func ImplListRainfallPost(
	getListRainfallPostGateway gateway.GetRainFallPostListGateway,
	getListRainfallDailyGateway gateway.GetRainfallDailyByPostID,
	getLatestRainfallTelemetriesGateway gateway.GetLatestTelemetryRainPostByIDGateway,
) ListRainfallPostUseCase {
	return func(ctx context.Context, request ListRainfallPostReq) (*ListRainfallPostRes, error) {

		if request.Date == "" {
			request.FilterDate = time.Now().Format(time.DateOnly)
		} else {
			request.FilterDate = request.Date
		}
		rainfallPosts, err := getListRainfallPostGateway(ctx, gateway.GetRainfallPostListReq{
			City:   request.City,
			Vendor: request.Vendor,
			Name:   request.Keyword,
		})
		if err != nil {
			return nil, err
		}

		rainfallPostIDs := make([]int64, 0)

		for _, rainfallPost := range rainfallPosts.RainFalls {
			rainfallPostIDs = append(rainfallPostIDs, rainfallPost.ExternalID)
		}

		// get and map rainfall daily
		rainfallPostDailies, err := getListRainfallDailyGateway(ctx,
			gateway.GetRainfallDailyReq{
				PostIDs: rainfallPostIDs,
				Date:    request.FilterDate,
			},
		)
		if err != nil {
			return nil, err
		}

		rainfallDailyMap := make(map[int64]model.HydrologyRainDaily)
		for _, telemetry := range rainfallPostDailies.RainfallDaily {
			rainfallDailyMap[int64(telemetry.RainPostID)] = telemetry
		}

		// get and map latest rainfall (from hourly)
		rainfallLatestTelemetries, err := getLatestRainfallTelemetriesGateway(ctx,
			gateway.GetLatestTelemetryRainPostByIDReq{
				PostIDs: rainfallPostIDs,
				Date:    request.FilterDate,
			},
		)
		if err != nil {
			return nil, err
		}

		latestTelemetryMap := make(map[int64]model.HydrologyRainHourly)
		for _, latestTelemetry := range rainfallLatestTelemetries.LatestTelemetry {
			latestTelemetryMap[int64(latestTelemetry.RainPostID)] = latestTelemetry
		}

		rainfallPostMap := make(map[int64]model.RainfallPost)
		for _, rainfallPost := range rainfallPosts.RainFalls {
			rainfallPostMap[rainfallPost.ExternalID] = rainfallPost
		}

		rainfallPostsData := filterRainfallPostsMap(
			rainfallPostMap,
			rainfallDailyMap,
			latestTelemetryMap,
			request,
		)

		if request.SortBy != "" {
			sortRainfallPosts(rainfallPostsData, request)
		}

		paginatedData := paginateRainfallPosts(rainfallPostsData, request.Page, request.PageSize)

		totalItems := len(rainfallPostsData)
		totalPages := (totalItems + request.PageSize - 1) / request.PageSize

		return &ListRainfallPostRes{
			RainFalls: paginatedData,
			Date:      request.FilterDate,
			Metadata: &usecase.Metadata{
				Pagination: usecase.Pagination{
					Page:       request.Page,
					Limit:      request.PageSize,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}

func paginateRainfallPosts(features []RainfallPost, page, pageSize int) []RainfallPost {
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

func filterRainfallPostsMap(
	rainfallPosts map[int64]model.RainfallPost,
	rainfallDailyMap map[int64]model.HydrologyRainDaily,
	rainfallLatestTelemetries map[int64]model.HydrologyRainHourly,
	// rainfallHourlyMap map[int64]model.RainfallHourlyCalculation,
	// officerMap map[string]model.PostOfficer,
	request ListRainfallPostReq) []RainfallPost {
	filteredFeatures := make([]RainfallPost, 0)

	for _, rainfallPost := range rainfallPosts {
		rainfallPostTelemetry := rainfallDailyMap[rainfallPost.ExternalID]
		// RainfallHourly := rainfallHourlyMap[rainfallPost.ExternalID]
		// officer := officerMap[rainfallPost.Officer]

		if request.Telemetry != 0 && float64(rainfallPostTelemetry.Rain) != request.Telemetry {
			continue
		}

		if request.Date != "" && rainfallPostTelemetry.Sampling != request.FilterDate {
			continue
		}

		lastUpdate := "no data"
		if !rainfallPostTelemetry.UpdatedAt.IsZero() {
			lastUpdate = rainfallPostTelemetry.UpdatedAt.Format(time.DateTime)
		}

		telemetry := float64(rainfallPostTelemetry.Rain)
		latestTelemetry := rainfallLatestTelemetries[rainfallPost.ExternalID]
		if request.Date == "" {
			//use latest telemetry
			telemetry = float64(latestTelemetry.Rain)
			if !latestTelemetry.UpdatedAt.IsZero() {
				lastUpdate = latestTelemetry.UpdatedAt.Format(time.DateTime)
			}
		}

		filteredFeatures = append(filteredFeatures, RainfallPost{
			ID:             rainfallPost.ID,
			ExternalID:     rainfallPost.ExternalID,
			Name:           rainfallPost.Name,
			Officer:        rainfallPost.Officer,
			City:           rainfallPost.City,
			Manual:         rainfallPostTelemetry.Manual,
			Telemetry:      telemetry,
			AttendanceData: float64(rainfallPostTelemetry.Count * 100 / (60 / 5 * 24)),
			Vendor:         strings.TrimSpace(rainfallPost.Vendor),
			LatestUpdate:   lastUpdate,
		})

	}

	return filteredFeatures
}

func sortRainfallPosts(features []RainfallPost, request ListRainfallPostReq) {
	sort.Slice(features, func(i, j int) bool {
		switch request.SortBy {
		case "name":
			if request.SortOrder == "asc" {
				return features[i].Name < features[j].Name
			}
			return features[i].Name > features[j].Name
		case "officer":
			if request.SortOrder == "asc" {
				return features[i].Officer < features[j].Officer
			}
			return features[i].Officer > features[j].Officer
		// case "whatsapp_number":
		// 	if request.SortOrder == "asc" {
		// 		return features[i].WhatsappNumber < features[j].WhatsappNumber
		// 	}
		// 	return features[i].WhatsappNumber > features[j].WhatsappNumber
		case "city":
			if request.SortOrder == "asc" {
				return features[i].City < features[j].City
			}
			return features[i].City > features[j].City
		case "manual":
			if request.SortOrder == "asc" {
				return compareFloatPointers(features[i].Manual, features[j].Manual)
			}
			return compareFloatPointers(features[j].Manual, features[i].Manual)
		case "telemetry":
			if request.SortOrder == "asc" {
				return features[i].Telemetry < features[j].Telemetry
			}
			return features[i].Telemetry > features[j].Telemetry
		case "attendance_data":
			if request.SortOrder == "asc" {
				return features[i].AttendanceData < features[j].AttendanceData
			}
			return features[i].AttendanceData > features[j].AttendanceData
		case "vendor":
			if request.SortOrder == "asc" {
				return features[i].Vendor < features[j].Vendor
			}
			return features[i].Vendor > features[j].Vendor
		// case "morning":
		// 	if request.SortOrder == "asc" {
		// 		return compareFloatPointers(&features[i].Morning, &features[j].Morning)
		// 	}
		// 	return compareFloatPointers(&features[j].Morning, &features[i].Morning)
		// case "afternoon":
		// 	if request.SortOrder == "asc" {
		// 		return compareFloatPointers(&features[i].Afternoon, &features[j].Afternoon)
		// 	}
		// 	return compareFloatPointers(&features[j].Afternoon, &features[i].Afternoon)
		// case "evening":
		// 	if request.SortOrder == "asc" {
		// 		return compareFloatPointers(&features[i].Evening, &features[j].Evening)
		// 	}
		// 	return compareFloatPointers(&features[j].Evening, &features[i].Evening)
		// case "dawn":
		// 	if request.SortOrder == "asc" {
		// 		return compareFloatPointers(&features[i].Dawn, &features[j].Dawn)
		// 	}
		// 	return compareFloatPointers(&features[j].Dawn, &features[i].Dawn)
		case "latest_update":
			if request.SortOrder == "asc" {
				return features[i].LatestUpdate < features[j].LatestUpdate
			}
			return features[i].LatestUpdate > features[j].LatestUpdate
		}

		return false
	})
}

func compareFloatPointers(a, b *float64) bool {
	if a == nil && b == nil {
		return false
	}
	if a == nil {
		return true
	}
	if b == nil {
		return false
	}
	return *a < *b
}
