package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"shared/usecase"
	"sort"
)

type GetListClimatologyReq struct {
	Keyword   string
	Page      int
	PageSize  int
	SortBy    string
	SortOrder string
}

type ListClimatologyResp struct {
	Total       int               `json:"total"`
	Climatology []ClimatologyPost `json:"climatology"`
	Metadata    *usecase.Metadata `json:"metadata"`
}

type ClimatologyPost struct {
	ID                 string  `json:"id"`
	ExternalID         int64   `json:"external_id"`
	Name               string  `json:"name"`
	Rainfall           float64 `json:"rainfall"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	MaximumTemperature float64 `json:"maximum_temperature"`
	Humidity           float64 `json:"humidity"`
	WindSpeed          float64 `json:"wind_speed"`
	WindDirection      float64 `json:"wind_direction"`
	SunshineDuration   float64 `json:"sunshine_duration"`
}

type GetListClimatologyUseCase = core.ActionHandler[GetListClimatologyReq, ListClimatologyResp]

func ImplGetListClimatologyUseCase(getListClimatology gateway.GetClimatologyListGateway) GetListClimatologyUseCase {
	return func(ctx context.Context, req GetListClimatologyReq) (*ListClimatologyResp, error) {
		climatologyData, err := getListClimatology(ctx, gateway.GetClimatologyListReq{
			Keyword: req.Keyword,
		})
		if err != nil {
			return nil, err
		}

		climatologies := make([]ClimatologyPost, 0, len(climatologyData.Climatology))
		for _, climatology := range climatologyData.Climatology {
			climatologies = append(climatologies, ClimatologyPost{
				ID:                 climatology.ID,
				ExternalID:         climatology.ExternalID,
				Name:               climatology.Name,
				Rainfall:           0,
				MinimumTemperature: 0,
				MaximumTemperature: 0,
				Humidity:           0,
				WindSpeed:          0,
				WindDirection:      0,
			})

		}

		if req.SortBy != "" {
			sortClimatologyPost(climatologies, req)
		}

		paginatedData := paginateClimatologyPosts(climatologies, req.Page, req.PageSize)

		totalItems := len(climatologies)
		totalPages := (totalItems + req.PageSize - 1) / req.PageSize

		return &ListClimatologyResp{
			Total:       len(climatologies),
			Climatology: paginatedData,
			Metadata: &usecase.Metadata{
				Pagination: usecase.Pagination{
					Page:       req.Page,
					Limit:      req.PageSize,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}

func paginateClimatologyPosts(features []ClimatologyPost, page, pageSize int) []ClimatologyPost {
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

func sortClimatologyPost(features []ClimatologyPost, request GetListClimatologyReq) {
	sort.Slice(features, func(i, j int) bool {
		switch request.SortBy {
		case "name":
			if request.SortOrder == "asc" {
				return features[i].Name < features[j].Name
			}
			return features[i].Name > features[j].Name
		case "rainfall":
			if request.SortOrder == "asc" {
				return features[i].Rainfall < features[j].Rainfall
			}
			return features[i].Rainfall > features[j].Rainfall
		case "minimum_temperature":
			if request.SortOrder == "asc" {
				return features[i].MinimumTemperature < features[j].MinimumTemperature
			}
			return features[i].MinimumTemperature > features[j].MinimumTemperature
		case "maximum_temperature":
			if request.SortOrder == "asc" {
				return features[i].MaximumTemperature < features[j].MaximumTemperature
			}
			return features[i].MaximumTemperature > features[j].MaximumTemperature
		case "humidity":
			if request.SortOrder == "asc" {
				return features[i].Humidity < features[j].Humidity
			}
			return features[i].Humidity > features[j].Humidity
		case "wind_speed":
			if request.SortOrder == "asc" {
				return features[i].WindSpeed < features[j].WindSpeed
			}
			return features[i].WindSpeed > features[j].WindSpeed
		case "wind_direction":
			if request.SortOrder == "asc" {
				return features[i].WindDirection < features[j].WindDirection
			}
			return features[i].WindDirection > features[j].WindDirection
		case "sunshine_duration":
			if request.SortOrder == "asc" {
				return features[i].SunshineDuration < features[j].SunshineDuration
			}
			return features[i].SunshineDuration > features[j].SunshineDuration
		}
		return false
	})
}
