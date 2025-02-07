package usecase

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"shared/core"
)

type GetCitanduyAreaReq struct{}

type GetCitanduyAreaResp struct {
	Type     string         `json:"type"`
	Features []CitanduyArea `json:"features"`
}

type CitanduyArea struct {
	Type       string               `json:"type"`
	Geometry   CitanduyAreaGeometry `json:"geometry"`
	Properties CitanduyAreaProperty `json:"properties"`
}

type CitanduyAreaGeometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type CitanduyAreaProperty struct {
	ID      string  `json:"id"`
	Luas    float64 `json:"luas"`
	Name    string  `json:"name"`
	Color   string  `json:"color"`
	WSID    string  `json:"ws_id"`
	KodeDAS string  `json:"kode_das"`
}

type GetCitanduyAreaUseCase = core.ActionHandler[GetCitanduyAreaReq, GetCitanduyAreaResp]

func ImplGetCitanduyAreaUseCase() GetCitanduyAreaUseCase {
	return func(ctx context.Context, req GetCitanduyAreaReq) (*GetCitanduyAreaResp, error) {

		file, err := os.Open(filepath.Join("static", "citanduy-area.json"))
		if err != nil {
			return nil, err
		}
		defer file.Close()

		var response GetCitanduyAreaResp

		if err := json.NewDecoder(file).Decode(&response); err != nil {
			return nil, err
		}

		return &response, nil
	}
}
