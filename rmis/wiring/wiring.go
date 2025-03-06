package wiring

import (
	"net/http"
	"rmis/controller"
	"rmis/gateway"
	"rmis/usecase"
	"shared/helper"
	"shared/helper/cronjob"
	"shared/middleware"

	"gorm.io/gorm"
)

func SetupDependency(mariaDB *gorm.DB, mux *http.ServeMux, jwtToken helper.JWTTokenizer, printer *helper.ApiPrinter, cj *cronjob.CronJob, sseDashboard *helper.SSE) {
	// =================================================================
	// Gateway
	// =================================================================
	exampleGetAllGateway := gateway.ImplExampleGateway(mariaDB)
	generateIdGateway := gateway.ImplGenerateId()

	// Gateway SPIP
	spipGetAllGateway := gateway.ImplSpipGetAll(mariaDB)
	spipGetOneGateway := gateway.ImplSpipGetByID(mariaDB)
	spipDeleteGateway := gateway.ImplSpipDelete(mariaDB)
	spipCreateGateway := gateway.ImplSpipSave(mariaDB)

	// Gateway Kategori Risiko
	kategoriRisikoGetAllGateway := gateway.ImplKategoriRisikoGetAll(mariaDB)
	kategoriRisikoGetOneGateway := gateway.ImplKategoriRisikoGetByID(mariaDB)
	kategoriRisikoDeleteGateway := gateway.ImplKategoriRisikoDelete(mariaDB)
	kategoriRisikoCreateGateway := gateway.ImplKategoriRisikoSave(mariaDB)

	// Gateway Rekapitulasi Hasil Kuesioner
	rekapitulasiHasilKuesionerCreateGateway := gateway.ImplRekapitulasiHasilKuesionerSave(mariaDB)
	rekapitulasiHasilKuesionerGetAllGateway := gateway.ImplRekapitulasiHasilkuesionerGetAll(mariaDB)
	rekapitulasiHasilKuesionerGetOneGateway := gateway.ImplRekapitulasiHasilKuesionerGetByID(mariaDB)
	rekapitulasiHasilKuesionerDeleteGateway := gateway.ImplRekapitulasiHasilKuesionerDelete(mariaDB)

	// Gateway Penetapan Konteks Risiko Strategis Pemda
	penetapanKonteksRisikoStrategisPemdaGetAllGateway := gateway.ImplPenetapanKonteksRisikoStrategisPemdaGetAll(mariaDB)
	penetapanKonteksRisikoStrategisPemdaGetOneGateway := gateway.ImplPenetapanKonteksRisikoStrategisPemdaGetByID(mariaDB)
	penetapanKonteksRisikoStrategisPemdaDeleteGateway := gateway.ImplPenetepanKonteksRisikoStrategisPemdaDelete(mariaDB)
	penetapanKonteksRisikoStrategisPemdaCreateGateway := gateway.ImplPenetepanKonteksRisikoStrategisPemdaSave(mariaDB)

	// Gateway Root Cause Analysis (RCA)
	rcaGetAllGateway := gateway.ImplRcaGetAll(mariaDB)
	rcaGetOneGateway := gateway.ImplRcaGetByID(mariaDB)
	rcaDeleteGateway := gateway.ImplRcaDelete(mariaDB)
	rcaCreateGateway := gateway.ImplRcaSave(mariaDB)

	// Gateway Identifikasi Risiko Strategis Pemda
	identifikasiRisikoStrategisPemdaGetAllGateway := gateway.ImplIdentifikasiRisikoStrategisPemdaGetAll(mariaDB)
	identifikasiRisikoStrategisPemdaGetOneGateway := gateway.ImplIdentifikasiRisikoStrategisPemdaGetByID(mariaDB)
	identifikasiRisikoStrategisPemdaDeleteGateway := gateway.ImplIdentifikasiRisikoStrategisPemdaDelete(mariaDB)
	identifikasiRisikoStrategisPemdaCreateGateway := gateway.ImplIdentifikasiRisikoStrategisPemdaSave(mariaDB)

	// Gateway Identifikasi Risiko Strategis OPD
	identifikasiRisikoStrategisOPDGetAllGateway := gateway.ImplIdentifikasiRisikoStrategisOPDGetAll(mariaDB)
	identifikasiRisikoStrategisOPDGetOneGateway := gateway.ImplIdentifikasiRisikoStrategisOPDGetByID(mariaDB)
	identifikasiRisikoStrategisOPDDeleteGateway := gateway.ImplIdentifikasiRisikoStrategisOPDDelete(mariaDB)
	identifikasiRisikoStrategisOPDCreateGateway := gateway.ImplIdentifikasiRisikoStrategisOPDSave(mariaDB)

	// Gateway Identifikasi Risiko Operasional OPD
	identifikasiRisikoOperasionalOPDGetAllGateway := gateway.ImplIdentifikasiRisikoOperasionalOPDGetAll(mariaDB)
	identifikasiRisikoOperasionalOPDGetOneGateway := gateway.ImplIdentifikasiRisikoOperasionalOPDGetByID(mariaDB)
	identifikasiRisikoOperasionalOPDDeleteGateway := gateway.ImplIdentifikasiRisikoOperasionalOPDDelete(mariaDB)
	identifikasiRisikoOperasionalOPDCreateGateway := gateway.ImplIdentifikasiRisikoOperasionalOPDSave(mariaDB)

	// Gateway Simpulan Kondisi Kelemahan Lingkungan
	simpulanKondisiKelemahanLingkunganGetAllGateway := gateway.ImplSimpulanKondisiKelemahanLingkunganGetAll(mariaDB)
	simpulanKondisiKelemahanLingkunganGetOneGateway := gateway.ImplSimpulanKondisiKelemahanLingkunganGetByID(mariaDB)
	simpulanKondisiKelemahanLingkunganDeleteGateway := gateway.ImplSimpulanKondisiKelemahanLingkunganDelete(mariaDB)
	simpulanKondisiKelemahanLingkunganCreateGateway := gateway.ImplSimpulanKondisiKelemahanLingkunganSave(mariaDB)

	// Gateway IKU
	ikuGetAllGateway := gateway.ImplIKUGetAll(mariaDB)
	ikuGetOneGateway := gateway.ImplIKUGetByID(mariaDB)
	ikuDeleteGateway := gateway.ImplIKUDelete(mariaDB)
	ikuCreateGateway := gateway.ImplIKUSave(mariaDB)
	// Gateway Penyebab Risiko
	penyebabRisikoGetAllGateway := gateway.ImplPenyebabRisikoGetAll(mariaDB)
	penyebabRisikoGetOneGateway := gateway.ImplPenyebabRisikoGetByID(mariaDB)
	penyebabRisikoDeleteGateway := gateway.ImplPenyebabRisikoDelete(mariaDB)
	penyebabRisikoCreateGateway := gateway.ImplPenyebabRisikoSave(mariaDB)

	// Gateway Hasil Analisis Risiko
	hasilAnalisisRisikoCreateGateway := gateway.ImplHasilAnalisisRisikoSave(mariaDB)
	hasilAnalisisRisikoGetAllGateway := gateway.ImplHasilAnalisisRisikoGetAll(mariaDB)
	hasilAnalisisRisikoGetOneGateway := gateway.ImplHasilAnalisisRisikoGetByID(mariaDB)
	hasilAnalisisRisikoDeleteGateway := gateway.ImplHasilAnalisisRisikoDelete(mariaDB)

	// Gateway Penilaian Kegiatan Pengendalian
	penilaianKegiatanPengendalianCreateGateway := gateway.ImplPenilaianKegiatanPengendalianSave(mariaDB)
	penilaianKegiatanPengendalianGetAllGateway := gateway.ImplPenilaianKegiatanPengendalianGetAll(mariaDB)
	penilaianKegiatanPengendalianGetOneGateway := gateway.ImplPenilaianKegiatanPengendalianGetByID(mariaDB)
	penilaianKegiatanPengendalianDeleteGateway := gateway.ImplPenilaianKegiatanPengendalianDelete(mariaDB)

	// Gateway Penetapan Konteks Risiko Strategis Renstra OPD
	penetapanKonteksRisikoStrategisRenstraOPDGetAllGateway := gateway.ImplPenetapanKonteksRisikoStrategisRenstraOPDGetAll(mariaDB)
	penetapanKonteksRisikoStrategisRenstraOPDGetOneGateway := gateway.ImplPenetapanKonteksRisikoStrategisRenstraOPDGetByID(mariaDB)
	penetapanKonteksRisikoStrategisRenstraOPDDeleteGateway := gateway.ImplPenetepanKonteksRisikoStrategisRenstraOPDDelete(mariaDB)
	penetapanKonteksRisikoStrategisRenstraOPDCreateGateway := gateway.ImplPenetepanKonteksRisikoStrategisRenstraOPDSave(mariaDB)

	// Gateway OPD
	opdGetOneGateway := gateway.ImplOPDGetByID(mariaDB)
	// Gateway Daftar Risiko Prioritas
	daftarRisikoPrioritasCreateGateway := gateway.ImplDaftarRisikoPrioritasSave(mariaDB)
	daftarRisikoPrioritasGetAllGateway := gateway.ImplDaftarRisikoPrioritasGetAll(mariaDB)
	daftarRisikoPrioritasGetOneGateway := gateway.ImplDaftarRisikoPrioritasGetByID(mariaDB)
	daftarRisikoPrioritasDeleteGateway := gateway.ImplDaftarRisikoPrioritasDelete(mariaDB)

	// Gateway Penetapan Konteks Risiko Operasional
	penetapanKonteksRisikoOperasionalGetAllGateway := gateway.ImplPenetapanKonteksRisikoOperasionalGetAll(mariaDB)
	penetapanKonteksRisikoOperasionalGetOneGateway := gateway.ImplPenetapanKonteksRisikoOperasionalGetByID(mariaDB)
	penetapanKonteksRisikoOperasionalDeleteGateway := gateway.ImplPenetepanKonteksRisikoOperasionalDelete(mariaDB)
	penetapanKonteksRisikoOperasionalCreateGateway := gateway.ImplPenetepanKonteksRisikoOperasionalSave(mariaDB)
	// Gateway Penilaian Risiko
	penilaianRisikoCreateGateway := gateway.ImplPenilaianRisikoSave(mariaDB)
	penilaianRisikoGetAllGateway := gateway.ImplPenilaianRisikoGetAll(mariaDB)
	penilaianRisikoGetOneGateway := gateway.ImplPenilaianRisikoGetByID(mariaDB)
	penilaianRisikoDeleteGateway := gateway.ImplPenilaianRisikoDelete(mariaDB)

	// Gateway Kriteria Dampak
	kriteriaDampakGetAllGateway := gateway.ImplKriteriaDampakGetAll(mariaDB)
	kriteriaDampakGetOneGateway := gateway.ImplKriteriaDampakGetByID(mariaDB)
	// Gateway Pengkomunikasian Pengendalian
	pengkomunikasianPengendalianCreateGateway := gateway.ImplPengkomunikasianPengendalianSave(mariaDB)
	pengkomunikasianPengendalianGetAllGateway := gateway.ImplPengkomunikasianPengendalianGetAll(mariaDB)
	pengkomunikasianPengendalianGetOneGateway := gateway.ImplPengkomunikasianPengendalianGetByID(mariaDB)
	pengkomunikasianPengendalianDeleteGateway := gateway.ImplPengkomunikasianPengendalianDelete(mariaDB)

	// Gateway Rancangan Pemantauan
	rancanganPemantauanCreateGateway := gateway.ImplRancanganPemantauanSave(mariaDB)
	rancanganPemantauanGetAllGateway := gateway.ImplRancanganPemantauanGetAll(mariaDB)
	rancanganPemantauanGetOneGateway := gateway.ImplRancanganPemantauanGetByID(mariaDB)
	rancanganPemantauanDeleteGateway := gateway.ImplRancanganPemantauanDelete(mariaDB)

	// Gateway Pencatatan Kejadian Risiko
	pencatatanKejadianRisikoCreateGateway := gateway.ImplPencatatanKejadianRisikoSave(mariaDB)
	pencatatanKejadianRisikoGetAllGateway := gateway.ImplPencatatanKejadianRisikoGetAll(mariaDB)
	pencatatanKejadianRisikoGetOneGateway := gateway.ImplPencatatanKejadianRisikoGetByID(mariaDB)
	pencatatanKejadianRisikoDeleteGateway := gateway.ImplPencatatanKejadianRisikoDelete(mariaDB)

	// Gateway Indeks Peringkat Prioritas
	indeksPeringkatPrioritasCreateGateway := gateway.ImplIndeksPeringkatPrioritasSave(mariaDB)
	indeksPeringkatPrioritasGetAllGateway := gateway.ImplIndeksPeringkatPrioritasGetAll(mariaDB)
	indeksPeringkatPrioritasGetOneGateway := gateway.ImplIndeksPeringkatPrioritasGetByID(mariaDB)
	// indeksPeringkatPrioritasDeleteGateway := gateway.ImplIndeksPeringkatPrioritasDelete(mariaDB)

	// =================================================================
	// Usecase
	// =================================================================
	exampleGetAllUsecase := usecase.ImplExampleGetAllUseCase(exampleGetAllGateway)

	// Usecase SPIP
	spipGetAllUseCase := usecase.ImplSpipGetAllUseCase(spipGetAllGateway)
	spipGetOneUseCase := usecase.ImplSpipGetByIDUseCase(spipGetOneGateway)
	spipDeleteUseCase := usecase.ImplSpipDeleteUseCase(spipDeleteGateway)
	SpipCreateUseCase := usecase.ImplSpipCreateUseCase(generateIdGateway, spipCreateGateway)
	spipUpdateUseCase := usecase.ImplSpipUpdateUseCase(spipGetOneGateway, spipCreateGateway)

	// Usecase Kategori Risiko
	kategoriRisikoGetAllUseCase := usecase.ImplKategoriRisikoGetAllUseCase(kategoriRisikoGetAllGateway)
	kategoriRisikoGetOneUseCase := usecase.ImplKategoriRisikoGetByIDUseCase(kategoriRisikoGetOneGateway)
	kategoriRisikoDeleteUseCase := usecase.ImplKategoriRisikoDeleteUseCase(kategoriRisikoDeleteGateway)
	kategoriRisikoCreateUseCase := usecase.ImplKategoriRisikoCreateUseCase(generateIdGateway, kategoriRisikoCreateGateway)
	kategoriRisikoUpdateUseCase := usecase.ImplKategoriRisikoUpdateUseCase(kategoriRisikoGetOneGateway, kategoriRisikoCreateGateway)

	// Usecase Rekapitulasi Hasil Kuesioner
	rekapitulasiHasilKuesionerCreateUseCase := usecase.ImplRekapitulasiHasilKuesionerCreateUseCase(generateIdGateway, rekapitulasiHasilKuesionerCreateGateway, spipGetOneGateway)
	rekapitulasiHasilKuesionerGetAllUseCase := usecase.ImplRekapitulasiHasilKuesionerGetAllUseCase(rekapitulasiHasilKuesionerGetAllGateway)
	rekapitulasiHasilKuesionerGetOneUseCase := usecase.ImplRekapitulasiHasilKuesionerGetByIDUseCase(rekapitulasiHasilKuesionerGetOneGateway)
	rekapitulasiHasilKuesionerDeleteUseCase := usecase.ImplRekapitulasiHasilKuesionerDeleteUseCase(rekapitulasiHasilKuesionerDeleteGateway)

	// Usecase Penetapan Konteks Risiko Strategis Pemda
	penetapanKonteksRisikoStrategisPemdaGetAllUseCase := usecase.ImplPenetapanKonteksRisikoGetAllUseCase(penetapanKonteksRisikoStrategisPemdaGetAllGateway, ikuGetAllGateway)
	penetapanKonteksRisikoStrategisPemdaGetOneUseCase := usecase.ImplPenetapanKonteksRisikoGetByIDUseCase(penetapanKonteksRisikoStrategisPemdaGetOneGateway, ikuGetAllGateway)
	penetapanKonteksRisikoStrategisPemdaDeleteUseCase := usecase.ImplPenetapanKonteksRisikoDeleteUseCase(penetapanKonteksRisikoStrategisPemdaDeleteGateway)
	penetapanKonteksRisikoStrategisPemdaCreateUseCase := usecase.ImplPenetapanKonteksRisikoStrategisPemdaCreateUseCase(generateIdGateway, penetapanKonteksRisikoStrategisPemdaCreateGateway)
	penetapanKonteksRisikoStrategisPemdaUpdateUseCase := usecase.ImplPenetapanKonteksRisikoStrategisPemdaUpdateUseCase(penetapanKonteksRisikoStrategisPemdaGetOneGateway, penetapanKonteksRisikoStrategisPemdaCreateGateway)

	// Usecase Root Cause Analysis (RCA)
	rcaGetAllUseCase := usecase.ImplRcaGetAllUseCase(rcaGetAllGateway)
	rcaGetOneUseCase := usecase.ImplRcaGetByIDUseCase(rcaGetOneGateway)
	rcaDeleteUseCase := usecase.ImplRcaDeleteUseCase(rcaDeleteGateway)
	rcaCreateUseCase := usecase.ImplRcaCreateUseCase(generateIdGateway, rcaCreateGateway, identifikasiRisikoStrategisPemdaGetOneGateway, penyebabRisikoGetOneGateway)
	rcaUpdateUseCase := usecase.ImplRcaUpdateUseCase(rcaGetOneGateway, rcaCreateGateway, identifikasiRisikoStrategisPemdaGetOneGateway, penyebabRisikoGetOneGateway)

	// Usecase Identifikasi Risiko Strategis Pemda
	identifikasiRisikoStrategisPemdaGetAllUseCase := usecase.ImplIdentifikasiRisikoStrategisPemdaGetAllUseCase(identifikasiRisikoStrategisPemdaGetAllGateway)
	identifikasiRisikoStrategisPemdaGetOneUseCase := usecase.ImplIdentifikasiRisikoStrategisPemdaGetByIDUseCase(identifikasiRisikoStrategisPemdaGetOneGateway)
	identifikasiRisikoStrategisPemdaDeleteUseCase := usecase.ImplIdentifikasiRisikoStrategisPemdaDeleteUseCase(identifikasiRisikoStrategisPemdaDeleteGateway)
	identifikasiRisikoStrategisPemdaCreateUseCase := usecase.ImplIdentifikasiRisikoStrategisPemdaCreateUseCase(generateIdGateway, identifikasiRisikoStrategisPemdaCreateGateway, kategoriRisikoGetOneGateway, penetapanKonteksRisikoStrategisPemdaGetOneGateway)
	identifikasiRisikoStrategisPemdaUpdateUseCase := usecase.ImplIdentifikasiRisikoStrategisPemdaUpdateUseCase(identifikasiRisikoStrategisPemdaGetOneGateway, identifikasiRisikoStrategisPemdaCreateGateway, kategoriRisikoGetOneGateway, rcaGetOneGateway, penetapanKonteksRisikoStrategisPemdaGetOneGateway)

	// Usecase Identifikasi Risiko Strategis OPD
	identifikasiRisikoStrategisOPDGetAllUseCase := usecase.ImplIdentifikasiRisikoStrategisOPDGetAllUseCase(identifikasiRisikoStrategisOPDGetAllGateway)
	identifikasiRisikoStrategisOPDGetOneUseCase := usecase.ImplIdentifikasiRisikoStrategisOPDGetByIDUseCase(identifikasiRisikoStrategisOPDGetOneGateway, opdGetOneGateway)
	identifikasiRisikoStrategisOPDDeleteUseCase := usecase.ImplIdentifikasiRisikoStrategisOPDDeleteUseCase(identifikasiRisikoStrategisOPDDeleteGateway)
	identifikasiRisikoStrategisOPDCreateUseCase := usecase.ImplIdentifikasiRisikoStrategisOPDCreateUseCase(generateIdGateway, identifikasiRisikoStrategisOPDCreateGateway, kategoriRisikoGetOneGateway, penetapanKonteksRisikoStrategisRenstraOPDGetOneGateway, opdGetOneGateway)
	identifikasiRisikoStrategisOPDUpdateUseCase := usecase.ImplIdentifikasiRisikoStrategisOPDUpdateUseCase(identifikasiRisikoStrategisOPDGetOneGateway, identifikasiRisikoStrategisOPDCreateGateway, kategoriRisikoGetOneGateway, rcaGetOneGateway, opdGetOneGateway, penetapanKonteksRisikoStrategisRenstraOPDGetOneGateway)
	identifikasiRisikoStrategisOPDApprovalUsecase := usecase.ImplIdentifikasiRisikoStrategisOPDApprovalUseCase(identifikasiRisikoStrategisOPDGetOneGateway, identifikasiRisikoStrategisOPDCreateGateway)

	// Usecase Identifikasi Risiko Operasional OPD
	identifikasiRisikoOperasionalOPDGetAllUseCase := usecase.ImplIdentifikasiRisikoOperasionalOPDGetAllUseCase(identifikasiRisikoOperasionalOPDGetAllGateway, opdGetOneGateway)
	identifikasiRisikoOperasionalOPDGetOneUseCase := usecase.ImplIdentifikasiRisikoOperasionalOPDGetByIDUseCase(identifikasiRisikoOperasionalOPDGetOneGateway, opdGetOneGateway)
	identifikasiRisikoOperasionalOPDDeleteUseCase := usecase.ImplIdentifikasiRisikoOperasionalOPDDeleteUseCase(identifikasiRisikoOperasionalOPDDeleteGateway)
	identifikasiRisikoOperasionalOPDCreateUseCase := usecase.ImplIdentifikasiRisikoOperasionalOPDCreateUseCase(generateIdGateway, identifikasiRisikoOperasionalOPDCreateGateway, kategoriRisikoGetOneGateway, opdGetOneGateway)
	identifikasiRisikoOperasionalOPDUpdateUseCase := usecase.ImplIdentifikasiRisikoOperasionalOPDUpdateUseCase(identifikasiRisikoOperasionalOPDGetOneGateway, identifikasiRisikoOperasionalOPDCreateGateway, kategoriRisikoGetOneGateway, rcaGetOneGateway, opdGetOneGateway)

	// Usecase Simpulan Kondisi Kelemahan Lingkungan
	simpulanKondisiKelemahanLingkunganGetAllUseCase := usecase.ImplSimpulanKondisiKelemahanLingkunganGetAllUseCase(simpulanKondisiKelemahanLingkunganGetAllGateway)
	simpulanKondisiKelemahanLingkunganGetOneUseCase := usecase.ImplSimpulanKondisiKelemahanLingkunganGetByIDUseCase(simpulanKondisiKelemahanLingkunganGetOneGateway)
	simpulanKondisiKelemahanLingkunganDeleteUseCase := usecase.ImplSimpulanKondisiKelemahanLingkunganDeleteUseCase(simpulanKondisiKelemahanLingkunganDeleteGateway)
	simpulanKondisiKelemahanLingkunganCreateUseCase := usecase.ImplSimpulanKondisiKelemahanLingkunganCreateUseCase(generateIdGateway, simpulanKondisiKelemahanLingkunganCreateGateway)
	simpulanKondisiKelemahanLingkunganUpdateUseCase := usecase.ImplSimpulanKondisiKelemahanLingkunganUpdateUseCase(simpulanKondisiKelemahanLingkunganGetOneGateway, simpulanKondisiKelemahanLingkunganCreateGateway)

	// Usecase IKU
	ikuGetAllUseCase := usecase.ImplIKUGetAllUseCase(ikuGetAllGateway)
	ikuGetOneUseCase := usecase.ImplIKUGetByIDUseCase(ikuGetOneGateway)
	ikuDeleteUseCase := usecase.ImplIKUDeleteUseCase(ikuDeleteGateway)
	ikuCreateUseCase := usecase.ImplIKUCreateUseCase(generateIdGateway, ikuCreateGateway)
	ikuUpdateUseCase := usecase.ImplIKUUpdateUseCase(ikuGetOneGateway, ikuCreateGateway)
	// Usecase Penyebab Risiko
	penyebabRisikoGetAllUseCase := usecase.ImplPenyebabRisikoGetAllUseCase(penyebabRisikoGetAllGateway)
	penyebabRisikoGetOneUseCase := usecase.ImplPenyebabRisikoGetByIDUseCase(penyebabRisikoGetOneGateway)
	penyebabRisikoDeleteUseCase := usecase.ImplPenyebabRisikoDeleteUseCase(penyebabRisikoDeleteGateway)
	penyebabRisikoCreateUseCase := usecase.ImplPenyebabRisikoCreateUseCase(generateIdGateway, penyebabRisikoCreateGateway)
	penyebabRisikoUpdateUseCase := usecase.ImplPenyebabRisikoUpdateUseCase(penyebabRisikoGetOneGateway, penyebabRisikoCreateGateway)

	// usecase Hasil Analisis Risiko
	hasilAnalisisRisikoCreateUseCase := middleware.TransactionMiddleware(usecase.ImplHasilAnalisisRisikoCreateUseCase(generateIdGateway, hasilAnalisisRisikoCreateGateway, identifikasiRisikoStrategisPemdaGetOneGateway, indeksPeringkatPrioritasCreateGateway, penetapanKonteksRisikoStrategisPemdaGetOneGateway, kategoriRisikoGetOneGateway), mariaDB)
	// hasilAnalisisRisikoCreateUseCase := usecase.ImplHasilAnalisisRisikoCreateUseCase(generateIdGateway, hasilAnalisisRisikoCreateGateway, identifikasiRisikoStrategisPemdaGetOneGateway, indeksPeringkatPrioritasCreateGateway)
	hasilAnalisisRisikoGetAllUseCase := usecase.ImplHasilAnalisisRisikoGetAllUseCase(hasilAnalisisRisikoGetAllGateway)
	hasilAnalisisRisikoGetOneUseCase := usecase.ImplHasilAnalisisRisikoGetByIDUseCase(hasilAnalisisRisikoGetOneGateway)
	hasilAnalisisRisikoDeleteUsecase := usecase.ImplHasilAnalisisRisikoDeleteUseCase(hasilAnalisisRisikoDeleteGateway)
	hasilAnalisisRisikoUpdateUsecase := middleware.TransactionMiddleware(usecase.ImplHasilAnalisisRisikoUpdateUseCase(hasilAnalisisRisikoGetOneGateway, hasilAnalisisRisikoCreateGateway, indeksPeringkatPrioritasGetOneGateway, indeksPeringkatPrioritasCreateGateway, identifikasiRisikoStrategisPemdaGetOneGateway, penetapanKonteksRisikoOperasionalGetOneGateway, kategoriRisikoGetOneGateway), mariaDB)

	// Usecase Penilaian Kegiatan Pengendalian
	penilaianKegiatanPengendalianCreateUseCase := usecase.ImplPenilaianKegiatanPengendalianCreateUseCase(generateIdGateway, penilaianKegiatanPengendalianCreateGateway, spipGetOneGateway)
	penilaianKegiatanPengendalianGetAllUseCase := usecase.ImplPenilaianKegiatanPengendalianGetAllUseCase(penilaianKegiatanPengendalianGetAllGateway)
	penilaianKegiatanPengendalianGetOneUseCase := usecase.ImplPenilaianKegiatanPengendalianGetByIDUseCase(penilaianKegiatanPengendalianGetOneGateway)
	penilaianKegiatanPengendalianDeleteUseCase := usecase.ImplPenilaianKegiatanPengendalianDeleteUseCase(penilaianKegiatanPengendalianDeleteGateway)
	penilaianKegiatanPengendalianUpdateUseCase := usecase.ImplPenilaianKegiatanPengendalianUpdateUseCase(penilaianKegiatanPengendalianGetOneGateway, penilaianKegiatanPengendalianCreateGateway, spipGetOneGateway)

	// Usecase Penetapan Konteks Risiko Strategis Renstra OPD
	penetapanKonteksRisikoStrategisRenstraOPDGetAllUseCase := usecase.ImplPenetapanKonteksRisikoRenstraOPDGetAllUseCase(penetapanKonteksRisikoStrategisRenstraOPDGetAllGateway, ikuGetAllGateway, opdGetOneGateway)
	penetapanKonteksRisikoStrategisRenstraOPDGetOneUseCase := usecase.ImplPenetapanKonteksRisikoRenstraOPDGetByIDUseCase(penetapanKonteksRisikoStrategisRenstraOPDGetOneGateway, ikuGetAllGateway, opdGetOneGateway)
	penetapanKonteksRisikoStrategisRenstraOPDDeleteUseCase := usecase.ImplPenetapanKonteksRisikoRenstraOPDDeleteUseCase(penetapanKonteksRisikoStrategisRenstraOPDDeleteGateway)
	penetapanKonteksRisikoStrategisRenstraOPDCreateUseCase := usecase.ImplPenetapanKonteksRisikoStrategisRenstraOPDCreateUseCase(generateIdGateway, penetapanKonteksRisikoStrategisRenstraOPDCreateGateway, opdGetOneGateway)
	penetapanKonteksRisikoStrategisRenstraOPDUpdateUseCase := usecase.ImplPenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCase(penetapanKonteksRisikoStrategisRenstraOPDGetOneGateway, penetapanKonteksRisikoStrategisRenstraOPDCreateGateway, opdGetOneGateway)
	// Usecase Daftar Risiko Prioritas
	daftarRisikoPrioritasCreateUseCase := usecase.ImplDaftarRisikoPrioritasCreateUseCase(generateIdGateway, daftarRisikoPrioritasCreateGateway, hasilAnalisisRisikoGetOneGateway, identifikasiRisikoStrategisPemdaGetOneGateway, penetapanKonteksRisikoStrategisPemdaGetOneGateway)
	daftarRisikoPrioritasGetAllUseCase := usecase.ImplDaftarRisikoPrioritasGetAllUseCase(daftarRisikoPrioritasGetAllGateway)
	daftarRisikoPrioritasGetOneUseCase := usecase.ImplDaftarRisikoPrioritasGetByIDUseCase(daftarRisikoPrioritasGetOneGateway)
	daftarRisikoPrioritasDeleteUseCase := usecase.ImplDaftarRisikoPrioritasDeleteUseCase(daftarRisikoPrioritasDeleteGateway)
	daftarRisikoPrioritasUpdateUsecase := usecase.ImplDaftarRisikoPrioritasUpdateUseCase(daftarRisikoPrioritasGetOneGateway, daftarRisikoPrioritasCreateGateway, hasilAnalisisRisikoGetOneGateway, identifikasiRisikoStrategisPemdaGetOneGateway, penetapanKonteksRisikoStrategisPemdaGetOneGateway)

	// Usecase Penetapan Konteks Risiko Strategis Renstra OPD
	penetapanKonteksRisikoOperasionalGetAllUseCase := usecase.ImplPenetapanKonteksRisikoOperasionalGetAllUseCase(penetapanKonteksRisikoOperasionalGetAllGateway, ikuGetAllGateway, opdGetOneGateway)
	penetapanKonteksRisikoOperasionalGetOneUseCase := usecase.ImplPenetapanKonteksRisikoOperasionalGetByIDUseCase(penetapanKonteksRisikoOperasionalGetOneGateway, ikuGetAllGateway, opdGetOneGateway)
	penetapanKonteksRisikoOperasionalDeleteUseCase := usecase.ImplPenetapanKonteksRisikoOperasionalDeleteUseCase(penetapanKonteksRisikoOperasionalDeleteGateway)
	penetapanKonteksRisikoOperasionalCreateUseCase := usecase.ImplPenetapanKonteksRisikoOperasionalCreateUseCase(generateIdGateway, penetapanKonteksRisikoOperasionalCreateGateway, opdGetOneGateway)
	penetapanKonteksRisikoOperasionalUpdateUseCase := usecase.ImplPenetapanKonteksRisikoOperasionalUpdateUseCase(penetapanKonteksRisikoOperasionalGetOneGateway, penetapanKonteksRisikoOperasionalCreateGateway, opdGetOneGateway)
	penetapanKonteksRisikoOperasionalApprovalUsecase := usecase.ImplPenetapanKonteksRisikoOperasionalApprovalUseCase(penetapanKonteksRisikoOperasionalGetOneGateway, penetapanKonteksRisikoOperasionalCreateGateway)
	// Usecase Penilaian Risiko
	penilaianRisikoCreateUseCase := usecase.ImplPenilaianRisikoCreateUseCase(generateIdGateway, penilaianRisikoCreateGateway, daftarRisikoPrioritasGetOneGateway, hasilAnalisisRisikoGetOneGateway)
	penilaianRisikoGetAllUseCase := usecase.ImplPenilaianRisikoGetAllUseCase(penilaianRisikoGetAllGateway)
	penilaianRisikoGetOneUseCase := usecase.ImplPenilaianRisikoGetByIDUseCase(penilaianRisikoGetOneGateway)
	penilaianRisikoDeleteUseCase := usecase.ImplPenilaianRisikoDeleteUseCase(penilaianRisikoDeleteGateway)
	penilaianRisikoUpdateUseCase := usecase.ImplPenilaianRisikoUpdateUseCase(penilaianRisikoGetOneGateway, penilaianRisikoCreateGateway, daftarRisikoPrioritasGetOneGateway, hasilAnalisisRisikoGetOneGateway)

	// Usecase Kriteria Dampak
	kriteriaDampakGetAllUseCase := usecase.ImplKriteriaDampakGetAllUseCase(kriteriaDampakGetAllGateway)
	kriteriaDampakGetOneUseCase := usecase.ImplKriteriaDampakGetByIDUseCase(kriteriaDampakGetOneGateway)
	// Usecase Pengkomunikasian Pengendalian
	pengkomunikasianPengendalianCreateUseCase := usecase.ImplPengkomunikasianPengendalianCreateUseCase(generateIdGateway, pengkomunikasianPengendalianCreateGateway, penilaianRisikoGetOneGateway)
	pengkomunikasianPengendalianGetAllUseCase := usecase.ImplPengkomunikasianPengendalianGetAllUseCase(pengkomunikasianPengendalianGetAllGateway)
	pengkomunikasianPengendalianGetOneUseCase := usecase.ImplPengkomunikasianPengendalianGetByIDUseCase(pengkomunikasianPengendalianGetOneGateway)
	pengkomunikasianPengendalianDeleteUseCase := usecase.ImplPengkomunikasianPengendalianDeleteUseCase(pengkomunikasianPengendalianDeleteGateway)
	pengkomunikasianPengendalianUpdateUseCase := usecase.ImplPengkomunikasianPengendalianUpdateUseCase(pengkomunikasianPengendalianGetOneGateway, pengkomunikasianPengendalianCreateGateway, penilaianRisikoGetOneGateway)

	// Usecase Rancangan Pemantauan
	rancanganPemantauanCreateUseCase := usecase.ImplRancanganPemantauanCreateUseCase(generateIdGateway, rancanganPemantauanCreateGateway, penilaianRisikoGetOneGateway)
	rancanganPemantauanGetAllUseCase := usecase.ImplRancanganPemantauanGetAllUseCase(rancanganPemantauanGetAllGateway)
	rancanganPemantauanGetOneUseCase := usecase.ImplRancanganPemantauanGetByIDUseCase(rancanganPemantauanGetOneGateway)
	rancanganPemantauanDeleteUseCase := usecase.ImplRancanganPemantauanDeleteUseCase(rancanganPemantauanDeleteGateway)
	rancanganPemantauanUpdateUseCase := usecase.ImplRancanganPemantauanUpdateUseCase(rancanganPemantauanGetOneGateway, rancanganPemantauanCreateGateway, penilaianRisikoGetOneGateway)

	// Usecase Pencatatan Kejadian Risiko
	pencatatanKejadianRisikoCreateUseCase := usecase.ImplPencatatanKejadianRisikoCreateUseCase(generateIdGateway, pencatatanKejadianRisikoCreateGateway, penetapanKonteksRisikoStrategisPemdaGetOneGateway, identifikasiRisikoStrategisPemdaGetOneGateway)
	pencatatanKejadianRisikoGetAllUsecase := usecase.ImplPencatatanKejadianRisikoGetAllUseCase(pencatatanKejadianRisikoGetAllGateway)
	pencatatanKejadianRisikoGetOneUseCase := usecase.ImplPencatatanKejadianRisikoGetByIDUseCase(pencatatanKejadianRisikoGetOneGateway)
	pencatatanKejadianRisikoDeleteUseCase := usecase.ImplPencatatanKejadianRisikoDeleteUseCase(pencatatanKejadianRisikoDeleteGateway)
	pencatatanKejadianRisikoUpdateUseCase := usecase.ImplPencatatanKejadianRisikoUpdateUseCase(pencatatanKejadianRisikoGetOneGateway, pencatatanKejadianRisikoCreateGateway, penetapanKonteksRisikoStrategisPemdaGetOneGateway, identifikasiRisikoStrategisPemdaGetOneGateway)

	// Usecase Indeks Peringkat Prioritas
	indeksPeringkatPrioritasGetAllUseCase := usecase.ImplIndeksPeringkatPrioritasGetAllUseCase(indeksPeringkatPrioritasGetAllGateway)
	indeksPeringkatPrioritasGetOneUseCase := usecase.ImplIndeksPeringkatPrioritasGetByIDUseCase(indeksPeringkatPrioritasGetOneGateway)
	c := controller.Controller{
		Mux: mux,
		JWT: jwtToken,
	}

	// =================================================================
	// Controllers
	// =================================================================
	printer.
		Add(c.ExampleGetAllHandler(exampleGetAllUsecase)).
		Add(c.SpipGetAllHandler(spipGetAllUseCase)).
		Add(c.SpipGetByIDHandler(spipGetOneUseCase)).
		Add(c.SpipCreateHandler(SpipCreateUseCase)).
		Add(c.SpipDeleteHandler(spipDeleteUseCase)).
		Add(c.SpipUpdateHandler(spipUpdateUseCase)).
		Add(c.KategoriRisikoGetAllHandler(kategoriRisikoGetAllUseCase)).
		Add(c.KategoriRisikoGetByIDHandler(kategoriRisikoGetOneUseCase)).
		Add(c.KategoriRisikoCreateHandler(kategoriRisikoCreateUseCase)).
		Add(c.KategoriRisikoDeleteHandler(kategoriRisikoDeleteUseCase)).
		Add(c.KategoriRisikoUpdateHandler(kategoriRisikoUpdateUseCase)).
		Add(c.RekapitulasiHasilKuesionerCreateHandler(rekapitulasiHasilKuesionerCreateUseCase)).
		Add(c.RekapitulasiHasilKuesionerGetAllHandler(rekapitulasiHasilKuesionerGetAllUseCase)).
		Add(c.RekapitulasiHasilKuesionerGetByIDHandler(rekapitulasiHasilKuesionerGetOneUseCase)).
		Add(c.RekapitulasiHasilKuesionerDeleteHandler(rekapitulasiHasilKuesionerDeleteUseCase)).
		Add(c.PenetapanKonteksRisikoStrategisPemdaCreateHandler(penetapanKonteksRisikoStrategisPemdaCreateUseCase)).
		Add(c.PenetapanKonteksRisikoStrategisPemdaGetAllHandler(penetapanKonteksRisikoStrategisPemdaGetAllUseCase)).
		Add(c.PenetapanKonteksRisikoStrategisPemdaGetOneHandler(penetapanKonteksRisikoStrategisPemdaGetOneUseCase)).
		Add(c.PenetapanKonteksRisikoStrategisPemdaDeleteHandler(penetapanKonteksRisikoStrategisPemdaDeleteUseCase)).
		Add(c.PenetapanKonteksRisikoStrategisPemdaUpdateHandler(penetapanKonteksRisikoStrategisPemdaUpdateUseCase)).
		Add(c.RcaCreateHandler(rcaCreateUseCase)).
		Add(c.RcaGetAllHandler(rcaGetAllUseCase)).
		Add(c.RcaGetByIDHandler(rcaGetOneUseCase)).
		Add(c.RcaDeleteHandler(rcaDeleteUseCase)).
		Add(c.RcaUpdateHandler(rcaUpdateUseCase)).
		Add(c.IdentifikasiRisikoStrategisPemdaCreateHandler(identifikasiRisikoStrategisPemdaCreateUseCase)).
		Add(c.IdentifikasiRisikoStrategisPemdaGetAllHandler(identifikasiRisikoStrategisPemdaGetAllUseCase)).
		Add(c.IdentifikasiRisikoStrategisPemdaGetByIDHandler(identifikasiRisikoStrategisPemdaGetOneUseCase)).
		Add(c.IdentifikasiRisikoStrategisPemdaDeleteHandler(identifikasiRisikoStrategisPemdaDeleteUseCase)).
		Add(c.IdentifikasiRisikoStrategisPemdaUpdateHandler(identifikasiRisikoStrategisPemdaUpdateUseCase)).
		Add(c.SimpulanKondisiKelemahanLingkunganCreateHandler(simpulanKondisiKelemahanLingkunganCreateUseCase)).
		Add(c.SimpulanKondisiKelemahanLingkunganGetAllHandler(simpulanKondisiKelemahanLingkunganGetAllUseCase)).
		Add(c.SimpulanKondisiKelemahanLingkunganGetByIDHandler(simpulanKondisiKelemahanLingkunganGetOneUseCase)).
		Add(c.SimpulanKondisiKelemahanLingkunganDeleteHandler(simpulanKondisiKelemahanLingkunganDeleteUseCase)).
		Add(c.SimpulanKondisiKelemahanLingkunganUpdateHandler(simpulanKondisiKelemahanLingkunganUpdateUseCase)).
		Add(c.IKUGetAllHandler(ikuGetAllUseCase)).
		Add(c.IKUGetByIDHandler(ikuGetOneUseCase)).
		Add(c.IKUCreateHandler(ikuCreateUseCase)).
		Add(c.IKUDeleteHandler(ikuDeleteUseCase)).
		Add(c.IKUUpdateHandler(ikuUpdateUseCase)).
		Add(c.PenyebabRisikoCreateHandler(penyebabRisikoCreateUseCase)).
		Add(c.PenyebabRisikoGetAllHandler(penyebabRisikoGetAllUseCase)).
		Add(c.PenyebabRisikoGetByIDHandler(penyebabRisikoGetOneUseCase)).
		Add(c.PenyebabRisikoDeleteHandler(penyebabRisikoDeleteUseCase)).
		Add(c.PenyebabRisikoUpdateHandler(penyebabRisikoUpdateUseCase)).
		Add(c.HasilAnalisisRisikoCreateHandler(hasilAnalisisRisikoCreateUseCase)).
		Add(c.HasilAnalisisRisikoDeleteHandler(hasilAnalisisRisikoDeleteUsecase)).
		Add(c.HasilAnalisisRisikoGetAllHandler(hasilAnalisisRisikoGetAllUseCase)).
		Add(c.HasilAnalisisRisikoGetByIDHandler(hasilAnalisisRisikoGetOneUseCase)).
		Add(c.HasilAnalisisRisikoUpdateHandler(hasilAnalisisRisikoUpdateUsecase)).
		Add(c.PenilaianKegiatanPengendalianCreateHandler(penilaianKegiatanPengendalianCreateUseCase)).
		Add(c.PenilaianKegiatanPengendalianDeleteHandler(penilaianKegiatanPengendalianDeleteUseCase)).
		Add(c.PenilaianKegiatanPengendalianGetAllHandler(penilaianKegiatanPengendalianGetAllUseCase)).
		Add(c.PenilaianKegiatanPengendalianGetByIDHandler(penilaianKegiatanPengendalianGetOneUseCase)).
		Add(c.PenilaianKegiatanPengendalianUpdateHandler(penilaianKegiatanPengendalianUpdateUseCase)).
		Add(c.PenetapanKonteksRisikoStrategisRenstraOPDCreateHandler(penetapanKonteksRisikoStrategisRenstraOPDCreateUseCase)).
		Add(c.PenetapanKonteksRisikoStrategisRenstraOPDGetAllHandler(penetapanKonteksRisikoStrategisRenstraOPDGetAllUseCase)).
		Add(c.PenetapanKonteksRisikoStrategisRenstraOPDGetOneHandler(penetapanKonteksRisikoStrategisRenstraOPDGetOneUseCase)).
		Add(c.PenetapanKonteksRisikoStrategisRenstraOPDDeleteHandler(penetapanKonteksRisikoStrategisRenstraOPDDeleteUseCase)).
		Add(c.PenetapanKonteksRisikoStrategisRenstraOPDUpdateHandler(penetapanKonteksRisikoStrategisRenstraOPDUpdateUseCase)).
		Add(c.DaftarRisikoPrioritasCreateHandler(daftarRisikoPrioritasCreateUseCase)).
		Add(c.DaftarRisikoPrioritasGetAllHandler(daftarRisikoPrioritasGetAllUseCase)).
		Add(c.DaftarRisikoPrioritasGetByIDHandler(daftarRisikoPrioritasGetOneUseCase)).
		Add(c.DaftarRisikoPrioritasDeleteHandler(daftarRisikoPrioritasDeleteUseCase)).
		Add(c.DaftarRisikoPrioritasUpdateHandler(daftarRisikoPrioritasUpdateUsecase)).
		Add(c.PenetapanKonteksRisikoOperasionalCreateHandler(penetapanKonteksRisikoOperasionalCreateUseCase)).
		Add(c.PenetapanKonteksRisikoOperasionalGetAllHandler(penetapanKonteksRisikoOperasionalGetAllUseCase)).
		Add(c.PenetapanKonteksRisikoOperasionalGetOneHandler(penetapanKonteksRisikoOperasionalGetOneUseCase)).
		Add(c.PenetapanKonteksRisikoOperasionalDeleteHandler(penetapanKonteksRisikoOperasionalDeleteUseCase)).
		Add(c.PenetapanKonteksRisikoOperasionalUpdateHandler(penetapanKonteksRisikoOperasionalUpdateUseCase)).
		Add(c.PenilaianRisikoCreateHandler(penilaianRisikoCreateUseCase)).
		Add(c.PenilaianRisikoGetAllHandler(penilaianRisikoGetAllUseCase)).
		Add(c.PenilaianRisikoGetByIDHandler(penilaianRisikoGetOneUseCase)).
		Add(c.PenilaianRisikoDeleteHandler(penilaianRisikoDeleteUseCase)).
		Add(c.PenilaianRisikoUpdateHandler(penilaianRisikoUpdateUseCase)).
		Add(c.KriteriaDampakGetAllHandler(kriteriaDampakGetAllUseCase)).
		Add(c.KriteriaDampakGetByIDHandler(kriteriaDampakGetOneUseCase)).
		Add(c.PengkomunikasianPengendalianCreateHandler(pengkomunikasianPengendalianCreateUseCase)).
		Add(c.PengkomunikasianPengendalianGetAllHandler(pengkomunikasianPengendalianGetAllUseCase)).
		Add(c.PengkomunikasianPengendalianGetByIDHandler(pengkomunikasianPengendalianGetOneUseCase)).
		Add(c.PengkomunikasianPengendalianDeleteHandler(pengkomunikasianPengendalianDeleteUseCase)).
		Add(c.PengkomunikasianPengendalianUpdateHandler(pengkomunikasianPengendalianUpdateUseCase)).
		Add(c.IdentifikasiRisikoStrategisOPDCreateHandler(identifikasiRisikoStrategisOPDCreateUseCase)).
		Add(c.IdentifikasiRisikoStrategisOPDGetAllHandler(identifikasiRisikoStrategisOPDGetAllUseCase)).
		Add(c.IdentifikasiRisikoStrategisOPDGetByIDHandler(identifikasiRisikoStrategisOPDGetOneUseCase)).
		Add(c.IdentifikasiRisikoStrategisOPDDeleteHandler(identifikasiRisikoStrategisOPDDeleteUseCase)).
		Add(c.IdentifikasiRisikoStrategisOPDUpdateHandler(identifikasiRisikoStrategisOPDUpdateUseCase)).
		Add(c.RancanganPemantauanCreateHandler(rancanganPemantauanCreateUseCase)).
		Add(c.RancanganPemantauanGetAllHandler(rancanganPemantauanGetAllUseCase)).
		Add(c.RancanganPemantauanGetByIDHandler(rancanganPemantauanGetOneUseCase)).
		Add(c.RancanganPemantauanDeleteHandler(rancanganPemantauanDeleteUseCase)).
		Add(c.RancanganPemantauanUpdateHandler(rancanganPemantauanUpdateUseCase)).
		Add(c.IdentifikasiRisikoOperasionalOPDCreateHandler(identifikasiRisikoOperasionalOPDCreateUseCase)).
		Add(c.IdentifikasiRisikoOperasionalOPDGetAllHandler(identifikasiRisikoOperasionalOPDGetAllUseCase)).
		Add(c.IdentifikasiRisikoOperasionalOPDGetByIDHandler(identifikasiRisikoOperasionalOPDGetOneUseCase)).
		Add(c.IdentifikasiRisikoOperasionalOPDDeleteHandler(identifikasiRisikoOperasionalOPDDeleteUseCase)).
		Add(c.IdentifikasiRisikoOperasionalOPDUpdateHandler(identifikasiRisikoOperasionalOPDUpdateUseCase)).
		Add(c.PencatatanKejadianRisikoCreateHandler(pencatatanKejadianRisikoCreateUseCase)).
		Add(c.PencatatanKejadianRisikoGetAllHandler(pencatatanKejadianRisikoGetAllUsecase)).
		Add(c.PencatatanKejadianRisikoGetByIDHandler(pencatatanKejadianRisikoGetOneUseCase)).
		Add(c.PencatatanKejadianRisikoDeleteHandler(pencatatanKejadianRisikoDeleteUseCase)).
		Add(c.PencatatanKejadianRisikoUpdateHandler(pencatatanKejadianRisikoUpdateUseCase)).
		Add(c.IndeksPeringkatPrioritasGetAllHandler(indeksPeringkatPrioritasGetAllUseCase)).
		Add(c.IndeksPeringkatPrioritasGetByIDHandler(indeksPeringkatPrioritasGetOneUseCase)).
		Add(c.PenetapanKonteksRisikoOperasionalApprovalHandler(penetapanKonteksRisikoOperasionalApprovalUsecase)).
		Add(c.IdentifikasiRisikoStrategisOPDApprovalHandler(identifikasiRisikoStrategisOPDApprovalUsecase))
}
