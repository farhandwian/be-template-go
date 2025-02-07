package wiring

import (
	"dashboard/controller"
	"dashboard/gateway"
	"dashboard/usecase"
	"net/http"
	sharedGateway "shared/gateway"
	"shared/helper"
	"shared/helper/cronjob"
	"shared/middleware"
	sharedUseCase "shared/usecase"

	iamgw "iam/gateway"

	"gorm.io/gorm"
)

func SetupDependency(mariaDB *gorm.DB, timeScaleDB *gorm.DB, mux *http.ServeMux, jwtToken helper.JWTTokenizer, printer *helper.ApiPrinter, cj *cronjob.CronJob, sseBigboard *helper.SSE, sseDashboard *helper.SSE) {

	// gateways

	generateIdGateway := gateway.ImplGenerateId()

	getConfigGateway := sharedGateway.ImplGetConfigGateway(mariaDB)

	getListLakeGateway := sharedGateway.ImplGetLakeList(mariaDB)
	getDetailLakeGateway := sharedGateway.ImplGetLakeDetailByID(mariaDB)

	getlistDamGateway := sharedGateway.ImplGetDamList(mariaDB)
	getDetailDamGateway := sharedGateway.ImplGetDamDetailByID(mariaDB)

	getListJDIHGateway := gateway.ImplGetListJDIH(mariaDB)
	createJDIHGateway := gateway.ImplJDIHSave(mariaDB)
	getJDIHByIDGateway := gateway.ImplGetJDIHByID(mariaDB)
	deleteJDIHGateway := gateway.ImplDeleteJDIH(mariaDB)

	projectCreate := gateway.ImplProjectSave(mariaDB)
	projectGetAll := gateway.ImplProjectGetAll(mariaDB)
	projectGetByID := gateway.ImplProjectGetByID(mariaDB)
	projectDelete := gateway.ImplProjectDelete(mariaDB)

	assetCreate := gateway.ImplAssetSave(mariaDB)
	assetGetAll := gateway.ImplAssetGetAll(mariaDB)
	assetGetByID := gateway.ImplAssetGetByID(mariaDB)
	assetDelete := gateway.ImplAssetDelete(mariaDB)

	employeeCreate := gateway.ImplEmployeeSave(mariaDB)
	employeeGetAll := gateway.ImplEmployeeGetAll(mariaDB)
	employeeGetByID := gateway.ImplEmployeeGetByID(mariaDB)
	employeeDelete := gateway.ImplEmployeeDelete(mariaDB)

	getOneWaterGatesByDoorID := gateway.ImplGetOneWaterGatesByDoorID(timeScaleDB)

	weirGetAll := sharedGateway.ImplGetWeirList(mariaDB)
	weirGetByID := sharedGateway.ImplWeirGetDetailGateway(mariaDB)

	groundwaterGetAll := sharedGateway.ImplGetGroundWaterList(mariaDB)
	groundwaterGetByID := sharedGateway.ImplGroundWaterGetDetailGateway(mariaDB)

	waterReservoirGetAll := sharedGateway.ImplGetWaterReservoirList(mariaDB)
	waterReservoirGetByID := sharedGateway.ImplWaterReservoirGetDetailGateway(mariaDB)

	rawWaterGetAll := sharedGateway.ImplGetRawWaterList(mariaDB)
	rawWaterGetByID := sharedGateway.ImplRawWaterGetDetailGateway(mariaDB)

	sedimentControlGetAll := sharedGateway.ImplGetSedimentControlList(mariaDB)
	sedimentControlGetByID := sharedGateway.ImplSedimentControlGetDetailGateway(mariaDB)

	coastalProtectionGetAll := sharedGateway.ImplGetCoastalProtectionList(mariaDB)
	coastalProtectionGetByID := sharedGateway.ImplCoastalProtectionGetDetailGateway(mariaDB)

	rainFallGetByID := sharedGateway.ImplRainfallGetDetailGateway(mariaDB)

	waterLevelGetByID := sharedGateway.ImplWaterLevelGetDetailGateway(mariaDB)

	climatologyGetAll := sharedGateway.ImplGetClimatologyList(mariaDB)
	climatologyGetByID := sharedGateway.ImplGetClimatologyDetailByID(mariaDB)

	pahAbsahGetAll := sharedGateway.ImplGetPahAbsahList(mariaDB)
	pahAbsahGetByID := sharedGateway.ImplPahAbsahGetDetailGateway(mariaDB)

	wellGetAll := sharedGateway.ImplGetWellList(mariaDB)
	wellGetByID := sharedGateway.ImplWellGetDetailGateway(mariaDB)

	intakeGetAll := sharedGateway.ImplGetIntakeList(mariaDB)
	intakeGetByID := sharedGateway.ImplIntakeGetDetailGateway(mariaDB)

	alarmConfigSaveGateway := sharedGateway.ImplAlarmConfigSave(mariaDB)
	alarmConfigDeleteGateway := sharedGateway.ImplAlarmConfigDelete(mariaDB)
	alarmConfigGetAllGateway := sharedGateway.ImplAlarmConfigGetAll(mariaDB)
	alarmConfigGetOneGateway := sharedGateway.ImplAlarmConfigGetOne(mariaDB)

	doorControlSaveGateway := gateway.ImplDoorControlSave(mariaDB)
	doorControlGetAllGateway := sharedGateway.ImplDoorControlGetAll(mariaDB)
	doorControlGetOneGateway := gateway.ImplDoorControlGetOne(mariaDB)

	doorControlHistoryGetAllGateway := sharedGateway.ImplDoorControlHistoryGetAll(mariaDB)
	doorControlHistorySaveGateway := gateway.ImplDoorControlHistorySave(mariaDB)

	getListHydrologyRainPost := sharedGateway.ImplGetRainfallPostList(mariaDB)
	getListHydrologyPostOfficer := sharedGateway.ImplGetOfficerPostByNameList(mariaDB)

	getListHydrologyRainDaily := sharedGateway.ImplGetRainfallDailyByPostIDGateway(timeScaleDB)
	getDetailRainfallHourly := sharedGateway.ImplGetRainfallHourlyGateway(timeScaleDB)
	getLatestRainfallTelemetryGateway := sharedGateway.ImplGetLatestTelemetryRainPostByIDGateway(timeScaleDB)

	getListWaterLevelPost := sharedGateway.ImplGetWaterLevelPostList(mariaDB)

	getDetailWaterLevelSummaryTelemetry := sharedGateway.ImplGetWaterLevelDetailSummaryTelemetryGateway(timeScaleDB)
	getDetailWaterLevelManual := sharedGateway.ImplGetWaterLevelDetailManualGateway(timeScaleDB)
	getDetailWaterLevelTelemetry := sharedGateway.ImplGetWaterLevelDetailTelemetryGateway(timeScaleDB)

	cronjobCancel := gateway.ImplCronjobCancelDoorControl(cj)
	cronjobSet := gateway.ImplCronjobSetDoorControl(cj)

	loginMangantiGateway := gateway.ImplLoginMangantiEmpty() // TODO temporary empty login until ready

	doorControlAPIGateway := gateway.ImplDoorControlAPI_(loginMangantiGateway)
	doorControlSecurityRelayGw := gateway.ImplDoorControlSecurityRelay(loginMangantiGateway)
	doorControlSensorGw := gateway.ImplDoorControlSensor(loginMangantiGateway)

	getOneUser := iamgw.ImplUserGetOneByIDWithDatabase(mariaDB)
	getListWaterQualityPost := gateway.ImplGetWaterQualityPostList(mariaDB)

	passwordValidate := iamgw.ImplPasswordValidate()

	getOneWaterChannelDeviceByDoorID := gateway.ImplGetOneWaterChannelDeviceByDoorID(mariaDB)

	getListWaterChannelDoorGateway := sharedGateway.ImplGetListWaterChannelDoor(mariaDB)
	getListWaterChannelTMAGateway := sharedGateway.ImplGetListWaterChannelTMA(timeScaleDB)
	getListWaterChannelActualDebitGateway := sharedGateway.ImplGetListWaterChannelActualDebit(timeScaleDB)
	getListCCTVCountGateway := gateway.ImplGetListCCTVCountGateway(mariaDB)

	deviceByWaterChannelDoorIdGateway := gateway.ImplDeviceByWaterChannelDoorId(mariaDB)
	waterChannelDoorByKeywordGateway := gateway.ImplWaterChannelDoorByKeyword(mariaDB)
	getCCTBDevices := sharedGateway.ImplGetCCTVDevices(mariaDB)

	waterChannelDoorByIDGateway := gateway.ImplWaterChannelDoorByID(mariaDB)
	waterChannelDoorOfficerGateway := gateway.ImplGetListWaterChannelDoorOfficer(mariaDB)

	waterChannelListGateway := gateway.ImplGetWaterChannelList(mariaDB)
	getLatestWaterGatesByDoorID := sharedGateway.ImplGetLatestWaterGatesByDoorID(timeScaleDB)
	getLatestWaterSurfaceElevation := sharedGateway.ImplGetLatestWaterSurfaceElevationByDoorID(timeScaleDB)

	getDebitGateway := sharedGateway.ImplGetDebit(timeScaleDB)
	getWaterSurfaceElevationGateway := sharedGateway.ImplGetWaterSurfaceElevation(timeScaleDB)
	getGateGateway := sharedGateway.ImplGetGate(timeScaleDB)
	getDevicesGateway := sharedGateway.ImplGetDevice(mariaDB)

	getListWaterLevelManualPost := sharedGateway.ImplGetListLatestManualWaterPostHandler(timeScaleDB)
	getListWaterLevelTelemetryPost := sharedGateway.ImplGetListLatestTelemetryWaterPostGateway(timeScaleDB)

	fetchBMKGWeatherForecast := sharedGateway.ImplFetchBMKGWeatherForecast()
	generalInfoUseCase := sharedGateway.ImplGetGeneralInfoGateway(mariaDB)

	waterAvailabilityForecastGateway := sharedGateway.ImplGetWaterAvailabilityForecastGateway(timeScaleDB)
	activityMonitoringGetAllGateway := sharedGateway.ImplGetActivityMonitoringGateway(mariaDB)

	getRequiredDebitMainWaterChannelDoorGateway := sharedGateway.ImplGetTotalRequirementDebitMainWaterChannelDoorGateway(mariaDB)
	getAlarmCountGateway := sharedGateway.ImplGetAlarmCountGateway(mariaDB)

	getBMKGLocationGateway := sharedGateway.ImplFetchBMKGWeatherLocationGateway()

	cctvImageProcessingCreate := gateway.ImplCctvImageProcessingSave(mariaDB, timeScaleDB)
	getWaterChannelDoorById := sharedGateway.ImplGetWaterChannelDoorByID(mariaDB)

	getLaporanPerizinanStatusGateway := sharedGateway.ImplGetLaporanPerizinanStatusCount(mariaDB)
	// usecase
	listWaterChannelDoorUseCase := usecase.ImplListWaterChannel(
		getListWaterChannelDoorGateway,
		getListWaterChannelTMAGateway,
		getListWaterChannelActualDebitGateway,
		getListCCTVCountGateway,
		waterChannelDoorOfficerGateway,
		getCCTBDevices,
	)

	getWaterChannelDoor := sharedGateway.ImplGetWaterChannelDoorByID(mariaDB)
	getWaterChannelDevices := sharedGateway.ImplGetWaterChannelDevicesByDoorID(mariaDB)
	getWaterChannelOfficers := sharedGateway.ImplGetWaterChannelOfficersByDoorID(mariaDB)
	getLatestDebitGateway := sharedGateway.ImplGetLatestDebit(timeScaleDB)

	getDeviceStatusGateway := sharedGateway.ImplGetDeviceStatusGateway()
	getSpeedtestStatusGateway := sharedGateway.ImplGetSpeedtestStatusGateway()
	getServiceStatusGateway := sharedGateway.ImplGetServiceStatusGateway()

	getTmaDeviceCountGateway := sharedGateway.ImplGetTmaDeviceCount(timeScaleDB)

	createMonitoringGateway := sharedGateway.ImplCreateActivityMonitoringGateway(mariaDB)

	projectCreateUseCase := usecase.ImplProjectCreateUseCase(generateIdGateway, projectCreate)
	projectGetAllUseCase := usecase.ImplProjectGetAllUseCase(projectGetAll)
	projectGetByIDUseCase := usecase.ImplProjectGetByIDUseCase(projectGetByID)
	projectUpdateUseCase := usecase.ImplProjectUpdateUseCase(projectGetByID, projectCreate)
	projectDeleteUseCase := usecase.ImplProjectDeleteUseCase(projectDelete)

	assetCreateUseCase := usecase.ImplAssetCreateUseCase(generateIdGateway, assetCreate)
	assetGetAllUseCase := usecase.ImplAssetGetAllUseCase(assetGetAll)
	assetGetByIDUseCase := usecase.ImplAssetGetByIDUseCase(assetGetByID)
	assetUpdateUseCase := usecase.ImplAssetUpdateUseCase(assetGetByID, assetCreate)
	assetDeleteUseCase := usecase.ImplAssetDeleteUseCase(assetDelete)

	employeeCreateUseCase := usecase.ImplEmployeeCreateUseCase(generateIdGateway, employeeCreate)
	employeeGetAllUseCase := usecase.ImplEmployeeGetAllUseCase(employeeGetAll)
	employeeGetByIDUseCase := usecase.ImplEmployeeGetByIDUseCase(employeeGetByID)
	employeeUpdateUseCase := usecase.ImplEmployeeUpdateUseCase(employeeGetByID, employeeCreate)
	employeeDeleteUseCase := usecase.ImplEmployeeDeleteUseCase(employeeDelete)

	getJDIHUseCase := usecase.ImplGetListDocumentAndLawUseCase(getListJDIHGateway)
	createJDIHUseCase := usecase.ImplCreateJDIHUseCase(generateIdGateway, createJDIHGateway)
	getJDIHByIDUseCase := usecase.ImplGetJDIHByIDUseCase(getJDIHByIDGateway)
	updateJDIHUseCase := usecase.ImplUpdateJDIHUseCase(getJDIHByIDGateway, createJDIHGateway)
	deleteJDIHUseCase := usecase.ImplDeleteJDIHUseCase(deleteJDIHGateway)

	getListLakeUseCase := sharedUseCase.ImplGetListLakeUseCase(getListLakeGateway)
	getDetailLakeUseCase := sharedUseCase.ImplGetLakeDetail(getDetailLakeGateway)

	getListDamUseCase := sharedUseCase.ImplGetListDamUseCase(getlistDamGateway)
	getDetailDamUseCase := sharedUseCase.ImplGetDamDetail(getDetailDamGateway)

	getListWeirUseCase := sharedUseCase.ImplGetListWeirUseCase(weirGetAll)
	getDetailWeirUseCase := sharedUseCase.ImplGetWeirDetail(weirGetByID)

	getListWaterReservoir := sharedUseCase.ImplGetListWaterReservoirUseCase(waterReservoirGetAll)
	getDetailWaterReservoir := sharedUseCase.ImplGetWaterReservoirDetail(waterReservoirGetByID)

	getListGroundWater := sharedUseCase.ImplGetListGroundWaterUseCase(groundwaterGetAll)
	getDetailGroundWater := sharedUseCase.ImplGetGroundwaterDetail(groundwaterGetByID)

	getListRawWater := sharedUseCase.ImplGetListRawWaterUseCase(rawWaterGetAll)
	getDetailRawWater := sharedUseCase.ImplGetRawWaterDetail(rawWaterGetByID)

	getListSedimentControl := sharedUseCase.ImplGetListSedimentControlUseCase(sedimentControlGetAll)
	getDetailSedimentControl := sharedUseCase.ImplGetSedimentControlDetail(sedimentControlGetByID)

	getListCoastalProtection := sharedUseCase.ImplGetListCoastalProtectionUseCase(coastalProtectionGetAll)
	getDetailCoastalProtection := sharedUseCase.ImplGetCoastalProtectionDetail(coastalProtectionGetByID)

	getDetailRainFall := sharedUseCase.ImplGetRainFallDetail(rainFallGetByID, getListHydrologyPostOfficer, getLatestRainfallTelemetryGateway)
	getListRainFall := sharedUseCase.ImplGetListRainFallUseCase(getListHydrologyRainPost, getLatestRainfallTelemetryGateway)

	getListWaterLevel := sharedUseCase.ImplGetListWaterLevelUseCase(getListWaterLevelPost, getListWaterLevelManualPost, getListWaterLevelTelemetryPost)
	getDetailWaterLevel := sharedUseCase.ImplGetWaterLevelDetail(waterLevelGetByID, getListHydrologyPostOfficer)

	getListPahAbsah := sharedUseCase.ImplGetListPahAbsahUseCase(pahAbsahGetAll)
	getDetailPahAbsah := sharedUseCase.ImplGetPahAbsahDetail(pahAbsahGetByID)

	getListWell := sharedUseCase.ImplGetListWellUseCase(wellGetAll)
	getDetailWell := sharedUseCase.ImplGetWellDetail(wellGetByID)

	getListIntake := sharedUseCase.ImplGetListIntakeUseCase(intakeGetAll)
	getDetailIntake := sharedUseCase.ImplGetIntakeDetail(intakeGetByID)

	grafanaUpsert := gateway.ImplGrafanaUpsert(nil)
	grafanaDelete := gateway.ImplAlarmConfigXDelete(nil)

	alarmHistoryGetAll := sharedGateway.ImplAlarmHistoryGetAll(mariaDB)

	alarmConfigCreateUsecase := middleware.TransactionMiddleware(usecase.ImplAlarmConfigCreate(generateIdGateway, alarmConfigSaveGateway, grafanaUpsert, waterChannelDoorByIDGateway, deviceByWaterChannelDoorIdGateway), mariaDB)
	alarmConfigDeleteUsecase := middleware.TransactionMiddleware(usecase.ImplAlarmConfigDelete(alarmConfigDeleteGateway, grafanaDelete), mariaDB)
	alarmConfigGetAllUsecase := usecase.ImplAlarmConfigGetAll(alarmConfigGetAllGateway)
	alarmConfigGetOneUsecase := usecase.ImplAlarmConfigGetOne(alarmConfigGetOneGateway)
	alarmConfigUpdateUsecase := middleware.TransactionMiddleware(usecase.ImplAlarmConfigUpdate(alarmConfigGetOneGateway, alarmConfigSaveGateway, grafanaUpsert, waterChannelDoorByIDGateway, deviceByWaterChannelDoorIdGateway), mariaDB)

	doorControlCreateUsecase := usecase.ImplDoorControlCreate(generateIdGateway, doorControlSaveGateway, getOneUser, cronjobSet, getOneWaterChannelDeviceByDoorID, createMonitoringGateway, getWaterChannelDoor)
	doorControlCancelUsecase := usecase.ImplDoorControlCancel(doorControlGetOneGateway, doorControlSaveGateway, cronjobCancel, generateIdGateway, getOneWaterGatesByDoorID, doorControlHistorySaveGateway)
	doorControlGetAllUsecase := sharedUseCase.ImplDoorControlGetAll(doorControlGetAllGateway)
	doorControlRunUsecase := usecase.ImplDoorControlRunScheduled(doorControlGetOneGateway, doorControlAPIGateway, doorControlSaveGateway, generateIdGateway, doorControlHistorySaveGateway, getOneWaterGatesByDoorID, getOneWaterChannelDeviceByDoorID, createMonitoringGateway, getWaterChannelDoor)
	doorControlRunDirectly := usecase.ImplDoorControlRunDirectly(passwordValidate, doorControlAPIGateway, generateIdGateway, doorControlHistorySaveGateway, getOneWaterGatesByDoorID, getOneUser, getOneWaterChannelDeviceByDoorID, createMonitoringGateway, getWaterChannelDoor)

	doorControlRunSecurityRelay := usecase.ImplDoorControlRunSecurityRelay(passwordValidate, doorControlSecurityRelayGw, generateIdGateway, doorControlHistorySaveGateway, getOneWaterGatesByDoorID, getOneUser, getOneWaterChannelDeviceByDoorID, createMonitoringGateway, getWaterChannelDoor)
	doorControlRunSensor := usecase.ImplDoorControlRunSensor(passwordValidate, doorControlSensorGw, generateIdGateway, doorControlHistorySaveGateway, getOneWaterGatesByDoorID, getOneUser, getOneWaterChannelDeviceByDoorID, createMonitoringGateway, getWaterChannelDoor)

	doorControlHistoryGetAllUsecase := sharedUseCase.ImplDoorControlHistoryGetAll(doorControlHistoryGetAllGateway)

	getListHydrologyRainPostUseCase := usecase.ImplListRainfallPost(
		getListHydrologyRainPost,
		// getListHydrologyPostOfficer,
		getListHydrologyRainDaily,
		getLatestRainfallTelemetryGateway,
	)

	getListWaterLevelPostUseCase := sharedUseCase.ImplListWaterLevelPost(
		getListWaterLevelPost,
		getListWaterLevelManualPost,
		getListWaterLevelTelemetryPost,
	)

	getListClimatologyInfra := usecase.ImplGetListClimatologyUseCase(climatologyGetAll)

	getDetailClimatologyInfra := sharedUseCase.ImplGetClimatologyDetailPostUseCase(climatologyGetByID)

	getListClimatologyPost := sharedUseCase.ImplGetClimatologyPostUseCase(climatologyGetAll)

	getCitanduyAreaMap := usecase.ImplGetCitanduyAreaUseCase()
	getCitanduyDasMap := usecase.ImplGetCitanduyDASUseCase()

	getDetailRainfallPostWithCalculation := usecase.ImplRainfallDetailWithCalculationUseCase(getLatestRainfallTelemetryGateway, rainFallGetByID, getDetailRainfallHourly)
	getDetailWaterLevelPost := sharedUseCase.ImplDetailWaterLevelPost(
		getDetailWaterLevelSummaryTelemetry,
		waterLevelGetByID,
		getDetailWaterLevelManual,
		getDetailWaterLevelTelemetry,
	)

	getListWaterQualityPostUseCase := usecase.ImplListWaterQualityPost(getListWaterQualityPost)

	waterChannelDoorByKeywordUsecase := usecase.ImplWaterChannelDoorByKeyword(waterChannelDoorByKeywordGateway)
	deviceByWaterChannelDoorIdUsecase := usecase.ImplDeviceByWaterChannelDoorId(deviceByWaterChannelDoorIdGateway)
	listWaterChannelUseCase := usecase.ImplListWaterChannelUseCase(waterChannelListGateway)

	detailPintuAirUseCase := sharedUseCase.ImplDetailPintuAir(
		getWaterChannelDoor,
		getWaterChannelDevices,
		getWaterChannelOfficers,
		// getLatestWaterGatesByDoorID,
		// getLatestWaterSurfaceElevation,
	)

	getGateStatusUsecase := sharedUseCase.ImplGetGateStatus(getWaterChannelDevices, getLatestWaterGatesByDoorID)
	getStatusDebitUseCase := sharedUseCase.ImplGetStatusDebit(getWaterChannelDoor, getLatestWaterSurfaceElevation, getLatestDebitGateway)

	chartDebitUseCase := sharedUseCase.ImplChartDebitUseCase(getDebitGateway)
	chartTMAUseCase := sharedUseCase.ImplChartTMAUseCase(getWaterSurfaceElevationGateway)
	chartPintuAirUseCase := sharedUseCase.ImplChartPintuAirUseCase(getDevicesGateway, getGateGateway)

	listMainWaterChannelDoorUseCase := usecase.ImplListMainWaterChannelDoorUseCase(
		getListWaterChannelDoorGateway,
		getListWaterChannelTMAGateway,
		getListWaterChannelActualDebitGateway,
	)

	waterAvailabilityForecastUsecase := sharedUseCase.ImplWaterAvailabilityForecastUseCase(waterAvailabilityForecastGateway)
	activityMonitoringUsecase := sharedUseCase.ImpActivityMonitorGetAllUseCase(activityMonitoringGetAllGateway)

	sseBigboardGW := sharedGateway.ImplSendSSEMessage(sseBigboard)
	sseDashboardGW := sharedGateway.ImplSendSSEMessage(sseDashboard)

	email := gateway.ImplSendEmailUsingGmail()
	wa := gateway.ImplSendWhatsApp()

	alarmHistorySaveGateway := gateway.ImplAlarmHistorySave(mariaDB)

	alarmConfigWebhookUseCase := usecase.ImplAlarmConfigWebhook(alarmConfigGetOneGateway, sseBigboardGW, sseDashboardGW, email, wa, alarmHistorySaveGateway, createMonitoringGateway, generateIdGateway)

	alarmHistoryUsecase := sharedUseCase.ImplAlarmHistoryGetAll(alarmHistoryGetAll)

	getWeatherForecast := sharedUseCase.ImplGetWeatherForecast(fetchBMKGWeatherForecast)

	generalInfoUsecase := sharedUseCase.ImplGeneralInfo(generalInfoUseCase, getRequiredDebitMainWaterChannelDoorGateway)

	healthCheckUsecase := usecase.ImplHealthCheckUseCase()

	getDeviceStatusUseCase := sharedUseCase.ImplGetDeviceStatusUseCase(getDeviceStatusGateway, getTmaDeviceCountGateway, getConfigGateway)
	getSpeedtestStatusUseCase := sharedUseCase.ImplGetSpeedtestStatusUseCase(getSpeedtestStatusGateway)
	getPerizinanStatusUseCase := sharedUseCase.ImplGetPerizinanStatusUseCase(getLaporanPerizinanStatusGateway)
	getDroneStatusUseCase := sharedUseCase.ImplGetDroneStatusUseCase()
	getServiceStatusUseCase := sharedUseCase.ImplGetServiceStatusUseCase(getServiceStatusGateway)
	getMonitoringStatusUseCase := sharedUseCase.ImplGetMonitoringStatusUseCase(getAlarmCountGateway)

	getBMKGLocationUseCase := sharedUseCase.ImplGetBMKGLocationUseCase(getBMKGLocationGateway)

	cctvImageProcessingUseCase := usecase.ImplCctvImageProcessing(generateIdGateway, cctvImageProcessingCreate, getWaterChannelDoorById)

	c := controller.Controller{
		Mux: mux,
		JWT: jwtToken,
	}

	controller.RunDoorControl(cj, doorControlRunUsecase)

	printer.
		Add(c.ProjectCreateHandler(projectCreateUseCase)).
		Add(c.ProjectGetAllHandler(projectGetAllUseCase)).
		Add(c.ProjectGetByIDHandler(projectGetByIDUseCase)).
		Add(c.ProjectUpdateHandler(projectUpdateUseCase)).
		Add(c.ProjectDeleteHandler(projectDeleteUseCase)).
		Add(c.AssetCreateHandler(assetCreateUseCase)).
		Add(c.AssetGetAllHandler(assetGetAllUseCase)).
		Add(c.AssetGetByIDHandler(assetGetByIDUseCase)).
		Add(c.AssetUpdateHandler(assetUpdateUseCase)).
		Add(c.AssetDeleteHandler(assetDeleteUseCase)).
		Add(c.EmployeeCreateHandler(employeeCreateUseCase)).
		Add(c.EmployeeGetAllHandler(employeeGetAllUseCase)).
		Add(c.EmployeeGetByIDHandler(employeeGetByIDUseCase)).
		Add(c.EmployeeUpdateHandler(employeeUpdateUseCase)).
		Add(c.EmployeeDeleteHandler(employeeDeleteUseCase)).
		Add(c.GetListJDIH(getJDIHUseCase)).
		Add(c.CreateJDIH(createJDIHUseCase)).
		Add(c.GetJDIHByIDHandler(getJDIHByIDUseCase)).
		Add(c.UpdateJDIHHandler(updateJDIHUseCase)).
		Add(c.DeleteJDIH(deleteJDIHUseCase)).
		Add(c.AlarmConfigCreateHandler(alarmConfigCreateUsecase)).
		Add(c.AlarmConfigUpdateHandler(alarmConfigUpdateUsecase)).
		Add(c.AlarmConfigGetAllHandler(alarmConfigGetAllUsecase)).
		Add(c.AlarmConfigGetOneHandler(alarmConfigGetOneUsecase)).
		Add(c.AlarmConfigDeleteHandler(alarmConfigDeleteUsecase)).
		Add(c.AlarmHistoryGetAllHandler(alarmHistoryUsecase)).
		Add(c.DoorControlCreateHandler(doorControlCreateUsecase)).
		Add(c.DoorControlGetAllHandler(doorControlGetAllUsecase)).
		Add(c.DoorControlCancelHandler(doorControlCancelUsecase)).
		Add(c.DoorControlRunHandler(doorControlRunDirectly)).
		Add(c.DoorControlRunSecurityRelayHandler(doorControlRunSecurityRelay)).
		Add(c.DoorControlRunSensorHandler(doorControlRunSensor)).
		Add(c.DoorControlHistoryGetAllHandler(doorControlHistoryGetAllUsecase)).
		Add(c.GetListLake(getListLakeUseCase)).
		Add(c.GetLakeDetailHandler(getDetailLakeUseCase)).
		Add(c.GetListDam(getListDamUseCase)).
		Add(c.GetDamDetailHandler(getDetailDamUseCase)).
		Add(c.GetWeirDetailHandler(getDetailWeirUseCase)).
		Add(c.GetListWeir(getListWeirUseCase)).
		Add(c.GetListWaterReservoir(getListWaterReservoir)).
		Add(c.GetWaterReservoirDetailHandler(getDetailWaterReservoir)).
		Add(c.GetListGroundWater(getListGroundWater)).
		Add(c.GetGroundWaterDetailHandler(getDetailGroundWater)).
		Add(c.GetListRawWater(getListRawWater)).
		Add(c.GetRawWaterDetailHandler(getDetailRawWater)).
		Add(c.GetSedimentControl(getListSedimentControl)).
		Add(c.GetSedimentControlDetailHandler(getDetailSedimentControl)).
		Add(c.GetListCoastalProtection(getListCoastalProtection)).
		Add(c.GetCoastalProtectionDetailHandler(getDetailCoastalProtection)).
		Add(c.GetListPahAbsah(getListPahAbsah)).
		Add(c.GetPahAbsahDetailHandler(getDetailPahAbsah)).
		Add(c.GetListWell(getListWell)).
		Add(c.GetWellDetailHandler(getDetailWell)).
		Add(c.GetListIntake(getListIntake)).
		Add(c.GetIntakeDetailHandler(getDetailIntake)).
		Add(c.GetListWaterChannelDoor(listWaterChannelDoorUseCase)).
		Add(c.WaterChannelDoorByKeywordHandler(waterChannelDoorByKeywordUsecase)).
		Add(c.DeviceByWaterChannelDoorIdHandler(deviceByWaterChannelDoorIdUsecase)).
		Add(c.RainfallPostGetAllHandler(getListHydrologyRainPostUseCase)).
		Add(c.WaterLevelPostGetAllHandler(getListWaterLevelPostUseCase)).
		Add(c.GetListClimatologyMap(getListClimatologyPost)).
		Add(c.GetCitanduyArea(getCitanduyAreaMap)).
		Add(c.GetCitanduyDas(getCitanduyDasMap)).
		Add(c.RainfallPostCalculationGetDetailHandler(getDetailRainfallPostWithCalculation)).
		Add(c.WaterLevelPostGetDetailHandler(getDetailWaterLevelPost)).
		Add(c.WaterQualityPostGetAllHandler(getListWaterQualityPostUseCase)).
		Add(c.GetListRainfall(getListRainFall)).
		Add(c.GetRainfallDetailHandler(getDetailRainFall)).
		Add(c.GetListWaterLevel(getListWaterLevel)).
		Add(c.GetWaterLevelDetailHandler(getDetailWaterLevel)).
		Add(c.GetClimatologyDetailHandler(getDetailClimatologyInfra)).
		Add(c.GetListClimatology(getListClimatologyInfra)).
		Add(c.GetListWaterChannel(listWaterChannelUseCase)).
		Add(c.DetailWaterChannelDoorHandler(detailPintuAirUseCase)).
		Add(c.GetGateStatusHandler(getGateStatusUsecase)).
		Add(c.GetStatusDebitHandler(getStatusDebitUseCase)).
		Add(c.GetChartPintuAirHandler(chartPintuAirUseCase)).
		Add(c.GetChartDebitHandler(chartDebitUseCase)).
		Add(c.GetChartTMAHandler(chartTMAUseCase)).
		Add(c.AlarmConfigWebhookHandler(alarmConfigWebhookUseCase)).
		Add(c.GetWeatherForecastHandler(getWeatherForecast)).
		Add(c.GeneralInfoHandler(generalInfoUsecase)).
		Add(c.GetListMainWaterChannelDoor(listMainWaterChannelDoorUseCase)).
		Add(c.GetWaterAvailabilityForecastHandler(waterAvailabilityForecastUsecase)).
		Add(c.ActivityMonitorGetAllHandler(activityMonitoringUsecase)).
		Add(c.HealthCheckHandler(healthCheckUsecase)).
		Add(c.GetDeviceStatusHandler(getDeviceStatusUseCase)).
		Add(c.GetSpeedtestStatusHandler(getSpeedtestStatusUseCase)).
		Add(c.GetPerizinanStatusHandler(getPerizinanStatusUseCase)).
		Add(c.GetDroneStatusHandler(getDroneStatusUseCase)).
		Add(c.GetServiceStatusHandler(getServiceStatusUseCase)).
		Add(c.GetMonitoringStatusHandler(getMonitoringStatusUseCase)).
		Add(c.GetBMKGLocationHandler(getBMKGLocationUseCase)).
		Add(c.CctvImageProcessingHandler(cctvImageProcessingUseCase))
}
