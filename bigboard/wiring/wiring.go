package wiring

import (
	"bigboard/controller"
	"bigboard/gateway"
	"bigboard/usecase"
	"net/http"
	"path/filepath"
	sharedGateway "shared/gateway"
	"shared/helper"
	"shared/middleware"
	sharedUsecase "shared/usecase"

	dashboardGateway "dashboard/gateway"
	iamgw "iam/gateway"

	"gorm.io/gorm"
)

func SetupDependency(mariaDB *gorm.DB, timeScaleDB *gorm.DB, mux *http.ServeMux, jwt helper.JWTTokenizer, printer *helper.ApiPrinter, sse *helper.SSE) {

	staticPath := filepath.Join("static", "assets")
	fileServer := http.FileServer(http.Dir(staticPath))
	mux.Handle("GET /images/", http.StripPrefix("/images/", fileServer))

	sendSSEMessage := sharedGateway.ImplSendSSEMessage(sse)

	// gateways
	getConfigGateway := sharedGateway.ImplGetConfigGateway(mariaDB)
	getListWaterChannelDoorGateway := sharedGateway.ImplGetListWaterChannelDoor(mariaDB)
	getListWaterChannelDoorWithCCTVAndOfficerGateway := gateway.ImplGetListWaterChannelDoorWithCCTVAndOfficer(mariaDB)
	sensitiveJobsSave := gateway.ImplSensitiveJobsSave(mariaDB)
	generateId := gateway.ImplGenerateId()
	generateIdIam := iamgw.ImplGenerateId()

	generateJWT := iamgw.ImplGenerateJWT(jwt)               //
	passwordValidate := iamgw.ImplPasswordValidate()        //
	userGetAll := iamgw.ImplUserGetAllWithDatabase(mariaDB) // USE DATABASE IMPLEMENTATION
	userSave := iamgw.ImplUserSaveWithDatabse(mariaDB)

	getWaterChannelDoor := sharedGateway.ImplGetWaterChannelDoorByID(mariaDB)
	getWaterChannelDevices := sharedGateway.ImplGetWaterChannelDevicesByDoorID(mariaDB)
	getWaterChannelOfficers := sharedGateway.ImplGetWaterChannelOfficersByDoorID(mariaDB)
	getLatestWaterGatesByDoorID := sharedGateway.ImplGetLatestWaterGatesByDoorID(timeScaleDB)
	getLatestWaterSurfaceElevation := sharedGateway.ImplGetLatestWaterSurfaceElevationByDoorID(timeScaleDB)

	fetchBMKGWeatherForecast := sharedGateway.ImplFetchBMKGWeatherForecast()

	rawWaterList := sharedGateway.ImplGetRawWaterList(mariaDB)
	rawWaterDetail := sharedGateway.ImplRawWaterGetDetailGateway(mariaDB)

	groundWaterList := sharedGateway.ImplGetGroundWaterList(mariaDB)
	groundWaterDetail := sharedGateway.ImplGroundWaterGetDetailGateway(mariaDB)

	weirList := sharedGateway.ImplGetWeirList(mariaDB)
	weirDetail := sharedGateway.ImplWeirGetDetailGateway(mariaDB)

	damList := sharedGateway.ImplGetDamList(mariaDB)
	damDetail := sharedGateway.ImplGetDamDetailByID(mariaDB)

	climatologyList := sharedGateway.ImplGetClimatologyList(mariaDB)
	climatologyDetail := sharedGateway.ImplGetClimatologyDetailByID(mariaDB)

	lakeList := sharedGateway.ImplGetLakeList(mariaDB)
	lakeDetail := sharedGateway.ImplGetLakeDetailByID(mariaDB)

	waterReservoirList := sharedGateway.ImplGetWaterReservoirList(mariaDB)
	waterReservoirDetail := sharedGateway.ImplWaterReservoirGetDetailGateway(mariaDB)

	intakeList := sharedGateway.ImplGetIntakeList(mariaDB)
	intakeDetail := sharedGateway.ImplIntakeGetDetailGateway(mariaDB)

	pahAbsahList := sharedGateway.ImplGetPahAbsahList(mariaDB)
	pahAbsahDetail := sharedGateway.ImplPahAbsahGetDetailGateway(mariaDB)

	coastalProtectionList := sharedGateway.ImplGetCoastalProtectionList(mariaDB)
	coastalProtectionDetail := sharedGateway.ImplCoastalProtectionGetDetailGateway(mariaDB)

	sedimentControlList := sharedGateway.ImplGetSedimentControlList(mariaDB)
	sedimentControlDetail := sharedGateway.ImplSedimentControlGetDetailGateway(mariaDB)

	wellList := sharedGateway.ImplGetWellList(mariaDB)
	wellDetail := sharedGateway.ImplWellGetDetailGateway(mariaDB)

	rainFallGetByID := sharedGateway.ImplRainfallGetDetailGateway(mariaDB)
	getListHydrologyPostOfficer := sharedGateway.ImplGetOfficerPostByNameList(mariaDB)
	getListHydrologyRainPost := sharedGateway.ImplGetRainfallPostList(mariaDB)

	getListWaterLevelPost := sharedGateway.ImplGetWaterLevelPostList(mariaDB)

	waterLevelGetByID := sharedGateway.ImplWaterLevelGetDetailGateway(mariaDB)

	getDebitGateway := sharedGateway.ImplGetDebit(timeScaleDB)
	getWaterSurfaceElevationGateway := sharedGateway.ImplGetWaterSurfaceElevation(timeScaleDB)
	getGateGateway := sharedGateway.ImplGetGate(timeScaleDB)
	getDevicesGateway := sharedGateway.ImplGetDevice(mariaDB)
	getLaporanPerizinanStatusGateway := sharedGateway.ImplGetLaporanPerizinanStatusCount(mariaDB)

	// getListWaterChannelCCTVCountGateway := gateway.ImplGetListWaterChannelCCTVCount(mariaDB)
	// getListWaterChannelOfficerCountGateway := gateway.ImplGetListWaterChannelOfficerCount(mariaDB)
	getListWaterChannelTMAGateway := sharedGateway.ImplGetListWaterChannelTMA(timeScaleDB)
	getListWaterChannelActualDebitGateway := sharedGateway.ImplGetListWaterChannelActualDebit(timeScaleDB)

	// gateway middleware
	userGetAllToUseCase := middleware.Logging(userGetAll, 4)
	generateJWTToUseCase := middleware.Logging(generateJWT, 4)
	generateIdToUseCase := middleware.Logging(generateIdIam, 4)
	userSaveToUseCase := middleware.Logging(userSave, 4)

	// SSE gateway
	getLatestDebitGateway := sharedGateway.ImplGetLatestDebit(timeScaleDB)

	generalInfoGateway := sharedGateway.ImplGetGeneralInfoGateway(mariaDB)
	waterAvailabilityForecastGateway := sharedGateway.ImplGetWaterAvailabilityForecastGateway(timeScaleDB)

	getListWaterLevelPostManualGateway := sharedGateway.ImplGetListLatestManualWaterPostHandler(timeScaleDB)
	getListWaterLevelPostTelemetryGateway := sharedGateway.ImplGetListLatestTelemetryWaterPostGateway(timeScaleDB)

	getLatestWaterLevelPostTelemetryByIDGateway := sharedGateway.ImplGetLatestTelemetryWaterPostByIDGateway(timeScaleDB)
	getLatestRainfallDailyByIDGateway := sharedGateway.ImplGetLatestTelemetryRainPostByIDGateway(timeScaleDB)
	getLatestRainfallTelemetryGateway := sharedGateway.ImplGetLatestTelemetryRainPostByIDGateway(timeScaleDB)

	createActivityMonitoringGateway := sharedGateway.ImplCreateActivityMonitoringGateway(mariaDB)

	passwordValidateToUseCase := middleware.Logging(passwordValidate, 4)
	// usecase
	listWaterChannelUseCase := middleware.Timing(usecase.ImplListWaterChannel(
		getListWaterChannelDoorGateway,
		getListWaterChannelTMAGateway,
		getListWaterChannelActualDebitGateway,
		sharedGateway.ImplGetCCTVDevices(mariaDB),
		getConfigGateway,
	), "ListWaterChannel")

	getServiceStatusGateway := sharedGateway.ImplGetServiceStatusGateway()
	getTmaDeviceCountGateway := sharedGateway.ImplGetTmaDeviceCount(timeScaleDB)

	getRequiredDebitMainWaterChannelDoorGateway := sharedGateway.ImplGetTotalRequirementDebitMainWaterChannelDoorGateway(mariaDB)
	getAlarmCountGateway := sharedGateway.ImplGetAlarmCountGateway(mariaDB)

	getBMKGLocationGateway := sharedGateway.ImplFetchBMKGWeatherLocationGateway()

	generateIdGateway := dashboardGateway.ImplGenerateId()
	loginMangantiGateway := dashboardGateway.ImplLoginMangantiEmpty() // TODO temporary empty login until ready
	doorControlAPIGateway := dashboardGateway.ImplDoorControlAPI_(loginMangantiGateway)
	doorControlHistorySaveGateway := dashboardGateway.ImplDoorControlHistorySave(mariaDB)
	getOneWaterGatesByDoorID := dashboardGateway.ImplGetOneWaterGatesByDoorID(timeScaleDB)
	getOneUser := iamgw.ImplUserGetOneByIDWithDatabase(mariaDB)
	getOneWaterChannelDeviceByDoorID := dashboardGateway.ImplGetOneWaterChannelDeviceByDoorID(mariaDB)
	getSensitiveJob := gateway.ImplGetOneSensitiveJob(mariaDB)
	// loginOTPSubmit := iamUsecase.ImplLoginOTPSubmit(passwordValidateToUseCase, userGetAllToUseCase, generateJWTToUseCase, generateIdToUseCase, userSaveToUseCase, createActivityMonitoringGateway)

	fcmTokenSaveGateaway := gateway.ImplFCMTokenSave(mariaDB)
	loginOTPSubmit := usecase.ImplLoginOTPSubmitAuthenticator(passwordValidateToUseCase, userGetAllToUseCase, generateJWTToUseCase, generateIdToUseCase, userSaveToUseCase, createActivityMonitoringGateway, fcmTokenSaveGateaway)

	detailPintuAirUseCase := sharedUsecase.ImplDetailPintuAir(
		getWaterChannelDoor,
		getWaterChannelDevices,
		getWaterChannelOfficers,
		// getLatestWaterGatesByDoorID,
		// getLatestWaterSurfaceElevation,
	)

	listWaterChannelDoorCCTVUseCase := usecase.ImplListWaterChannelDoorCCTV(
		getListWaterChannelDoorWithCCTVAndOfficerGateway,
	)

	listWaterChannelDoorOfficerUseCase := usecase.ImplListWaterChannelDoorOfficer(
		getListWaterChannelDoorWithCCTVAndOfficerGateway,
	)

	listWaterChannelDoorWithWaterSurfaceElevationUseCase := usecase.ImplListWaterChannelDoorWithWaterSurfaceElevation(
		getListWaterChannelDoorGateway,
		getListWaterChannelTMAGateway,
	)

	listWaterChannelDoorWithActualDebitUseCase := usecase.ImplListWaterChannelDoorWithActualDebit(
		getListWaterChannelDoorGateway,
		getListWaterChannelActualDebitGateway,
	)

	getDeviceStatusGateway := sharedGateway.ImplGetDeviceStatusGateway()
	getSpeedtestStatusGateway := sharedGateway.ImplGetSpeedtestStatusGateway()

	getOneFCMToken := gateway.ImplGetOneFCMToken(mariaDB)

	generalInfo := sharedUsecase.ImplGeneralInfo(generalInfoGateway, getRequiredDebitMainWaterChannelDoorGateway)

	getWeatherForecast := sharedUsecase.ImplGetWeatherForecast(fetchBMKGWeatherForecast)

	listRawWaterUseCase := sharedUsecase.ImplGetListRawWaterUseCase(rawWaterList)
	detailRawWaterUseCase := sharedUsecase.ImplGetRawWaterDetail(rawWaterDetail)

	listGroundWaterUseCase := sharedUsecase.ImplGetListGroundWaterUseCase(groundWaterList)
	detailGroundWaterUseCase := sharedUsecase.ImplGetGroundwaterDetail(groundWaterDetail)

	listWeirUseCase := sharedUsecase.ImplGetListWeirUseCase(weirList)
	detailWeirUseCase := sharedUsecase.ImplGetWeirDetail(weirDetail)

	listDamUseCase := sharedUsecase.ImplGetListDamUseCase(damList)
	detailDamUseCase := sharedUsecase.ImplGetDamDetail(damDetail)

	climatologyListUseCase := sharedUsecase.ImplGetClimatologyPostUseCase(climatologyList)
	climatologyDetailUseCase := sharedUsecase.ImplGetClimatologyDetailPostUseCase(climatologyDetail)

	lakeListUseCase := sharedUsecase.ImplGetListLakeUseCase(lakeList)
	lakeDetailUseCase := sharedUsecase.ImplGetLakeDetail(lakeDetail)

	waterReservoirListUseCase := sharedUsecase.ImplGetListWaterReservoirUseCase(waterReservoirList)
	waterReservoirDetailUseCase := sharedUsecase.ImplGetWaterReservoirDetail(waterReservoirDetail)

	intakeListUseCase := sharedUsecase.ImplGetListIntakeUseCase(intakeList)
	intakeDetailUseCase := sharedUsecase.ImplGetIntakeDetail(intakeDetail)

	pahAbsahListUseCase := sharedUsecase.ImplGetListPahAbsahUseCase(pahAbsahList)
	pahAbsahDetailUseCase := sharedUsecase.ImplGetPahAbsahDetail(pahAbsahDetail)

	coastalProtectionListUseCase := sharedUsecase.ImplGetListCoastalProtectionUseCase(coastalProtectionList)
	coastalProtectionDetailUseCase := sharedUsecase.ImplGetCoastalProtectionDetail(coastalProtectionDetail)

	sedimentControlListUseCase := sharedUsecase.ImplGetListSedimentControlUseCase(sedimentControlList)
	sedimentControlDetailUseCase := sharedUsecase.ImplGetSedimentControlDetail(sedimentControlDetail)

	wellListUseCase := sharedUsecase.ImplGetListWellUseCase(wellList)
	wellDetailUseCase := sharedUsecase.ImplGetWellDetail(wellDetail)

	getListRainFall := sharedUsecase.ImplGetListRainFallUseCase(getListHydrologyRainPost, getLatestRainfallTelemetryGateway)
	getDetailRainFall := sharedUsecase.ImplGetRainFallDetail(rainFallGetByID, getListHydrologyPostOfficer, getLatestRainfallDailyByIDGateway)

	getListWaterLevel := sharedUsecase.ImplGetListWaterLevelUseCase(getListWaterLevelPost, getListWaterLevelPostManualGateway, getListWaterLevelPostTelemetryGateway)
	getDetailWaterLevel := sharedUsecase.ImplGetWaterLevelDetail(waterLevelGetByID, getListHydrologyPostOfficer)

	chartDebitUseCase := sharedUsecase.ImplChartDebitUseCase(getDebitGateway)
	chartTMAUseCase := sharedUsecase.ImplChartTMAUseCase(getWaterSurfaceElevationGateway)
	chartPintuAirUseCase := sharedUsecase.ImplChartPintuAirUseCase(getDevicesGateway, getGateGateway)

	getStatusDebitUseCase := sharedUsecase.ImplGetStatusDebit(getWaterChannelDoor, getLatestWaterSurfaceElevation, getLatestDebitGateway)

	aiGetAllCCTVUseCase := usecase.ImplAIGetAllCCTVFromWaterChannel(getWaterChannelDevices, getWaterChannelDoor, sendSSEMessage)
	aiGetCertainCCTVUseCase := usecase.ImplAIGetCertainCCTVFromWaterChannel(getWaterChannelDevices, getWaterChannelDoor, sendSSEMessage)
	aiGetDetailWaterChannelUseCase := usecase.ImplWaterChannelDetailUseCase(getWaterChannelDoor, getWaterChannelDevices, getWaterChannelOfficers, getLatestDebitGateway, getLatestWaterSurfaceElevation, sendSSEMessage)
	aiOpenLayerUseCase := usecase.ImplAiOpenLayer(sendSSEMessage)
	aiPlayAudioUseCase := usecase.ImplAiPlayAudioUseCase(sendSSEMessage)
	aiOpenPhotoUseCase := usecase.ImplAiOpenPhotoUseCase(sendSSEMessage)
	aiGetPhotoPathUseCase := usecase.ImplGetPhotoPath()
	aiBackToHomeUseCase := usecase.ImplAiBackToHome(sendSSEMessage)
	aiGetRainFallDetailUseCase := usecase.ImplAiGetRainFallDetail(rainFallGetByID, getListHydrologyPostOfficer, getLatestRainfallDailyByIDGateway, sendSSEMessage)
	aiGetWaterLevelPostDetailUseCase := usecase.ImplAiGetWaterLevelDetail(waterLevelGetByID, getListHydrologyPostOfficer, getLatestWaterLevelPostTelemetryByIDGateway, sendSSEMessage)
	aiOpenVideoUseCase := usecase.ImplAiOpenVideoUseCase(sendSSEMessage)
	aiGetVideoPathUseCase := usecase.ImplGetVideoPath()
	aiShowGraphUseCase := usecase.ImplAiShowGraphUseCase(sendSSEMessage)
	aiZoomInUseCase := usecase.ImplAiZoomInUseCase(getWaterChannelDoor, sendSSEMessage)
	aiDoorControlUseCase := usecase.ImplAiDoorControlRunDirectly(generateId, sensitiveJobsSave, sendSSEMessage, getWaterChannelDoor, getWaterChannelDevices, getLatestWaterGatesByDoorID, getOneFCMToken)
	aiDoorControlRunUseCase := usecase.ImplAIDoorControlRunDirectly(passwordValidate, doorControlAPIGateway, generateIdGateway, getSensitiveJob, sensitiveJobsSave, doorControlHistorySaveGateway, getOneWaterGatesByDoorID, getOneUser, getOneWaterChannelDeviceByDoorID, getWaterChannelDoor, getWaterChannelDevices, getWaterChannelOfficers, getLatestDebitGateway, getLatestWaterSurfaceElevation, sendSSEMessage, createActivityMonitoringGateway)
	aiCancelSensitiveJobs := usecase.ImplAiCancelSensitiveJobs(getSensitiveJob, sensitiveJobsSave, sendSSEMessage)

	citanduyAreaUseCase := usecase.ImplGetCitanduyAreaUseCase()
	citanduyDasUseCase := usecase.ImplGetCitanduyDASUseCase()

	doorControlGetAllGateway := sharedGateway.ImplDoorControlGetAll(mariaDB)
	doorControlHistoryGetAllGateway := sharedGateway.ImplDoorControlHistoryGetAll(mariaDB)

	doorControlGetAllUsecase := sharedUsecase.ImplDoorControlGetAll(doorControlGetAllGateway)
	doorControlHistoryGetAllUsecase := sharedUsecase.ImplDoorControlHistoryGetAll(doorControlHistoryGetAllGateway)

	getGateStatusUsecase := sharedUsecase.ImplGetGateStatus(getWaterChannelDevices, getLatestWaterGatesByDoorID)
	waterAvailabilityForecastUsecase := sharedUsecase.ImplWaterAvailabilityForecastUseCase(waterAvailabilityForecastGateway)

	webHookSpeedTestUseCase := usecase.ImplWebhookSpeedTestUseCase(createActivityMonitoringGateway)
	getDeviceStatusUseCase := sharedUsecase.ImplGetDeviceStatusUseCase(getDeviceStatusGateway, getTmaDeviceCountGateway, getConfigGateway)
	getSpeedtestStatusUseCase := sharedUsecase.ImplGetSpeedtestStatusUseCase(getSpeedtestStatusGateway)
	getPerizinanStatusUseCase := sharedUsecase.ImplGetPerizinanStatusUseCase(getLaporanPerizinanStatusGateway)
	getDroneStatusUseCase := sharedUsecase.ImplGetDroneStatusUseCase()
	getServiceStatusUseCase := sharedUsecase.ImplGetServiceStatusUseCase(getServiceStatusGateway)
	getMonitoringStatusUseCase := sharedUsecase.ImplGetMonitoringStatusUseCase(getAlarmCountGateway)
	getBMKGLocationUseCase := sharedUsecase.ImplGetBMKGLocationUseCase(getBMKGLocationGateway)

	loginOTPSubmitToHandler := middleware.Logging(loginOTPSubmit, 0) //

	aiSystemSummaryUseCase := usecase.ImplAISystemSummary(
		generalInfo,
		getDeviceStatusUseCase,
		getDroneStatusUseCase,
		getMonitoringStatusUseCase,
		getPerizinanStatusUseCase,
		getServiceStatusUseCase,
		getSpeedtestStatusUseCase,
		getWeatherForecast,
		chartTMAUseCase)

	// controller
	printer.
		Add(controller.GetGateStatusHandler(mux, getGateStatusUsecase)).
		Add(controller.DoorControlGetAllHandler(mux, doorControlGetAllUsecase)).
		Add(controller.DoorControlHistoryGetAllHandler(mux, doorControlHistoryGetAllUsecase)).
		Add(controller.GetChartPintuAirHandler(mux, chartPintuAirUseCase)).
		Add(controller.GetChartDebitHandler(mux, chartDebitUseCase)).
		Add(controller.GetChartTMAHandler(mux, chartTMAUseCase)).
		Add(controller.GetStatusDebitHandler(mux, getStatusDebitUseCase)).
		Add(controller.GetListWaterChannelDoor(mux, listWaterChannelUseCase)).
		Add(controller.DetailPintuAirHandler(mux, detailPintuAirUseCase)).
		Add(controller.GeneralInfoHandler(mux, generalInfo)).
		Add(controller.GetWeatherForecastHandler(mux, getWeatherForecast)).
		Add(controller.GetListWaterChannelDoorCCTV(mux, listWaterChannelDoorCCTVUseCase)).
		Add(controller.GetListWaterChannelDoorOfficer(mux, listWaterChannelDoorOfficerUseCase)).
		Add(controller.GetListWaterChannelDoorWaterSurfaceElevation(mux, listWaterChannelDoorWithWaterSurfaceElevationUseCase)).
		Add(controller.GetListWaterChannelDoorDebit(mux, listWaterChannelDoorWithActualDebitUseCase)).
		Add(controller.GetListRawWater(mux, listRawWaterUseCase)).
		Add(controller.GetRawWaterDetailHandler(mux, detailRawWaterUseCase)).
		Add(controller.GetListGroundWater(mux, listGroundWaterUseCase)).
		Add(controller.GetGroundWaterDetailHandler(mux, detailGroundWaterUseCase)).
		Add(controller.GetListWeir(mux, listWeirUseCase)).
		Add(controller.GetWeirDetailHandler(mux, detailWeirUseCase)).
		Add(controller.GetListDam(mux, listDamUseCase)).
		Add(controller.GetDamDetailHandler(mux, detailDamUseCase)).
		Add(controller.GetListClimatologyMap(mux, climatologyListUseCase)).
		Add(controller.GetClimatologyDetailHandler(mux, climatologyDetailUseCase)).
		Add(controller.GetListLake(mux, lakeListUseCase)).
		Add(controller.GetLakeDetailHandler(mux, lakeDetailUseCase)).
		Add(controller.GetListWaterReservoir(mux, waterReservoirListUseCase)).
		Add(controller.GetWaterReservoirDetailHandler(mux, waterReservoirDetailUseCase)).
		Add(controller.GetListIntake(mux, intakeListUseCase)).
		Add(controller.GetIntakeDetailHandler(mux, intakeDetailUseCase)).
		Add(controller.GetListPahAbsah(mux, pahAbsahListUseCase)).
		Add(controller.GetPahAbsahDetailHandler(mux, pahAbsahDetailUseCase)).
		Add(controller.GetListCoastalProtection(mux, coastalProtectionListUseCase)).
		Add(controller.GetCoastalProtectionDetailHandler(mux, coastalProtectionDetailUseCase)).
		Add(controller.GetSedimentControl(mux, sedimentControlListUseCase)).
		Add(controller.GetSedimentControlDetailHandler(mux, sedimentControlDetailUseCase)).
		Add(controller.GetListWell(mux, wellListUseCase)).
		Add(controller.GetWellDetailHandler(mux, wellDetailUseCase)).
		Add(controller.GetListRainfall(mux, getListRainFall)).
		Add(controller.GetRainfallDetailHandler(mux, getDetailRainFall)).
		Add(controller.GetListWaterLevel(mux, getListWaterLevel)).
		Add(controller.GetWaterLevelDetailHandler(mux, getDetailWaterLevel)).
		Add(controller.GetAllCCTVHandler(mux, aiGetAllCCTVUseCase)).
		Add(controller.GetCertainCCTVHandler(mux, aiGetCertainCCTVUseCase)).
		Add(controller.GetWaterChannelDetailAIHandler(mux, aiGetDetailWaterChannelUseCase)).
		Add(controller.OpenLayerHandler(mux, aiOpenLayerUseCase)).
		Add(controller.PlayAudioHandler(mux, aiPlayAudioUseCase)).
		Add(controller.GetCitanduyArea(mux, citanduyAreaUseCase)).
		Add(controller.GetCitanduyDas(mux, citanduyDasUseCase)).
		Add(controller.OpenPhotoHandler(mux, aiOpenPhotoUseCase)).
		Add(controller.GetPhotoPathHandler(mux, aiGetPhotoPathUseCase)).
		Add(controller.BackToHomeHandler(mux, aiBackToHomeUseCase)).
		Add(controller.GetWaterRainPostDetailAIHandler(mux, aiGetRainFallDetailUseCase)).
		Add(controller.GetWaterLevelPostDetailAIHandler(mux, aiGetWaterLevelPostDetailUseCase)).
		Add(controller.GetWaterAvailabilityForecastHandler(mux, waterAvailabilityForecastUsecase)).
		Add(controller.OpenVideoHandler(mux, aiOpenVideoUseCase)).
		Add(controller.GetVideoPathHandler(mux, aiGetVideoPathUseCase)).
		Add(controller.WebhookSpeedTestHandler(mux, webHookSpeedTestUseCase)).
		Add(controller.GetDeviceStatusHandler(mux, getDeviceStatusUseCase)).
		Add(controller.GetSpeedtestStatusHandler(mux, getSpeedtestStatusUseCase)).
		Add(controller.GetPerizinanStatusHandler(mux, getPerizinanStatusUseCase)).
		Add(controller.GetDroneStatusHandler(mux, getDroneStatusUseCase)).
		Add(controller.GetServiceStatusHandler(mux, getServiceStatusUseCase)).
		Add(controller.GetMonitoringStatusHandler(mux, getMonitoringStatusUseCase)).
		Add(controller.ShowGraphHandler(mux, aiShowGraphUseCase)).
		Add(controller.AiZoomInHandler(mux, aiZoomInUseCase)).
		Add(controller.AiDoorControllRunStore(mux, aiDoorControlUseCase)).
		Add(controller.AiDoorControlRunHandler(mux, jwt, aiDoorControlRunUseCase)).
		Add(controller.AISystemSummaryHandler(mux, aiSystemSummaryUseCase)).
		Add(controller.AiCancelSensitiveJobs(mux, aiCancelSensitiveJobs)).
		Add(controller.LoginBigboardOTPSubmitHandler(mux, loginOTPSubmitToHandler)).
		Add(controller.GetBMKGLocationHandler(mux, getBMKGLocationUseCase))
}
