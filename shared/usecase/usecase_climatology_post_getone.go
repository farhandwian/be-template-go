package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"time"
)

type GetClimatologyDetailPostReq struct {
	ID string
}

type GetClimatologyDetailPostResp struct {
	ID         string  `json:"id"`
	ExternalID int64   `json:"external_id"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	PostVendor string  `json:"post_vendor"`
	Officer    string  `json:"officer"`
	Elevation  float64 `json:"elevation"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type GetClimatologyDetailPostUseCase = core.ActionHandler[GetClimatologyDetailPostReq, GetClimatologyDetailPostResp]

func ImplGetClimatologyDetailPostUseCase(climatologyDetail gateway.GetDetailClimatologyByIDGateway) GetClimatologyDetailPostUseCase {
	return func(ctx context.Context, req GetClimatologyDetailPostReq) (*GetClimatologyDetailPostResp, error) {
		data, err := climatologyDetail(ctx, gateway.GetClimatologyDetailByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &GetClimatologyDetailPostResp{
			ID:         data.Climatology.ID,
			ExternalID: data.Climatology.ExternalID,
			Name:       data.Climatology.Name,
			Type:       data.Climatology.Type,
			PostVendor: data.Climatology.PostVendor,
			Officer:    data.Climatology.Officer,
			Elevation:  data.Climatology.Elevation,
			Latitude:   data.Climatology.Latitude,
			Longitude:  data.Climatology.Longitude,
			CreatedAt:  data.Climatology.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  data.Climatology.UpdatedAt.Format(time.RFC3339),
		}, nil
	}
}
