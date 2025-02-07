package wiring

import (
	"gorm.io/gorm"
	"monitoring/controller"
	"monitoring/gateway"
	"monitoring/usecase"
	"net/http"
	gateway2 "shared/gateway"
	"shared/helper"
	usecase3 "shared/usecase"
)

func SetupDependency(mariaDB *gorm.DB, mux *http.ServeMux, printer *helper.ApiPrinter) {

	// gateways

	createActivityMonitor := gateway2.ImplCreateActivityMonitoringGateway(mariaDB)
	getActivityMonitor := gateway2.ImplGetActivityMonitoringGateway(mariaDB)
	getActivityMonitorByID := gateway.ImplActivityMonitorGetOneGateway(mariaDB)
	deleteActivityMonitor := gateway.ImplActivityMonitorDelete(mariaDB)

	// usecase
	createActivityMonitorUseCase := usecase.ImplActivityMonitorCreateUseCase(createActivityMonitor)
	getAllActivityMonitorUseCase := usecase3.ImpActivityMonitorGetAllUseCase(getActivityMonitor)
	getActivityMonitorUseCase := usecase.ImplActivityMonitorGetOne(getActivityMonitorByID)
	deleteActivityMonitorUseCase := usecase.ImplActivityMonitorDelete(deleteActivityMonitor)

	// controller
	printer.
		Add(controller.ActivityMonitorCreateHandler(mux, createActivityMonitorUseCase)).
		Add(controller.ActivityMonitorGetAllHandler(mux, getAllActivityMonitorUseCase)).
		Add(controller.ActivityMonitorDetailHandler(mux, getActivityMonitorUseCase)).
		Add(controller.ActivityMonitorDeleteHandler(mux, deleteActivityMonitorUseCase))
}
