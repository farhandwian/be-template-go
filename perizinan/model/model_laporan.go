package model

import (
	"fmt"
	"time"

	"gorm.io/datatypes"
)

func LaporanPerizinanEmpty(noSK NomorSK, periode Periode) *LaporanPerizinan {
	return &LaporanPerizinan{
		Status:                StatusLaporDraft,
		PeriodePengambilanSDA: periode,
		NoSK:                  noSK,
	}
}

type LaporanPerizinanID string

type Periode string

// validatePeriodePengambilanSDA checks if the period is in correct format (YYYY-MM)
func (p Periode) Validate(min, now time.Time) error {

	// Check format using time.Parse
	periodTime, err := time.Parse("2006-01", string(p))
	if err != nil {
		return fmt.Errorf("salah format periode. Harus YYYY-MM (contoh: 2024-08)")
	}

	_ = periodTime

	// Set the input times to the first day of their respective months for consistent comparison
	// minTime := time.Date(min.Year(), min.Month(), 1, 0, 0, 0, 0, time.UTC)
	// maxTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)

	// Subtract one month from current time to get the latest allowed period
	// maxTime = maxTime.AddDate(0, -1, 0)

	// // Check if period is within allowed range
	// if periodTime.Before(minTime) {
	// 	return fmt.Errorf("periode tidak bisa lebih awal dari %s", minTime.Format("2006-01"))
	// }

	// if periodTime.After(maxTime) {
	// 	return fmt.Errorf("periode tidak boleh lebih lama dari %s", maxTime.Format("2006-01"))
	// }

	return nil
}

type LaporanPerizinan struct {
	// meta

	ID            LaporanPerizinanID     `json:"id"`
	Status        LaporanPerizinanStatus `json:"status"`
	SKPerizinanID SKPerizinanID          `json:"-" gorm:"type:char(36);index:idx_lapor_bulan"`
	NoSK          NomorSK                `json:"no_sk"`

	// part 1
	KoordinatDiLapangan           Geometry       `json:"koordinat_di_lapangan"`
	CaraPengambilanSDADiLapangan  string         `json:"cara_pengambilan_sda_di_lapangan"`
	JenisTipeKonstruksiDiLapangan string         `json:"jenis_tipe_konstruksi_di_lapangan"`
	PeriodePengambilanSDA         Periode        `json:"periode_pengambilan_sda" gorm:"type:varchar(20);index:idx_lapor_bulan"`
	DebitPengambilan              float64        `json:"debit_pengambilan"`
	LaporanPengambilanAirBulanan  datatypes.JSON `json:"laporan_pengambilan_air_bulanan"`

	// part 2
	JadwalPengambilanDiLapangan         string `json:"jadwal_pengambilan_di_lapangan"`
	JadwalPembangunanDiLapangan         string `json:"jadwal_pembangunan_di_lapangan"`
	TanggalPemegangDilarangMengambilAir string `json:"tanggal_pemegang_dilarang_mengambil_air"`
	RealisasiDiLapangan                 string `json:"realisasi_di_lapangan"`

	// part 3
	LaporanHasilPemeriksaanTimVerifikasi datatypes.JSON `json:"laporan_hasil_pemeriksaan_tim_verifikasi"`
	FileKeberadaanAlatUkurDebit          datatypes.JSON `json:"file_keberadaan_alat_ukur_debit"`
	FileKeberadaanSistemTelemetri        datatypes.JSON `json:"file_keberadaan_sistem_telemetri"`

	KeberadaanAlatUkurDebit   bool `json:"keberadaan_alat_ukur_debit"`
	KeberadaanSistemTelemetri bool `json:"keberadaan_sistem_telemetri"`

	TerdapatAirDibuangKeSumber     bool           `json:"terdapat_air_dibuang_ke_sumber"`
	DebitAirBuangan                float64        `json:"debit_air_buangan"`
	LaporanHasilPemeriksaanBuangan datatypes.JSON `json:"laporan_hasil_pemeriksaan_buangan"`

	// part 4
	DokumenBuktiBayar               datatypes.JSON `json:"dokumen_bukti_bayar"`
	DokumenKewajibanKeuanganLainnya datatypes.JSON `json:"dokumen_kewajiban_keuangan_lainnya"`

	BuktiKerusakanSumberAir datatypes.JSON `json:"bukti_kerusakan_sumber_air"`
	PerbaikanKerusakan      bool           `json:"perbaikan_kerusakan"`

	BuktiUsahaPengendalianPencemaran datatypes.JSON `json:"bukti_usaha_pengendalian_pencemaran"`
	BentukUsahaPengendalian          string         `json:"bentuk_usaha_pengendalian"`

	BuktiPenggunaanAir   datatypes.JSON `json:"bukti_penggunaan_air"`
	DebitAirDihasilkan   float64        `json:"debit_air_dihasilkan"`
	ManfaatUntukKegiatan string         `json:"manfaat_untuk_kegiatan"`

	KegiatanOP               string `json:"kegiatan_op"`
	KoordinasiBBWSRencanaOP  string `json:"koordinasi_bbws_rencana_op"`
	KoordinasiBBWSKonstruksi string `json:"koordinasi_bbws_konstruksi"`

	PLDebitYangDigunakan      float64 `json:"pl_debit_yang_digunakan"`
	PLKualiatsAirDigunakan    string  `json:"pl_kualitas_air_digunakan"`
	PLDebitYangDikembalikan   float64 `json:"pl_debit_yang_dikembalikan"`
	PLKualitasAirDikembalikan string  `json:"pl_kualitas_air_dikembalikan"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LaporanPerizinanStatus string

// enum for status
const (
	StatusLaporSubmitted LaporanPerizinanStatus = "submitted"
	StatusLaporDraft     LaporanPerizinanStatus = "draft"
)
