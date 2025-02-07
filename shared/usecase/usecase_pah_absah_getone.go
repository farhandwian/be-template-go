package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetPahAbsahDetailReq struct {
	ID string
}

type GetPahAbsahDetailResp struct {
	ID          string         `json:"id"`
	Geometry    datatypes.JSON `json:"geometry"`
	DasID       string         `json:"das_id"`
	DasName     string         `json:"das_name"`
	IsPUPR      bool           `json:"is_pupr"`
	Province    string         `json:"province"`
	City        string         `json:"city"`
	District    string         `json:"district"`
	Authority   string         `json:"authority"`
	SubDistrict string         `json:"sub_district"`
	Name        string         `json:"name"`
	ManagerID   int            `json:"manager_id"`
	ManagerName string         `json:"manager_name"`
	WSID        string         `json:"ws_id"`
	WSName      string         `json:"ws_name"`
	RefID       string         `json:"ref_id"`
	TypeID      int            `json:"type_id"`
	TypeName    string         `json:"type_name"`
}

type GetPahAbsahDetailUseCase = core.ActionHandler[GetPahAbsahDetailReq, GetPahAbsahDetailResp]

func ImplGetPahAbsahDetail(getPahAbsahDetail gateway.PahAbsahGetDetailGateway) GetPahAbsahDetailUseCase {
	return func(ctx context.Context, req GetPahAbsahDetailReq) (*GetPahAbsahDetailResp, error) {
		data, err := getPahAbsahDetail(ctx, gateway.PahAbsahGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return &GetPahAbsahDetailResp{
			ID:          data.PahAbsah.ID,
			Geometry:    data.PahAbsah.Geometry,
			DasID:       data.PahAbsah.DASID,
			DasName:     data.PahAbsah.DASName,
			IsPUPR:      data.PahAbsah.IsPUPR,
			Province:    data.PahAbsah.Provinsi,
			City:        data.PahAbsah.Kabupaten,
			District:    data.PahAbsah.Kecamatan,
			SubDistrict: data.PahAbsah.Kelurahan,
			Name:        data.PahAbsah.Name,
			ManagerID:   data.PahAbsah.PengelolaID,
			ManagerName: data.PahAbsah.PengelolaName,
			WSID:        data.PahAbsah.WSID,
			WSName:      data.PahAbsah.WSName,
			RefID:       data.PahAbsah.RefID,
			TypeID:      data.PahAbsah.TypeID,
			TypeName:    data.PahAbsah.TypeName,
			Authority:   data.PahAbsah.Kewenangan,
		}, err
	}
}
