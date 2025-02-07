package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func ShowGraphHandler(mux *http.ServeMux, u usecase.AiShowGraphResUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/show-graph",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Ask AI to open graph",
		Tag:     "Bigboard AI",
		Body:    usecase.AiShowGraphReq{},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		request, ok := controller.ParseJSON[usecase.AiShowGraphReq](w, r)
		if !ok {
			return
		}
		req := usecase.AiShowGraphReq{
			WaterChannelDoorID: request.WaterChannelDoorID,
			GraphType:          request.GraphType,
			StartDate:          request.StartDate,
			EndDate:            request.EndDate,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
