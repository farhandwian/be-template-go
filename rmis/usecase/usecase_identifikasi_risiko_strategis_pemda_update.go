package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"shared/core"
)

type IdentifikasiRisikoStrategisPemdaUpdateUseCaseReq struct {
	ID                                     string `json:"id"`
	PenetapanKonteksRisikoStrategisPemdaID string `json:"penetapan_konteks_risiko_strategis_id"`
	TujuanStrategis                        string `json:"tujuan_strategis"`
	IndikatorKinerja                       string `json:"indikator_kinerja"`
	UraianRisiko                           string `json:"uraian_resiko"`
	PemilikRisiko                          string `json:"pemilik_resiko"`
	Controllable                           string `json:"controllable"`
	UraianDampak                           string `json:"uraian_dampak"`
	PihakDampak                            string `json:"pihak_dampak"`
	KategoriRisikoID                       string `json:"kategori_resiko_id"`
	RcaID                                  string `json:"rca_id"`
}

type IdentifikasiRisikoStrategisPemdaUpdateUseCaseRes struct{}

type IdentifikasiRisikoStrategisPemdaUpdateUseCase = core.ActionHandler[IdentifikasiRisikoStrategisPemdaUpdateUseCaseReq, IdentifikasiRisikoStrategisPemdaUpdateUseCaseRes]

func ImplIdentifikasiRisikoStrategisPemdaUpdateUseCase(
	getIdentifikasiRisikoStrategisPemdaById gateway.IdentifikasiRisikoStrategisPemdaGetByID,
	updateIdentifikasiRisikoStrategisPemda gateway.IdentifikasiRisikoStrategisPemdaSave,
	kodeRisikoByID gateway.KategoriRisikoGetByID,
	RcaByID gateway.RcaGetByID,
	penetapanKonteksRisikoStrategisPemdaID gateway.PenetapanKonteksRisikoStrategisPemdaGetByID,
) IdentifikasiRisikoStrategisPemdaUpdateUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaUpdateUseCaseReq) (*IdentifikasiRisikoStrategisPemdaUpdateUseCaseRes, error) {

		res, err := getIdentifikasiRisikoStrategisPemdaById(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		identifikasiRisikoStrategisPemda := res.IdentifikasiRisikoStrategisPemda
		identifikasiRisikoStrategisPemda.UraianRisiko = &req.UraianRisiko
		identifikasiRisikoStrategisPemda.PemilikRisiko = &req.PemilikRisiko
		identifikasiRisikoStrategisPemda.Controllable = &req.Controllable
		identifikasiRisikoStrategisPemda.UraianDampak = &req.UraianDampak
		identifikasiRisikoStrategisPemda.PihakDampak = &req.PihakDampak

		if identifikasiRisikoStrategisPemda.KategoriRisikoID == nil || *identifikasiRisikoStrategisPemda.KategoriRisikoID == "" {
			return nil, fmt.Errorf("KategoriRisikoID is missing in the database record")
		}

		kategoriRisikoRes, err := kodeRisikoByID(ctx, gateway.KategoriRisikoGetByIDReq{
			ID: *identifikasiRisikoStrategisPemda.KategoriRisikoID,
		})

		penetapanKonteksRisikoStrategisPemdaRes, err := penetapanKonteksRisikoStrategisPemdaID(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetByIDReq{ID: req.PenetapanKonteksRisikoStrategisPemdaID})
		if err != nil {
			return nil, fmt.Errorf("failed to get PenetapanKonteksRisikoStrategisPemda: %v", err)
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get KategoriRisikoName: %v", err)
		}

		if identifikasiRisikoStrategisPemda.KodeRisiko == nil || *identifikasiRisikoStrategisPemda.KodeRisiko == "" {
			identifikasiRisikoStrategisPemda.GenerateKodeRisiko(*penetapanKonteksRisikoStrategisPemdaRes.PenetapanKonteksRisikoStrategisPemda.TahunPenilaian, *kategoriRisikoRes.KategoriRisiko.Kode)
		}

		// rcaRes, err := RcaByID(ctx, gateway.RcaGetByIDReq{ID: req.RcaID})
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to get RcaName: %v", err)
		// }
		// identifikasiRisikoStrategisPemda.RcaID = rcaRes.Rca.ID
		// identifikasiRisikoStrategisPemda.UraianSebab = rcaRes.Rca.AkarPenyebab
		// identifikasiRisikoStrategisPemda.SumberSebab = rcaRes.Rca.PernyataanRisiko

		if _, err := updateIdentifikasiRisikoStrategisPemda(ctx, gateway.IdentifikasiRisikoStrategisPemdaSaveReq{IdentifikasiRisikoStrategisPemda: identifikasiRisikoStrategisPemda}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoStrategisPemdaUpdateUseCaseRes{}, nil
	}
}
