package gateway

import (
	"context"
	"shared/core"

	"gorm.io/gorm"
)

type GetLaporanPerizinanStatusCountReq struct{}

type GetLaporanPerizinanStatusCountRes struct {
	SubmittedCount    int
	NotSubmittedCount int
}

type GetLaporanPerizinanStatusCountGateway = core.ActionHandler[GetLaporanPerizinanStatusCountReq, GetLaporanPerizinanStatusCountRes]

func ImplGetLaporanPerizinanStatusCount(db *gorm.DB) GetLaporanPerizinanStatusCountGateway {
	return func(ctx context.Context, request GetLaporanPerizinanStatusCountReq) (*GetLaporanPerizinanStatusCountRes, error) {
		var submittedCount int64
		var notSubmittedCount int64

		var prevMonth string
		err := db.Raw("SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 1 MONTH), '%Y-%m')").Scan(&prevMonth).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		err = db.Table("laporan_perizinans").
			Where("status = ?", "submitted").
			Where("LEFT(periode_pengambilan_sda, 7) = ?", prevMonth).
			Count(&submittedCount).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		err = db.Table("sk_perizinans").
			Joins("LEFT JOIN laporan_perizinans ON laporan_perizinans.no_sk = sk_perizinans.no_sk AND laporan_perizinans.status = 'submitted' AND LEFT(laporan_perizinans.periode_pengambilan_sda, 7) = ?", prevMonth).
			Where("sk_perizinans.status = ?", "Berlaku").
			Where("laporan_perizinans.id IS NULL").
			Count(&notSubmittedCount).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetLaporanPerizinanStatusCountRes{
			SubmittedCount:    int(submittedCount),
			NotSubmittedCount: int(notSubmittedCount),
		}, nil
	}
}
