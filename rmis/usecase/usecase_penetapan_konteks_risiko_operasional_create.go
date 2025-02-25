package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/helper"
)

type PenetapanKonteksRisikoOperasionalCreateUseCaseReq struct {
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

type PenetapanKonteksRisikoOperasionalCreateUseCaseRes struct {
	ID string `json:"id"`
}

type PenetapanKonteksRisikoOperasionalCreateUseCase = core.ActionHandler[PenetapanKonteksRisikoOperasionalCreateUseCaseReq, PenetapanKonteksRisikoOperasionalCreateUseCaseRes]

func ImplPenetapanKonteksRisikoOperasionalCreateUseCase(
	generateId gateway.GenerateId,
	createPenetapanKonteksRisikoOperasional gateway.PenetepanKonteksRisikoOperasionalSave,
) PenetapanKonteksRisikoOperasionalCreateUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoOperasionalCreateUseCaseReq) (*PenetapanKonteksRisikoOperasionalCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		programInspektorat := helper.ToDataTypeJSONPtr(req.ProgramInspektorat)

		obj := model.PenetapanKonteksRisikoOperasional{
			ID:                           &genObj.RandomId,
			NamaPemda:                    &req.NamaPemda,
			Periode:                      &req.Periode,
			TahunPenilaian:               &req.TahunPenilaian,
			UrusanPemerintahan:           &req.UrusanPemerintahan,
			OPDID:                        &req.OPDID,
			TujuanStrategis:              &req.TujuanStrategis,
			ProgramInspektorat:           programInspektorat,
			InformasiLain:                &req.InformasiLain,
			KegiatanDanIndikatorKeluaran: &req.KegiatanDanIndikatorKeluaran,
			NamaYBS:                      &req.NamaYBS,
		}

		if _, err = createPenetapanKonteksRisikoOperasional(ctx, gateway.PenetapanKonteksRisikoOperasionalSaveReq{PenetepanKonteksRisikoOperasional: obj}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoOperasionalCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
