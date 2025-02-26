package controller

import (
	"iam/controller"
	"iam/model"
	"net/http"
	"rmis/usecase"
	"shared/helper"
)

// RancanganPemantauanDeleteHandler handles deleting a RancanganPemantauan
func (c Controller) RancanganPemantauanDeleteHandler(u usecase.RancanganPemantauanDeleteUseCase) helper.APIData {
	apiData := helper.APIData{
		Method: http.MethodDelete,
		Url:    "/api/rancangan-pemantauan/{id}",
		AccessKeto: model.AccessKetoStruct{
			Namespace: "rmis",
			Object:    "rancangan-pemantauan",
			Relation:  "delete",
		},
		Summary: "Delete a Rancangan Pemantauan",
		Tag:     "Rancangan Pemantauan",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		req := usecase.RancanganPemantauanDeleteUseCaseReq{ID: id}
		controller.HandleUsecase(r.Context(), w, u, req)
	}

	c.Mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
