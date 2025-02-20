package model

// Form 2A

type PenetapanKonteksRisikoStrategisPemda struct {
	ID                     *string `json:"id"`
	NamaPemda              *string `json:"nama_pemda"`
	Periode                *string `json:"periode"`
	SumberData             *string `json:"sumber_data"`
	TujuanStrategis        *string `json:"tujuan_strategis"`
	PenetapanKonteksRisiko *string `json:"penetapan_konteks_resiko"`
	NamaDinas              *string `json:"nama_dinas"`
	Sasaran                *string `json:"sasaran"`
	IKUSasaran             *string `json:"iku_sasaran"`
	PrioritasPembangunan   *string `json:"prioritas_pembangunan"`
	Penilaian              *string `json:"penilaian"`
	NamaYBS                *string `json:"nama_ybs"`
}
