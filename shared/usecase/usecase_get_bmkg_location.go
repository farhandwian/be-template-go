package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
)

type GetBMKGLocationReq struct {
	Location string
}

type GetBMKGLocationRes struct {
	Location []Location `json:"location"`
}

type Location struct {
	Adm4        string `json:"adm4"`
	Province    string `json:"province"`
	City        string `json:"City"`
	District    string `json:"district"`
	Subdistrict string `json:"sub_district"`
}

type GetBMKGLocationUseCase = core.ActionHandler[GetBMKGLocationReq, GetBMKGLocationRes]

func ImplGetBMKGLocationUseCase(getBMKGLocationGateway gateway.FetchBMKGWeatherLocationGateway) GetBMKGLocationUseCase {
	return func(ctx context.Context, req GetBMKGLocationReq) (*GetBMKGLocationRes, error) {
		res, err := getBMKGLocationGateway(ctx, gateway.FetchBMKGWeatherLocationReq{
			Location: req.Location,
		})
		if err != nil {
			return nil, err
		}
		var location []Location
		for _, l := range res.Location {
			location = append(location, Location{
				Adm4:        l.Adm4,
				Province:    l.Province,
				City:        l.City,
				District:    l.District,
				Subdistrict: l.Subdistrict,
			})
		}
		return &GetBMKGLocationRes{
			Location: location,
		}, nil
	}
}
