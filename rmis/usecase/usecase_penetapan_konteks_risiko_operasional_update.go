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

type PenetapanKonteksRisikoOperasionalUpdateUseCaseReq struct {
	ID                        string                          `json:"id"`
	NamaPemda                 string                          `json:"nama_pemda"`
	TahunPenilaian            string                          `json:"tahun_penilaian"`
	Periode                   string                          `json:"periode"`
	SumberData                string                          `json:"sumber_data"`
	UrusanPemerintahan        string                          `json:"urusan_pemerintahan"`
	OpdID                     string                          `json:"opd_id"`
	TujuanStrategis           string                          `json:"tujuan_strategis"`
	KegiatanUtama             string                          `json:"kegiatan_utama"`
	InformasiLain             string                          `json:"informasi_lain"`
	KeluaranAtauHasilKegiatan model.KeluaranAtauHasilKegiatan `json:"keluaran_atau_hasil_kegiatan"`
}

type PenetapanKonteksRisikoOperasionalUpdateUseCaseRes struct{}

type PenetapanKonteksRisikoOperasionalUpdateUseCase = core.ActionHandler[PenetapanKonteksRisikoOperasionalUpdateUseCaseReq, PenetapanKonteksRisikoOperasionalUpdateUseCaseRes]

func ImplPenetapanKonteksRisikoOperasionalUpdateUseCase(
	getPenetapanKonteksRisikoOperasionalById gateway.PenetapanKonteksRisikoOperasionalGetByID,
	updatePenetapanKonteksRisikoOperasional gateway.PenetepanKonteksRisikoOperasionalSave,
	OpdByID gateway.OPDGetByID,
) PenetapanKonteksRisikoOperasionalUpdateUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoOperasionalUpdateUseCaseReq) (*PenetapanKonteksRisikoOperasionalUpdateUseCaseRes, error) {

		res, err := getPenetapanKonteksRisikoOperasionalById(ctx, gateway.PenetapanKonteksRisikoOperasionalGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		_, err = OpdByID(ctx, gateway.OPDGetByIDReq{ID: req.OpdID})
		if err != nil {
			return nil, err
		}

		keluaranAtauHasilKegiatan := helper.ToDataTypeJSONPtr(req.KeluaranAtauHasilKegiatan)

		tahunPenilaian, err := extractYear(req.TahunPenilaian)
		if err != nil {
			return nil, fmt.Errorf("invalid TahunPenilaian format: %v", err)
		}

		res.PenetapanKonteksRisikoOperasional.NamaPemda = &req.NamaPemda
		res.PenetapanKonteksRisikoOperasional.Periode = &req.Periode
		res.PenetapanKonteksRisikoOperasional.TahunPenilaian = &tahunPenilaian
		res.PenetapanKonteksRisikoOperasional.UrusanPemerintahan = &req.UrusanPemerintahan
		res.PenetapanKonteksRisikoOperasional.OpdID = &req.OpdID
		res.PenetapanKonteksRisikoOperasional.TujuanStrategis = &req.TujuanStrategis
		res.PenetapanKonteksRisikoOperasional.KeluaranAtauHasilKegiatan = keluaranAtauHasilKegiatan
		res.PenetapanKonteksRisikoOperasional.InformasiLain = &req.InformasiLain
		res.PenetapanKonteksRisikoOperasional.SumberData = &req.SumberData
		res.PenetapanKonteksRisikoOperasional.Status = sharedModel.StatusMenungguVerifikasi

		if _, err := updatePenetapanKonteksRisikoOperasional(ctx, gateway.PenetapanKonteksRisikoOperasionalSaveReq{PenetepanKonteksRisikoOperasional: res.PenetapanKonteksRisikoOperasional}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoOperasionalUpdateUseCaseRes{}, nil
	}
}
