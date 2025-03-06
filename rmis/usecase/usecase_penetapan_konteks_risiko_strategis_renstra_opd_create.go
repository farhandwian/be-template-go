package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	sharedModel "shared/model"
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
	PenetapanTujuan    string `json:"penetapan_tujuan"`
	PenetapanSasaran   string `json:"penetapan_sasaran"`
	PenetapanIku       string `json:"penetapan_iku"`
}

type PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseRes struct {
	ID string `json:"id"`
}

type PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCase = core.ActionHandler[PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseReq, PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseRes]

func ImplPenetapanKonteksRisikoStrategisRenstraOPDCreateUseCase(
	generateId gateway.GenerateId,
	createPenetapanKonteksRisikoStrategisRenstraOPD gateway.PenetepanKonteksRisikoStrategisRenstraOPDSave,
	OpdByID gateway.OPDGetByID,
) PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseReq) (*PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
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
		obj := model.PenetapanKonteksRisikoStrategisRenstraOPD{
			ID:                 &genObj.RandomId,
			NamaPemda:          &req.NamaPemda,
			TahunPenilaian:     &tahunPenilaian,
			Periode:            &req.Periode,
			UrusanPemerintahan: &req.UrusanPemerintahan,
			OpdID:              &req.OPDID,
			TujuanStrategis:    &req.TujuanStrategis,
			SasaranStrategis:   &req.SasaranStrategis,
			InformasiLain:      &req.InformasiLain,
			PenetapanTujuan:    &req.PenetapanTujuan,
			PenetapanSasaran:   &req.PenetapanSasaran,
			PenetapanIku:       &req.PenetapanIku,
			Status:             sharedModel.StatusMenungguVerifikasi,
		}

		if _, err = createPenetapanKonteksRisikoStrategisRenstraOPD(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDSaveReq{PenetepanKonteksRisikoStrategisRenstraOPD: obj}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoStrategisRenstraOPDCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
