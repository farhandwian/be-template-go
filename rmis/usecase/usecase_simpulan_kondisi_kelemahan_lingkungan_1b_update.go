package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"shared/core"
	"shared/helper"
)

type SimpulanKondisiKelemahanLingkunganUpdateUseCaseReq struct {
	ID                 string `json:"id"`
	NamaPemda          string `json:"nama_pemda"`
	TahunPenilaian     string `json:"tahun_penilaian"`
	UrusanPemerintahan string `json:"urusan_pemerintahan"`
	SumberData         string `json:"sumber_data"`
	UraianKelamahan    string `json:"uraian_kelemahan"`
	Klasifikasi        string `json:"klasifikasi"`
}

type SimpulanKondisiKelemahanLingkunganUpdateUseCaseRes struct{}

type SimpulanKondisiKelemahanLingkunganUpdateUseCase = core.ActionHandler[SimpulanKondisiKelemahanLingkunganUpdateUseCaseReq, SimpulanKondisiKelemahanLingkunganUpdateUseCaseRes]

func ImplSimpulanKondisiKelemahanLingkunganUpdateUseCase(
	getSimpulanKondisiKelemahanLingkunganById gateway.SimpulanKondisiKelemahanLingkunganGetByID,
	updateSimpulanKondisiKelemahanLingkungan gateway.SimpulanKondisiKelemahanLingkunganSave,
) SimpulanKondisiKelemahanLingkunganUpdateUseCase {
	return func(ctx context.Context, req SimpulanKondisiKelemahanLingkunganUpdateUseCaseReq) (*SimpulanKondisiKelemahanLingkunganUpdateUseCaseRes, error) {

		res, err := getSimpulanKondisiKelemahanLingkunganById(ctx, gateway.SimpulanKondisiKelemahanLingkunganGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		tahunPenilaian, err := extractYear(req.TahunPenilaian)
		if err != nil {
			return nil, fmt.Errorf("invalid TahunPenilaian format: %v", err)
		}

		uraianKelemahanJSON := helper.ToDataTypeJSONPtr(req.UraianKelamahan)

		res.SimpulanKondisiKelemahanLingkungan.NamaPemda = &req.NamaPemda
		res.SimpulanKondisiKelemahanLingkungan.TahunPenilaian = &tahunPenilaian
		res.SimpulanKondisiKelemahanLingkungan.UrusanPemerintahan = &req.UrusanPemerintahan
		res.SimpulanKondisiKelemahanLingkungan.SumberData = &req.SumberData
		res.SimpulanKondisiKelemahanLingkungan.UraianKelamahan = uraianKelemahanJSON
		res.SimpulanKondisiKelemahanLingkungan.Klasifikasi = &req.Klasifikasi

		if _, err := updateSimpulanKondisiKelemahanLingkungan(ctx, gateway.SimpulanKondisiKelemahanLingkunganSaveReq{SimpulanKondisiKelemahanLingkungan: res.SimpulanKondisiKelemahanLingkungan}); err != nil {
			return nil, err
		}

		return &SimpulanKondisiKelemahanLingkunganUpdateUseCaseRes{}, nil
	}
}
