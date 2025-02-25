package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
	"shared/helper"
)

type PenetapanKonteksRisikoOperasionalUpdateUseCaseReq struct {
	ID                           string `json:"id"`
	NamaPemda                    string `json:"nama_pemda"`
	TahunPenilaian               string `json:"tahun_penilaian"`
	Periode                      string `json:"periode"`
	UrusanPemerintahan           string `json:"urusan_pemerintahan"`
	OPDID                        string `json:"opd_id"` // references opd
	TujuanStrategis              string `json:"tujuan_strategis"`
	ProgramInspektorat           string `json:"program_inspektorat"`
	InformasiLain                string `json:"informasi_lain"`
	KegiatanDanIndikatorKeluaran string `json:"kegiatan_dan_indikator_keluaran"` // Kegiatan,dan indikator keluaran yang akan dilakukan penilaian risiko
	NamaYBS                      string `json:"nama_ybs"`
}

type PenetapanKonteksRisikoOperasionalUpdateUseCaseRes struct{}

type PenetapanKonteksRisikoOperasionalUpdateUseCase = core.ActionHandler[PenetapanKonteksRisikoOperasionalUpdateUseCaseReq, PenetapanKonteksRisikoOperasionalUpdateUseCaseRes]

func ImplPenetapanKonteksRisikoOperasionalUpdateUseCase(
	getPenetapanKonteksRisikoOperasionalById gateway.PenetapanKonteksRisikoOperasionalGetByID,
	updatePenetapanKonteksRisikoOperasional gateway.PenetepanKonteksRisikoOperasionalSave,
) PenetapanKonteksRisikoOperasionalUpdateUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoOperasionalUpdateUseCaseReq) (*PenetapanKonteksRisikoOperasionalUpdateUseCaseRes, error) {

		res, err := getPenetapanKonteksRisikoOperasionalById(ctx, gateway.PenetapanKonteksRisikoOperasionalGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		programInspektorat := helper.ToDataTypeJSONPtr(req.ProgramInspektorat)

		res.PenetapanKonteksRisikoOperasional.NamaPemda = &req.NamaPemda
		res.PenetapanKonteksRisikoOperasional.Periode = &req.Periode
		res.PenetapanKonteksRisikoOperasional.TahunPenilaian = &req.TahunPenilaian
		res.PenetapanKonteksRisikoOperasional.UrusanPemerintahan = &req.UrusanPemerintahan
		res.PenetapanKonteksRisikoOperasional.OPDID = &req.OPDID
		res.PenetapanKonteksRisikoOperasional.TujuanStrategis = &req.TujuanStrategis
		res.PenetapanKonteksRisikoOperasional.ProgramInspektorat = programInspektorat
		res.PenetapanKonteksRisikoOperasional.InformasiLain = &req.InformasiLain
		res.PenetapanKonteksRisikoOperasional.KegiatanDanIndikatorKeluaran = &req.KegiatanDanIndikatorKeluaran
		res.PenetapanKonteksRisikoOperasional.NamaYBS = &req.NamaYBS

		if _, err := updatePenetapanKonteksRisikoOperasional(ctx, gateway.PenetapanKonteksRisikoOperasionalSaveReq{PenetepanKonteksRisikoOperasional: res.PenetapanKonteksRisikoOperasional}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoOperasionalUpdateUseCaseRes{}, nil
	}
}
