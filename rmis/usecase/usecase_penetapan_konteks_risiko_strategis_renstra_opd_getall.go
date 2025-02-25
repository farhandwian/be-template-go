package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type PenetapanKonteksRisikoRenstraOPDGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type PenetapanKonteksRisikoRenstraOPDGetAllUseCaseRes struct {
	PenetapanKonteksRisikoRenstraOPDs []model.PenetapanKonteksRisikoStrategisRenstraOPDGet `json:"penetapan_konteks_risiko_strategis_resntra_opds"`
	Metadata                          *usecase.Metadata                                    `json:"metadata"`
}

type PenetapanKonteksRisikoRenstraOPDGetAllUseCase = core.ActionHandler[PenetapanKonteksRisikoRenstraOPDGetAllUseCaseReq, PenetapanKonteksRisikoRenstraOPDGetAllUseCaseRes]

func ImplPenetapanKonteksRisikoRenstraOPDGetAllUseCase(getAllPenetapanKonteksRisikoRenstraOPDs gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetAll, getAllIKUs gateway.IKUGetAll, getOneOPD gateway.OPDGetByID) PenetapanKonteksRisikoRenstraOPDGetAllUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoRenstraOPDGetAllUseCaseReq) (*PenetapanKonteksRisikoRenstraOPDGetAllUseCaseRes, error) {

		penetapanKonteksRisikos, err := getAllPenetapanKonteksRisikoRenstraOPDs(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(penetapanKonteksRisikos.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		penetapanKonteksRisikoRes := make([]model.PenetapanKonteksRisikoStrategisRenstraOPDGet, len(penetapanKonteksRisikos.PenetapanKonteksRisikoStrategisRenstraOPD))
		for i, penetapanKonteksRisiko := range penetapanKonteksRisikos.PenetapanKonteksRisikoStrategisRenstraOPD {
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

			penetapanKonteksRisikoRes[i] = model.PenetapanKonteksRisikoStrategisRenstraOPDGet{
				PenetapanKonteksRisikoStrategisRenstraOPD: penetapanKonteksRisiko,
				IKUs: ikus.IKU,
				OPD:  opd.OPD,
			}

		}

		return &PenetapanKonteksRisikoRenstraOPDGetAllUseCaseRes{
			PenetapanKonteksRisikoRenstraOPDs: penetapanKonteksRisikoRes,
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
