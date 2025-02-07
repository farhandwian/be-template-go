package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetWellDetailReq struct {
	ID string
}

type GetWellDetailResp struct {
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
	RefID       string         `json:"ref_id"`
	WSID        string         `json:"ws_id"`
	WSName      string         `json:"ws_name"`
	TypeID      int            `json:"type_id"`
	TypeName    string         `json:"type_name"`
}

type GetWellDetailUseCase = core.ActionHandler[GetWellDetailReq, GetWellDetailResp]

func ImplGetWellDetail(getWellDetail gateway.WellGetDetailGateway) GetWellDetailUseCase {
	return func(ctx context.Context, req GetWellDetailReq) (*GetWellDetailResp, error) {
		data, err := getWellDetail(ctx, gateway.WellGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return &GetWellDetailResp{
			ID:          data.Well.ID,
			Geometry:    data.Well.Geometry,
			DasID:       data.Well.DASID,
			DasName:     data.Well.DASName,
			IsPUPR:      data.Well.IsPUPR,
			Province:    data.Well.Provinsi,
			City:        data.Well.Kabupaten,
			District:    data.Well.Kecamatan,
			SubDistrict: data.Well.Kelurahan,
			Authority:   data.Well.Kewenangan,
			Name:        data.Well.Name,
			ManagerID:   data.Well.PengelolaID,
			ManagerName: data.Well.PengelolaName,
			RefID:       data.Well.RefID,
			WSID:        data.Well.WSID,
			WSName:      data.Well.WSName,
			TypeID:      data.Well.TypeID,
			TypeName:    data.Well.TypeName,
		}, err
	}
}
