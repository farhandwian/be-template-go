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

	// Usecase
	exampleGetAllUsecase := usecase.ImplExampleGetAllUseCase(exampleGetAllGateway)

	c := controller.Controller{
		Mux: mux,
		JWT: jwtToken,
	}

	// Controllers
	printer.
		Add(c.ExampleGetAllHandler(exampleGetAllUsecase))
}
