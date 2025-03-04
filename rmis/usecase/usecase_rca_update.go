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
	ID                                 string   `json:"id"`
	PemilikRisiko                      string   `json:"pemilik_risiko"`
	TahunPenilaian                     string   `json:"tahun_penilaian"`
	Why                                []string `json:"why"`
	PenyebabRisikoID                   string   `json:"penyebab_risiko_id"`
	IdentifikasiRisikoStrategisPemdaId string   `json:"identifikasi_risiko_strategis_pemda_id"`
	KegiatanPengendalian               string   `json:"kegiatan_pengendalian"`
}

type RcaUpdateUseCaseRes struct{}

type RcaUpdateUseCase = core.ActionHandler[RcaUpdateUseCaseReq, RcaUpdateUseCaseRes]

func ImplRcaUpdateUseCase(
	getRcaById gateway.RcaGetByID,
	updateRca gateway.RcaSave,
	IdentifikasiRisikoStrategisPemdaGetByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
	PenyebabRisikoGetByID gateway.PenyebabRisikoGetByID,
) RcaUpdateUseCase {
	return func(ctx context.Context, req RcaUpdateUseCaseReq) (*RcaUpdateUseCaseRes, error) {

		res, err := getRcaById(ctx, gateway.RcaGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		rca := res.Rca

		rca.PemilikRisiko = &req.PemilikRisiko

		if req.TahunPenilaian != "" {
			year, err := extractYear(req.TahunPenilaian)
			if err != nil {
				return nil, fmt.Errorf("invalid TahunPenilaian format: %v", err)
			}
			rca.TahunPenilaian = &year
		}

		// Penyebab Risiko
		_, err = PenyebabRisikoGetByID(ctx, gateway.PenyebabRisikoGetByIDReq{ID: req.PenyebabRisikoID})
		if err != nil {
			return nil, fmt.Errorf("error getting penyebab risiko table: %v", err)
		}

		// Identifikasi Risiko Strategis Pemda
		_, err = IdentifikasiRisikoStrategisPemdaGetByID(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{ID: req.IdentifikasiRisikoStrategisPemdaId})
		if err != nil {
			return nil, fmt.Errorf("error getting identifikasi risiko strategis pemda table: %v", err)
		}

		rca.IdentifikasiRisikoStrategisPemdaID = &req.IdentifikasiRisikoStrategisPemdaId
		rca.PenyebabRisikoID = &req.PenyebabRisikoID
		rca.Why = helper.ToDataTypeJSONPtr(req.Why...)
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
