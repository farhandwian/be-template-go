package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PenetapanKonteksRisikoStrategisPemdaCreateUseCaseReq struct {
	NamaPemda              string `json:"nama_pemda"`
	Periode                string `json:"periode"`
	SumberData             string `json:"sumber_data"`
	TujuanStrategis        string `json:"tujuan_strategis"`
	NamaDinas              string `json:"nama_dinas"`
	Sasaran                string `json:"sasaran"`
	IKUSasaran             string `json:"iku_sasaran"`
	PrioritasPembangunan   string `json:"prioritas_pembangunan"`
	Penilaian              string `json:"penilaian"`
	PenetapanKonteksRisiko string `json:"penetapan_konteks_resiko"`
}

type PenetapanKonteksRisikoStrategisPemdaCreateUseCaseRes struct {
	ID string `json:"id"`
}

type PenetapanKonteksRisikoStrategisPemdaCreateUseCase = core.ActionHandler[PenetapanKonteksRisikoStrategisPemdaCreateUseCaseReq, PenetapanKonteksRisikoStrategisPemdaCreateUseCaseRes]

func ImplPenetapanKonteksRisikoStrategisPemdaCreateUseCase(
	generateId gateway.GenerateId,
	createPenetapanKonteksRisikoStrategisPemda gateway.PenetepanKonteksRisikoStrategisPemdaSave,
) PenetapanKonteksRisikoStrategisPemdaCreateUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisPemdaCreateUseCaseReq) (*PenetapanKonteksRisikoStrategisPemdaCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		obj := model.PenetapanKonteksRisikoStrategisPemda{
			ID:                     &genObj.RandomId,
			NamaPemda:              &req.NamaPemda,
			Periode:                &req.Periode,
			SumberData:             &req.SumberData,
			TujuanStrategis:        &req.TujuanStrategis,
			PenetapanKonteksRisiko: &req.PenetapanKonteksRisiko,
			NamaDinas:              &req.NamaDinas,
			Sasaran:                &req.Sasaran,
			PrioritasPembangunan:   &req.PrioritasPembangunan,
			Penilaian:              &req.Penilaian,
			NamaYBS:                &req.NamaYBS,
		}

		if _, err = createPenetapanKonteksRisikoStrategisPemda(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaSaveReq{PenetepanKonteksRisikoStrategisPemda: obj}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoStrategisPemdaCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
