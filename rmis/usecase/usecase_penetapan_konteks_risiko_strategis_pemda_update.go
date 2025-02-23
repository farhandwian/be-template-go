package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseReq struct {
	ID                     string `json:"id"`
	NamaPemda              string `json:"nama_pemda"`
	Periode                string `json:"periode"`
	SumberData             string `json:"sumber_data"`
	TujuanStrategis        string `json:"tujuan_strategis"`
	PenetapanKonteksRisiko string `json:"penetapan_konteks_resiko"`
	NamaDinas              string `json:"nama_dinas"`
	Sasaran                string `json:"sasaran"`
	IKUSasaran             string `json:"iku_sasaran"`
	PrioritasPembangunan   string `json:"prioritas_pembangunan"`
	Penilaian              string `json:"penilaian"`
	NamaYBS                string `json:"nama_ybs"`
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

		res.PenetapanKonteksRisikoStrategisPemda.NamaPemda = &req.NamaPemda
		res.PenetapanKonteksRisikoStrategisPemda.Periode = &req.Periode
		res.PenetapanKonteksRisikoStrategisPemda.SumberData = &req.SumberData
		res.PenetapanKonteksRisikoStrategisPemda.TujuanStrategis = &req.TujuanStrategis
		res.PenetapanKonteksRisikoStrategisPemda.PenetapanKonteksRisiko = &req.PenetapanKonteksRisiko
		res.PenetapanKonteksRisikoStrategisPemda.NamaDinas = &req.NamaDinas
		res.PenetapanKonteksRisikoStrategisPemda.Sasaran = &req.Sasaran
		res.PenetapanKonteksRisikoStrategisPemda.IKUSasaran = &req.IKUSasaran
		res.PenetapanKonteksRisikoStrategisPemda.PrioritasPembangunan = &req.PrioritasPembangunan
		res.PenetapanKonteksRisikoStrategisPemda.Penilaian = &req.Penilaian
		res.PenetapanKonteksRisikoStrategisPemda.NamaYBS = &req.NamaYBS

		if _, err := updatePenetapanKonteksRisikoStrategisPemda(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaSaveReq{PenetepanKonteksRisikoStrategisPemda: res.PenetapanKonteksRisikoStrategisPemda}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseRes{}, nil
	}
}
