package usecase

import (
	"context"
	"errors"
	"fmt"
	"perizinan/gateway"
	"perizinan/model"
	"shared/core"
	"shared/helper"
)

type LngLat struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type LaporanPerizinanSaveUseCaseReq struct {
	//

	Page    int           `json:"page"`
	NomorSK model.NomorSK `json:"nomor_sk"`

	// part 1
	KoordinatDiLapangan           LngLat              `json:"koordinat_di_lapangan"`
	CaraPengambilanSDADiLapangan  string              `json:"cara_pengambilan_sda_di_lapangan"`
	JenisTipeKonstruksiDiLapangan string              `json:"jenis_tipe_konstruksi_di_lapangan"`
	PeriodePengambilanSDA         model.Periode       `json:"periode_pengambilan_sda"`
	DebitPengambilan              float64             `json:"debit_pengambilan"`
	LaporanPengambilanAirBulanan  []model.Attachments `json:"laporan_pengambilan_air_bulanan"`

	// part 2
	JadwalPengambilanDiLapangan         string `json:"jadwal_pengambilan_di_lapangan"`
	JadwalPembangunanDiLapangan         string `json:"jadwal_pembangunan_di_lapangan"`
	TanggalPemegangDilarangMengambilAir string `json:"tanggal_pemegang_dilarang_mengambil_air"`
	RealisasiDiLapangan                 string `json:"realisasi_di_lapangan"`

	// part 3
	LaporanHasilPemeriksaanTimVerifikasi []model.Attachments `json:"laporan_hasil_pemeriksaan_tim_verifikasi"`
	FileKeberadaanAlatUkurDebit          []model.Attachments `json:"file_keberadaan_alat_ukur_debit"`
	FileKeberadaanSistemTelemetri        []model.Attachments `json:"file_keberadaan_sistem_telemetri"`

	TerdapatAirDibuangKeSumber     bool                `json:"terdapat_air_dibuang_ke_sumber"`
	DebitAirBuangan                float64             `json:"debit_air_buangan"`
	LaporanHasilPemeriksaanBuangan []model.Attachments `json:"laporan_hasil_pemeriksaan_buangan"`

	// part 4
	DokumenBuktiBayar               []model.Attachments `json:"dokumen_bukti_bayar"`
	DokumenKewajibanKeuanganLainnya []model.Attachments `json:"dokumen_kewajiban_keuangan_lainnya"`

	BuktiKerusakanSumberAir []model.Attachments `json:"bukti_kerusakan_sumber_air"`
	PerbaikanKerusakan      bool                `json:"perbaikan_kerusakan"`

	BuktiUsahaPengendalianPencemaran []model.Attachments `json:"bukti_usaha_pengendalian_pencemaran"`
	BentukUsahaPengendalian          string              `json:"bentuk_usaha_pengendalian"`

	BuktiPenggunaanAir   []model.Attachments `json:"bukti_penggunaan_air"`
	DebitAirDihasilkan   float64             `json:"debit_air_dihasilkan"`
	ManfaatUntukKegiatan string              `json:"manfaat_untuk_kegiatan"`

	KegiatanOP               string `json:"kegiatan_op"`
	KoordinasiBBWSRencanaOP  string `json:"koordinasi_bbws_rencana_op"`
	KoordinasiBBWSKonstruksi string `json:"koordinasi_bbws_konstruksi"`

	PLDebitYangDigunakan      float64 `json:"pl_debit_yang_digunakan"`
	PLKualiatsAirDigunakan    string  `json:"pl_kualitas_air_digunakan"`
	PLDebitYangDikembalikan   float64 `json:"pl_debit_yang_dikembalikan"`
	PLKualitasAirDikembalikan string  `json:"pl_kualitas_air_dikembalikan"`
}

type LaporanPerizinanSaveUseCaseRes struct {
	ID model.LaporanPerizinanID `json:"id"`
}

type LaporanPerizinanSaveUseCase = core.ActionHandler[
	LaporanPerizinanSaveUseCaseReq,
	LaporanPerizinanSaveUseCaseRes,
]

