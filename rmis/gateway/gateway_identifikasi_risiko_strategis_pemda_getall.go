package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/helper"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoStrategisPemdaGetAllReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
	Periode   string
}

type IdentifikasiRisikoStrategisPemdaGetAllRes struct {
	IdentifikasiRisikoStrategisPemda []model.IdentifikasiRisikoStrategisPemdaResponse `json:"identifikasi_risiko_strategis_pemda"`
	Count                            int64                                            `json:"count"`
}

type IdentifikasiRisikoStrategisPemdaGetAll = core.ActionHandler[IdentifikasiRisikoStrategisPemdaGetAllReq, IdentifikasiRisikoStrategisPemdaGetAllRes]

func ImplIdentifikasiRisikoStrategisPemdaGetAll(db *gorm.DB) IdentifikasiRisikoStrategisPemdaGetAll {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaGetAllReq) (*IdentifikasiRisikoStrategisPemdaGetAllRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		query = query.
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON identifikasi_risiko_strategis_pemdas.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id").
			Joins("LEFT JOIN kategori_risikos ON identifikasi_risiko_strategis_pemdas.kategori_risiko_id = kategori_risikos.id").
			Joins("LEFT JOIN rcas ON identifikasi_risiko_strategis_pemdas.rca_id = rcas.id")

		if req.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", req.Keyword)
			query = query.
				Where("penetapan_konteks_risiko_strategis_pemdas.nama_pemda LIKE ?", keyword).
				Or("penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian LIKE ?", keyword).
				Or("penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan LIKE ?", keyword).
				Or("penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintahan LIKE ?", keyword)
		}

		if req.Status != "" {
			query = query.Where("identifikasi_risiko_strategis_pemdas.status = ?", req.Status)
		}

		if req.Periode != "" {
			query = query.Where("penetapan_konteks_risiko_strategis_pemdas.periode = ?", req.Periode)
		}

		var count int64
		if err := query.
			Model(&model.IdentifikasiRisikoStrategisPemda{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}
		allowedSortBy := map[string]bool{
			"status": true,
		}

		allowerdForeignSortBy := map[string]string{
			"nama_pemda":        "penetapan_konteks_risiko_strategis_pemdas.nama_pemda",
			"tahun":             "penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian",
			"penetapan_konteks": "penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan",
			"urusan_pemerintah": "penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintahan",
		}

		sortBy, sortOrder, err := helper.ValidateSortParamsWithForeignKey(allowedSortBy, allowerdForeignSortBy, req.SortBy, req.SortOrder, "status")
		if err != nil {
			return nil, err
		}

		// Apply sorting
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)

		page, size := ValidatePageSize(req.Page, req.Size)

		var objs []model.IdentifikasiRisikoStrategisPemdaResponse

		if err := query.
			Select(`identifikasi_risiko_strategis_pemdas.*,
			    penetapan_konteks_risiko_strategis_pemdas.nama_pemda AS nama_pemda,
		        penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian AS tahun,
				penetapan_konteks_risiko_strategis_pemdas.periode AS periode,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS tujuan,
		        penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintahan AS urusan_pemerintah,
				penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS penetapan_konteks
			`).
			Offset((page - 1) * size).
			Order(orderClause).
			Limit(size).
			Scan(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}
		// 		if err := db.Raw(`
		//     SELECT identifikasi_risiko_strategis_pemdas.*,
		//         penetapan_konteks_risiko_strategis_pemdas.nama_pemda AS nama_pemda,
		//         penetapan_konteks_risiko_strategis_pemdas.tahun_penilaian AS tahun,
		//         penetapan_konteks_risiko_strategis_pemdas.periode AS periode,
		//         penetapan_konteks_risiko_strategis_pemdas.penetapan_tujuan AS penetapan_konteks,
		//         penetapan_konteks_risiko_strategis_pemdas.urusan_pemerintahan AS urusan_pemerintah
		//     FROM identifikasi_risiko_strategis_pemdas
		//     LEFT JOIN penetapan_konteks_risiko_strategis_pemdas
		//         ON identifikasi_risiko_strategis_pemdas.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id
		//     LIMIT ? OFFSET ?
		// `, size, (page-1)*size).Scan(&objs).Error; err != nil {
		// 			return nil, core.NewInternalServerError(err)
		// 		}

		return &IdentifikasiRisikoStrategisPemdaGetAllRes{
			IdentifikasiRisikoStrategisPemda: objs,
			Count:                            count,
		}, nil
	}
}
