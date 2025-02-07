package gateway

import (
	"context"
	"fmt"
	"iam/model"
	"shared/core"
	"shared/helper"
	"shared/middleware"

	"gorm.io/gorm"
)

type UserGetAllReq struct {
	Page        int               `json:"page"`
	Size        int               `json:"size"`
	UserID      model.UserID      `json:"user_id"`
	Email       model.Email       `json:"email"`
	PhoneNumber model.PhoneNumber `json:"phone_number"`
	NameLike    string            `json:"name_like"`
	SortOrder   string
	SortBy      string
}

type UserGetAllRes struct {
	Count int64        `json:"count"`
	Items []model.User `json:"items"`
}

type UserGetAll = core.ActionHandler[UserGetAllReq, UserGetAllRes]

func ImplUserGetAll() UserGetAll {
	return func(ctx context.Context, request UserGetAllReq) (*UserGetAllRes, error) {

		return &UserGetAllRes{
			Count: 0,
			Items: []model.User{},
		}, nil
	}
}

func ImplUserGetAllWithDatabase(db *gorm.DB) UserGetAll {
	return func(ctx context.Context, req UserGetAllReq) (*UserGetAllRes, error) {

		// Start the query chain with Model(&model.User{})
		query := middleware.
			GetDBFromContext(ctx, db).
			Model(&model.User{})

		var count int64

		if req.NameLike != "" {
			query = query.Where("name LIKE ?", "%"+req.NameLike+"%")
		}

		if req.Email != "" {
			query = query.Where("email = ?", req.Email)
		}

		if req.PhoneNumber != "" {
			query = query.Where("phone_number = ?", req.PhoneNumber)
		}

		if req.UserID != "" {
			query = query.Where("id = ?", req.UserID)
		}

		// Count total matching records
		if err := query.
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		// Validate sortBy
		allowedSortBy := map[string]bool{
			"name":         true,
			"phone_number": true,
			"email":        true,
		}

		// Validate and get sorting parameters
		sortBy, sortOrder, err := helper.ValidateSortParams(allowedSortBy, req.SortBy, req.SortOrder, "name")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := helper.ValidatePageSize(req.Page, req.Size)

		var users []model.User

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Order(orderClause).
			Find(&users).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &UserGetAllRes{
			Count: count,
			Items: users,
		}, nil
	}
}
