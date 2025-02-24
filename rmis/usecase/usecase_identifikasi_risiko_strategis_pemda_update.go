package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"shared/core"
)

type IdentifikasiRisikoStrategisPemdaUpdateUseCaseReq struct {
	ID                 string `json:"id"`
	TahunPenilaian     string `json:"tahun_penilaian"`
	Periode            string `json:"periode"`
	UrusanPemerintahan string `json:"urusan_pemerintahan"`
	TujuanStrategis    string `json:"tujuan_strategis"`
	IndikatorKinerja   string `json:"indikator_kinerja"`
	UraianRisiko       string `json:"uraian_resiko"`
	PemilikRisiko      string `json:"pemilik_resiko"`
	Controllable       string `json:"controllable"`
	UraianDampak       string `json:"uraian_dampak"`
	PihakDampak        string `json:"pihak_dampak"`
	KategoriRisikoID   string `json:"kategori_resiko_id"`
	RcaID              string `json:"rca_id"`
}

type IdentifikasiRisikoStrategisPemdaUpdateUseCaseRes struct{}

type IdentifikasiRisikoStrategisPemdaUpdateUseCase = core.ActionHandler[IdentifikasiRisikoStrategisPemdaUpdateUseCaseReq, IdentifikasiRisikoStrategisPemdaUpdateUseCaseRes]

func ImplIdentifikasiRisikoStrategisPemdaUpdateUseCase(
	getIdentifikasiRisikoStrategisPemdaById gateway.IdentifikasiRisikoStrategisPemdaGetByID,
	updateIdentifikasiRisikoStrategisPemda gateway.IdentifikasiRisikoStrategisPemdaSave,
	kodeRisikoByID gateway.KategoriRisikoGetByID,
	RcaByID gateway.RcaGetByID,
) IdentifikasiRisikoStrategisPemdaUpdateUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaUpdateUseCaseReq) (*IdentifikasiRisikoStrategisPemdaUpdateUseCaseRes, error) {

		res, err := getIdentifikasiRisikoStrategisPemdaById(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		tahunPenilaian, err := extractYear(req.TahunPenilaian)
		if err != nil {
			return nil, fmt.Errorf("invalid TahunPenilaian format: %v", err)
		}

		res.IdentifikasiRisikoStrategisPemda.TahunPenilaian = &tahunPenilaian
		res.IdentifikasiRisikoStrategisPemda.Periode = &req.Periode
		res.IdentifikasiRisikoStrategisPemda.UrusanPemerintahan = &req.UrusanPemerintahan
		res.IdentifikasiRisikoStrategisPemda.TujuanStrategis = &req.TujuanStrategis
		res.IdentifikasiRisikoStrategisPemda.IndikatorKinerja = &req.IndikatorKinerja
		res.IdentifikasiRisikoStrategisPemda.UraianRisiko = &req.UraianRisiko
		res.IdentifikasiRisikoStrategisPemda.PemilikRisiko = &req.PemilikRisiko
		res.IdentifikasiRisikoStrategisPemda.Controllable = &req.Controllable
		res.IdentifikasiRisikoStrategisPemda.UraianDampak = &req.UraianDampak
		res.IdentifikasiRisikoStrategisPemda.PihakDampak = &req.PihakDampak

		res.IdentifikasiRisikoStrategisPemda.KategoriRisikoID = &req.KategoriRisikoID
		kategoriRisikoRes, err := kodeRisikoByID(ctx, gateway.KategoriRisikoGetByIDReq{ID: req.KategoriRisikoID})
		if err != nil {
			return nil, fmt.Errorf("failed to get KategoriRisikoName: %v", err)
		}
		res.IdentifikasiRisikoStrategisPemda.KategoriRisikoName = kategoriRisikoRes.KategoriRisiko.Nama
		res.IdentifikasiRisikoStrategisPemda.GenerateKodeRisiko(*kategoriRisikoRes.KategoriRisiko.Kode)

		rcaRes, err := RcaByID(ctx, gateway.RcaGetByIDReq{ID: req.RcaID})
		if err != nil {
			return nil, fmt.Errorf("failed to get RcaName: %v", err)
		}
		res.IdentifikasiRisikoStrategisPemda.RcaID = rcaRes.Rca.ID
		res.IdentifikasiRisikoStrategisPemda.UraianSebab = rcaRes.Rca.AkarPenyebab
		res.IdentifikasiRisikoStrategisPemda.SumberSebab = rcaRes.Rca.PernyataanRisiko

		if _, err := updateIdentifikasiRisikoStrategisPemda(ctx, gateway.IdentifikasiRisikoStrategisPemdaSaveReq{IdentifikasiRisikoStrategisPemda: res.IdentifikasiRisikoStrategisPemda}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoStrategisPemdaUpdateUseCaseRes{}, nil
	}
}
