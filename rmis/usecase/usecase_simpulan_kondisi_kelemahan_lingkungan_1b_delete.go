package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type SimpulanKondisiKelemahanLingkunganDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type SimpulanKondisiKelemahanLingkunganDeleteUseCaseRes struct{}

type SimpulanKondisiKelemahanLingkunganDeleteUseCase = core.ActionHandler[SimpulanKondisiKelemahanLingkunganDeleteUseCaseReq, SimpulanKondisiKelemahanLingkunganDeleteUseCaseRes]

func ImplSimpulanKondisiKelemahanLingkunganDeleteUseCase(deleteSimpulanKondisiKelemahanLingkungan gateway.SimpulanKondisiKelemahanLingkunganDelete) SimpulanKondisiKelemahanLingkunganDeleteUseCase {
	return func(ctx context.Context, req SimpulanKondisiKelemahanLingkunganDeleteUseCaseReq) (*SimpulanKondisiKelemahanLingkunganDeleteUseCaseRes, error) {

		if _, err := deleteSimpulanKondisiKelemahanLingkungan(ctx, gateway.SimpulanKondisiKelemahanLingkunganDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &SimpulanKondisiKelemahanLingkunganDeleteUseCaseRes{}, nil
	}
}
