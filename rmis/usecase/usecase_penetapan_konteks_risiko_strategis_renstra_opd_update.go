package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
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
	IKUStrategis       string `json:"iku_strategis"`
	NamaYBS            string `json:"nama_ybs"`
}

type PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCaseRes struct{}

type PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCase = core.ActionHandler[PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCaseReq, PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCaseRes]

func ImplPenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCase(
	getPenetapanKonteksRisikoStrategisRenstraOPDById gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByID,
	updatePenetapanKonteksRisikoStrategisRenstraOPD gateway.PenetepanKonteksRisikoStrategisRenstraOPDSave,
) PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCaseReq) (*PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCaseRes, error) {

		res, err := getPenetapanKonteksRisikoStrategisRenstraOPDById(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.PenetapanKonteksRisikoStrategisRenstraOPD.NamaPemda = &req.NamaPemda
		res.PenetapanKonteksRisikoStrategisRenstraOPD.Periode = &req.Periode
		res.PenetapanKonteksRisikoStrategisRenstraOPD.TahunPenilaian = &req.TahunPenilaian
		res.PenetapanKonteksRisikoStrategisRenstraOPD.UrusanPemerintahan = &req.UrusanPemerintahan
		res.PenetapanKonteksRisikoStrategisRenstraOPD.OPDID = &req.OPDID
		res.PenetapanKonteksRisikoStrategisRenstraOPD.TujuanStrategis = &req.TujuanStrategis
		res.PenetapanKonteksRisikoStrategisRenstraOPD.SasaranStrategis = &req.SasaranStrategis
		res.PenetapanKonteksRisikoStrategisRenstraOPD.InformasiLain = &req.InformasiLain
		res.PenetapanKonteksRisikoStrategisRenstraOPD.IKUStrategis = &req.IKUStrategis
		res.PenetapanKonteksRisikoStrategisRenstraOPD.NamaYBS = &req.NamaYBS

		if _, err := updatePenetapanKonteksRisikoStrategisRenstraOPD(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDSaveReq{PenetepanKonteksRisikoStrategisRenstraOPD: res.PenetapanKonteksRisikoStrategisRenstraOPD}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCaseRes{}, nil
	}
}
