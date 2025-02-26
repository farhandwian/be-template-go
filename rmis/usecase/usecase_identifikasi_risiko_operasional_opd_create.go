package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type IdentifikasiRisikoOperasionalOPDCreateUseCaseReq struct {
	NamaPemda          string `json:"nama_pemda"`
	OPDID              string `json:"opd_id"`
	TahunPenilaian     string `json:"tahun_penilaian"`
	Periode            string `json:"periode"`
	UrusanPemerintahan string `json:"urusan_pemerintahan"`
	IndikatorKeluaran  string `json:"indikator_keluaran"`
	TahapRisiko        string `json:"tahap_risiko"`
	KategoriRisikoID   string `json:"kategori_resiko_id"`
	NomorUraianRisiko  int    `json:"nomor_uraian_risiko"`
	UraianRisiko       string `json:"uraian_resiko"`
	PemilikRisiko      string `json:"pemilik_resiko"`
	Controllable       string `json:"controllable"`
	UraianDampak       string `json:"uraian_dampak"`
	PihakDampak        string `json:"pihak_dampak"`
}

type IdentifikasiRisikoOperasionalOPDCreateUseCaseRes struct {
	ID string `json:"id"`
}

type IdentifikasiRisikoOperasionalOPDCreateUseCase = core.ActionHandler[
	IdentifikasiRisikoOperasionalOPDCreateUseCaseReq,
	IdentifikasiRisikoOperasionalOPDCreateUseCaseRes,
]

func ImplIdentifikasiRisikoOperasionalOPDCreateUseCase(
	generateId gateway.GenerateId,
	createIdentifikasiRisikoOperasionalOPD gateway.IdentifikasiRisikoOperasionalOPDSave,
	kodeRisikoByID gateway.KategoriRisikoGetByID,
	getOneOPD gateway.OPDGetByID,
) IdentifikasiRisikoOperasionalOPDCreateUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDCreateUseCaseReq) (*IdentifikasiRisikoOperasionalOPDCreateUseCaseRes, error) {
		fmt.Println("IdentifikasiRisikoOperasionalOPDCreateUseCase")
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

		opd, err := getOneOPD(ctx, gateway.OPDGetByIDReq{ID: req.OPDID})
		if err != nil {
			return nil, err
		}

		obj := model.IdentifikasiRisikoOperasionalOPD{
			ID:                 &genObj.RandomId,
			NamaPemda:          &req.NamaPemda,
			OPDID:              &req.OPDID,
			TahunPenilaian:     &tahunPenilaian,
			Periode:            &req.Periode,
			UrusanPemerintahan: &req.UrusanPemerintahan,
			IndikatorKeluaran:  &req.IndikatorKeluaran,
			TahapRisiko:        &req.TahapRisiko,
			KategoriRisikoID:   &req.KategoriRisikoID,
			NomorUraianRisiko:  &req.NomorUraianRisiko,
			UraianRisiko:       &req.UraianRisiko,
			PemilikRisiko:      &req.PemilikRisiko,
			Controllable:       &req.Controllable,
			UraianDampak:       &req.UraianDampak,
			PihakDampak:        &req.PihakDampak,
		}

		obj.GenerateKodeRisiko(*kategoriRisikoRes.KategoriRisiko.Kode, *opd.OPD.Kode)
		// Save the new entry
		if _, err = createIdentifikasiRisikoOperasionalOPD(ctx, gateway.IdentifikasiRisikoOperasionalOPDSaveReq{
			IdentifikasiRisikoOperasionalOPD: obj,
		}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoOperasionalOPDCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
