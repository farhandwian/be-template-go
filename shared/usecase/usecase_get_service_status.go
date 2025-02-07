package usecase

import (
	"context"
	"shared/core"
	sharedGateway "shared/gateway"
)

type GetServiceStatusReq struct {
}

type GetServiceStatusResp struct {
	Services []ServiceItem `json:"services"`
}

type ServiceItem struct {
	Name     string `json:"name"`
	IsOnline bool   `json:"is_online"`
}

type GetServiceStatusUseCase = core.ActionHandler[GetServiceStatusReq, GetServiceStatusResp]

func ImplGetServiceStatusUseCase(prometheusGateway sharedGateway.GetServiceStatusGateway) GetServiceStatusUseCase {
	return func(ctx context.Context, req GetServiceStatusReq) (*GetServiceStatusResp, error) {
		res, err := prometheusGateway(ctx, sharedGateway.GetServiceStatusReq{})
		if err != nil {
			return nil, err
		}
		services := []ServiceItem{
			{
				Name:     "Sistem Deteksi Sampah",
				IsOnline: true,
			},
			{
				Name:     "Sistem Deteksi Keamanan Pintu Air",
				IsOnline: true,
			},
			{
				Name:     "Sistem Monitoring & Peringatan",
				IsOnline: true,
			},
			{
				Name:     "Data Warehouse",
				IsOnline: res.Service.AdapterHydrology && res.Service.AdapterManganti,
			}, {
				Name:     "Adam & Hawa",
				IsOnline: res.Service.AdamHawa,
			},
			{
				Name:     "Autonomous Drone Patrolling System",
				IsOnline: false,
			},
			{
				Name:     "Dashboard Manganti",
				IsOnline: res.Service.Dashboard,
			},
			{
				Name:     "Sihka",
				IsOnline: res.Service.Sihka,
			},
			{
				Name:     "Si JagaCai",
				IsOnline: res.Service.SiJagaCai,
			},
		}
		return &GetServiceStatusResp{
			Services: services,
		}, nil

	}
}
