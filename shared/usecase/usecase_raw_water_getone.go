package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetRawWaterDetailReq struct {
	ID string
}

type GetRawWaterDetailResp struct {
	ID          string         `json:"id"`
	Geometry    datatypes.JSON `json:"geometry"`
	DasID       string         `json:"das_id"`
	DasName     string         `json:"das_name"`
	IsPUPR      bool           `json:"is_pupr"`
	Province    string         `json:"province"`
	City        string         `json:"city"`
	District    string         `json:"district"`
	SubDistrict string         `json:"sub_district"`
	Authority   string         `json:"authority"`
	Name        string         `json:"name"`
	ManagerID   int            `json:"manager_id"`
	ManagerName string         `json:"manager_name"`
	RefID       *string        `json:"ref_id"`
	WSID        string         `json:"ws_id"`
	WSName      string         `json:"ws_name"`
	TypeID      int            `json:"type_id"`
	TypeName    string         `json:"type_name"`
}

type GetRawWaterDetailUseCase = core.ActionHandler[GetRawWaterDetailReq, GetRawWaterDetailResp]

func ImplGetRawWaterDetail(getRawWaterDetail gateway.RawWaterGetDetailGateway) GetRawWaterDetailUseCase {
	return func(ctx context.Context, req GetRawWaterDetailReq) (*GetRawWaterDetailResp, error) {
		data, err := getRawWaterDetail(ctx, gateway.RawWaterGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return &GetRawWaterDetailResp{
			ID:          data.RawWater.ID,
			Geometry:    data.RawWater.Geometry,
			DasID:       data.RawWater.DASID,
			DasName:     data.RawWater.DASName,
			IsPUPR:      data.RawWater.IsPUPR,
			Province:    data.RawWater.Provinsi,
			City:        data.RawWater.Kabupaten,
			District:    data.RawWater.Kecamatan,
			SubDistrict: data.RawWater.Kelurahan,
			Authority:   data.RawWater.Kewenangan,
			Name:        data.RawWater.Name,
			ManagerID:   data.RawWater.PengelolaID,
			ManagerName: data.RawWater.PengelolaName,
			RefID:       data.RawWater.RefID,
			WSID:        data.RawWater.WSID,
			WSName:      data.RawWater.WSName,
			TypeID:      data.RawWater.TypeID,
			TypeName:    data.RawWater.TypeName,
		}, err
	}
}
