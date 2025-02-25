package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenilaianRisikoGetByIDReq struct {
	ID string
}

type PenilaianRisikoGetByIDRes struct {
	PenilaianRisiko model.PenilaianRisiko
}

type PenilaianRisikoGetByID = core.ActionHandler[PenilaianRisikoGetByIDReq, PenilaianRisikoGetByIDRes]

func ImplPenilaianRisikoGetByID(db *gorm.DB) PenilaianRisikoGetByID {
	return func(ctx context.Context, req PenilaianRisikoGetByIDReq) (*PenilaianRisikoGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var PenilaianRisiko model.PenilaianRisiko
		if err := query.First(&PenilaianRisiko, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("PenilaianRisiko id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &PenilaianRisikoGetByIDRes{PenilaianRisiko: PenilaianRisiko}, nil
	}
}
