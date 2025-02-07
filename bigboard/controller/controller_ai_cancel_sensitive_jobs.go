package controller

import (
	"bigboard/usecase"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"shared/helper"
)

func AiCancelSensitiveJobs(mux *http.ServeMux, u usecase.AiAiCancelSensitiveJobsUseCase) helper.APIData {
	type Body struct {
		IDSensitveJob string `json:"id_sensitive_job"`
	}
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/bigboard/ai/cancel-sensitive-jobs",
		Body:    Body{},
		Access:  iammodel.PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_PINTU_AIR_UPDATE,
		Summary: "Cancel sensitive Jobs",
		Tag:     "Bigboard AI",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		request, ok := controller.ParseJSON[usecase.AiCancelSensitiveJobsReq](w, r)
		if !ok {
			return
		}

		req := usecase.AiCancelSensitiveJobsReq{
			IDSensitveJob: request.IDSensitveJob,
		}

		controller.HandleUsecase(r.Context(), w, u, req)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