func ImplLaporanPerizinanSaveUseCase(
	//
	generateId gateway.GenerateId,
	getOneSK gateway.SKPerizinanGetOne,
	getOneLapor gateway.LaporanPerizinanGetOne,
	saveLaporanPerizinan gateway.LaporanPerizinanSave, //
) LaporanPerizinanSaveUseCase {
	return func(ctx context.Context, req LaporanPerizinanSaveUseCaseReq) (*LaporanPerizinanSaveUseCaseRes, error) {

		if err := req.Validate(); err != nil {
			return nil, err
		}

		resSK, err := getOneSK(ctx, gateway.SKPerizinanGetOneReq{NomorSK: req.NomorSK})
		if err != nil {
			return nil, err
		}

		if resSK.SKPerizinan == nil {
			return nil, fmt.Errorf("SK no '%s' tidak ditemukan", req.NomorSK)
		}

		resLapor, err := getOneLapor(ctx, gateway.LaporanPerizinanGetOneReq{
			NomorSK:               req.NomorSK,
			PeriodePengambilanSDA: req.PeriodePengambilanSDA,
		})
		if err != nil {
			return nil, err
		}

		var laporanPerizinanID model.LaporanPerizinanID
		if resLapor.LaporanPerizinan == nil {
			gen, _ := generateId(ctx, gateway.GenerateIdReq{})
			laporanPerizinanID = model.LaporanPerizinanID(gen.UUID)
		} else {

			if resLapor.LaporanPerizinan.Status == model.StatusLaporSubmitted {
				return nil, fmt.Errorf("laporan periode '%s' sudah di submit", resLapor.LaporanPerizinan.PeriodePengambilanSDA)
			}

			laporanPerizinanID = resLapor.LaporanPerizinan.ID
		}

		if _, err = saveLaporanPerizinan(ctx, gateway.LaporanPerizinanSaveReq{LaporanPerizinan: &model.LaporanPerizinan{
			//
			ID:            laporanPerizinanID,
			SKPerizinanID: resSK.SKPerizinan.ID,

			NoSK:                  req.NomorSK,
			PeriodePengambilanSDA: req.PeriodePengambilanSDA,
			Status:                model.StatusLaporDraft,

			KoordinatDiLapangan: model.NewGeometryPoint(
				req.KoordinatDiLapangan.Longitude,
				req.KoordinatDiLapangan.Latitude,
			),
			DebitPengambilan: req.DebitPengambilan,

			CaraPengambilanSDADiLapangan:        req.CaraPengambilanSDADiLapangan,
			JenisTipeKonstruksiDiLapangan:       req.JenisTipeKonstruksiDiLapangan,
			JadwalPengambilanDiLapangan:         req.JadwalPengambilanDiLapangan,
			JadwalPembangunanDiLapangan:         req.JadwalPembangunanDiLapangan,
			TanggalPemegangDilarangMengambilAir: req.TanggalPemegangDilarangMengambilAir,
			RealisasiDiLapangan:                 req.RealisasiDiLapangan,
			TerdapatAirDibuangKeSumber:          req.TerdapatAirDibuangKeSumber,
			DebitAirBuangan:                     req.DebitAirBuangan,
			PerbaikanKerusakan:                  req.PerbaikanKerusakan,
			BentukUsahaPengendalian:             req.BentukUsahaPengendalian,
			DebitAirDihasilkan:                  req.DebitAirDihasilkan,
			ManfaatUntukKegiatan:                req.ManfaatUntukKegiatan,
			KegiatanOP:                          req.KegiatanOP,
			KoordinasiBBWSRencanaOP:             req.KoordinasiBBWSRencanaOP,
			KoordinasiBBWSKonstruksi:            req.KoordinasiBBWSKonstruksi,
			PLDebitYangDigunakan:                req.PLDebitYangDigunakan,
			PLKualiatsAirDigunakan:              req.PLKualiatsAirDigunakan,
			PLDebitYangDikembalikan:             req.PLDebitYangDikembalikan,
			PLKualitasAirDikembalikan:           req.PLKualitasAirDikembalikan,

			LaporanPengambilanAirBulanan:         helper.ToDataTypeJSON(req.LaporanPengambilanAirBulanan...),
			FileKeberadaanAlatUkurDebit:          helper.ToDataTypeJSON(req.FileKeberadaanAlatUkurDebit...),
			FileKeberadaanSistemTelemetri:        helper.ToDataTypeJSON(req.FileKeberadaanSistemTelemetri...),
			LaporanHasilPemeriksaanTimVerifikasi: helper.ToDataTypeJSON(req.LaporanHasilPemeriksaanTimVerifikasi...),
			LaporanHasilPemeriksaanBuangan:       helper.ToDataTypeJSON(req.LaporanHasilPemeriksaanBuangan...),
			DokumenBuktiBayar:                    helper.ToDataTypeJSON(req.DokumenBuktiBayar...),
			DokumenKewajibanKeuanganLainnya:      helper.ToDataTypeJSON(req.DokumenKewajibanKeuanganLainnya...),
			BuktiKerusakanSumberAir:              helper.ToDataTypeJSON(req.BuktiKerusakanSumberAir...),
			BuktiUsahaPengendalianPencemaran:     helper.ToDataTypeJSON(req.BuktiUsahaPengendalianPencemaran...),
			BuktiPenggunaanAir:                   helper.ToDataTypeJSON(req.BuktiPenggunaanAir...),
		}}); err != nil {
			return nil, err
		}

		return &LaporanPerizinanSaveUseCaseRes{
			ID: laporanPerizinanID,
		}, nil
	}
}

type ErrorField struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

