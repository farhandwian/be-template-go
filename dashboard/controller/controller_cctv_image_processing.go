package controller

import (
	"dashboard/usecase"
	"iam/controller"
	"iam/model"
	"net/http"
	"shared/helper"
)

func (c Controller) CctvImageProcessingHandler(u usecase.CctvImageProcessing) helper.APIData {

	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/cctv/image-processing",
		Access:  model.DEFAULT_OPERATION,
		Body:    usecase.CctvImageProcessingReq{},
		Summary: "Create CCTV Image Processing",
		Tag:     "CCTV Image Processing",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.CctvImageProcessingReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
