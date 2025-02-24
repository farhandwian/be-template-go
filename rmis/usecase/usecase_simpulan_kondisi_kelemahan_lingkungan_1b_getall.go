package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type SimpulanKondisiKelemahanLingkunganGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type SimpulanKondisiKelemahanLingkunganGetAllUseCaseRes struct {
	SimpulanKondisiKelemahanLingkungan []model.SimpulanKondisiKelemahanLingkungan `json:"simpulan_kondisi_kelemahan_lingkungans"`
	Metadata                           *usecase.Metadata                          `json:"metadata"`
}

type SimpulanKondisiKelemahanLingkunganGetAllUseCase = core.ActionHandler[SimpulanKondisiKelemahanLingkunganGetAllUseCaseReq, SimpulanKondisiKelemahanLingkunganGetAllUseCaseRes]

func ImplSimpulanKondisiKelemahanLingkunganGetAllUseCase(getAllSimpulanKondisiKelemahanLingkungans gateway.SimpulanKondisiKelemahanLingkunganGetAll) SimpulanKondisiKelemahanLingkunganGetAllUseCase {
	return func(ctx context.Context, req SimpulanKondisiKelemahanLingkunganGetAllUseCaseReq) (*SimpulanKondisiKelemahanLingkunganGetAllUseCaseRes, error) {

		res, err := getAllSimpulanKondisiKelemahanLingkungans(ctx, gateway.SimpulanKondisiKelemahanLingkunganGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &SimpulanKondisiKelemahanLingkunganGetAllUseCaseRes{
			SimpulanKondisiKelemahanLingkungan: res.SimpulanKondisiKelemahanLingkungan,
			Metadata: &usecase.Metadata{
				Pagination: usecase.Pagination{
					Page:       req.Page,
					Limit:      req.Size,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}
