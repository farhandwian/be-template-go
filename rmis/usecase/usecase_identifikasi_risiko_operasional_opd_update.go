package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"shared/core"
)

type IdentifikasiRisikoOperasionalOPDUpdateUseCaseReq struct {
	ID                 string `json:"id"`
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

type IdentifikasiRisikoOperasionalOPDUpdateUseCaseRes struct{}

type IdentifikasiRisikoOperasionalOPDUpdateUseCase = core.ActionHandler[IdentifikasiRisikoOperasionalOPDUpdateUseCaseReq, IdentifikasiRisikoOperasionalOPDUpdateUseCaseRes]

func ImplIdentifikasiRisikoOperasionalOPDUpdateUseCase(
	getIdentifikasiRisikoOperasionalOPDById gateway.IdentifikasiRisikoOperasionalOPDGetByID,
	updateIdentifikasiRisikoOperasionalOPD gateway.IdentifikasiRisikoOperasionalOPDSave,
	kodeRisikoByID gateway.KategoriRisikoGetByID,
	RcaByID gateway.RcaGetByID,
	getOneOPD gateway.OPDGetByID,
) IdentifikasiRisikoOperasionalOPDUpdateUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDUpdateUseCaseReq) (*IdentifikasiRisikoOperasionalOPDUpdateUseCaseRes, error) {

		res, err := getIdentifikasiRisikoOperasionalOPDById(ctx, gateway.IdentifikasiRisikoOperasionalOPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		if req.TahunPenilaian != "" {
			year, err := extractYear(req.TahunPenilaian)
			if err != nil {
				return nil, fmt.Errorf("invalid TahunPenilaian format: %v", err)
			}
			res.IdentifikasiRisikoOperasionalOPD.TahunPenilaian = &year
		}
		res.IdentifikasiRisikoOperasionalOPD.NamaPemda = &req.NamaPemda
		res.IdentifikasiRisikoOperasionalOPD.OPDID = &req.OPDID
		res.IdentifikasiRisikoOperasionalOPD.Periode = &req.Periode
		res.IdentifikasiRisikoOperasionalOPD.UrusanPemerintahan = &req.UrusanPemerintahan
		res.IdentifikasiRisikoOperasionalOPD.IndikatorKeluaran = &req.IndikatorKeluaran
		res.IdentifikasiRisikoOperasionalOPD.TahapRisiko = &req.TahapRisiko
		res.IdentifikasiRisikoOperasionalOPD.KategoriRisikoID = &req.KategoriRisikoID
		res.IdentifikasiRisikoOperasionalOPD.NomorUraianRisiko = &req.NomorUraianRisiko
		res.IdentifikasiRisikoOperasionalOPD.UraianRisiko = &req.UraianRisiko
		res.IdentifikasiRisikoOperasionalOPD.PemilikRisiko = &req.PemilikRisiko
		res.IdentifikasiRisikoOperasionalOPD.Controllable = &req.Controllable
		res.IdentifikasiRisikoOperasionalOPD.UraianDampak = &req.UraianDampak
		res.IdentifikasiRisikoOperasionalOPD.PihakDampak = &req.PihakDampak

		if res.IdentifikasiRisikoOperasionalOPD.KategoriRisikoID == nil || *res.IdentifikasiRisikoOperasionalOPD.KategoriRisikoID == "" {
			return nil, fmt.Errorf("KategoriRisikoID is missing in the database record")
		}

		kategoriRisikoRes, err := kodeRisikoByID(ctx, gateway.KategoriRisikoGetByIDReq{
			ID: *res.IdentifikasiRisikoOperasionalOPD.KategoriRisikoID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get KategoriRisiko: %v", err)
		}

		opd, err := getOneOPD(ctx, gateway.OPDGetByIDReq{ID: req.OPDID})
		if err != nil {
			return nil, fmt.Errorf("failed to get OPDName: %v", err)
		}

		if res.IdentifikasiRisikoOperasionalOPD.KodeRisiko == nil || *res.IdentifikasiRisikoOperasionalOPD.KodeRisiko == "" {
			res.IdentifikasiRisikoOperasionalOPD.GenerateKodeRisiko(*kategoriRisikoRes.KategoriRisiko.Kode, *opd.OPD.Kode)
		}

		// rcaRes, err := RcaByID(ctx, gateway.RcaGetByIDReq{ID: req.RcaID})
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to get RcaName: %v", err)
		// }

		if _, err := updateIdentifikasiRisikoOperasionalOPD(ctx, gateway.IdentifikasiRisikoOperasionalOPDSaveReq{IdentifikasiRisikoOperasionalOPD: res.IdentifikasiRisikoOperasionalOPD}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoOperasionalOPDUpdateUseCaseRes{}, nil
	}
}
