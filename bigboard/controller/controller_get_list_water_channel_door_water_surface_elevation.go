package controller

import (
	"bigboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
)

func GetListWaterChannelDoorWaterSurfaceElevation(mux *http.ServeMux, u usecase.ListWaterChannelDoorWithWaterSurfaceElevationUseCase) helper.APIData {
	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/bigboard/water-channel-doors/tma",
		Summary: "Get list water channel door with water surface elevation",
		Tag:     "Bigboard",
	}
	handler := func(w http.ResponseWriter, r *http.Request) {

		req := usecase.ListWaterChannelDoorWithWaterSurfaceElevationReq{}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData
}
