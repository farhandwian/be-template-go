package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IdentifikasiRisikoStrategisOPDGetByIDReq struct {
	ID string
}

type IdentifikasiRisikoStrategisOPDGetByIDRes struct {
	IdentifikasiRisikoStrategisOPD model.IdentifikasiRisikoStrategisOPD
}

type IdentifikasiRisikoStrategisOPDGetByID = core.ActionHandler[IdentifikasiRisikoStrategisOPDGetByIDReq, IdentifikasiRisikoStrategisOPDGetByIDRes]

func ImplIdentifikasiRisikoStrategisOPDGetByID(db *gorm.DB) IdentifikasiRisikoStrategisOPDGetByID {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDGetByIDReq) (*IdentifikasiRisikoStrategisOPDGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var IdentifikasiRisikoStrategisOPD model.IdentifikasiRisikoStrategisOPD
		if err := query.First(&IdentifikasiRisikoStrategisOPD, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("IdentifikasiRisikoStrategisOPD id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &IdentifikasiRisikoStrategisOPDGetByIDRes{IdentifikasiRisikoStrategisOPD: IdentifikasiRisikoStrategisOPD}, nil
	}
}
