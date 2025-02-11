package wiring

import (
	"iam/controller"
	"iam/gateway"
	"iam/usecase"
	"net/http"
	sharedGateway "shared/gateway"
	"shared/helper"
	ketoHelper "shared/helper/ory/keto"
	"shared/middleware"

	"gorm.io/gorm"
)

func SetupDependencyWithDatabase(apiPrinter *helper.ApiPrinter, mux *http.ServeMux, keto *ketoHelper.KetoGRPCClient, jwtToken helper.JWTTokenizer, db *gorm.DB) {

	// gateways
	generateId := gateway.ImplGenerateId()                                 //
	generateRandom := gateway.ImplGenerateRandomButStaticForMock("123456") //
	generateJWT := gateway.ImplGenerateJWT(jwtToken)                       //
	validateJWT := gateway.ImplValidateJWT(jwtToken)                       //
	passwordEncrypt := gateway.ImplPasswordEncrypt()                       //
	passwordValidate := gateway.ImplPasswordValidate()                     //
	sendEmail := gateway.ImplSendEmail()                                   //
	sendOTP := gateway.ImplSendOTP()                                       //
	userGetAll := gateway.ImplUserGetAllWithDatabase(db)                   // USE DATABASE IMPLEMENTATION
	userGetOneByID := gateway.ImplUserGetOneByIDWithDatabase(db)           // USE DATABASE IMPLEMENTATION
	userSave := gateway.ImplUserSaveWithDatabse(db)
	createActivityMonitoringGateway := sharedGateway.ImplCreateActivityMonitoringGateway(db)
	// USE DATABASE IMPLEMENTATION

	// gateway middleware
	generateIdToUseCase := middleware.Logging(generateId, 4)
	generateRandomToUseCase := middleware.Logging(generateRandom, 4)
	generateJWTToUseCase := middleware.Logging(generateJWT, 4)
	validateJWTToUseCase := middleware.Logging(validateJWT, 4)
	passwordEncryptToUseCase := middleware.Logging(passwordEncrypt, 4)
	passwordValidateToUseCase := middleware.Logging(passwordValidate, 4)
	sendEmailToUseCase := middleware.Logging(sendEmail, 4)
	sendOTPToUseCase := middleware.Logging(sendOTP, 4)
	userGetAllToUseCase := middleware.Logging(userGetAll, 4)
	userGetOneByIDToUseCase := middleware.Logging(userGetOneByID, 4)
	userSaveToUseCase := middleware.Logging(userSave, 4)

	// usecases
	accessReset := usecase.ImplAccessReset(userGetOneByIDToUseCase, userSaveToUseCase)
	userGetAccess := usecase.ImplUserGetAccess(userGetOneByIDToUseCase)
	checkAccessKeto := usecase.ImplCheckAccessKeto(keto)
	assignAccessKeto := usecase.ImplAssignAccess(keto)
	deleteAccessKeto := usecase.ImplDeleteAccessKeto(keto)
	emailActivationRequest := usecase.ImplEmailActivationRequest(userGetOneByIDToUseCase, generateJWTToUseCase, sendEmailToUseCase)
	emailActivationSubmit := usecase.ImplEmailActivationSubmit(validateJWTToUseCase, userGetOneByIDToUseCase, userSaveToUseCase, passwordEncryptToUseCase)
	loginOTPSubmit := usecase.ImplLoginOTPSubmit(passwordValidateToUseCase, userGetAllToUseCase, generateJWTToUseCase, generateIdToUseCase, userSaveToUseCase, createActivityMonitoringGateway)
	loginUseCase := usecase.ImplLogin(userGetAllToUseCase, sendOTPToUseCase, generateRandomToUseCase, passwordValidateToUseCase, passwordEncryptToUseCase, userSaveToUseCase)
	passwordChangeRequest := usecase.ImplPasswordChangeRequest(sendOTPToUseCase, userGetOneByIDToUseCase, generateRandomToUseCase, userSaveToUseCase, passwordEncryptToUseCase)
	passwordChangeSubmit := usecase.ImplPasswordChangeSubmit(passwordValidateToUseCase, passwordEncryptToUseCase, userGetOneByIDToUseCase, userSaveToUseCase)
	passwordResetRequest := usecase.ImplPasswordResetRequest(generateJWTToUseCase, sendEmailToUseCase, userGetOneByIDToUseCase)
	passwordResetSubmit := usecase.ImplPasswordResetSubmit(validateJWTToUseCase, userGetOneByIDToUseCase, userSaveToUseCase, passwordEncryptToUseCase)
	pinChangeRequest := usecase.ImplPinChangeRequest(sendOTPToUseCase, userGetOneByIDToUseCase, generateRandomToUseCase, userSaveToUseCase, passwordEncryptToUseCase)
	pinChangeSubmit := usecase.ImplPinChangeSubmit(passwordValidateToUseCase, passwordEncryptToUseCase, userGetOneByIDToUseCase, userSaveToUseCase)
	refreshToken := usecase.ImplRefreshToken(userGetOneByIDToUseCase, generateJWTToUseCase, validateJWTToUseCase)
	registerUser := usecase.ImplRegisterUser(generateIdToUseCase, userSaveToUseCase, userGetAllToUseCase, createActivityMonitoringGateway)
	userGetAllUsecase := usecase.ImplUserGetAll(userGetAllToUseCase)
	userGetOneUsecase := usecase.ImplUserGetOne(userGetOneByIDToUseCase)
	logoutUsecase := usecase.ImplLogout(userGetOneByIDToUseCase, userSaveToUseCase)

	// usecase middlewares
	accessResetToHandler := middleware.Logging(accessReset, 0)                                                           //
	userGetAccessToHandler := middleware.Logging(userGetAccess, 0)                                                       //
	emailActivationRequestToHandler := middleware.Logging(emailActivationRequest, 0)                                     //
	emailActivationSubmitToHandler := middleware.Logging(emailActivationSubmit, 0)                                       //
	loginOTPSubmitToHandler := middleware.Logging(loginOTPSubmit, 0)                                                     //
	loginToHandler := middleware.Logging(middleware.TransactionMiddleware(loginUseCase, db), 0)                          // USE TRANSACTION
	passwordChangeRequestToHandler := middleware.Logging(middleware.TransactionMiddleware(passwordChangeRequest, db), 0) // USE TRANSACTION
	passwordChangeSubmitToHandler := middleware.Logging(passwordChangeSubmit, 0)                                         //
	passwordResetRequestToHandler := middleware.Logging(passwordResetRequest, 0)                                         //
	passwordResetSubmitToHandler := middleware.Logging(passwordResetSubmit, 0)                                           //
	pinChangeRequestToHandler := middleware.Logging(middleware.TransactionMiddleware(pinChangeRequest, db), 0)           // USE TRANSACTION
	pinChangeSubmitToHandler := middleware.Logging(pinChangeSubmit, 0)                                                   //
	refreshTokenToHandler := middleware.Logging(refreshToken, 0)                                                         //
	registerUserToHandler := middleware.Logging(registerUser, 0)                                                         //
	userGetAllToHandler := middleware.Logging(userGetAllUsecase, 0)                                                      //
	userGetOneToHandler := middleware.Logging(userGetOneUsecase, 0)                                                      //
	logoutToHandler := middleware.Logging(logoutUsecase, 0)

	c := controller.Controller{
		Mux: mux,
		JWT: jwtToken,
	}

	// controllers
	apiPrinter.
		Add(c.UserGetAllHandler(userGetAllToHandler)).
		Add(c.CheckAccessKetoHandler(checkAccessKeto)).
		Add(c.AssignAccessKetoHandler(assignAccessKeto)).
		Add(c.DeleteAccessHandler(deleteAccessKeto)).
		Add(c.UserGetMeHandler(userGetOneToHandler)).
		Add(c.UserGetOneHandler(userGetOneToHandler)).
		Add(c.UserGetAccessHandler(userGetAccessToHandler)).
		Add(c.AccessResetHandler(accessResetToHandler)).
		Add(c.RegisterUserHandler(registerUserToHandler)).
		Add(c.EmailActivationRequestHandler(emailActivationRequestToHandler)).
		Add(c.EmailActivationSubmitHandler(emailActivationSubmitToHandler)).
		Add(c.LoginHandler(loginToHandler)).
		Add(c.LoginOTPSubmitHandler(loginOTPSubmitToHandler)).
		Add(c.RefreshTokenHandler(refreshTokenToHandler)).
		Add(c.LogoutHandler(logoutToHandler)).
		Add(c.PinChangeRequestHandler(pinChangeRequestToHandler)).
		Add(c.PinChangeSubmitHandler(pinChangeSubmitToHandler)).
		Add(c.PasswordChangeRequestHandler(passwordChangeRequestToHandler)).
		Add(c.PasswordChangeSubmitHandler(passwordChangeSubmitToHandler)).
		Add(c.PasswordResetRequestHandler(passwordResetRequestToHandler)).
		Add(c.PasswordResetSubmitHandler(passwordResetSubmitToHandler))

}
