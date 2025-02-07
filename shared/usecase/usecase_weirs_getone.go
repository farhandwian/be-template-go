package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetWeirDetailReq struct {
	ID string
}

type GetWeirDetailResp struct {
	ID                    string         `json:"id"`
	Geometry              datatypes.JSON `json:"geometry"`
	DasID                 string         `json:"das_id"`
	DasName               string         `json:"das_name"`
	DebitIntakeRainSeason *float64       `json:"debit_intake_rain_season"`
	DebitIntakeDrySeason  *float64       `json:"debit_intake_dry_season"`
	IsPUPR                bool           `json:"is_pupr"`
	WeirType              *string        `json:"weir_type"`
	Province              string         `json:"province"`
	City                  string         `json:"city"`
	District              string         `json:"district"`
	SubDistrict           string         `json:"sub_district"`
	Authority             string         `json:"authority"`
	Name                  string         `json:"name"`
	ManagerID             int            `json:"manager_id"`
	ManagerName           string         `json:"manager_name"`
	RefID                 *string        `json:"ref_id"`
	WeirHeight            *float64       `json:"weir_height"`
	WSID                  string         `json:"ws_id"`
	WSName                string         `json:"ws_name"`
	TypeID                int            `json:"type_id"`
	TypeName              string         `json:"type_name"`
}

type GetWeirDetailUseCase = core.ActionHandler[GetWeirDetailReq, GetWeirDetailResp]

func ImplGetWeirDetail(getWeirDetail gateway.WeirGetDetailGateway) GetWeirDetailUseCase {
	return func(ctx context.Context, req GetWeirDetailReq) (*GetWeirDetailResp, error) {
		data, err := getWeirDetail(ctx, gateway.WeirGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return &GetWeirDetailResp{
			ID:                    data.Weir.ID,
			Geometry:              data.Weir.Geometry,
			DasID:                 data.Weir.DASID,
			DasName:               data.Weir.DASName,
			DebitIntakeRainSeason: data.Weir.DebitIntakeMusimHujan,
			DebitIntakeDrySeason:  data.Weir.DebitIntakeMusimKemarau,
			IsPUPR:                data.Weir.IsPUPR,
			WeirType:              data.Weir.JenisBendung,
			Province:              data.Weir.Provinsi,
			City:                  data.Weir.Kabupaten,
			District:              data.Weir.Kecamatan,
			SubDistrict:           data.Weir.Kelurahan,
			Authority:             data.Weir.Kewenangan,
			Name:                  data.Weir.Name,
			ManagerID:             data.Weir.PengelolaID,
			ManagerName:           data.Weir.PengelolaName,
			RefID:                 data.Weir.RefID,
			WeirHeight:            data.Weir.TinggiBendung,
			WSID:                  data.Weir.WSID,
			WSName:                data.Weir.WSName,
			TypeID:                data.Weir.TypeID,
			TypeName:              data.Weir.TypeName,
		}, err
	}
}
