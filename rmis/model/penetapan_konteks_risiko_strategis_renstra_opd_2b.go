package model

// Form 2B

type PenetapanKonteksRisikoStrategisRenstraOPD struct {
	ID                 *string `json:"id"`
	NamaPemda          *string `json:"nama_pemda"`
	TahunPenilaian     *string `json:"tahun_penilaian"`
	Periode            *string `json:"periode"`
	UrusanPemerintahan *string `json:"urusan_pemerintahan"`
	OPDID              *string `json:"opd_id"` // references opd
	TujuanStrategis    *string `json:"tujuan_strategis"`
	SasaranStrategis   *string `json:"sasaran_strategis"`
	InformasiLain      *string `json:"informasi_lain"`
	IKUStrategis       *string `json:"iku_strategis"`
	NamaYBS            *string `json:"nama_ybs"`
}

type PenetapanKonteksRisikoStrategisRenstraOPDGet struct {
	PenetapanKonteksRisikoStrategisRenstraOPD PenetapanKonteksRisikoStrategisRenstraOPD `json:"penetapan_konteks_risiko_strategis_renstra_opd"`
	IKUs                                      []IKU                                     `json:"ikus"`
	OPD                                       OPD                                       `json:"opd"`
}
