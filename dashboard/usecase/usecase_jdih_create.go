package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"fmt"
	"shared/core"
	"time"
)

type CreateJDIHReq struct {
	Title       string           `json:"title"`
	PublishedAt time.Time        `json:"published_at"`
	Status      model.JDIHStatus `json:"status"`
	Now         time.Time
}

type CreateJDIHResp struct {
	ID string `json:"id"`
}

type CreateJDIHUseCase = core.ActionHandler[CreateJDIHReq, CreateJDIHResp]

func ImplCreateJDIHUseCase(
	generateId gateway.GenerateId,
	createJDIH gateway.JDIHSaveGateway,
) CreateJDIHUseCase {
	return func(ctx context.Context, req CreateJDIHReq) (*CreateJDIHResp, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		if !req.PublishedAt.Before(req.Now) {
			return nil, fmt.Errorf("tanggal publish tidak boleh di masa mendatang")
		}

		obj := model.JDIH{
			ID:          genObj.RandomId,
			Title:       req.Title,
			PublishedAt: req.PublishedAt,
			Status:      req.Status,
		}

		if _, err = createJDIH(ctx, gateway.JDIHSaveReq{JDIH: obj}); err != nil {
			return nil, err
		}

		return &CreateJDIHResp{
			ID: genObj.RandomId,
		}, nil
	}
}
