package usecase

import (
	"context"
	"gorm.io/datatypes"
	"shared/core"
	"shared/gateway"
)

type GetWaterReservoirDetailReq struct {
	ID string
}

type GetWaterReservoirDetailResp struct {
	ID                             string         `json:"id"`
	Geometry                       datatypes.JSON `json:"geometry"`
	DasID                          string         `json:"das_id"`
	DasName                        string         `json:"das_name"`
	IsPUPR                         bool           `json:"is_pupr"`
	WSID                           string         `json:"ws_id"`
	WSName                         string         `json:"ws_name"`
	Name                           string         `json:"name"`
	RefID                          string         `json:"ref_id"`
	WaterSurfaceElevationMax       float64        `json:"water_surface_elevation_max"`
	WaterSurfaceElevationMin       float64        `json:"water_surface_elevation_min"`
	WaterSurfaceElevationNormal    float64        `json:"water_surface_elevation_norma"`
	WaterSurfaceElevationPeak      float64        `json:"water_surface_elevation_peak"`
	WaterReservoirType             string         `json:"water_reservoir_type"`
	Province                       string         `json:"province"`
	City                           string         `json:"city"`
	District                       string         `json:"district"`
	SubDistrict                    string         `json:"sub_district"`
	Authority                      string         `json:"authority"`
	ManagerID                      int            `json:"manager_id"`
	ManagerName                    string         `json:"manager_name"`
	WaterReservoirWidth            float64        `json:"water_reservoir_width"`
	WaterReservoirLength           float64        `json:"water_reservoir_length"`
	FoundationWaterReservoirHeight float64        `json:"foundation_water_reservoir_height"`
	FoundationRiverHeight          float64        `json:"foundation_river_height"`
	WaterReservoirBodyType         string         `json:"water_reservoir_body_type"`
	WaterSurfaceElevation          float64        `json:"water_surface_elevation"`
	BodyVolume                     float64        `json:"body_volume"`
	TypeID                         int            `json:"type_id"`
}
type GetWaterReservoirDetailUseCase = core.ActionHandler[GetWaterReservoirDetailReq, GetWaterReservoirDetailResp]

func ImplGetWaterReservoirDetail(getWaterReservoirDetail gateway.WaterReservoirGetAllGatewayGateway) GetWaterReservoirDetailUseCase {
	return func(ctx context.Context, req GetWaterReservoirDetailReq) (*GetWaterReservoirDetailResp, error) {
		data, err := getWaterReservoirDetail(ctx, gateway.WaterReservoirGetDetailReq{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return &GetWaterReservoirDetailResp{
			ID:                             data.WaterReservoir.ID,
			Geometry:                       data.WaterReservoir.Geometry,
			DasID:                          data.WaterReservoir.DASID,
			DasName:                        data.WaterReservoir.DASName,
			IsPUPR:                         data.WaterReservoir.IsPUPR,
			WSID:                           data.WaterReservoir.WSID,
			WSName:                         data.WaterReservoir.WSName,
			Name:                           data.WaterReservoir.Name,
			RefID:                          data.WaterReservoir.RefID,
			WaterSurfaceElevationMax:       data.WaterReservoir.ElevasiMukaAirMax,
			WaterSurfaceElevationMin:       data.WaterReservoir.ElevasiMukaAirMin,
			WaterSurfaceElevationNormal:    data.WaterReservoir.ElevasiMukaAirNormal,
			WaterSurfaceElevationPeak:      data.WaterReservoir.ElevasiMukaAirPuncak,
			WaterReservoirType:             data.WaterReservoir.JenisEmbung,
			Province:                       data.WaterReservoir.Provinsi,
			City:                           data.WaterReservoir.Kabupaten,
			District:                       data.WaterReservoir.Kecamatan,
			SubDistrict:                    data.WaterReservoir.Kelurahan,
			Authority:                      data.WaterReservoir.Kewenangan,
			ManagerID:                      data.WaterReservoir.PengelolaID,
			ManagerName:                    data.WaterReservoir.PengelolaName,
			WaterReservoirWidth:            data.WaterReservoir.LebarTubuhEmbung,
			WaterReservoirLength:           data.WaterReservoir.PanjangTubuhEmbung,
			FoundationWaterReservoirHeight: data.WaterReservoir.TinggiTubuhEmbungPondasi,
			FoundationRiverHeight:          data.WaterReservoir.TinggiTubuhEmbungSungai,
			WaterReservoirBodyType:         data.WaterReservoir.TipeTubuhEmbung,
			WaterSurfaceElevation:          data.WaterReservoir.TMA,
			BodyVolume:                     data.WaterReservoir.VolumeTubuhEmbung,
			TypeID:                         data.WaterReservoir.TypeID,
		}, err
	}
}
