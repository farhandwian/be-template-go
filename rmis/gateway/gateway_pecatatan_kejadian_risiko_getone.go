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
		if err := query.First(&PencatatanKejadianRisiko, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("PencatatanKejadianRisiko id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &PencatatanKejadianRisikoGetByIDRes{PencatatanKejadianRisiko: PencatatanKejadianRisiko}, nil
	}
}
