package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PencatatanKejadianRisikoUpdateUseCaseReq struct {
	ID                      string `json:"id"`
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

type PencatatanKejadianRisikoUpdateUseCaseRes struct{}

type PencatatanKejadianRisikoUpdateUseCase = core.ActionHandler[PencatatanKejadianRisikoUpdateUseCaseReq, PencatatanKejadianRisikoUpdateUseCaseRes]

func ImplPencatatanKejadianRisikoUpdateUseCase(
	getPencatatanKejadianRisikoById gateway.PencatatanKejadianRisikoGetByID,
	updatePencatatanKejadianRisiko gateway.PencatatanKejadianRisikoSave,
) PencatatanKejadianRisikoUpdateUseCase {
	return func(ctx context.Context, req PencatatanKejadianRisikoUpdateUseCaseReq) (*PencatatanKejadianRisikoUpdateUseCaseRes, error) {

		res, err := getPencatatanKejadianRisikoById(ctx, gateway.PencatatanKejadianRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		pencatatanKejadianRisiko := res.PencatatanKejadianRisiko
		pencatatanKejadianRisiko.RisikoTeridentifikasi = &req.RisikoTeridentifikasi
		pencatatanKejadianRisiko.KodeRisiko = &req.KodeRisiko
		pencatatanKejadianRisiko.TanggalTerjadiRisiko = &req.TanggalTerjadiRisiko
		pencatatanKejadianRisiko.SebabRisiko = &req.SebabRisiko
		pencatatanKejadianRisiko.DampakRisiko = &req.DampakRisiko
		pencatatanKejadianRisiko.KeteranganRisiko = &req.KeteranganRisiko
		pencatatanKejadianRisiko.RTP = &req.RTP
		pencatatanKejadianRisiko.RencanaPelaksanaanRTP = &req.RencanaPelaksanaanRTP
		pencatatanKejadianRisiko.RealisasiPelaksanaanRTP = &req.RealisasiPelaksanaanRTP
		pencatatanKejadianRisiko.KeteranganRTP = &req.KeteranganRTP

		if _, err := updatePencatatanKejadianRisiko(ctx, gateway.PencatatanKejadianRisikoSaveReq{PencatatanKejadianRisiko: res.PencatatanKejadianRisiko}); err != nil {
			return nil, err
		}

		return &PencatatanKejadianRisikoUpdateUseCaseRes{}, nil
	}
}
