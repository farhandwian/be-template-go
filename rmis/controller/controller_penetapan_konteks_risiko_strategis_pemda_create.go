package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// PenetapanKonteksRisikoStrategisPemdaCreateHandler handles the creation of a new PenetapanKonteksRisikoStrategisPemda
func (c Controller) PenetapanKonteksRisikoStrategisPemdaCreateHandler(u usecase.PenetapanKonteksRisikoStrategisPemdaCreateUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodPost,
		Url:    "/api/penetapan-konteks-risiko-strategis-pemda",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "penetapan-konteks-risiko-strategis-pemda",
			Relation:  "create",
		},
		Body:    usecase.PenetapanKonteksRisikoStrategisPemdaCreateUseCaseReq{},
		Summary: "Create a new Penetapan Konteks Risiko Strategis Pemda",
		Tag:     "Penetapan Konteks Risiko Strategis Pemda",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		request, ok := controller.ParseJSON[usecase.PenetapanKonteksRisikoStrategisPemdaCreateUseCaseReq](w, r)
		if !ok {
			return
		}

		controller.HandleUsecase(r.Context(), w, u, request)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
