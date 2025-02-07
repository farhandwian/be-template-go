package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
	"shared/usecase"
)

func GetWeatherForecastHandler(mux *http.ServeMux, u usecase.GetWeatherForecastUseCase) helper.APIData {

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodGet,
		Url:     "/bigboard/weather-forecast/{adm4}",
		Summary: "Get weather forecast info",
		Tag:     "Bigboard",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		adm4 := r.PathValue("adm4")

		req := usecase.GetWeatherForecastReq{ADM4: adm4}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)

	return apiData

}
