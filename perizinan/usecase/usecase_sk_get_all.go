package usecase

import (
	"context"
	"perizinan/gateway"
	"perizinan/model"
	"shared/core"
	"sort"
	"time"
)

type SKGetAllUseCaseReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
}

type SKGetAllUseCaseRes struct {
	Items    []model.SKPerizinan `json:"items"`
	Metadata *Metadata           `json:"metadata"`
}

type SKGetAllUseCase = core.ActionHandler[SKGetAllUseCaseReq, SKGetAllUseCaseRes]

func ImplSKGetAllUseCase(
	getAll gateway.SKPerizinanGetAll,
) SKGetAllUseCase {
	return func(ctx context.Context, req SKGetAllUseCaseReq) (*SKGetAllUseCaseRes, error) {

		res, err := getAll(ctx, gateway.SKPerizinanGetAllReq{Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		// sorting
		if req.SortBy != "" {
			sortAllSk(res.Items, req)
		}
		SKPerizinan := paginateFeatures(res.Items, req.Page, req.Size)

		// Calculate pagination details
		totalItems := len(res.Items)
		totalPages := (totalItems + req.Page - 1) / req.Size

		return &SKGetAllUseCaseRes{
			Items: SKPerizinan,
			Metadata: &Metadata{
				Pagination: Pagination{
					Page:       req.Page,
					Limit:      req.Size,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}

func sortAllSk(items []model.SKPerizinan, request SKGetAllUseCaseReq) {
	sort.Slice(items, func(i, j int) bool {
		switch request.SortBy {
		case "status":
			if request.SortOrder == "asc" {
				return items[i].Status < items[j].Status
			}
			return items[i].Status > items[j].Status
		case "no_sk":
			if request.SortOrder == "asc" {
				return items[i].NoSK < items[j].NoSK
			}
			return items[i].NoSK > items[j].NoSK
		case "perusahaan_pemohon":
			if request.SortOrder == "asc" {
				return items[i].PerusahaanPemohon < items[j].PerusahaanPemohon
			}
			return items[i].PerusahaanPemohon > items[j].PerusahaanPemohon
		case "masa_berlaku_sk":
			return compareMasaBerlaku(items[i].MasaBerlakuSK, items[j].MasaBerlakuSK, request.SortOrder)
		case "kuota_air_dalam_sk":
			if request.SortOrder == "asc" {
				return items[i].KuotaAirDalamSK < items[j].KuotaAirDalamSK
			}
			return items[i].KuotaAirDalamSK > items[j].KuotaAirDalamSK
		}
		return false
	})
}

func compareMasaBerlaku(sk1, sk2, sortOrder string) bool {
	const (
		TidakBerlaku  = "Tidak Berlaku"
		SepanjangUmur = "Sepanjang umur layan konstruksi"
	)

	// Empty values always at the top
	if sk1 == "" && sk2 == "" {
		return false
	}
	if sk1 == "" {
		return true
	}
	if sk2 == "" {
		return false
	}

	// Handle "Tidak Berlaku"
	if sk1 == TidakBerlaku && sk2 == TidakBerlaku {
		return false
	}
	if sk1 == TidakBerlaku {
		return true
	}
	if sk2 == TidakBerlaku {
		return false
	}

	// Handle "Sepanjang umur layan konstruksi"
	if sk1 == SepanjangUmur && sk2 == SepanjangUmur {
		return false
	}
	if sk1 == SepanjangUmur {
		return false
	}
	if sk2 == SepanjangUmur {
		return true
	}

	// Parse as dates
	date1, isDate1 := parseDate(sk1)
	date2, isDate2 := parseDate(sk2)

	if isDate1 && isDate2 {
		if sortOrder == "asc" {
			return date2.Before(date1)
		}
		return date2.After(date1)
	}
	if isDate1 {
		return true
	}
	if isDate2 {
		return false
	}

	if sortOrder == "asc" {
		return sk1 < sk2
	}
	return sk1 > sk2
}

func parseDate(input string) (time.Time, bool) {
	layout := "02-Jan-06"
	date, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, false
	}
	return date, true
}

// Pagination function
func paginateFeatures(features []model.SKPerizinan, page, pageSize int) []model.SKPerizinan {
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(features) {
		start = len(features) // Ensure start does not exceed number of features
	}
	if end > len(features) {
		end = len(features) // Ensure end does not exceed number of features
	}

	return features[start:end]
}
