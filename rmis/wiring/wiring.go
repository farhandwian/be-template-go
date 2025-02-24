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
	rcaCreateUseCase := usecase.ImplRcaCreateUseCase(generateIdGateway, rcaCreateGateway, identifikasiRisikoStrategisPemdaGetOneGateway)
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
		Add(c.IKUUpdateHandler(ikuUpdateUseCase))
}
