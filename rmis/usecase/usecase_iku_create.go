package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type IKUCreateUseCaseReq struct {
	Nama       string `json:"nama"`
	Periode    string `json:"periode"`
	Target     string `json:"target"`
	ExternalID string `json:"external_id"` // nilai nya berupa id antara 3 tabel yang sesuai dengan typenya
	Type       string `json:"type"`        // PenetapanKonteksRisikoOperasionalInspektoratDaerah | PenetapanKonteksRisikoStrategisInspektoratDaerah |PenetapanKonteksRisikoStrategisPemda
}

type IKUCreateUseCaseRes struct {
	ID string `json:"id"`
}

type IKUCreateUseCase = core.ActionHandler[IKUCreateUseCaseReq, IKUCreateUseCaseRes]

func ImplIKUCreateUseCase(
	generateId gateway.GenerateId,
	createIKU gateway.IKUSave,
) IKUCreateUseCase {
	return func(ctx context.Context, req IKUCreateUseCaseReq) (*IKUCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		obj := model.IKU{
			ID:         &genObj.RandomId,
			Nama:       &req.Nama,
			Periode:    &req.Periode,
			Target:     &req.Target,
			ExternalID: &req.ExternalID,
			Type:       &req.Type,
		}

		if _, err = createIKU(ctx, gateway.IKUSaveReq{IKU: obj}); err != nil {
			return nil, err
		}

		return &IKUCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
