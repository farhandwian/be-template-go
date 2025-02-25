package wiring

import (
	"net/http"
	"rmis/controller"
	"rmis/gateway"
	"rmis/usecase"
	"shared/helper"
	"shared/helper/cronjob"

	"gorm.io/gorm"
)

func SetupDependency(mariaDB *gorm.DB, mux *http.ServeMux, jwtToken helper.JWTTokenizer, printer *helper.ApiPrinter, cj *cronjob.CronJob, sseDashboard *helper.SSE) {
	// =================================================================
	// Gateway
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

	// Gateway Penilaian Risiko
	penilaianRisikoCreateGateway := gateway.ImplPenilaianRisikoSave(mariaDB)
	penilaianRisikoGetAllGateway := gateway.ImplPenilaianRisikoGetAll(mariaDB)
	penilaianRisikoGetOneGateway := gateway.ImplPenilaianRisikoGetByID(mariaDB)
	penilaianRisikoDeleteGateway := gateway.ImplPenilaianRisikoDelete(mariaDB)

	// =================================================================
	// Usecase
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
	rcaUpdateUseCase := usecase.ImplRcaUpdateUseCase(rcaGetOneGateway, rcaCreateGateway)

	// Usecase Identifikasi Risiko Strategis Pemda
	identifikasiRisikoStrategisPemdaGetAllUseCase := usecase.ImplIdentifikasiRisikoStrategisPemdaGetAllUseCase(identifikasiRisikoStrategisPemdaGetAllGateway)
	identifikasiRisikoStrategisPemdaGetOneUseCase := usecase.ImplIdentifikasiRisikoStrategisPemdaGetByIDUseCase(identifikasiRisikoStrategisPemdaGetOneGateway)
	identifikasiRisikoStrategisPemdaDeleteUseCase := usecase.ImplIdentifikasiRisikoStrategisPemdaDeleteUseCase(identifikasiRisikoStrategisPemdaDeleteGateway)
	identifikasiRisikoStrategisPemdaCreateUseCase := usecase.ImplIdentifikasiRisikoStrategisPemdaCreateUseCase(generateIdGateway, identifikasiRisikoStrategisPemdaCreateGateway, kategoriRisikoGetOneGateway)
	identifikasiRisikoStrategisPemdaUpdateUseCase := usecase.ImplIdentifikasiRisikoStrategisPemdaUpdateUseCase(identifikasiRisikoStrategisPemdaGetOneGateway, identifikasiRisikoStrategisPemdaCreateGateway, kategoriRisikoGetOneGateway, rcaGetOneGateway)

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
	hasilAnalisisRisikoCreateUseCase := usecase.ImplHasilAnalisisRisikoCreateUseCase(generateIdGateway, hasilAnalisisRisikoCreateGateway, identifikasiRisikoStrategisPemdaGetOneGateway)
	hasilAnalisisRisikoGetAllUseCase := usecase.ImplHasilAnalisisRisikoGetAllUseCase(hasilAnalisisRisikoGetAllGateway)
	hasilAnalisisRisikoGetOneUseCase := usecase.ImplHasilAnalisisRisikoGetByIDUseCase(hasilAnalisisRisikoGetOneGateway)
	hasilAnalisisRisikoDeleteUsecase := usecase.ImplHasilAnalisisRisikoDeleteUseCase(hasilAnalisisRisikoDeleteGateway)
	hasilAnalisisRisikoUpdateUsecase := usecase.ImplHasilAnalisisRisikoUpdateUseCase(hasilAnalisisRisikoGetOneGateway, hasilAnalisisRisikoCreateGateway)

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
	penetapanKonteksRisikoStrategisRenstraOPDCreateUseCase := usecase.ImplPenetapanKonteksRisikoStrategisRenstraOPDCreateUseCase(generateIdGateway, penetapanKonteksRisikoStrategisRenstraOPDCreateGateway)
	penetapanKonteksRisikoStrategisRenstraOPDUpdateUseCase := usecase.ImplPenetapanKonteksRisikoStrategisRenstraOPDUpdateUseCase(penetapanKonteksRisikoStrategisRenstraOPDGetOneGateway, penetapanKonteksRisikoStrategisRenstraOPDCreateGateway)
	// Usecase Daftar Risiko Prioritas
	daftarRisikoPrioritasCreateUseCase := usecase.ImplDaftarRisikoPrioritasCreateUseCase(generateIdGateway, daftarRisikoPrioritasCreateGateway, hasilAnalisisRisikoGetOneGateway, identifikasiRisikoStrategisPemdaGetOneGateway)
	daftarRisikoPrioritasGetAllUseCase := usecase.ImplDaftarRisikoPrioritasGetAllUseCase(daftarRisikoPrioritasGetAllGateway)
	daftarRisikoPrioritasGetOneUseCase := usecase.ImplDaftarRisikoPrioritasGetByIDUseCase(daftarRisikoPrioritasGetOneGateway)
	daftarRisikoPrioritasDeleteUseCase := usecase.ImplDaftarRisikoPrioritasDeleteUseCase(daftarRisikoPrioritasDeleteGateway)
	daftarRisikoPrioritasUpdateUsecase := usecase.ImplDaftarRisikoPrioritasUpdateUseCase(daftarRisikoPrioritasGetOneGateway, daftarRisikoPrioritasCreateGateway, hasilAnalisisRisikoGetOneGateway, identifikasiRisikoStrategisPemdaGetOneGateway)

	// Usecase Penilaian Risiko
	penilaianRisikoCreateUseCase := usecase.ImplPenilaianRisikoCreateUseCase(generateIdGateway, penilaianRisikoCreateGateway, daftarRisikoPrioritasGetOneGateway, hasilAnalisisRisikoGetOneGateway)
	penilaianRisikoGetAllUseCase := usecase.ImplPenilaianRisikoGetAllUseCase(penilaianRisikoGetAllGateway)
	penilaianRisikoGetOneUseCase := usecase.ImplPenilaianRisikoGetByIDUseCase(penilaianRisikoGetOneGateway)
	penilaianRisikoDeleteUseCase := usecase.ImplPenilaianRisikoDeleteUseCase(penilaianRisikoDeleteGateway)
	penilaianRisikoUpdateUseCase := usecase.ImplPenilaianRisikoUpdateUseCase(penilaianRisikoGetOneGateway, penilaianRisikoCreateGateway, daftarRisikoPrioritasGetOneGateway, hasilAnalisisRisikoGetOneGateway)

	c := controller.Controller{
		Mux: mux,
		JWT: jwtToken,
	}

	// Controllers
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
		Add(c.PenilaianRisikoCreateHandler(penilaianRisikoCreateUseCase)).
		Add(c.PenilaianRisikoGetAllHandler(penilaianRisikoGetAllUseCase)).
		Add(c.PenilaianRisikoGetByIDHandler(penilaianRisikoGetOneUseCase)).
		Add(c.PenilaianRisikoDeleteHandler(penilaianRisikoDeleteUseCase)).
		Add(c.PenilaianRisikoUpdateHandler(penilaianRisikoUpdateUseCase))
}
