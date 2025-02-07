package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetIntakeDetailReq struct {
	ID string
}

type GetIntakeDetailResp struct {
	ID           string         `json:"id"`
	Geometry     datatypes.JSON `json:"geometry"`
	DasID        string         `json:"das_id"`
	DasName      string         `json:"das_name"`
	IsPUPR       bool           `json:"is_pupr"`
	IntakeObject string         `json:"intake_object"`
	Province     string         `json:"province"`
	City         string         `json:"city"`
	District     string         `json:"district"`
	SubDistrict  string         `json:"sub_district"`
	Authority    string         `json:"authority"`
	Name         string         `json:"name"`
	ManagerID    int            `json:"manager_id"`
	ManagerName  string         `json:"manager_name"`
	RefID        string         `json:"ref_id"`
	WSID         string         `json:"ws_id"`
	WSName       string         `json:"ws_name"`
	TypeID       int            `json:"type_id"`
	TypeName     string         `json:"type_name"`
}

type GetIntakeDetailUseCase = core.ActionHandler[GetIntakeDetailReq, GetIntakeDetailResp]

func ImplGetIntakeDetail(getIntakeDetail gateway.IntakeGetDetailGateway) GetIntakeDetailUseCase {
	return func(ctx context.Context, req GetIntakeDetailReq) (*GetIntakeDetailResp, error) {
		data, err := getIntakeDetail(ctx, gateway.IntakeGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return &GetIntakeDetailResp{
			ID:           data.Intake.ID,
			Geometry:     data.Intake.Geometry,
			DasID:        data.Intake.DASID,
			DasName:      data.Intake.DASName,
			IsPUPR:       data.Intake.IsPUPR,
			Province:     data.Intake.Provinsi,
			City:         data.Intake.Kabupaten,
			District:     data.Intake.Kecamatan,
			SubDistrict:  data.Intake.Kelurahan,
			Authority:    data.Intake.Kewenangan,
			Name:         data.Intake.Name,
			ManagerID:    data.Intake.PengelolaID,
			ManagerName:  data.Intake.PengelolaName,
			RefID:        data.Intake.RefID,
			WSID:         data.Intake.WSID,
			WSName:       data.Intake.WSName,
			TypeID:       data.Intake.TypeID,
			TypeName:     data.Intake.TypeName,
			IntakeObject: data.Intake.IntakeObject,
		}, err
	}
}
