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
		Add(c.KategoriRisikoUpdateHandler(kategoriRisikoUpdateUseCase))
}
