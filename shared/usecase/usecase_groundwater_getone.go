package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetGroundwaterDetailReq struct {
	ID string
}

type GetGroundwaterDetailResp struct {
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

type GetGroundwaterDetailUseCase = core.ActionHandler[GetGroundwaterDetailReq, GetGroundwaterDetailResp]

func ImplGetGroundwaterDetail(getGroundwaterDetail gateway.GroundWaterGetDetailGateway) GetGroundwaterDetailUseCase {
	return func(ctx context.Context, req GetGroundwaterDetailReq) (*GetGroundwaterDetailResp, error) {
		data, err := getGroundwaterDetail(ctx, gateway.GroundWaterGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return &GetGroundwaterDetailResp{
			ID:          data.GroundWater.ID,
			Geometry:    data.GroundWater.Geometry,
			DasID:       data.GroundWater.DASID,
			DasName:     data.GroundWater.DASName,
			IsPUPR:      data.GroundWater.IsPUPR,
			Province:    data.GroundWater.Provinsi,
			City:        data.GroundWater.Kabupaten,
			District:    data.GroundWater.Kecamatan,
			SubDistrict: data.GroundWater.Kelurahan,
			Authority:   data.GroundWater.Kewenangan,
			Name:        data.GroundWater.Name,
			ManagerID:   data.GroundWater.PengelolaID,
			ManagerName: data.GroundWater.PengelolaName,
			RefID:       data.GroundWater.RefID,
			WSID:        data.GroundWater.WSID,
			WSName:      data.GroundWater.WSName,
			TypeID:      data.GroundWater.TypeID,
			TypeName:    data.GroundWater.TypeName,
		}, err
	}
}
