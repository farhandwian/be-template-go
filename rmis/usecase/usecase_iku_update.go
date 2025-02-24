package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type IKUUpdateUseCaseReq struct {
	ID         string `json:"id"`
	Nama       string `json:"nama"`
	Periode    string `json:"periode"`
	Target     string `json:"target"`
	ExternalID string `json:"external_id"` // nilai nya berupa id antara 3 tabel yang sesuai dengan typenya
	Type       string `json:"type"`        // PenetapanKonteksRisikoOperasionalInspektoratDaerah | PenetapanKonteksRisikoStrategisInspektoratDaerah |PenetapanKonteksRisikoStrategisPemda
}

type IKUUpdateUseCaseRes struct{}

type IKUUpdateUseCase = core.ActionHandler[IKUUpdateUseCaseReq, IKUUpdateUseCaseRes]

func ImplIKUUpdateUseCase(
	getIKUById gateway.IKUGetByID,
	updateIKU gateway.IKUSave,
) IKUUpdateUseCase {
	return func(ctx context.Context, req IKUUpdateUseCaseReq) (*IKUUpdateUseCaseRes, error) {

		res, err := getIKUById(ctx, gateway.IKUGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.IKU.Nama = &req.Nama
		res.IKU.Periode = &req.Periode
		res.IKU.Target = &req.Target
		res.IKU.ExternalID = &req.ExternalID
		res.IKU.Type = &req.Type

		if _, err := updateIKU(ctx, gateway.IKUSaveReq{IKU: res.IKU}); err != nil {
			return nil, err
		}

		return &IKUUpdateUseCaseRes{}, nil
	}
}
