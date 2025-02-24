package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type PenetapanKonteksRisikoGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type PenetapanKonteksRisiko struct {
	PenetapanKonteksRisiko model.PenetapanKonteksRisikoStrategisPemda `json:"penetapan_konteks_risiko"`
	IKUs                   []model.IKU
}

type PenetapanKonteksRisikoGetAllUseCaseRes struct {
	PenetapanKonteksRisikos []PenetapanKonteksRisiko `json:"penetapan_konteks_risikos"`
	Metadata                *usecase.Metadata        `json:"metadata"`
}

type PenetapanKonteksRisikoGetAllUseCase = core.ActionHandler[PenetapanKonteksRisikoGetAllUseCaseReq, PenetapanKonteksRisikoGetAllUseCaseRes]

func ImplPenetapanKonteksRisikoGetAllUseCase(getAllPenetapanKonteksRisikos gateway.PenetapanKonteksRisikoStrategisPemdaGetAll, getAllIKUs gateway.IKUGetAll) PenetapanKonteksRisikoGetAllUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoGetAllUseCaseReq) (*PenetapanKonteksRisikoGetAllUseCaseRes, error) {

		penetapanKonteksRisikos, err := getAllPenetapanKonteksRisikos(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(penetapanKonteksRisikos.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		penetapanKonteksRisikoRes := make([]PenetapanKonteksRisiko, len(penetapanKonteksRisikos.PenetapanKonteksRisikoStrategisPemda))
		for i, penetapanKonteksRisiko := range penetapanKonteksRisikos.PenetapanKonteksRisikoStrategisPemda {
			// need improvement
			ikus, err := getAllIKUs(ctx, gateway.IKUGetAllReq{
				ExternalID: *penetapanKonteksRisiko.ID,
			})
			if err != nil {
				return nil, err
			}

			penetapanKonteksRisikoRes[i] = PenetapanKonteksRisiko{
				PenetapanKonteksRisiko: penetapanKonteksRisiko,
				IKUs:                   ikus.IKU,
			}

		}

		return &PenetapanKonteksRisikoGetAllUseCaseRes{
			PenetapanKonteksRisikos: penetapanKonteksRisikoRes,
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
