package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenetapanKonteksRisikoOperasionalGetByIDReq struct {
	ID string
}

type PenetapanKonteksRisikoOperasionalGetByIDRes struct {
	PenetapanKonteksRisikoOperasional model.PenetapanKonteksRisikoOperasional
}

type PenetapanKonteksRisikoOperasionalGetByID = core.ActionHandler[PenetapanKonteksRisikoOperasionalGetByIDReq, PenetapanKonteksRisikoOperasionalGetByIDRes]

func ImplPenetapanKonteksRisikoOperasionalGetByID(db *gorm.DB) PenetapanKonteksRisikoOperasionalGetByID {
	return func(ctx context.Context, req PenetapanKonteksRisikoOperasionalGetByIDReq) (*PenetapanKonteksRisikoOperasionalGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var PenetapanKonteksRisikoOperasional model.PenetapanKonteksRisikoOperasional
		if err := query.First(&PenetapanKonteksRisikoOperasional, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("PenetapanKonteksRisikoOperasional id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &PenetapanKonteksRisikoOperasionalGetByIDRes{PenetapanKonteksRisikoOperasional: PenetapanKonteksRisikoOperasional}, nil
	}
}
