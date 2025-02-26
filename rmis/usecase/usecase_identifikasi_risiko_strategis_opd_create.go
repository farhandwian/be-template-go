package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type IdentifikasiRisikoStrategisOPDCreateUseCaseReq struct {
	NamaOPD            string `json:"nama_pemda"`
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

type IdentifikasiRisikoStrategisOPDCreateUseCaseRes struct {
	ID string `json:"id"`
}

type IdentifikasiRisikoStrategisOPDCreateUseCase = core.ActionHandler[
	IdentifikasiRisikoStrategisOPDCreateUseCaseReq,
	IdentifikasiRisikoStrategisOPDCreateUseCaseRes,
]

func ImplIdentifikasiRisikoStrategisOPDCreateUseCase(
	generateId gateway.GenerateId,
	createIdentifikasiRisikoStrategisOPD gateway.IdentifikasiRisikoStrategisOPDSave,
	kodeRisikoByID gateway.KategoriRisikoGetByID,
) IdentifikasiRisikoStrategisOPDCreateUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDCreateUseCaseReq) (*IdentifikasiRisikoStrategisOPDCreateUseCaseRes, error) {

		// Generate a unique ID
		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		tahunPenilaian, err := extractYear(req.TahunPenilaian) // ex : 2024
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

		obj := model.IdentifikasiRisikoStrategisOPD{
			ID:                 &genObj.RandomId,
			NamaOPD:            &req.NamaOPD,
			TahunPenilaian:     &tahunPenilaian,
			Periode:            &req.Periode,
			UrusanPemerintahan: &req.UrusanPemerintahan,

			IndikatorKinerja: &req.IndikatorKinerja,
			KategoriRisikoID: &req.KategoriRisikoID,

			UraianRisiko:  &req.UraianRisiko,
			PemilikRisiko: &req.PemilikRisiko,
			Controllable:  &req.Controllable,
			UraianDampak:  &req.UraianDampak,
			PihakDampak:   &req.PihakDampak,
		}

		obj.GenerateKodeRisiko(*kategoriRisikoRes.KategoriRisiko.Kode)
		// Save the new entry
		if _, err = createIdentifikasiRisikoStrategisOPD(ctx, gateway.IdentifikasiRisikoStrategisOPDSaveReq{
			IdentifikasiRisikoStrategisOPD: obj,
		}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoStrategisOPDCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
