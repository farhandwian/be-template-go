package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PencatatanKejadianRisikoCreateUseCaseReq struct {
	RisikoTeridentifikasi   string `json:"risiko_teridentifikasi"`
	KodeRisiko              string `json:"kode_risiko"`
	TanggalTerjadiRisiko    string `json:"tanggal_terjadi_risiko"`
	SebabRisiko             string `json:"sebab_risiko"`
	DampakRisiko            string `json:"dampak_risiko"`
	KeteranganRisiko        string `json:"keterangan_risiko"`
	RTP                     string `json:"rtp"`
	RencanaPelaksanaanRTP   string `json:"rencana_pelaksanaan_rtp"`
	RealisasiPelaksanaanRTP string `json:"realisasi_pelaksanaan_rtp"`
	KeteranganRTP           string `json:"keterangan_rtp"`
}

type PencatatanKejadianRisikoCreateUseCaseRes struct {
	ID string `json:"id"`
}

type PencatatanKejadianRisikoCreateUseCase = core.ActionHandler[PencatatanKejadianRisikoCreateUseCaseReq, PencatatanKejadianRisikoCreateUseCaseRes]

func ImplPencatatanKejadianRisikoCreateUseCase(
	generateId gateway.GenerateId,
	createPencatatanKejadianRisiko gateway.PencatatanKejadianRisikoSave,
) PencatatanKejadianRisikoCreateUseCase {
	return func(ctx context.Context, req PencatatanKejadianRisikoCreateUseCaseReq) (*PencatatanKejadianRisikoCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		obj := model.PencatatanKejadianRisiko{
			ID:                      &genObj.RandomId,
			RisikoTeridentifikasi:   &req.RisikoTeridentifikasi,
			KodeRisiko:              &req.KodeRisiko,
			TanggalTerjadiRisiko:    &req.TanggalTerjadiRisiko,
			SebabRisiko:             &req.SebabRisiko,
			DampakRisiko:            &req.DampakRisiko,
			KeteranganRisiko:        &req.KeteranganRisiko,
			RTP:                     &req.RTP,
			RencanaPelaksanaanRTP:   &req.RencanaPelaksanaanRTP,
			RealisasiPelaksanaanRTP: &req.RealisasiPelaksanaanRTP,
			KeteranganRTP:           &req.KeteranganRTP,
		}

		if _, err = createPencatatanKejadianRisiko(ctx, gateway.PencatatanKejadianRisikoSaveReq{PencatatanKejadianRisiko: obj}); err != nil {
			return nil, err
		}

		return &PencatatanKejadianRisikoCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
