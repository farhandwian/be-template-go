package wiring

import (
	"net/http"
	"perizinan/controller"
	"perizinan/gateway"
	"perizinan/usecase"
	"shared/helper"

	"gorm.io/gorm"
)

func SetupDependency(mariaDB *gorm.DB, mux *http.ServeMux, printer *helper.ApiPrinter, jwt helper.JWTTokenizer) {

	// gateways
	generateId := gateway.ImplGenerateId()
	laporSaveGateway := gateway.ImplLaporanPerizinanSave(mariaDB)
	perizinanGetOneGateway := gateway.ImplLaporanPerizinanGetOne(mariaDB)
	perizinanGetAllGateway := gateway.ImplLaporanPerizinanGetAll(mariaDB)
	perizinanGetOneByIDGateway := gateway.ImplLaporanPerizinanGetOneByID(mariaDB)
	skGetOneGateway := gateway.ImplSKPerizinanGetOne(mariaDB)
	skPerizinanGetAllGateway := gateway.ImplSKPerizinanGetAll(mariaDB)

	// usecase
	laporSaveUsecase := usecase.ImplLaporanPerizinanSaveUseCase(generateId, skGetOneGateway, perizinanGetOneGateway, laporSaveGateway)
	perizinanGetSKPeriodeUsecase := usecase.ImplLaporanPerizinanSKPeriodeUseCase(skGetOneGateway, perizinanGetOneGateway)
	perizinanGetAllUsecase := usecase.ImplLaporGetAllUseCase(perizinanGetAllGateway, skPerizinanGetAllGateway)
	laporSubmitUsecase := usecase.ImplLaporanPerizinanSubmitUseCase(perizinanGetOneByIDGateway, laporSaveGateway)
	skGetAllUseCase := usecase.ImplSKGetAllUseCase(skPerizinanGetAllGateway)
	skExistUseCase := usecase.ImplLaporanPerizinanExistUseCase(skGetOneGateway, perizinanGetOneGateway)

	// controllers
	printer.
		Add(controller.SKGetAllHandler(mux, skGetAllUseCase, jwt)).
		Add(controller.LaporanPerizinanGetAllHandler(mux, perizinanGetAllUsecase, jwt)).
		Add(controller.LaporanPerizinanSKPeriodeHandler(mux, perizinanGetSKPeriodeUsecase)).
		Add(controller.LaporanPerizinanSaveHandler(mux, laporSaveUsecase)).
		Add(controller.LaporanPerizinanSubmitHandler(mux, laporSubmitUsecase)).
		Add(controller.UploadFileHandler(mux, skExistUseCase)).
		Add(controller.DownloadFileHandler(mux))
}
