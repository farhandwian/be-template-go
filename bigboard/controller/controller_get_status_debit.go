package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
	"strconv"
)

func GetStatusDebitHandler(mux *http.ServeMux, u usecase.GetStatusDebitUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/bigboard/status-debit/{id}",
		Summary: "Get status debit",
		Tag:     "Bigboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			controller.Fail(w, err)
			return
		}

		req := usecase.GetStatusDebitReq{
			WaterChannelDoorID: id,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData

}
