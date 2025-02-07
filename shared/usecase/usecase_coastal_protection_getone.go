package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetCoastalProtectionDetailReq struct {
	ID string
}

type GetCoastalProtectionDetailResp struct {
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

type GetCoastalProtectionDetailUseCase = core.ActionHandler[GetCoastalProtectionDetailReq, GetCoastalProtectionDetailResp]

func ImplGetCoastalProtectionDetail(getCoastalProtectionDetail gateway.CoastalProtectionGetDetailGateway) GetCoastalProtectionDetailUseCase {
	return func(ctx context.Context, req GetCoastalProtectionDetailReq) (*GetCoastalProtectionDetailResp, error) {
		data, err := getCoastalProtectionDetail(ctx, gateway.CoastalProtectionGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return &GetCoastalProtectionDetailResp{
			ID:          data.CoastalProtection.ID,
			Geometry:    data.CoastalProtection.Geometry,
			DasID:       data.CoastalProtection.DASID,
			DasName:     data.CoastalProtection.DASName,
			IsPUPR:      data.CoastalProtection.IsPUPR,
			Province:    data.CoastalProtection.Provinsi,
			City:        data.CoastalProtection.Kabupaten,
			District:    data.CoastalProtection.Kecamatan,
			SubDistrict: data.CoastalProtection.Kelurahan,
			Authority:   data.CoastalProtection.Kewenangan,
			Name:        data.CoastalProtection.Name,
			ManagerID:   data.CoastalProtection.PengelolaID,
			ManagerName: data.CoastalProtection.PengelolaName,
			RefID:       data.CoastalProtection.RefID,
			WSID:        data.CoastalProtection.WSID,
			WSName:      data.CoastalProtection.WSName,
			TypeID:      data.CoastalProtection.TypeID,
			TypeName:    data.CoastalProtection.TypeName,
		}, err
	}
}
