package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/helper"
)

type SimpulanKondisiKelemahanLingkunganCreateUseCaseReq struct {
	NamaPemda          string `json:"nama_pemda"`
	TahunPenilaian     string `json:"tahun_penilaian"`
	UrusanPemerintahan string `json:"urusan_pemerintahan"`
	SumberData         string `json:"sumber_data"`
	UraianKelamahan    string `json:"uraian_kelemahan"`
	Klasifikasi        string `json:"klasifikasi"`
}

type SimpulanKondisiKelemahanLingkunganCreateUseCaseRes struct {
	ID string `json:"id"`
}

type SimpulanKondisiKelemahanLingkunganCreateUseCase = core.ActionHandler[SimpulanKondisiKelemahanLingkunganCreateUseCaseReq, SimpulanKondisiKelemahanLingkunganCreateUseCaseRes]

func ImplSimpulanKondisiKelemahanLingkunganCreateUseCase(
	generateId gateway.GenerateId,
	createSimpulanKondisiKelemahanLingkungan gateway.SimpulanKondisiKelemahanLingkunganSave,
) SimpulanKondisiKelemahanLingkunganCreateUseCase {
	return func(ctx context.Context, req SimpulanKondisiKelemahanLingkunganCreateUseCaseReq) (*SimpulanKondisiKelemahanLingkunganCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		tahunPenilaian, err := extractYear(req.TahunPenilaian)
		if err != nil {
			return nil, fmt.Errorf("invalid TahunPenilaian format: %v", err)
		}

		uraianKelemahanJSON := helper.ToDataTypeJSONPtr(req.UraianKelamahan)

		obj := model.SimpulanKondisiKelemahanLingkungan{
			ID:                 &genObj.RandomId,
			NamaPemda:          &req.NamaPemda,
			TahunPenilaian:     &tahunPenilaian,
			UrusanPemerintahan: &req.UrusanPemerintahan,
			SumberData:         &req.SumberData,
			UraianKelamahan:    uraianKelemahanJSON,
			Klasifikasi:        &req.Klasifikasi,
		}

		if _, err = createSimpulanKondisiKelemahanLingkungan(ctx, gateway.SimpulanKondisiKelemahanLingkunganSaveReq{SimpulanKondisiKelemahanLingkungan: obj}); err != nil {
			return nil, err
		}

		return &SimpulanKondisiKelemahanLingkunganCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
