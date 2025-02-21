package wiring

import (
	"rmis/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedSpip(db *gorm.DB) error {
	spipNames := []string{
		"Penegakan Integritas dan Nilai Etika",
		"Komitmen Terhadap Kompetensi",
		"Kepemimpinan yang kondusif",
		"Struktur Organisasi Sesuai Kebutuhan",
		"Pendelegasian Wewenang dan Tanggung Jawab yang Tepat",
		"Penyusunan dan Penerapan Kebijakan yang Sehat tentang Pembinaan SDM",
		"Perwujudan Peran APIP yang Efektif",
		"Hubungan Kerja yang Baik dengan Instansi Pemerintah Terkait",
		"Identifikasi Risiko",
		"Analisis Risiko",
		"Reviu Kinerja",
		"Pembinaan Sumber Daya Manusia",
		"Pengendalian atas Pengelolaan Sistem Informasi",
		"Pengendalian Fisik atas Aset",
		"Penetapan dan Reviu Indikator",
		"Pemisahan Fungsi",
		"Otorisasi Transaksi dan Kejadian Penting",
		"Pencatatan yang Akurat dan Tepat Waktu",
		"Pembatasan Akses atas Sumber Daya dan Catatan",
		"Akuntabilitas Pencatatan dan Sumber Daya",
		"Dokumentasi yang baik atas Sistem Pengendalian Intern (SPI) serta transaksi dan kejadian penting",
		"Informasi",
		"Penyelenggaraan Komunikasi yang Efektif",
		"Pemantauan Berkelanjutan",
		"Evaluasi Terpisah",
	}

	for _, name := range spipNames {
		spip := model.SPIP{
			ID:   stringPtr(uuid.New().String()),
			Nama: stringPtr(name),
		}
		if err := db.Create(&spip).Error; err != nil {
			return err
		}
	}
	return nil
}

func stringPtr(s string) *string {
	return &s
}

func SeedKategoriRisiko(db *gorm.DB) error {
	// Define the data: key = name, value = kode.
	kategoriData := map[string]string{
		"Risiko Pendapatan":  "01",
		"Risiko Belanja":     "02",
		"Risiko Pembiayaan":  "03",
		"Risiko Strategis":   "04",
		"Risiko Fraud":       "05",
		"Risiko Kepatuhan":   "06",
		"Risiko Operasional": "07",
		"Risiko Reputasi":    "08",
	}

	// Iterate over the map and create each record.
	for nama, kode := range kategoriData {
		kategori := model.KategoriRisiko{
			ID:   stringPtr(uuid.New().String()),
			Nama: stringPtr(nama),
			Kode: stringPtr(kode),
		}
		if err := db.Create(&kategori).Error; err != nil {
			return err
		}
	}
	return nil
}

func SeedPenyebabRisiko(db *gorm.DB) error {
	// Define the list of Penyebab Risiko names.
	penyebabRisikoList := []string{
		"Man",
		"Money",
		"Method",
		"Machine",
		"Material",
		"Eksternal",
	}

	// Iterate over the list and create each record.
	for _, nama := range penyebabRisikoList {
		penyebab := model.PenyebabRisiko{
			ID:   stringPtr(uuid.New().String()),
			Nama: stringPtr(nama),
		}
		if err := db.Create(&penyebab).Error; err != nil {
			return err
		}
	}
	return nil
}

