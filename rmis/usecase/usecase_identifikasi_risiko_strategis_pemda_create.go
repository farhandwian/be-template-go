package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	sharedModel "shared/model"
)

type IdentifikasiRisikoStrategisPemdaCreateUseCaseReq struct {
	PenetapanKonteksRisikoStrategisPemdaID string             `json:"penetapan_konteks_risiko_strategis_pemda_id"`
	TujuanStrategis                        string             `json:"tujuan_strategis"`
	IndikatorKinerja                       string             `json:"indikator_kinerja"`
	UraianRisiko                           string             `json:"uraian_resiko"`
	PemilikRisiko                          string             `json:"pemilik_resiko"`
	Controllable                           string             `json:"controllable"`
	UraianDampak                           string             `json:"uraian_dampak"`
	PihakDampak                            string             `json:"pihak_dampak"`
	KategoriRisikoID                       string             `json:"kategori_risiko_id"`
	RcaID                                  *string            `json:"rca_id"`
	Status                                 sharedModel.Status `json:"status"`
}

type IdentifikasiRisikoStrategisPemdaCreateUseCaseRes struct {
	ID string `json:"id"`
}

type IdentifikasiRisikoStrategisPemdaCreateUseCase = core.ActionHandler[
	IdentifikasiRisikoStrategisPemdaCreateUseCaseReq,
	IdentifikasiRisikoStrategisPemdaCreateUseCaseRes,
]

func ImplIdentifikasiRisikoStrategisPemdaCreateUseCase(
	generateId gateway.GenerateId,
	createIdentifikasiRisikoStrategisPemda gateway.IdentifikasiRisikoStrategisPemdaSave,
	kodeRisikoByID gateway.KategoriRisikoGetByID,
	penetapanKonteksRisikoStrategisPemdaID gateway.PenetapanKonteksRisikoStrategisPemdaGetByID,
	// rcaByID gateway.RcaGetByID,
) IdentifikasiRisikoStrategisPemdaCreateUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaCreateUseCaseReq) (*IdentifikasiRisikoStrategisPemdaCreateUseCaseRes, error) {

		// Generate a unique ID
		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		kategoriRisikoRes, err := kodeRisikoByID(ctx, gateway.KategoriRisikoGetByIDReq{ID: req.KategoriRisikoID})
		if err != nil {
			return nil, fmt.Errorf("failed to get KategoriRisiko: %v", err)
		}

		penetapanKonteksRisikoStrategisPemdaRes, err := penetapanKonteksRisikoStrategisPemdaID(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetByIDReq{ID: req.PenetapanKonteksRisikoStrategisPemdaID})
		if err != nil {
			return nil, fmt.Errorf("failed to get PenetapanKonteksRisikoStrategisPemda: %v", err)
		}

		// if req.RcaID != nil {
		// 	rcaRes, err := rcaByID(ctx, gateway.RcaGetByIDReq{ID: *req.RcaID})
		// 	if err != nil {
		// 		return nil, fmt.Errorf("failed to get rca: %v", err)
		// 	}
		// }
		obj := model.IdentifikasiRisikoStrategisPemda{
			ID:                                     &genObj.RandomId,
			PenetapanKonteksRisikoStrategisPemdaID: &req.PenetapanKonteksRisikoStrategisPemdaID,
			TujuanStrategis:                        &req.TujuanStrategis,
			IndikatorKinerja:                       &req.IndikatorKinerja,
			KategoriRisikoID:                       &req.KategoriRisikoID,
			UraianRisiko:                           &req.UraianRisiko,
			PemilikRisiko:                          &req.PemilikRisiko,
			Controllable:                           &req.Controllable,
			UraianDampak:                           &req.UraianDampak,
			PihakDampak:                            &req.PihakDampak,
			Status:                                 sharedModel.StatusMenungguVerifikasi,
		}

		fmt.Println(penetapanKonteksRisikoStrategisPemdaRes.PenetapanKonteksRisikoStrategisPemda)
		fmt.Println("================================================================", *penetapanKonteksRisikoStrategisPemdaRes.PenetapanKonteksRisikoStrategisPemda.TahunPenilaian)
		obj.GenerateKodeRisiko(*penetapanKonteksRisikoStrategisPemdaRes.PenetapanKonteksRisikoStrategisPemda.TahunPenilaian, *kategoriRisikoRes.KategoriRisiko.Kode)
		fmt.Println("================================================================TETSUN")

		// Save the new entry
		if _, err = createIdentifikasiRisikoStrategisPemda(ctx, gateway.IdentifikasiRisikoStrategisPemdaSaveReq{
			IdentifikasiRisikoStrategisPemda: obj,
		}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoStrategisPemdaCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
