package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCaseReq struct {
	ID                 string `json:"id"`
	NamaPemda          string `json:"nama_pemda"`
	TahunPenilaian     string `json:"tahun_penilaian"`
	Periode            string `json:"periode"`
	UrusanPemerintahan string `json:"urusan_pemerintahan"`
	OPDID              string `json:"opd_id"`
	TujuanStrategis    string `json:"tujuan_strategis"`
	SasaranStrategis   string `json:"sasaran_strategis"`
	InformasiLain      string `json:"informasi_lain"`
	PenetapanTujuan    string `json:"penetapan_tujuan"`
	PenetapanSasaran   string `json:"penetapan_sasaran"`
	PenetapanIku       string `json:"penetapan_iku"`
}

type PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCaseRes struct{}

type PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCase = core.ActionHandler[PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCaseReq, PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCaseRes]

func ImplPenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCase(
	getPenetapanKonteksRisikoStrategisRenstraOPDById gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByID,
	updatePenetapanKonteksRisikoStrategisRenstraOPD gateway.PenetepanKonteksRisikoStrategisRenstraOPDSave,
	OpdByID gateway.OPDGetByID,
) PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCaseReq) (*PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCaseRes, error) {

		res, err := getPenetapanKonteksRisikoStrategisRenstraOPDById(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		_, err = OpdByID(ctx, gateway.OPDGetByIDReq{ID: req.OPDID})
		if err != nil {
			return nil, err
		}

		tahunPenilaian, err := extractYear(req.TahunPenilaian)
		if err != nil {
			return nil, fmt.Errorf("invalid TahunPenilaian format: %v", err)
		}
		penetapanKonteksRisikoStrategisRenstraOPD := res.PenetapanKonteksRisikoStrategisRenstraOPD

		penetapanKonteksRisikoStrategisRenstraOPD.NamaPemda = &req.NamaPemda
		penetapanKonteksRisikoStrategisRenstraOPD.Periode = &req.Periode
		penetapanKonteksRisikoStrategisRenstraOPD.TahunPenilaian = &tahunPenilaian
		penetapanKonteksRisikoStrategisRenstraOPD.UrusanPemerintahan = &req.UrusanPemerintahan
		penetapanKonteksRisikoStrategisRenstraOPD.OpdID = &req.OPDID
		penetapanKonteksRisikoStrategisRenstraOPD.TujuanStrategis = &req.TujuanStrategis
		penetapanKonteksRisikoStrategisRenstraOPD.SasaranStrategis = &req.SasaranStrategis
		penetapanKonteksRisikoStrategisRenstraOPD.InformasiLain = &req.InformasiLain
		penetapanKonteksRisikoStrategisRenstraOPD.PenetapanTujuan = &req.PenetapanTujuan
		penetapanKonteksRisikoStrategisRenstraOPD.PenetapanSasaran = &req.PenetapanSasaran
		penetapanKonteksRisikoStrategisRenstraOPD.PenetapanIku = &req.PenetapanIku
		penetapanKonteksRisikoStrategisRenstraOPD.Status = sharedModel.StatusMenungguVerifikasi

		if _, err := updatePenetapanKonteksRisikoStrategisRenstraOPD(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDSaveReq{PenetepanKonteksRisikoStrategisRenstraOPD: penetapanKonteksRisikoStrategisRenstraOPD}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCaseRes{}, nil
	}
}
