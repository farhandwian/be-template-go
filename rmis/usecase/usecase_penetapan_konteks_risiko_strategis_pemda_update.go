package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseReq struct {
	ID                              string             `json:"id"`
	NamaPemda                       string             `json:"nama_pemda"`
	Periode                         string             `json:"periode"`
	SumberData                      string             `json:"sumber_data"`
	TujuanStrategis                 string             `json:"tujuan_strategis"`
	SasaranStrategis                string             `json:"sasaran_strategis"`
	IKUSasaran                      string             `json:"iku_sasaran"`
	PrioritasPembangunan            string             `json:"prioritas_pembangunan"`
	PenetapanKonteksRisikoStrategis string             `json:"penetapan_konteks_resiko_strategis"`
	NamaDinas                       string             `json:"nama_dinas"`
	PenetapanTujuan                 string             `json:"penetapan_tujuan"`
	PenetapanSasaran                string             `json:"penetapan_sasaran"`
	PenetapanIku                    string             `json:"penetapan_iku"`
	Status                          sharedModel.Status `json:"status"`
}

type PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseRes struct{}

type PenetapanKonteksRisikoStrategisPemdaUpdateUseCase = core.ActionHandler[PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseReq, PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseRes]

func ImplPenetapanKonteksRisikoStrategisPemdaUpdateUseCase(
	getPenetapanKonteksRisikoStrategisPemdaById gateway.PenetapanKonteksRisikoStrategisPemdaGetByID,
	updatePenetapanKonteksRisikoStrategisPemda gateway.PenetepanKonteksRisikoStrategisPemdaSave,
) PenetapanKonteksRisikoStrategisPemdaUpdateUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseReq) (*PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseRes, error) {

		res, err := getPenetapanKonteksRisikoStrategisPemdaById(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		penetapanKonteksRisikoStrategisPemda := res.PenetapanKonteksRisikoStrategisPemda

		penetapanKonteksRisikoStrategisPemda.NamaPemda = &req.NamaPemda
		penetapanKonteksRisikoStrategisPemda.Periode = &req.Periode
		penetapanKonteksRisikoStrategisPemda.SumberData = &req.SumberData
		penetapanKonteksRisikoStrategisPemda.TujuanStrategis = &req.TujuanStrategis
		penetapanKonteksRisikoStrategisPemda.PenetapanKonteksRisikoStrategis = &req.PenetapanKonteksRisikoStrategis
		penetapanKonteksRisikoStrategisPemda.NamaDinas = &req.NamaDinas
		penetapanKonteksRisikoStrategisPemda.SasaranStrategis = &req.SasaranStrategis
		penetapanKonteksRisikoStrategisPemda.PrioritasPembangunan = &req.PrioritasPembangunan
		penetapanKonteksRisikoStrategisPemda.PenetapanTujuan = &req.PenetapanTujuan
		penetapanKonteksRisikoStrategisPemda.PenetapanSasaran = &req.PenetapanSasaran
		penetapanKonteksRisikoStrategisPemda.PenetapanIku = &req.PenetapanIku
		penetapanKonteksRisikoStrategisPemda.Status = sharedModel.StatusMenungguVerifikasi
		penetapanKonteksRisikoStrategisPemda.IkuSasaran = &req.IKUSasaran

		if _, err := updatePenetapanKonteksRisikoStrategisPemda(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaSaveReq{PenetepanKonteksRisikoStrategisPemda: penetapanKonteksRisikoStrategisPemda}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseRes{}, nil
	}
}
