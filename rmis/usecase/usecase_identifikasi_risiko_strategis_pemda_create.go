package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type IdentifikasiRisikoStrategisPemdaCreateUseCaseReq struct {
	NamaPemda          string `json:"nama_pemda"`
	TahunPenilaian     string `json:"tahun_penilaian"`
	Periode            string `json:"periode"`
	UrusanPemerintahan string `json:"urusan_pemerintahan"`
	TujuanStrategis    string `json:"tujuan_strategis"`
	IndikatorKinerja   string `json:"indikator_kinerja"`
	KategoriRisikoID   string `json:"kategori_resiko_id"`
	UraianRisiko       string `json:"uraian_resiko"`
	PemilikRisiko      string `json:"pemilik_resiko"`
	Controllable       string `json:"controllable"`
	UraianDampak       string `json:"uraian_dampak"`
	PihakDampak        string `json:"pihak_dampak"`
}

type IdentifikasiRisikoStrategisPemdaCreateUseCaseRes struct {
	ID string `json:"id"`
}

type IdentifikasiRisikoStrategisPemdaCreateUseCase = core.ActionHandler[
	IdentifikasiRisikoStrategisPemdaCreateUseCaseReq,
	IdentifikasiRisikoStrategisPemdaCreateUseCaseRes,
]

func ImplIdentifikasiRisikoStrategisPemdaCreateUseCase(
	generateId gateway.GenerateId,
	createIdentifikasiRisikoStrategisPemda gateway.IdentifikasiRisikoStrategisPemdaSave,
	kodeRisikoByID gateway.KategoriRisikoGetByID,
) IdentifikasiRisikoStrategisPemdaCreateUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaCreateUseCaseReq) (*IdentifikasiRisikoStrategisPemdaCreateUseCaseRes, error) {

		// Generate a unique ID
		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		tahunPenilaian, err := extractYear(req.TahunPenilaian)
		if err != nil {
			return nil, fmt.Errorf("invalid TahunPenilaian format: %v", err)
		}

		kategoriRisikoRes, err := kodeRisikoByID(ctx, gateway.KategoriRisikoGetByIDReq{ID: req.KategoriRisikoID})
		if err != nil {
			return nil, fmt.Errorf("failed to get KategoriRisikoName: %v", err)
		}

		var kategoriRisikoName *string
		if kategoriRisikoRes.KategoriRisiko.Nama != nil {
			kategoriRisikoName = kategoriRisikoRes.KategoriRisiko.Nama
		}

		obj := model.IdentifikasiRisikoStrategisPemerintahDaerah{
			ID:                 &genObj.RandomId,
			NamaPemda:          &req.NamaPemda,
			TahunPenilaian:     &tahunPenilaian,
			Periode:            &req.Periode,
			UrusanPemerintahan: &req.UrusanPemerintahan,
			TujuanStrategis:    &req.TujuanStrategis,
			IndikatorKinerja:   &req.IndikatorKinerja,
			KategoriRisikoID:   &req.KategoriRisikoID,
			KategoriRisikoName: kategoriRisikoName,
			UraianRisiko:       &req.UraianRisiko,
			PemilikRisiko:      &req.PemilikRisiko,
			Controllable:       &req.Controllable,
			UraianDampak:       &req.UraianDampak,
			PihakDampak:        &req.PihakDampak,
		}

		obj.GenerateKodeRisiko(*kategoriRisikoRes.KategoriRisiko.Kode)
		// Save the new entry
		if _, err = createIdentifikasiRisikoStrategisPemda(ctx, gateway.IdentifikasiRisikoStrategisPemdaSaveReq{
			IdentifikasiRisikoStrategisPemda: obj,
		}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoStrategisPemdaCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
