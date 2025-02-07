package usecase

import (
	"context"
	"dashboard/gateway"
	"shared/core"
	"shared/usecase"
	"sort"
)

type ListWaterQualityReq struct {
	Date      string
	River     string
	Keyword   string
	Page      int
	PageSize  int
	SortBy    string
	SortOrder string
}

type ListWaterQualityResp struct {
	Date              string             `json:"date"`
	WaterQualityPosts []WaterQualityPost `json:"water_level_posts"`
	Metadata          *usecase.Metadata  `json:"metadata"`
}

type WaterQualityPost struct {
	ID         string `json:"id"`
	ExternalID int64  `json:"external_id"`
	Name       string `json:"name"`
	River      string `json:"river"`
}

type ListWaterQualityPostUseCase = core.ActionHandler[ListWaterQualityReq, ListWaterQualityResp]

func ImplListWaterQualityPost(
	getListWaterQualityPostGateway gateway.GetWaterQualityPostListGateway,
) ListWaterQualityPostUseCase {
	return func(ctx context.Context, request ListWaterQualityReq) (*ListWaterQualityResp, error) {

		// get water quality post
		waterQualityPosts, err := getListWaterQualityPostGateway(ctx, gateway.GetWaterQualityPostListReq{
			River: request.River,
			Name:  request.Keyword,
		})
		if err != nil {
			return nil, err
		}

		waterQualityPostsData := make([]WaterQualityPost, 0)
		for _, waterQualityPost := range waterQualityPosts.WaterQualityPost {
			waterQualityPostsData = append(waterQualityPostsData, WaterQualityPost{
				ID:         waterQualityPost.ID,
				ExternalID: waterQualityPost.ExternalID,
				Name:       waterQualityPost.Name,
				River:      waterQualityPost.Location,
			})
		}

		if request.SortBy != "" {
			sortWaterQualityPost(waterQualityPostsData, request)
		}

		paginatedData := paginateWaterQualityposts(waterQualityPostsData, request.Page, request.PageSize)

		totalItems := len(waterQualityPostsData)
		totalPages := (totalItems + request.PageSize - 1) / request.PageSize

		return &ListWaterQualityResp{
			Date:              request.Date,
			WaterQualityPosts: paginatedData,
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

func paginateWaterQualityposts(features []WaterQualityPost, page, pageSize int) []WaterQualityPost {
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

func sortWaterQualityPost(waterQualityPostsData []WaterQualityPost, request ListWaterQualityReq) {
	sort.Slice(waterQualityPostsData, func(i, j int) bool {
		switch request.SortBy {
		case "name":
			if request.SortOrder == "asc" {
				return waterQualityPostsData[i].Name < waterQualityPostsData[j].Name
			}
			return waterQualityPostsData[i].Name > waterQualityPostsData[j].Name
		case "river":
			if request.SortOrder == "asc" {
				return waterQualityPostsData[i].River < waterQualityPostsData[j].River
			}
			return waterQualityPostsData[i].River > waterQualityPostsData[j].River
		}
		return false
	})
}
