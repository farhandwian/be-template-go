package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type IndeksPeringkatPrioritasGetByIDReq struct {
	ID string
}

type IndeksPeringkatPrioritasGetByIDRes struct {
	IndeksPeringkatPrioritas model.IndeksPeringkatPrioritas
}

type IndeksPeringkatPrioritasGetByID = core.ActionHandler[IndeksPeringkatPrioritasGetByIDReq, IndeksPeringkatPrioritasGetByIDRes]

func ImplIndeksPeringkatPrioritasGetByID(db *gorm.DB) IndeksPeringkatPrioritasGetByID {
	return func(ctx context.Context, req IndeksPeringkatPrioritasGetByIDReq) (*IndeksPeringkatPrioritasGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var IndeksPeringkatPrioritas model.IndeksPeringkatPrioritas
		if err := query.First(&IndeksPeringkatPrioritas, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("IndeksPeringkatPrioritas id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &IndeksPeringkatPrioritasGetByIDRes{IndeksPeringkatPrioritas: IndeksPeringkatPrioritas}, nil
	}
}
