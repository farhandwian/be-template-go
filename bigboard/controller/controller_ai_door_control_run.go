package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iamModel "iam/model"
	"net/http"
	"shared/core"
	"shared/helper"
)

func AiDoorControlRunHandler(mux *http.ServeMux, jwt helper.JWTTokenizer, u usecase.AIDoorControlRunDirectlyUseCase) helper.APIData {

	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/doorcontrol-run",
		Access:  iamModel.PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_PINTU_AIR_UPDATE,
		Summary: "Ask AI to run door control",
		Tag:     "Bigboard AI- Door Control",
		Body:    usecase.AIDoorControlRunDirectlyReq{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		request, ok := controller.ParseJSON[usecase.AIDoorControlRunDirectlyReq](w, r)
		if !ok {
			return
		}
		req := usecase.AIDoorControlRunDirectlyReq{
			IDSensitveJob: request.IDSensitveJob,
			Pin:           string(request.Pin),
			OfficerId:     core.GetDataFromContext[iamModel.UserID](r.Context(), controller.UserIDContext),
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	authorizationHandler := controller.Authorization(handler, apiData.Access)
	authenticatedHandler := controller.Authentication(authorizationHandler, jwt)

	mux.HandleFunc(apiData.GetMethodUrl(), authenticatedHandler)
	return apiData
}
