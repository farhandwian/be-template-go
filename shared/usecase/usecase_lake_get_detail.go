package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetLakeDetailReq struct {
	ID string
}

type GetLakeDetailResp struct {
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
	TypeID      int            `json:"type_id"`
	TypeName    string         `json:"type_name"`
	Type        *string        `json:"type"`
	WsID        string         `json:"ws_id"`
	WsName      string         `json:"ws_name"`
}

type GetLakeDetailUseCase = core.ActionHandler[GetLakeDetailReq, GetLakeDetailResp]

func ImplGetLakeDetail(getLakeDetail gateway.GetLakeDetailByIDGateway) GetLakeDetailUseCase {
	return func(ctx context.Context, req GetLakeDetailReq) (*GetLakeDetailResp, error) {
		data, err := getLakeDetail(ctx, gateway.GetLakeDetailByIDReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return &GetLakeDetailResp{
			ID:          data.Lake.ID,
			Geometry:    data.Lake.Geometry,
			DasID:       data.Lake.DasID,
			DasName:     data.Lake.DasName,
			IsPUPR:      data.Lake.IsPupr,
			Province:    data.Lake.Provinsi,
			City:        data.Lake.Kabupaten,
			District:    data.Lake.Kecamatan,
			SubDistrict: data.Lake.Kelurahan,
			Authority:   data.Lake.Kewenangan,
			Name:        data.Lake.Name,
			ManagerID:   data.Lake.PengelolaID,
			ManagerName: data.Lake.PengelolaName,
			RefID:       data.Lake.RefID,
			TypeID:      data.Lake.TypeID,
			TypeName:    data.Lake.TypeName,
			Type:        data.Lake.Jenis,
			WsID:        data.Lake.WsID,
			WsName:      data.Lake.WsName,
		}, err
	}
}
