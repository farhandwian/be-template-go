package controller

import (
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
	"strconv"
)

func GetGateStatusHandler(mux *http.ServeMux, u usecase.GetGateStatusUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/bigboard/gatestatus/{id}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Gates Status",
		Tag:     "Bigboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			controller.Fail(w, err)
			return
		}

		req := usecase.GetGateStatusReq{WaterChannelDoorID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
