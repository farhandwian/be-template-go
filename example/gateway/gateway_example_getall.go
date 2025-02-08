package gateway

import (
	"context"
	"example/model"
	"fmt"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type ExampleGetAllGatewayReq struct {
	Keyword string
	Page    int
	Size    int
}

type ExampleGetAllGatewayRes struct {
	Employees []model.Example `json:"employees"`
	Count     int64           `json:"count"`
}

type ExampleGateway = core.ActionHandler[ExampleGetAllGatewayReq, ExampleGetAllGatewayRes]

func ImplExampleGateway(db *gorm.DB) ExampleGateway {
	return func(ctx context.Context, req ExampleGetAllGatewayReq) (*ExampleGetAllGatewayRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("test_string LIKE ?", keyword)
		}

		var count int64

		if err := query.
			Model(&model.Example{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := ValidatePageSize(req.Page, req.Size)

		var examples []model.Example

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&examples).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &ExampleGetAllGatewayRes{
			Employees: examples,
			Count:     count,
		}, nil
	}
}