func SeedOpd(db *gorm.DB) error {
	// Define a slice of tuples with Name and Kode.
	opdData := []struct {
		Name string
		Kode string
	}{
		// First group
		{"DINAS PENDIDIKAN DAN KEBUDAYAAN", "DISDIKBUD"},
		{"DINAS KESEHATAN", "DINKES"},
		{"DINAS PEKERJAAN UMUM DAN PENATAAN RUANG", "PUPR"},
		{"DINAS PERUMAHAN, KAWASAN PERMUKIMAN DAN PERTANAHAN", "DPKPP"},
		{"DINAS SATUAN POLISI PAMONG PRAJA DAN PEMADAM KEBAKARAN", "SATPOLPPDAMKAR"},
		{"BADAN PENANGGULANGAN BENCANA DAERAH", "BPBD"},
		{"DINAS SOSIAL", "DINSOS"},

		// Second group
		{"DINAS TENAGA KERJA, TRANSMIGRASI, ENERGI DAN SUMBER DAYA MINERAL", "DISNAKERTRANS"},
		{"DINAS KETAHANAN PANGAN", "DISKETPAN"},
		{"DINAS LINGKUNGAN HIDUP", "DLH"},
		{"DINAS KEPENDUDUKAN DAN PENCATATAN SIPIL", "DISDUKCAPIL"},
		{"DINAS PEMBERDAYAAN MASYARAKAT DAN DESA", "DISPEMDES"},
		{"DINAS PENGENDALIAN PENDUDUK, KB, PEMBERDAYAAN PEREMPUAN DAN PERLINDUNGAN ANAK", "DPKBP3A"},
		{"DINAS PERHUBUNGAN", "DISHUB"},
		{"DINAS KOMUNIKASI DAN INFORMATIKA", "DISKOMINFO"},
		{"DINAS KOPERASI, UMKM, PERDAGANGAN DAN PERINDUSTRIAN", "DKUPP"},
		{"DINAS PENANAMAN MODAL DAN PELAYANAN TERPADU SATU PINTU", "DPMPTSP"},
		{"DINAS KEARSIPAN DAN PERPUSTAKAAN", "DISARPUS"},
		{"DINAS PERIKANAN", "DISPERIKAN"},
		{"DINAS PARIWISATA PEMUDA DAN OLAHRAGA", "DISPARPORA"},
		{"DINAS PERTANIAN", "DISPERTAN"},
		{"DINAS PETERNAKAN DAN KESEHATAN HEWAN", "DISNAKESWAN"},
		{"BADAN PERENCANAAN PEMBANGUNAN, PENELITIAN DAN PENGEMBANGAN DAERAH", "BP4D"},
		{"BADAN KEUANGAN DAN ASET DAERAH", "BKAD"},
		{"BADAN PENDAPATAN DAERAH", "BAPENDA"},
		{"BADAN KEPEGAWAIAN DAN PENGEMBANGAN SDM", "BKPSDM"},
		{"BADAN KESATUAN BANGSA DAN POLITIK", "BAKESBANGPOL"},
		{"SEKRETARIAT DAERAH", "SETDA"},
		{"SEKRETARIAT DPRD", "SETWAN"},
		{"INSPEKTORAT", "IRDA"},

		// Third group (Kecamatan)
		{"KECAMATAN SUBANG", "Kec1"},
		{"KECAMATAN CIBOGO", "Kec2"},
		{"KECAMATAN CIJAMBE", "Kec3"},
		{"KECAMATAN JALANCAGAK", "Kec4"},
		{"KECAMATAN SAGALAHERANG", "Kec5"},
		{"KECAMATAN CISALAK", "Kec6"},
		{"KECAMATAN TANJUNGSIANG", "Kec7"},
		{"KECAMATAN PAGADEN", "Kec8"},
		{"KECAMATAN BINONG", "Kec9"},
		{"KECAMATAN PAMANUKAN", "Kec10"},
		{"KECAMATAN LEGONKULON", "Kec11"},
		{"KECAMATAN CIPUNAGARA", "Kec12"},
		{"KECAMATAN COMPRENG", "Kec13"},
		{"KECAMATAN PUSAKANAGARA", "Kec14"},
		{"KECAMATAN CIASEM", "Kec15"},
		{"KECAMATAN BLANAKAN", "Kec16"},
		{"KECAMATAN PATOKBEUSI", "Kec17"},
		{"KECAMATAN PABUARAN", "Kec18"},
		{"KECAMATAN CIPEUNDEUY", "Kec19"},
		{"KECAMATAN PURWADADI", "Kec20"},
		{"KECAMATAN KALIJATI", "Kec21"},
		{"KECAMATAN CIKAUM", "Kec22"},
		{"KECAMATAN SERANGPANJANG", "Kec23"},
		{"KECAMATAN SUKASARI", "Kec24"},
		{"KECAMATAN TAMBAKDAHAN", "Kec25"},
		{"KECAMATAN KASOMALANG", "Kec26"},
		{"KECAMATAN DAWUAN", "Kec27"},
		{"KECAMATAN PAGADEN BARAT", "Kec28"},
		{"KECAMATAN CIATER", "Kec29"},
		{"KECAMATAN PUSAKAJAYA", "Kec30"},
	}

	// Iterate over the data slice and insert each record.
	for _, opd := range opdData {
		newOpd := model.OPD{
			ID:   stringPtr(uuid.New().String()),
			Nama: stringPtr(opd.Name),
			Kode: stringPtr(opd.Kode),
		}

		if err := db.Create(&newOpd).Error; err != nil {
			return err
		}
	}

	return nil
}
