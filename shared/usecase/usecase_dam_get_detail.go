package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetDamDetailReq struct {
	ID string
}

type GetDamDetailResp struct {
	ID                              string         `json:"id"`
	Geometry                        datatypes.JSON `json:"geometry"`
	DasID                           string         `json:"das_id"`
	DasName                         string         `json:"das_name"`
	DMI                             *float64       `json:"dmi"`
	DamPeakElevation                *float64       `json:"dam_peak_elevation"`
	Irrigation                      *float64       `json:"irrigation"`
	IsPUPR                          bool           `json:"is_pupr"`
	City                            string         `json:"city"`
	District                        string         `json:"district"`
	SubDistrict                     string         `json:"sub_district"`
	DamTypeDescription              *string        `json:"dam_type_description"`
	Authority                       string         `json:"authority"`
	BuildingCondition               *string        `json:"building_condition"`
	DamWidth                        *float64       `json:"dam_width"`
	PeakDamWidth                    *float64       `json:"peak_dam_width"`
	MinInundationArea               *float64       `json:"min_inundation_area"`
	NormalInundationArea            *float64       `json:"normal_inundation_area"`
	TotalInundationArea             *float64       `json:"total_inundation_area"`
	Name                            string         `json:"name"`
	DamLength                       *float64       `json:"dam_length"`
	PeakDamLength                   *float64       `json:"peak_dam_length"`
	SpillwayWithGate                *string        `json:"spillway_with_gate"`
	SpillwayCrestElevation          *float64       `json:"spillway_crest_elevation"`
	SpillwayWidth                   *float64       `json:"spillway_width"`
	SpillwayCrestWidth              *float64       `json:"spillway_crest_width"`
	SpillwayLength                  *float64       `json:"spillway_length"`
	SpillwayTransitionChannelLength *float64       `json:"spillway_transition_channel_length"`
	SpillwayChannel                 *string        `json:"spillway_channel"`
	SpillwayType                    *string        `json:"spillway_type"`
	SpillwayTypeDescription         *string        `json:"spillway_type_description"`
	ManagerID                       int            `json:"manager_id"`
	ManagerName                     string         `json:"manager_name"`
	FloodReductionVolume            *float64       `json:"flood_reduction_volume"`
	FloodReductionInundationArea    *float64       `json:"flood_reduction_inundation_area"`
	HydropowerPlant                 *float64       `json:"hydropower_plant"`
	Province                        string         `json:"province"`
	ReferenceID                     string         `json:"reference_id"`
	InfrastructureStatus            *string        `json:"infrastructure_status"`
	ConstructionYear                *string        `json:"construction_year"`
	DamHeight                       *float64       `json:"dam_height"`
	ExcavationBaseHeight            *float64       `json:"excavation_base_height"`
	RiverBaseHeight                 *float64       `json:"river_base_height"`
	DamType                         *string        `json:"dam_type"`
	TypeID                          int            `json:"type_id"`
	TypeName                        string         `json:"type_name"`
	DamVolume                       *float64       `json:"dam_volume"`
	MinReservoirVolume              *float64       `json:"min_reservoir_volume"`
	NormalReservoirVolume           *float64       `json:"normal_reservoir_volume"`
	TotalReservoirVolume            *float64       `json:"total_reservoir_volume"`
	WsID                            string         `json:"ws_id"`
	WsName                          string         `json:"ws_name"`
}

type GetDamDetailUseCase = core.ActionHandler[GetDamDetailReq, GetDamDetailResp]

func ImplGetDamDetail(getDamDetail gateway.GetDetailDamByIDGateway) GetDamDetailUseCase {
	return func(ctx context.Context, request GetDamDetailReq) (*GetDamDetailResp, error) {
		data, err := getDamDetail(ctx, gateway.GetDamDetailByIDReq{ID: request.ID})
		if err != nil {
			return nil, err
		}

		return &GetDamDetailResp{
			ID:                              data.Dam.ID,
			Geometry:                        data.Dam.Geometry,
			DasID:                           data.Dam.DasID,
			DasName:                         data.Dam.DasName,
			DMI:                             data.Dam.Dmi,
			DamPeakElevation:                data.Dam.ElevasiPuncakBendungan,
			Irrigation:                      data.Dam.Irigasi,
			IsPUPR:                          data.Dam.IsPupr,
			City:                            data.Dam.Kabupaten,
			District:                        data.Dam.Kecamatan,
			SubDistrict:                     data.Dam.Kelurahan,
			DamTypeDescription:              data.Dam.KeteranganTypeBendungan,
			Authority:                       data.Dam.Kewenangan,
			BuildingCondition:               data.Dam.KondisiBangunan,
			DamWidth:                        data.Dam.LebarBendungan,
			PeakDamWidth:                    data.Dam.LebarPuncakBendungan,
			MinInundationArea:               data.Dam.LuasGenanganMinimal,
			NormalInundationArea:            data.Dam.LuasGenanganNormal,
			TotalInundationArea:             data.Dam.LuasGenanganTotal,
			Name:                            data.Dam.Name,
			DamLength:                       data.Dam.PanjangBendungan,
			PeakDamLength:                   data.Dam.PanjangPuncakBendungan,
			SpillwayWithGate:                data.Dam.PelimpahDenganPintu,
			SpillwayCrestElevation:          data.Dam.PelimpahElevasiMercuAmbang,
			SpillwayWidth:                   data.Dam.PelimpahLebar,
			SpillwayCrestWidth:              data.Dam.PelimpahLebarAmbang,
			SpillwayLength:                  data.Dam.PelimpahPanjang,
			SpillwayTransitionChannelLength: data.Dam.PelimpahPanjangSaluranTransisi,
			SpillwayChannel:                 data.Dam.PelimpahSaluran,
			SpillwayType:                    data.Dam.PelimpahType,
			SpillwayTypeDescription:         data.Dam.PelimpahTypeKeterangan,
			ManagerID:                       data.Dam.PengelolaID,
			ManagerName:                     data.Dam.PengelolaName,
			FloodReductionVolume:            data.Dam.PenguranganDebitBanjir,
			FloodReductionInundationArea:    data.Dam.PenguranganLuasGenanganBanjirHilir,
			HydropowerPlant:                 data.Dam.Plta,
			Province:                        data.Dam.Provinsi,
			ReferenceID:                     data.Dam.RefID,
			InfrastructureStatus:            data.Dam.StatusInfrastructure,
			ConstructionYear:                data.Dam.TahunPembangunan,
			DamHeight:                       data.Dam.TinggiBendungan,
			ExcavationBaseHeight:            data.Dam.TinggiDasarGalian,
			RiverBaseHeight:                 data.Dam.TinggiDasarSungai,
			DamType:                         data.Dam.TypeBendungan,
			TypeID:                          data.Dam.TypeID,
			TypeName:                        data.Dam.TypeName,
			DamVolume:                       data.Dam.VolumeBendungan,
			MinReservoirVolume:              data.Dam.VolumeTampungMinimal,
			NormalReservoirVolume:           data.Dam.VolumeTampungNormal,
			TotalReservoirVolume:            data.Dam.VolumeTampungTotal,
			WsID:                            data.Dam.WsID,
			WsName:                          data.Dam.WsName,
		}, nil
	}
}
