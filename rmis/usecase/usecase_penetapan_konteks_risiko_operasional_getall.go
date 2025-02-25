package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type PenetapanKonteksRisikoOperasionalGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type PenetapanKonteksRisikoOperasional struct {
	PenetapanKonteksRisikoOperasional model.PenetapanKonteksRisikoOperasional `json:"penetapan_konteks_risiko_operasional"`
	IKUs                              []model.IKU                             `json:"ikus"`
	OPD                               model.OPD                               `json:"opd"`
}

type PenetapanKonteksRisikoOperasionalGetAllUseCaseRes struct {
	PenetapanKonteksRisikoOperasionals []PenetapanKonteksRisikoOperasional `json:"penetapan_konteks_risiko_operasionals"`
	Metadata                           *usecase.Metadata                   `json:"metadata"`
}

type PenetapanKonteksRisikoOperasionalGetAllUseCase = core.ActionHandler[PenetapanKonteksRisikoOperasionalGetAllUseCaseReq, PenetapanKonteksRisikoOperasionalGetAllUseCaseRes]

func ImplPenetapanKonteksRisikoOperasionalGetAllUseCase(getAllPenetapanKonteksRisikoOperasionals gateway.PenetapanKonteksRisikoOperasionalGetAll, getAllIKUs gateway.IKUGetAll, getOneOPD gateway.OPDGetByID) PenetapanKonteksRisikoOperasionalGetAllUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoOperasionalGetAllUseCaseReq) (*PenetapanKonteksRisikoOperasionalGetAllUseCaseRes, error) {

		penetapanKonteksRisikos, err := getAllPenetapanKonteksRisikoOperasionals(ctx, gateway.PenetapanKonteksRisikoOperasionalGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(penetapanKonteksRisikos.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		penetapanKonteksRisikoRes := make([]PenetapanKonteksRisikoOperasional, len(penetapanKonteksRisikos.PenetapanKonteksRisikoOperasional))
		for i, penetapanKonteksRisiko := range penetapanKonteksRisikos.PenetapanKonteksRisikoOperasional {
			// need improvement
			ikus, err := getAllIKUs(ctx, gateway.IKUGetAllReq{
				ExternalID: *penetapanKonteksRisiko.ID,
			})
			if err != nil {
				return nil, err
			}

			opd, err := getOneOPD(ctx, gateway.OPDGetByIDReq{ID: *penetapanKonteksRisiko.OPDID})
			if err != nil {
				return nil, err
			}

			penetapanKonteksRisikoRes[i] = PenetapanKonteksRisikoOperasional{
				PenetapanKonteksRisikoOperasional: penetapanKonteksRisiko,
				IKUs:                              ikus.IKU,
				OPD:                               opd.OPD,
			}

		}

		return &PenetapanKonteksRisikoOperasionalGetAllUseCaseRes{
			PenetapanKonteksRisikoOperasionals: penetapanKonteksRisikoRes,
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
