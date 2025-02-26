package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoOperasionalOPDGetByIDReq struct {
	ID string
}

type IdentifikasiRisikoOperasionalOPDGetByIDRes struct {
	IdentifikasiRisikoOperasionalOPD model.IdentifikasiRisikoOperasionalOPD
}

type IdentifikasiRisikoOperasionalOPDGetByID = core.ActionHandler[IdentifikasiRisikoOperasionalOPDGetByIDReq, IdentifikasiRisikoOperasionalOPDGetByIDRes]

func ImplIdentifikasiRisikoOperasionalOPDGetByID(db *gorm.DB) IdentifikasiRisikoOperasionalOPDGetByID {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDGetByIDReq) (*IdentifikasiRisikoOperasionalOPDGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var IdentifikasiRisikoOperasionalOPD model.IdentifikasiRisikoOperasionalOPD
		if err := query.First(&IdentifikasiRisikoOperasionalOPD, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("IdentifikasiRisikoOperasionalOPD id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoOperasionalOPDGetByIDRes{IdentifikasiRisikoOperasionalOPD: IdentifikasiRisikoOperasionalOPD}, nil
	}
}
