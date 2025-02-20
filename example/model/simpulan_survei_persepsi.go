package model

import "time"

// Form 1C
// Simpulan Survei Persepsi atas Lingkungan Pengendalian Intern Pemerintah Daerah

type SimpulanSurveiPersepsi struct {
	ID                      *string       `json:"id"`
	NamaPemda               *string       `json:"nama_pemda"`
	NamaPD                  *string       `json:"nama_pd"`
	TahunPenilaian          *time.Time    `json:"tahun_penilaian"`
	SPIPId                  *string       `json:"spip_id"`
	SimpulanHasilDokumen1A  hasilSimpulan `json:"simpulan_hasil_dokumen_1a"`
	SimpulanUraianDokumen1A hasilSimpulan `json:"simpulan_uraian_dokumen_1a"`
	SimpulanHasilDokumen1B  hasilSimpulan `json:"simpulan_hasil_dokumen_1b"`
	SimpulanUraianDokumen1B hasilSimpulan `json:"simpulan_uraian_dokumen_1b"`
	Simpulan                *string       `json:"simpulan"`
	Penjelasan              *string       `json:"penjelasan"`
}

type hasilSimpulan string

const (
	hasilSimpulanMemadai       hasilSimpulan = "Memadai"
	hasilSimpulanKurangMemadai hasilSimpulan = "Kurang Memadai"
)
