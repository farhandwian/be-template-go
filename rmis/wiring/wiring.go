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

	// Usecase
	exampleGetAllUsecase := usecase.ImplExampleGetAllUseCase(exampleGetAllGateway)

	// Usecase SPIP
	spipGetAllUseCase := usecase.ImplSpipGetAllUseCase(spipGetAllGateway)
	spipGetOneUseCase := usecase.ImplSpipGetByIDUseCase(spipGetOneGateway)
	spipDeleteUseCase := usecase.ImplSpipDeleteUseCase(spipDeleteGateway)
	SpipCreateUseCase := usecase.ImplSpipCreateUseCase(generateIdGateway, spipCreateGateway)
	spipUpdateUseCase := usecase.ImplSpipUpdateUseCase(spipGetOneGateway, spipCreateGateway)

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
		Add(c.SpipUpdateHandler(spipUpdateUseCase))
}
