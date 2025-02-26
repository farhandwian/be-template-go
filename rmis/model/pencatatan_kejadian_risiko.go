package model

import "time"

// form 10

type PencatatanKejadianRisiko struct {
	ID                      *string    `json:"id"`
	NamaPemda               *string    `json:"nama_pemda"`
	TahunPenilaian          *time.Time `json:"tahun_penilaian"`
	TujuanStrategis         *string    `json:"tujuaN_strategis"` // corrected to "TujuanStrategis" in struct; adjust tag if necessary
	UrusanPemerintahan      *string    `json:"urusan_pemerintahan"`
	RisikoTeridentifikasi   *string    `json:"risiko_teridentifikasi"`
	KodeRisiko              *string    `json:"kode_risiko"`
	TanggalTerjadiRisiko    *string    `json:"tanggal_terjadi_risiko"`
	SebabRisiko             *string    `json:"sebab_risiko"`
	DampakRisiko            *string    `json:"dampak_risiko"`
	KeteranganRisiko        *string    `json:"keterangan_risiko"`
	RTP                     *string    `json:"rtp"`
	RencanaPelaksanaanRTP   *string    `json:"rencana_pelaksanaan_rtp"`
	RealisasiPelaksanaanRTP *string    `json:"realisasi_pelaksanaan_rtp"`
	KeteranganRTP           *string    `json:"keterangan_rtp"`
}
