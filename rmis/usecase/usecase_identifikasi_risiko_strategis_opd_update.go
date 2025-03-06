package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"shared/core"
)

type IdentifikasiRisikoStrategisOPDUpdateUseCaseReq struct {
	ID               string `json:"-"`
	TujuanStrategis  string `json:"tujuan_strategis"`
	IndikatorKinerja string `json:"indikator_kinerja"`
	UraianRisiko     string `json:"uraian_resiko"`
	PemilikRisiko    string `json:"pemilik_resiko"`
	Controllable     string `json:"controllable"`
	UraianDampak     string `json:"uraian_dampak"`
	PihakDampak      string `json:"pihak_dampak"`

	PenetapanKonteksRisikoStrategisRenstraID string  `json:"penetapan_konteks_risiko_strategis_renstra_id"`
	KategoriRisikoID                         string  `json:"kategori_risiko_id"`
	RcaID                                    *string `json:"rca_id"`
}

type IdentifikasiRisikoStrategisOPDUpdateUseCaseRes struct{}

type IdentifikasiRisikoStrategisOPDUpdateUseCase = core.ActionHandler[IdentifikasiRisikoStrategisOPDUpdateUseCaseReq, IdentifikasiRisikoStrategisOPDUpdateUseCaseRes]

func ImplIdentifikasiRisikoStrategisOPDUpdateUseCase(
	getIdentifikasiRisikoStrategisOPDById gateway.IdentifikasiRisikoStrategisOPDGetByID,
	updateIdentifikasiRisikoStrategisOPD gateway.IdentifikasiRisikoStrategisOPDSave,
	kategoriRIsikoByID gateway.KategoriRisikoGetByID,
	RcaByID gateway.RcaGetByID,
	getOneOPD gateway.OPDGetByID,
	penetapanKonteksRisikoStrategisRenstraID gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByID,
) IdentifikasiRisikoStrategisOPDUpdateUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDUpdateUseCaseReq) (*IdentifikasiRisikoStrategisOPDUpdateUseCaseRes, error) {

		res, err := getIdentifikasiRisikoStrategisOPDById(ctx, gateway.IdentifikasiRisikoStrategisOPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.IdentifikasiRisikoStrategisOPD.UraianRisiko = &req.UraianRisiko
		res.IdentifikasiRisikoStrategisOPD.PemilikRisiko = &req.PemilikRisiko
		res.IdentifikasiRisikoStrategisOPD.Controllable = &req.Controllable
		res.IdentifikasiRisikoStrategisOPD.UraianDampak = &req.UraianDampak
		res.IdentifikasiRisikoStrategisOPD.PihakDampak = &req.PihakDampak
		res.IdentifikasiRisikoStrategisOPD.KategoriRisikoID = &req.KategoriRisikoID
		res.IdentifikasiRisikoStrategisOPD.PenetapanKonteksRisikoStrategisRenstraID = &req.PenetapanKonteksRisikoStrategisRenstraID

		kategoriRisikoRes, err := kategoriRIsikoByID(ctx, gateway.KategoriRisikoGetByIDReq{
			ID: *res.IdentifikasiRisikoStrategisOPD.KategoriRisikoID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get KategoriRisiko: %v", err)
		}

		penetapanKonteksRisikoStrategisRenstraIDRes, err := penetapanKonteksRisikoStrategisRenstraID(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByIDReq{ID: *res.IdentifikasiRisikoStrategisOPD.PenetapanKonteksRisikoStrategisRenstraID})
		if err != nil {
			return nil, fmt.Errorf("failed to get PenetapanKonteksRisikoStrategisRenstraID: %v", err)
		}

		opdRes, err := getOneOPD(ctx, gateway.OPDGetByIDReq{ID: *penetapanKonteksRisikoStrategisRenstraIDRes.PenetapanKonteksRisikoStrategisRenstraOPD.OpdID})
		if err != nil {
			return nil, fmt.Errorf("failed to get OPDName: %v", err)
		}

		if res.IdentifikasiRisikoStrategisOPD.KodeRisiko == nil || *res.IdentifikasiRisikoStrategisOPD.KodeRisiko == "" {
			res.IdentifikasiRisikoStrategisOPD.GenerateKodeRisiko(*penetapanKonteksRisikoStrategisRenstraIDRes.PenetapanKonteksRisikoStrategisRenstraOPD.TahunPenilaian, *kategoriRisikoRes.KategoriRisiko.Kode, *opdRes.OPD.Kode)
		}

		// rcaRes, err := RcaByID(ctx, gateway.RcaGetByIDReq{ID: req.RcaID})
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to get RcaName: %v", err)
		// }

		if _, err := updateIdentifikasiRisikoStrategisOPD(ctx, gateway.IdentifikasiRisikoStrategisOPDSaveReq{IdentifikasiRisikoStrategisOPD: res.IdentifikasiRisikoStrategisOPD}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoStrategisOPDUpdateUseCaseRes{}, nil
	}
}
