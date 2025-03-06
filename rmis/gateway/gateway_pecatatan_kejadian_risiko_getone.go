package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PencatatanKejadianRisikoGetByIDReq struct {
	ID string
}

type PencatatanKejadianRisikoGetByIDRes struct {
	PencatatanKejadianRisiko model.PencatatanKejadianRisiko
}

type PencatatanKejadianRisikoGetByID = core.ActionHandler[PencatatanKejadianRisikoGetByIDReq, PencatatanKejadianRisikoGetByIDRes]

func ImplPencatatanKejadianRisikoGetByID(db *gorm.DB) PencatatanKejadianRisikoGetByID {
	return func(ctx context.Context, req PencatatanKejadianRisikoGetByIDReq) (*PencatatanKejadianRisikoGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var PencatatanKejadianRisiko model.PencatatanKejadianRisiko
		if err := query.
			Joins("LEFT JOIN identifikasi_risiko_strategis_pemdas ON pencatatan_kejadian_risikos.identifikasi_risiko_strategis_pemda_id = identifikasi_risiko_strategis_pemdas.id").
			Joins("LEFT JOIN penetapan_konteks_risiko_strategis_pemdas ON pencatatan_kejadian_risikos.penetapan_konteks_risiko_strategis_pemda_id = penetapan_konteks_risiko_strategis_pemdas.id").
			Where("pencatatan_kejadian_risikos.id =?", req.ID).
			First(&PencatatanKejadianRisiko).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("PencatatanKejadianRisiko id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &PencatatanKejadianRisikoGetByIDRes{PencatatanKejadianRisiko: PencatatanKejadianRisiko}, nil
	}
}
