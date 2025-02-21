package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"

	"gorm.io/datatypes"
)

type RekapitulasiHasilKuesionerCreateUseCaseReq struct {
	NamaPemda        string `json:"nama_pemda"`
	SpipID           string `json:"spip_id"`
	Pertanyaan       string `json:"pertanyaan"`
	JawabanResponden string `json:"jawaban_responden"`
}

type RekapitulasiHasilKuesionerCreateUseCaseRes struct {
	ID string `json:"id"`
}

type RekapitulasiHasilKuesionerCreateUseCase = core.ActionHandler[RekapitulasiHasilKuesionerCreateUseCaseReq, RekapitulasiHasilKuesionerCreateUseCaseRes]

func ImplRekapitulasiHasilKuesionerCreateUseCase(
	generateId gateway.GenerateId,
	createRekapitulasiHasilKuesioner gateway.RekapitulasiHasilKuesionerSave,
	getSpipById gateway.SpipGetByID,
) RekapitulasiHasilKuesionerCreateUseCase {
	return func(ctx context.Context, req RekapitulasiHasilKuesionerCreateUseCaseReq) (*RekapitulasiHasilKuesionerCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		var jawabanResponden datatypes.JSON
		if err := json.Unmarshal([]byte(req.JawabanResponden), &jawabanResponden); err != nil {
			return nil, fmt.Errorf("failed to parse JawabanResponden: %v", err)
		}

		spip, err := getSpipById(ctx, gateway.SpipGetByIDReq{ID: req.SpipID})
		if err != nil {
			return nil, err
		}
		var namaSpip *string
		if spip.SPIP.Nama != nil {
			namaSpip = spip.SPIP.Nama
		}

		obj := model.RekapitulasiHasilKuesioner{
			ID:               &genObj.RandomId,
			NamaPemda:        &req.NamaPemda,
			SpipID:           &req.SpipID,
			NamaSpip:         namaSpip,
			Pertanyaan:       &req.Pertanyaan,
			JawabanResponden: &jawabanResponden,
		}

		if err := obj.CalculateModus(); err != nil {
			return nil, fmt.Errorf("failed to calculate Modus: %v", err)
		}

		obj.SetSimpulanKuesioner()

		if _, err = createRekapitulasiHasilKuesioner(ctx, gateway.RekapitulasiHasilKuesionerSaveReq{RekapitulasiHasilKuesioner: obj}); err != nil {
			return nil, err
		}

		return &RekapitulasiHasilKuesionerCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
