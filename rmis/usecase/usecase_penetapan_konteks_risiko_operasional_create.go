package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/helper"
	sharedModel "shared/model"
)

type PenetapanKonteksRisikoOperasionalCreateUseCaseReq struct {
	NamaPemda                 string                            `json:"nama_pemda"`
	TahunPenilaian            string                            `json:"tahun_penilaian"`
	Periode                   string                            `json:"periode"`
	SumberData                string                            `json:"sumber_data"`
	UrusanPemerintahan        string                            `json:"urusan_pemerintahan"`
	OpdID                     string                            `json:"opd_id"`
	TujuanStrategis           string                            `json:"tujuan_strategis"`
	KegiatanUtama             string                            `json:"kegiatan_utama"`
	InformasiLain             string                            `json:"informasi_lain"`
	PenetapanKegiatan         string                            `json:"penetapan_kegiatan"`
	KeluaranAtauHasilKegiatan []model.KeluaranAtauHasilKegiatan `json:"keluaran_atau_hasil_kegiatan"`
}

type PenetapanKonteksRisikoOperasionalCreateUseCaseRes struct {
	ID string `json:"id"`
}

type PenetapanKonteksRisikoOperasionalCreateUseCase = core.ActionHandler[PenetapanKonteksRisikoOperasionalCreateUseCaseReq, PenetapanKonteksRisikoOperasionalCreateUseCaseRes]

func ImplPenetapanKonteksRisikoOperasionalCreateUseCase(
	generateId gateway.GenerateId,
	createPenetapanKonteksRisikoOperasional gateway.PenetepanKonteksRisikoOperasionalSave,
	OpdByID gateway.OPDGetByID,
) PenetapanKonteksRisikoOperasionalCreateUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoOperasionalCreateUseCaseReq) (*PenetapanKonteksRisikoOperasionalCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}
		_, err = OpdByID(ctx, gateway.OPDGetByIDReq{ID: req.OpdID})
		if err != nil {
			return nil, err
		}

		tahunPenilaian, err := extractYear(req.TahunPenilaian)
		if err != nil {
			return nil, fmt.Errorf("invalid TahunPenilaian format: %v", err)
		}

		keluaranAtauHasilKegiatan := helper.ToDataTypeJSONPtr(req.KeluaranAtauHasilKegiatan)

		obj := model.PenetapanKonteksRisikoOperasional{
			ID:                        &genObj.RandomId,
			NamaPemda:                 &req.NamaPemda,
			Periode:                   &req.Periode,
			TahunPenilaian:            &tahunPenilaian,
			UrusanPemerintahan:        &req.UrusanPemerintahan,
			OpdID:                     &req.OpdID,
			SumberData:                &req.SumberData,
			TujuanStrategis:           &req.TujuanStrategis,
			InformasiLain:             &req.InformasiLain,
			KeluaranAtauHasilKegiatan: keluaranAtauHasilKegiatan,
			PenetapanKegiatan:         &req.PenetapanKegiatan,

			Status: sharedModel.StatusMenungguVerifikasi,
		}

		if _, err = createPenetapanKonteksRisikoOperasional(ctx, gateway.PenetapanKonteksRisikoOperasionalSaveReq{PenetepanKonteksRisikoOperasional: obj}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoOperasionalCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
