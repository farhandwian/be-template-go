package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type PenyebabRisikoGetByIDReq struct {
	ID string
}

type PenyebabRisikoGetByIDRes struct {
	PenyebabRisiko model.PenyebabRisiko
}

type PenyebabRisikoGetByID = core.ActionHandler[PenyebabRisikoGetByIDReq, PenyebabRisikoGetByIDRes]

func ImplPenyebabRisikoGetByID(db *gorm.DB) PenyebabRisikoGetByID {
	return func(ctx context.Context, req PenyebabRisikoGetByIDReq) (*PenyebabRisikoGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var PenyebabRisiko model.PenyebabRisiko
		if err := query.First(&PenyebabRisiko, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("PenyebabRisiko id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &PenyebabRisikoGetByIDRes{PenyebabRisiko: PenyebabRisiko}, nil
	}
}
