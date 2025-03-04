package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseReq struct {
	ID                              string             `json:"id"`
	NamaPemda                       string             `json:"nama_pemda"`
	Periode                         string             `json:"periode"`
	SumberData                      string             `json:"sumber_data"`
	TahunPenilaian                  string             `json:"tahun_penilaian"`
	TujuanStrategis                 string             `json:"tujuan_strategis"`
	SasaranStrategis                string             `json:"sasaran_strategis"`
	IKUSasaran                      string             `json:"iku_sasaran"`
	PrioritasPembangunan            string             `json:"prioritas_pembangunan"`
	PenetapanKonteksRisikoStrategis string             `json:"penetapan_konteks_resiko_strategis"`
	NamaDinas                       string             `json:"nama_dinas"`
	PenetapanTujuan                 string             `json:"penetapan_tujuan"`
	PenetapanSasaran                string             `json:"penetapan_sasaran"`
	PenetapanIku                    string             `json:"penetapan_iku"`
	UrusanPemerintahan              string             `json:"urusan_pemerintahan"`
	Status                          sharedModel.Status `json:"status"`
}

type PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseRes struct{}

type PenetapanKonteksRisikoStrategisPemdaUpdateUseCase = core.ActionHandler[PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseReq, PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseRes]

func ImplPenetapanKonteksRisikoStrategisPemdaUpdateUseCase(
	getPenetapanKonteksRisikoStrategisPemdaById gateway.PenetapanKonteksRisikoStrategisPemdaGetByID,
	updatePenetapanKonteksRisikoStrategisPemda gateway.PenetapanKonteksRisikoStrategisPemdaSave,
) PenetapanKonteksRisikoStrategisPemdaUpdateUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseReq) (*PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseRes, error) {

		res, err := getPenetapanKonteksRisikoStrategisPemdaById(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		penetapanKonteksRisikoStrategisPemda := res.PenetapanKonteksRisikoStrategisPemda

		tahunPenilaian, err := extractYear(req.TahunPenilaian)
		if err != nil {
			return nil, fmt.Errorf("invalid TahunPenilaian format: %v", err)
		}

		penetapanKonteksRisikoStrategisPemda.TahunPenilaian = &tahunPenilaian
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
		penetapanKonteksRisikoStrategisPemda.UrusanPemerintahan = &req.UrusanPemerintahan
		penetapanKonteksRisikoStrategisPemda.IkuSasaran = &req.IKUSasaran

		if _, err := updatePenetapanKonteksRisikoStrategisPemda(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaSaveReq{PenetepanKonteksRisikoStrategisPemda: penetapanKonteksRisikoStrategisPemda}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseRes{}, nil
	}
}
