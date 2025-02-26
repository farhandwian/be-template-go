package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PencatatanKejadianRisikoDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type PencatatanKejadianRisikoDeleteUseCaseRes struct{}

type PencatatanKejadianRisikoDeleteUseCase = core.ActionHandler[PencatatanKejadianRisikoDeleteUseCaseReq, PencatatanKejadianRisikoDeleteUseCaseRes]

func ImplPencatatanKejadianRisikoDeleteUseCase(deletePencatatanKejadianRisiko gateway.PencatatanKejadianRisikoDelete) PencatatanKejadianRisikoDeleteUseCase {
	return func(ctx context.Context, req PencatatanKejadianRisikoDeleteUseCaseReq) (*PencatatanKejadianRisikoDeleteUseCaseRes, error) {

		if _, err := deletePencatatanKejadianRisiko(ctx, gateway.PencatatanKejadianRisikoDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &PencatatanKejadianRisikoDeleteUseCaseRes{}, nil
	}
}
