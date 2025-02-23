package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"shared/core"
	"shared/helper"
	"strconv"
	"time"
)

type RcaUpdateUseCaseReq struct {
	ID                    string   `json:"id"`
	NamaUnitPemilikRisiko string   `json:"nama_unit_pemilik_risiko"`
	TahunPenilaian        string   `json:"tahun_penilaian"`
	PernyataanRisiko      string   `json:"pernyataan_risiko"`
	Why                   []string `json:"why"`
	JenisPenyebab         string   `json:"jenis_penyebab"`
	KegiatanPengendalian  string   `json:"kegiatan_pengendalian"`
}

type RcaUpdateUseCaseRes struct{}

type RcaUpdateUseCase = core.ActionHandler[RcaUpdateUseCaseReq, RcaUpdateUseCaseRes]

func ImplRcaUpdateUseCase(
	getRcaById gateway.RcaGetByID,
	updateRca gateway.RcaSave,
) RcaUpdateUseCase {
	return func(ctx context.Context, req RcaUpdateUseCaseReq) (*RcaUpdateUseCaseRes, error) {

		res, err := getRcaById(ctx, gateway.RcaGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		rca := res.Rca

		rca.NamaUnitPemilikRisiko = &req.NamaUnitPemilikRisiko

		if req.TahunPenilaian != "" {
			year, err := extractYear(req.TahunPenilaian)
			if err != nil {
				return nil, fmt.Errorf("invalid TahunPenilaian format: %v", err)
			}
			rca.TahunPenilaian = &year
		}

		rca.PenyebabRisiko = &req.PernyataanRisiko
		rca.Why = helper.ToDataTypeJSONPtr(req.Why...)
		rca.JenisPenyebab = &req.JenisPenyebab
		rca.KegiatanPengendalian = &req.KegiatanPengendalian
		rca.SetAkarPenyebab()

		if _, err := updateRca(ctx, gateway.RcaSaveReq{Rca: rca}); err != nil {
			return nil, err
		}

		return &RcaUpdateUseCaseRes{}, nil
	}
}

func extractYear(dateStr string) (time.Time, error) {
	// Try parsing different formats
	parsedDate, err := time.Parse("2006-01-02", dateStr) // "YYYY-MM-DD"
	if err != nil {
		parsedDate, err = time.Parse("2006-01", dateStr) // "YYYY-MM"
		if err != nil {
			parsedDate, err = time.Parse("2006", dateStr) // "YYYY"
			if err != nil {
				return time.Time{}, fmt.Errorf("unsupported date format: %s", dateStr)
			}
		}
	}

	// Convert year to time.Time (use January 1st as default)
	year := parsedDate.Year()
	yearTime, _ := time.Parse("2006", strconv.Itoa(year)) // Convert back to time.Time

	return yearTime, nil
}
