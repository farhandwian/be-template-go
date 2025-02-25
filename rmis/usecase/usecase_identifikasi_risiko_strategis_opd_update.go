package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"shared/core"
)

type IdentifikasiRisikoStrategisOPDUpdateUseCaseReq struct {
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

type IdentifikasiRisikoStrategisOPDUpdateUseCaseRes struct{}

type IdentifikasiRisikoStrategisOPDUpdateUseCase = core.ActionHandler[IdentifikasiRisikoStrategisOPDUpdateUseCaseReq, IdentifikasiRisikoStrategisOPDUpdateUseCaseRes]

func ImplIdentifikasiRisikoStrategisOPDUpdateUseCase(
	getIdentifikasiRisikoStrategisOPDById gateway.IdentifikasiRisikoStrategisOPDGetByID,
	updateIdentifikasiRisikoStrategisOPD gateway.IdentifikasiRisikoStrategisOPDSave,
	kodeRisikoByID gateway.KategoriRisikoGetByID,
	RcaByID gateway.RcaGetByID,
) IdentifikasiRisikoStrategisOPDUpdateUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDUpdateUseCaseReq) (*IdentifikasiRisikoStrategisOPDUpdateUseCaseRes, error) {

		res, err := getIdentifikasiRisikoStrategisOPDById(ctx, gateway.IdentifikasiRisikoStrategisOPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		if req.TahunPenilaian != "" {
			year, err := extractYear(req.TahunPenilaian)
			if err != nil {
				return nil, fmt.Errorf("invalid TahunPenilaian format: %v", err)
			}
			res.IdentifikasiRisikoStrategisOPD.TahunPenilaian = &year
		}
		res.IdentifikasiRisikoStrategisOPD.Periode = &req.Periode
		res.IdentifikasiRisikoStrategisOPD.UrusanPemerintahan = &req.UrusanPemerintahan
		res.IdentifikasiRisikoStrategisOPD.TujuanStrategis = &req.TujuanStrategis
		res.IdentifikasiRisikoStrategisOPD.IndikatorKinerja = &req.IndikatorKinerja
		res.IdentifikasiRisikoStrategisOPD.UraianRisiko = &req.UraianRisiko
		res.IdentifikasiRisikoStrategisOPD.PemilikRisiko = &req.PemilikRisiko
		res.IdentifikasiRisikoStrategisOPD.Controllable = &req.Controllable
		res.IdentifikasiRisikoStrategisOPD.UraianDampak = &req.UraianDampak
		res.IdentifikasiRisikoStrategisOPD.PihakDampak = &req.PihakDampak

		if res.IdentifikasiRisikoStrategisOPD.KategoriRisikoID == nil || *res.IdentifikasiRisikoStrategisOPD.KategoriRisikoID == "" {
			return nil, fmt.Errorf("KategoriRisikoID is missing in the database record")
		}

		kategoriRisikoRes, err := kodeRisikoByID(ctx, gateway.KategoriRisikoGetByIDReq{
			ID: *res.IdentifikasiRisikoStrategisOPD.KategoriRisikoID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get KategoriRisikoName: %v", err)
		}

		res.IdentifikasiRisikoStrategisOPD.KategoriRisikoName = kategoriRisikoRes.KategoriRisiko.Nama

		if res.IdentifikasiRisikoStrategisOPD.KodeRisiko == nil || *res.IdentifikasiRisikoStrategisOPD.KodeRisiko == "" {
			res.IdentifikasiRisikoStrategisOPD.GenerateKodeRisiko(*kategoriRisikoRes.KategoriRisiko.Kode)
		}

		rcaRes, err := RcaByID(ctx, gateway.RcaGetByIDReq{ID: req.RcaID})
		if err != nil {
			return nil, fmt.Errorf("failed to get RcaName: %v", err)
		}
		res.IdentifikasiRisikoStrategisOPD.RcaID = rcaRes.Rca.ID
		res.IdentifikasiRisikoStrategisOPD.UraianSebab = rcaRes.Rca.AkarPenyebab
		res.IdentifikasiRisikoStrategisOPD.SumberSebab = rcaRes.Rca.PernyataanRisiko

		if _, err := updateIdentifikasiRisikoStrategisOPD(ctx, gateway.IdentifikasiRisikoStrategisOPDSaveReq{IdentifikasiRisikoStrategisOPD: res.IdentifikasiRisikoStrategisOPD}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoStrategisOPDUpdateUseCaseRes{}, nil
	}
}
