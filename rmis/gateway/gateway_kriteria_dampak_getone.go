package gateway

import (
	"context"
	"fmt"
	"rmis/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type KriteriaDampakGetByIDReq struct {
	ID string
}

type KriteriaDampakGetByIDRes struct {
	KriteriaDampak model.KriteriaDampak
}

type KriteriaDampakGetByID = core.ActionHandler[KriteriaDampakGetByIDReq, KriteriaDampakGetByIDRes]

func ImplKriteriaDampakGetByID(db *gorm.DB) KriteriaDampakGetByID {
	return func(ctx context.Context, req KriteriaDampakGetByIDReq) (*KriteriaDampakGetByIDRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var KriteriaDampak model.KriteriaDampak
		if err := query.First(&KriteriaDampak, "id = ?", req.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("KriteriaDampak id %v is not found", req.ID)
			}
			return nil, core.NewInternalServerError(err)
		}

		return &KriteriaDampakGetByIDRes{KriteriaDampak: KriteriaDampak}, nil
	}
}
