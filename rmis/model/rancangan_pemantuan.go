package model

// form 9

type RancanganPemantauan struct {
	ID                   *string           `json:"id"`
	NamaPemda            *string           `json:"nama_pemda"`
	TahunPenilaian       *string           `json:"tahun_penilaian"`
	TujuanStrategis      *string           `json:"tujuan_strategis"`
	UrusanPemerintahan   *string           `json:"urusan_pemerintahan"`
	KegiatanPengendalian *string           `json:"kegiatan_pengendalian"`
	MetodePemantauan     *MetodePemantauan `json:"metode_pemantuan"`
	PenanggungJawab      *string           `json:"penanggung_jawab"`
	RencanaPenyelesaian  *string           `json:"rencana_penyelesaian"`
	RealisasiPelaksanaan *string           `json:"realisasi_pelaksanaan"`
	Keterangan           *string           `json:"keterangan"`
}

type MetodePemantauan string

const (
	MetodePemantauanPersiapanDanLaporan MetodePemantauan = "Konfirmasi persiapan dan laporan pelaksanaan kegiatan"
	MetodePemantauanPelaksanakan        MetodePemantauan = "Konfirmasi pelaksanaan Laporan pelaksanaan kegiatan"
)
