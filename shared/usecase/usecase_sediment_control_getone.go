package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetSedimentControlDetailReq struct {
	ID string
}

type GetSedimentControlDetailResp struct {
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

type GetSedimentControlDetailUseCase = core.ActionHandler[GetSedimentControlDetailReq, GetSedimentControlDetailResp]

func ImplGetSedimentControlDetail(getSedimentControlDetail gateway.SedimentControlGetDetailGateway) GetSedimentControlDetailUseCase {
	return func(ctx context.Context, req GetSedimentControlDetailReq) (*GetSedimentControlDetailResp, error) {
		data, err := getSedimentControlDetail(ctx, gateway.SedimentControlGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return &GetSedimentControlDetailResp{
			ID:          data.SedimentControl.ID,
			Geometry:    data.SedimentControl.Geometry,
			DasID:       data.SedimentControl.DASID,
			DasName:     data.SedimentControl.DASName,
			IsPUPR:      data.SedimentControl.IsPUPR,
			Province:    data.SedimentControl.Provinsi,
			City:        data.SedimentControl.Kabupaten,
			District:    data.SedimentControl.Kecamatan,
			SubDistrict: data.SedimentControl.Kelurahan,
			Authority:   data.SedimentControl.Kewenangan,
			Name:        data.SedimentControl.Name,
			ManagerID:   data.SedimentControl.PengelolaID,
			ManagerName: data.SedimentControl.PengelolaName,
			RefID:       data.SedimentControl.RefID,
			WSID:        data.SedimentControl.WSID,
			WSName:      data.SedimentControl.WSName,
			TypeID:      data.SedimentControl.TypeID,
			TypeName:    data.SedimentControl.TypeName,
		}, err
	}
}
