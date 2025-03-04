package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	sharedModel "shared/model"
)

type PenetapanKonteksRisikoStrategisPemdaCreateUseCaseReq struct {
	NamaPemda                       string `json:"nama_pemda"`
	Periode                         string `json:"periode"`
	SumberData                      string `json:"sumber_data"`
	TujuanStrategis                 string `json:"tujuan_strategis"`
	SasaranStrategis                string `json:"sasaran_strategis"`
	IKUSasaran                      string `json:"iku_sasaran"`
	PrioritasPembangunan            string `json:"prioritas_pembangunan"`
	PenetapanKonteksRisikoStrategis string `json:"penetapan_konteks_resiko_strategis"`
	NamaDinas                       string `json:"nama_dinas"`
	PenetapanTujuan                 string `json:"penetapan_tujuan"`
	PenetapanSasaran                string `json:"penetapan_sasaran"`
	PenetapanIku                    string `json:"penetapan_iku"`
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
			ID:                              &genObj.RandomId,
			NamaPemda:                       &req.NamaPemda,
			Periode:                         &req.Periode,
			SumberData:                      &req.SumberData,
			TujuanStrategis:                 &req.TujuanStrategis,
			NamaDinas:                       &req.NamaDinas,
			SasaranStrategis:                &req.SasaranStrategis,
			PrioritasPembangunan:            &req.PrioritasPembangunan,
			PenetapanTujuan:                 &req.PenetapanTujuan,
			PenetapanSasaran:                &req.PenetapanSasaran,
			PenetapanIku:                    &req.PenetapanIku,
			IkuSasaran:                      &req.IKUSasaran,
			PenetapanKonteksRisikoStrategis: &req.PenetapanKonteksRisikoStrategis,
			Status:                          sharedModel.StatusMenungguVerifikasi,
		}

		if _, err = createPenetapanKonteksRisikoStrategisPemda(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaSaveReq{PenetepanKonteksRisikoStrategisPemda: obj}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoStrategisPemdaCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
