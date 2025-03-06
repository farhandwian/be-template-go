package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	sharedModel "shared/model"
)

type IdentifikasiRisikoStrategisOPDCreateUseCaseReq struct {
	PenetapanKonteksRisikoStrategisRenstraID string  `json:"penetapan_konteks_risiko_strategis_renstra_id"`
	TujuanStrategis                          string  `json:"tujuan_strategis"`
	IndikatorKinerja                         string  `json:"indikator_kinerja"`
	UraianRisiko                             string  `json:"uraian_resiko"`
	PemilikRisiko                            string  `json:"pemilik_resiko"`
	Controllable                             string  `json:"controllable"`
	UraianDampak                             string  `json:"uraian_dampak"`
	PihakDampak                              string  `json:"pihak_dampak"`
	UraianSebab                              string  `json:"uraian_sebab"`
	SumberSebab                              string  `json:"sumber_sebab"`
	KategoriRisikoID                         string  `json:"kategori_risiko_id"`
	RcaID                                    *string `json:"rca_id"`
}

type IdentifikasiRisikoStrategisOPDCreateUseCaseRes struct {
	ID string `json:"id"`
}

type IdentifikasiRisikoStrategisOPDCreateUseCase = core.ActionHandler[
	IdentifikasiRisikoStrategisOPDCreateUseCaseReq,
	IdentifikasiRisikoStrategisOPDCreateUseCaseRes,
]

func ImplIdentifikasiRisikoStrategisOPDCreateUseCase(
	generateId gateway.GenerateId,
	createIdentifikasiRisikoStrategisOPD gateway.IdentifikasiRisikoStrategisOPDSave,
	kodeRisikoByID gateway.KategoriRisikoGetByID,
	penetapanKonteksRisikoStrategisRenstraID gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByID,
	OpdByID gateway.OPDGetByID,
) IdentifikasiRisikoStrategisOPDCreateUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDCreateUseCaseReq) (*IdentifikasiRisikoStrategisOPDCreateUseCaseRes, error) {
		fmt.Println("IdentifikasiRisikoStrategisOPDCreateUseCase")
		// Generate a unique ID
		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		kategoriRisikoRes, err := kodeRisikoByID(ctx, gateway.KategoriRisikoGetByIDReq{ID: req.KategoriRisikoID})
		if err != nil {
			return nil, fmt.Errorf("failed to get KategoriRisikoName: %v", err)
		}

		penetapanKonteksRisikoStrategisRenstraRes, err := penetapanKonteksRisikoStrategisRenstraID(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByIDReq{ID: req.PenetapanKonteksRisikoStrategisRenstraID})
		if err != nil {
			return nil, fmt.Errorf("failed to get PenetapanKonteksRisikoStrategisPemda: %v", err)
		}

		opdRes, err := OpdByID(ctx, gateway.OPDGetByIDReq{ID: *penetapanKonteksRisikoStrategisRenstraRes.PenetapanKonteksRisikoStrategisRenstraOPD.OpdID})
		if err != nil {
			return nil, fmt.Errorf("failed to get OPD: %v", err)
		}

		// if req.RcaID != nil {
		// 	rcaRes, err := rcaByID(ctx, gateway.RcaGetByIDReq{ID: *req.RcaID})
		// 	if err != nil {
		// 		return nil, fmt.Errorf("failed to get rca: %v", err)
		// 	}
		// }
		obj := model.IdentifikasiRisikoStrategisOPD{
			ID:                                       &genObj.RandomId,
			PenetapanKonteksRisikoStrategisRenstraID: &req.PenetapanKonteksRisikoStrategisRenstraID,
			KategoriRisikoID:                         &req.KategoriRisikoID,
			UraianRisiko:                             &req.UraianRisiko,
			PemilikRisiko:                            &req.PemilikRisiko,
			Controllable:                             &req.Controllable,
			UraianDampak:                             &req.UraianDampak,
			PihakDampak:                              &req.PihakDampak,
			UraianSebab:                              &req.UraianSebab,
			SumberSebab:                              &req.SumberSebab,
			Status:                                   sharedModel.StatusMenungguVerifikasi,
		}

		fmt.Println("TEST")
		obj.GenerateKodeRisiko(
			*penetapanKonteksRisikoStrategisRenstraRes.PenetapanKonteksRisikoStrategisRenstraOPD.TahunPenilaian, *kategoriRisikoRes.KategoriRisiko.Kode, *opdRes.OPD.Kode,
		)
		// Save the new entry
		if _, err = createIdentifikasiRisikoStrategisOPD(ctx, gateway.IdentifikasiRisikoStrategisOPDSaveReq{
			IdentifikasiRisikoStrategisOPD: obj,
		}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoStrategisOPDCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
