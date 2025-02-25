package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseReq struct {
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

type PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseRes struct {
	ID string `json:"id"`
}

type PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCase = core.ActionHandler[PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseReq, PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseRes]

func ImplPenetapanKonteksRisikoStrategisRenstraOPDCreateUseCase(
	generateId gateway.GenerateId,
	createPenetapanKonteksRisikoStrategisRenstraOPD gateway.PenetepanKonteksRisikoStrategisRenstraOPDSave,
) PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseReq) (*PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		obj := model.PenetapanKonteksRisikoStrategisRenstraOPD{
			ID:                 &genObj.RandomId,
			NamaPemda:          &req.NamaPemda,
			TahunPenilaian:     &req.TahunPenilaian,
			Periode:            &req.Periode,
			UrusanPemerintahan: &req.UrusanPemerintahan,
			OPDID:              &req.OPDID,
			TujuanStrategis:    &req.TujuanStrategis,
			SasaranStrategis:   &req.SasaranStrategis,
			InformasiLain:      &req.InformasiLain,
			IKUStrategis:       &req.IKUStrategis,
			NamaYBS:            &req.NamaYBS,
		}

		if _, err = createPenetapanKonteksRisikoStrategisRenstraOPD(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDSaveReq{PenetepanKonteksRisikoStrategisRenstraOPD: obj}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
