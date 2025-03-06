package model

import (
	sharedModel "shared/model"
	"time"
)

// form 10

type PencatatanKejadianRisiko struct {
	ID                                     *string `json:"id"`
	PenetapanKonteksRisikoStrategisPemdaID *string `json:"-" gorm:"type:VARCHAR(255)"`
	IdentifikasiRisikoStrategisPemdaID     *string `json:"-" gorm:"type:VARCHAR(255)"`

	TanggalTerjadiRisiko    *string `json:"tanggal_terjadi_risiko"`
	SebabRisiko             *string `json:"sebab_risiko"`
	DampakRisiko            *string `json:"dampak_risiko"`
	KeteranganRisiko        *string `json:"keterangan_risiko"`
	RTP                     *string `json:"rtp"`
	RencanaPelaksanaanRTP   *string `json:"rencana_pelaksanaan_rtp"`
	RealisasiPelaksanaanRTP *string `json:"realisasi_pelaksanaan_rtp"`
	KeteranganRTP           *string `json:"keterangan_rtp"`

	Status    sharedModel.Status `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type PencatatanKejadianRisikoResponse struct {
	ID                                     *string `json:"id"`
	PenetapanKonteksRisikoStrategisPemdaID *string `json:"-" gorm:"type:VARCHAR(255)"`
	IdentifikasiRisikoStrategisPemdaID     *string `json:"-" gorm:"type:VARCHAR(255)"`

	TanggalTerjadiRisiko    *string `json:"tanggal_terjadi_risiko"`
	SebabRisiko             *string `json:"sebab_risiko"`
	DampakRisiko            *string `json:"dampak_risiko"`
	KeteranganRisiko        *string `json:"keterangan_risiko"`
	RTP                     *string `json:"rtp"`
	RencanaPelaksanaanRTP   *string `json:"rencana_pelaksanaan_rtp"`
	RealisasiPelaksanaanRTP *string `json:"realisasi_pelaksanaan_rtp"`
	KeteranganRTP           *string `json:"keterangan_rtp"`

	// From Join
	NamaPemda          *string    `json:"nama_pemda"`
	TahunPenilaian     *time.Time `json:"tahun_penilaian"`
	UrusanPemerintahan *string    `json:"urusan_pemerintahan"`
	TujuanStrategis    *string    `json:"tujuan_strategis"`
	KodeRisiko         *string    `json:"kode_risiko"`

	Status    sharedModel.Status `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
