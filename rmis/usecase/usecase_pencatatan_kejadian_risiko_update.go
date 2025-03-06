package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type PencatatanKejadianRisikoUpdateUseCaseReq struct {
	ID                                     string `json:"-"`
	PenetapanKonteksRisikoStrategisPemdaID string `json:"penetapan_konteks_risiko_strategis_pemda_id"`
	IdentifikasiRisikoStrategisPemdaID     string `json:"identifikasi_risiko_strategis_pemda_id"`

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
	PenetapanKonteksRisikoStrategisPemdaByID gateway.PenetapanKonteksRisikoStrategisPemdaGetByID,
	IdentifikasiRisikoStrategisPemdaByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
) PencatatanKejadianRisikoUpdateUseCase {
	return func(ctx context.Context, req PencatatanKejadianRisikoUpdateUseCaseReq) (*PencatatanKejadianRisikoUpdateUseCaseRes, error) {

		res, err := getPencatatanKejadianRisikoById(ctx, gateway.PencatatanKejadianRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		_, err = PenetapanKonteksRisikoStrategisPemdaByID(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetByIDReq{ID: req.PenetapanKonteksRisikoStrategisPemdaID})
		if err != nil {
			return nil, err
		}

		_, err = IdentifikasiRisikoStrategisPemdaByID(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{ID: req.IdentifikasiRisikoStrategisPemdaID})
		if err != nil {
			return nil, err
		}

		pencatatanKejadianRisiko := res.PencatatanKejadianRisiko
		pencatatanKejadianRisiko.PenetapanKonteksRisikoStrategisPemdaID = &req.PenetapanKonteksRisikoStrategisPemdaID
		pencatatanKejadianRisiko.IdentifikasiRisikoStrategisPemdaID = &req.IdentifikasiRisikoStrategisPemdaID
		pencatatanKejadianRisiko.TanggalTerjadiRisiko = &req.TanggalTerjadiRisiko
		pencatatanKejadianRisiko.SebabRisiko = &req.SebabRisiko
		pencatatanKejadianRisiko.DampakRisiko = &req.DampakRisiko
		pencatatanKejadianRisiko.KeteranganRisiko = &req.KeteranganRisiko
		pencatatanKejadianRisiko.RTP = &req.RTP
		pencatatanKejadianRisiko.RencanaPelaksanaanRTP = &req.RencanaPelaksanaanRTP
		pencatatanKejadianRisiko.RealisasiPelaksanaanRTP = &req.RealisasiPelaksanaanRTP
		pencatatanKejadianRisiko.KeteranganRTP = &req.KeteranganRTP
		pencatatanKejadianRisiko.Status = sharedModel.StatusMenungguVerifikasi

		if _, err := updatePencatatanKejadianRisiko(ctx, gateway.PencatatanKejadianRisikoSaveReq{PencatatanKejadianRisiko: pencatatanKejadianRisiko}); err != nil {
			return nil, err
		}

		return &PencatatanKejadianRisikoUpdateUseCaseRes{}, nil
	}
}
