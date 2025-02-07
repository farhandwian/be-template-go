package gateway

import (
	"context"
	"fmt"
	"perizinan/model"
	"shared/core"
	"shared/middleware"
	"strings"

	"gorm.io/gorm"
)

type LaporanPerizinanGetAllReq struct {
	Keyword               string
	PeriodePengambilanSda string
	Page                  int
	Size                  int
	SortBy                string
	SortOrder             string
}

type LaporanPerizinanGetAllRes struct {
	Items []model.LaporanPerizinan `json:"items"`
	Count int64                    `json:"count"`
}

type LaporanPerizinanGetAll = core.ActionHandler[LaporanPerizinanGetAllReq, LaporanPerizinanGetAllRes]

// Function to map Indonesian month names to English
func mapIndonesianMonthsToEnglish(column string) string {
	return `
        CASE
            WHEN ` + column + ` LIKE 'Januari %' THEN REPLACE(` + column + `, 'Januari', 'January')
            WHEN ` + column + ` LIKE 'Februari %' THEN REPLACE(` + column + `, 'Februari', 'February')
            WHEN ` + column + ` LIKE 'Maret %' THEN REPLACE(` + column + `, 'Maret', 'March')
            WHEN ` + column + ` LIKE 'April %' THEN REPLACE(` + column + `, 'April', 'April')
            WHEN ` + column + ` LIKE 'Mei %' THEN REPLACE(` + column + `, 'Mei', 'May')
            WHEN ` + column + ` LIKE 'Juni %' THEN REPLACE(` + column + `, 'Juni', 'June')
            WHEN ` + column + ` LIKE 'Juli %' THEN REPLACE(` + column + `, 'Juli', 'July')
            WHEN ` + column + ` LIKE 'Agustus %' THEN REPLACE(` + column + `, 'Agustus', 'August')
            WHEN ` + column + ` LIKE 'September %' THEN REPLACE(` + column + `, 'September', 'September')
            WHEN ` + column + ` LIKE 'Oktober %' THEN REPLACE(` + column + `, 'Oktober', 'October')
            WHEN ` + column + ` LIKE 'November %' THEN REPLACE(` + column + `, 'November', 'November')
            WHEN ` + column + ` LIKE 'Desember %' THEN REPLACE(` + column + `, 'Desember', 'December')
            ELSE ` + column + `
        END
    `
}

func ImplLaporanPerizinanGetAll(db *gorm.DB) LaporanPerizinanGetAll {
	return func(ctx context.Context, req LaporanPerizinanGetAllReq) (*LaporanPerizinanGetAllRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var filters []string
		var args []interface{}

		if req.Keyword != "" {
			filters = append(filters, "no_sk LIKE ?")
			args = append(args, "%"+req.Keyword+"%")
		}

		if req.PeriodePengambilanSda != "" {
			filters = append(filters, "periode_pengambilan_sda LIKE ?")
			args = append(args, "%"+req.PeriodePengambilanSda+"%")
		}

		query = query.Where(strings.Join(filters, " AND "), args...)

		var count int64
		if err := query.
			Model(&model.LaporanPerizinan{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		validSortColumns := map[string]bool{
			"created_at":              true,
			"updated_at":              true,
			"no_sk":                   true,
			"perusahaan":              true,
			"periode_pengambilan_sda": true,
		}
		sortBy := "created_at"
		if req.SortBy != "" && validSortColumns[req.SortBy] {
			if req.SortBy == "periode_pengambilan_sda" {
				// Use CASE statement for Indonesian month conversion
				sortBy = fmt.Sprintf("STR_TO_DATE(%s, '%%M %%Y')", mapIndonesianMonthsToEnglish("periode_pengambilan_sda"))
			} else {
				sortBy = req.SortBy
			}
		}
		sortOrder := "ASC"
		if strings.ToLower(req.SortOrder) == "desc" {
			sortOrder = "DESC"
		}

		query = query.Order(sortBy + " " + sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var laporanPerizinan []model.LaporanPerizinan
		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&laporanPerizinan).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &LaporanPerizinanGetAllRes{
			Count: count,
			Items: laporanPerizinan,
		}, nil
	}
}
