package usecase

import (
	"context"
	"perizinan/gateway"
	"perizinan/model"
	"shared/core"
)

type LaporGetAllUseCaseReq struct {
	Keyword               string
	PeriodePengambilanSda string
	Page                  int
	Size                  int
	SortBy                string
	SortOrder             string
}

type LaporGetAllUseCaseRes struct {
	Items    []ListLapor `json:"items"`
	Metadata *Metadata   `json:"metadata"`
}

type ListLapor struct {
	Perusahaan       string                 `json:"perusahaan"`
	LaporanPerizinan model.LaporanPerizinan `json:"laporan_perizinan"`
}

type LaporGetAllUseCase = core.ActionHandler[LaporGetAllUseCaseReq, LaporGetAllUseCaseRes]

func ImplLaporGetAllUseCase(
	getAllLaporan gateway.LaporanPerizinanGetAll,
	getAllSK gateway.SKPerizinanGetAll,
) LaporGetAllUseCase {
	return func(ctx context.Context, req LaporGetAllUseCaseReq) (*LaporGetAllUseCaseRes, error) {
		resAllLaporan, err := getAllLaporan(ctx, gateway.LaporanPerizinanGetAllReq{
			Page:                  req.Page,
			Size:                  req.Size,
			Keyword:               req.Keyword,
			PeriodePengambilanSda: req.PeriodePengambilanSda,
			SortBy:                req.SortBy,
			SortOrder:             req.SortOrder,
		})
		if err != nil {
			return nil, err
		}

		resAllSK, err := getAllSK(ctx, gateway.SKPerizinanGetAllReq{})
		if err != nil {
			return nil, err
		}

		AllSKMap := make(map[model.SKPerizinanID]model.SKPerizinan)
		for _, item := range resAllSK.Items {
			AllSKMap[item.ID] = item
		}

		listLapor := make([]ListLapor, 0)
		for _, item := range resAllLaporan.Items {
			SKMap, ok := AllSKMap[item.SKPerizinanID]
			if ok {
				Lapor := ListLapor{
					Perusahaan:       SKMap.PerusahaanPemohon,
					LaporanPerizinan: item,
				}
				listLapor = append(listLapor, Lapor)
			} else {
				Lapor := ListLapor{
					Perusahaan:       "",
					LaporanPerizinan: item,
				}
				listLapor = append(listLapor, Lapor)
			}
		}

		totalItems := int(resAllLaporan.Count)
		totalPages := (totalItems + req.Size - 1) / req.Size

		return &LaporGetAllUseCaseRes{
			Items: listLapor,
			Metadata: &Metadata{
				Pagination: Pagination{
					Page:       req.Page,
					Limit:      req.Size,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}