func (x LaporanPerizinanSaveUseCaseReq) Validate() error {

	var errs []ErrorField

	if x.Page != 1 && x.Page != 2 && x.Page != 3 && x.Page != 4 {
		return errors.New("page must 1, 2, 3, 4")
	}

	switch x.Page {
	case 1:
		// Coordinate validation
		if x.KoordinatDiLapangan.Latitude == 0 || x.KoordinatDiLapangan.Longitude == 0 {
			errs = append(errs, ErrorField{
				Message: "Mohon masukkan koordinat yang valid",
				Field:   "koordinat_di_lapangan",
			})
		}

		// if x.CaraPengambilanSDADiLapangan == "" {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan cara pengambilan SDA di lapangan",
		// 		Field:   "cara_pengambilan_sda_di_lapangan",
		// 	})
		// }

		// if x.JenisTipeKonstruksiDiLapangan == "" {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan jenis tipe konstruksi di lapangan",
		// 		Field:   "jenis_tipe_konstruksi_di_lapangan",
		// 	})
		// }

		// if len(x.LaporanPengambilanAirBulanan) == 0 {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan laporan pengambilan air bulanan",
		// 		Field:   "laporan_pengambilan_air_bulanan",
		// 	})
		// }

		// if x.DebitPengambilan <= 0 {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan debit pengambilan yang valid",
		// 		Field:   "debit_pengambilan",
		// 	})
		// }

	case 2:
		// if x.JadwalPembangunanDiLapangan == "" {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan jadwal pembangunan di lapangan",
		// 		Field:   "jadwal_pembangunan_di_lapangan",
		// 	})
		// }

		// if x.JadwalPengambilanDiLapangan == "" {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan jadwal pengambilan di lapangan",
		// 		Field:   "jadwal_pengambilan_di_lapangan",
		// 	})
		// }

		// if x.RealisasiDiLapangan == "" {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan realisasi di lapangan",
		// 		Field:   "realisasi_di_lapangan",
		// 	})
		// }

		// Optional field - no validation needed
		// tanggal_pemegang_dilarang_mengambil_air

	case 3:
		// if x.KeberadaanAlatUkurDebit == false {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan keberadaan alat ukur debit",
		// 		Field:   "keberadaan_alat_ukur_debit",
		// 	})
		// }

		// if x.KeberadaanSistemTelemetri == false {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan keberadaan sistem telemetri",
		// 		Field:   "keberadaan_sistem_telemetri",
		// 	})
		// }

		// if x.DebitAirBuangan <= 0 {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan debit air buangan yang valid",
		// 		Field:   "debit_air_buangan",
		// 	})
		// }

		// if !x.TerdapatAirDibuangKeSumber {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan terdapat air dibuang ke sumber",
		// 		Field:   "terdapat_air_dibuang_ke_sumber",
		// 	})
		// }

		// Optional array fields - no validation needed
		// laporan_hasil_pemeriksaan_buangan
		// laporan_hasil_pemeriksaan_tim_verifikasi

	case 4:
		// Optional array fields - no validation needed
		// dokumen_bukti_bayar
		// dokumen_kewajiban_keuangan_lainnya
		// bukti_kerusakan_sumber_air
		// bukti_usaha_pengendalian_pencemaran
		// bukti_penggunaan_air

		// if !x.PerbaikanKerusakan {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan perbaikan kerusakan",
		// 		Field:   "perbaikan_kerusakan",
		// 	})
		// }

		// if x.BentukUsahaPengendalian == "" {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan bentuk usaha pengendalian",
		// 		Field:   "bentuk_usaha_pengendalian",
		// 	})
		// }

		// if x.KegiatanOP == "" {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan kegiatan operasional",
		// 		Field:   "kegiatan_op",
		// 	})
		// }

		// if x.KoordinasiBBWSRencanaOP == "" {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan koordinasi BBWS rencana operasional",
		// 		Field:   "koordinasi_bbws_rencana_op",
		// 	})
		// }

		// if x.KoordinasiBBWSKonstruksi == "" {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan koordinasi BBWS konstruksi",
		// 		Field:   "koordinasi_bbws_konstruksi",
		// 	})
		// }

		// if x.PLDebitYangDigunakan <= 0 {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan debit yang digunakan",
		// 		Field:   "pl_debit_yang_digunakan",
		// 	})
		// }

		// if x.PLKualiatsAirDigunakan == "" {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan kualitas air yang digunakan",
		// 		Field:   "pl_kualitas_air_digunakan",
		// 	})
		// }

		// if x.PLDebitYangDikembalikan <= 0 {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan debit yang dikembalikan",
		// 		Field:   "pl_debit_yang_dikembalikan",
		// 	})
		// }

		// if x.PLKualitasAirDikembalikan == "" {
		// 	errs = append(errs, ErrorField{
		// 		Message: "Mohon masukkan kualitas air yang dikembalikan",
		// 		Field:   "pl_kualitas_air_dikembalikan",
		// 	})
		// }

		// Optional fields - no validation needed
		// debit_air_dihasilkan
		// manfaat_untuk_kegiatan
		// bentuk_koordinasi_bbws_1
		// bentuk_koordinasi_bbws_2
	}

	if len(errs) > 0 {
		return core.NewErrorWithData(errors.New("ada field yang masih invalid"), errs)
	}

	return nil
}
