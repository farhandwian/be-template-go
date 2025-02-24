package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type SimpulanKondisiKelemahanLingkunganGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type SimpulanKondisiKelemahanLingkunganGetByIDUseCaseRes struct {
	SimpulanKondisiKelemahanLingkungan model.SimpulanKondisiKelemahanLingkungan `json:"rekapitulasi_hasil_kuesioner"`
}

type SimpulanKondisiKelemahanLingkunganGetByIDUseCase = core.ActionHandler[SimpulanKondisiKelemahanLingkunganGetByIDUseCaseReq, SimpulanKondisiKelemahanLingkunganGetByIDUseCaseRes]

func ImplSimpulanKondisiKelemahanLingkunganGetByIDUseCase(getSimpulanKondisiKelemahanLingkunganByID gateway.SimpulanKondisiKelemahanLingkunganGetByID) SimpulanKondisiKelemahanLingkunganGetByIDUseCase {
	return func(ctx context.Context, req SimpulanKondisiKelemahanLingkunganGetByIDUseCaseReq) (*SimpulanKondisiKelemahanLingkunganGetByIDUseCaseRes, error) {
		res, err := getSimpulanKondisiKelemahanLingkunganByID(ctx, gateway.SimpulanKondisiKelemahanLingkunganGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &SimpulanKondisiKelemahanLingkunganGetByIDUseCaseRes{SimpulanKondisiKelemahanLingkungan: res.SimpulanKondisiKelemahanLingkungan}, nil
	}
}
