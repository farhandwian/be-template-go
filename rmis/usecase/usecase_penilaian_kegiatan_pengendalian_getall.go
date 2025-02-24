package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type PenilaianKegiatanPengendalianGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type PenilaianKegiatanPengendalianGetAllUseCaseRes struct {
	PenilaianKegiatanPengendalian []model.PenilaianKegiatanPengendalian `json:"penilaian_kegiatan_pengendalians"`
	Metadata                      *usecase.Metadata                     `json:"metadata"`
}

type PenilaianKegiatanPengendalianGetAllUseCase = core.ActionHandler[PenilaianKegiatanPengendalianGetAllUseCaseReq, PenilaianKegiatanPengendalianGetAllUseCaseRes]

func ImplPenilaianKegiatanPengendalianGetAllUseCase(getAllPenilaianKegiatanPengendalians gateway.PenilaianKegiatanPengendalianGetAll) PenilaianKegiatanPengendalianGetAllUseCase {
	return func(ctx context.Context, req PenilaianKegiatanPengendalianGetAllUseCaseReq) (*PenilaianKegiatanPengendalianGetAllUseCaseRes, error) {

		res, err := getAllPenilaianKegiatanPengendalians(ctx, gateway.PenilaianKegiatanPengendalianGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &PenilaianKegiatanPengendalianGetAllUseCaseRes{
			PenilaianKegiatanPengendalian: res.PenilaianKegiatanPengendalian,
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
