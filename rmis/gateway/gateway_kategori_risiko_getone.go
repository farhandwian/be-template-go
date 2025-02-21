package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type KategoriRisikoGetByIDReq struct {
	ID string
}

type KategoriRisikoGetByIDRes struct {
	KategoriRisiko model.KategoriRisiko
}

type KategoriRisikoGetByID = core.ActionHandler[KategoriRisikoGetByIDReq, KategoriRisikoGetByIDRes]

func ImplKategoriRisikoGetByID(db *gorm.DB) KategoriRisikoGetByID {
	return func(ctx context.Context, req KategoriRisikoGetByIDReq) (*KategoriRisikoGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var kategoriRisiko model.KategoriRisiko
		if err := query.First(&kategoriRisiko, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("KategoriRisiko id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &KategoriRisikoGetByIDRes{KategoriRisiko: kategoriRisiko}, nil
	}
}
