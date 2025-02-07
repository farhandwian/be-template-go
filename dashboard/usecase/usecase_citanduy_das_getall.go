package usecase

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"shared/core"
)

type GetCitanduyDASReq struct{}

type GetCitanduyDASResp struct {
	Type     string        `json:"type"`
	Features []CitanduyDAS `json:"features"`
}

type CitanduyDAS struct {
	Type       string      `json:"type"`
	Geometry   DASGeometry `json:"geometry"`
	Properties DASProperty `json:"properties"`
}

type DASGeometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"` // Use interface{} to support multiple geometry types
}

type DASProperty struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Ordo  string `json:"ordo"`
	WspID string `json:"wsp_id"`
}

// Use case type definition for Citanduy DAS
type GetCitanduyDASUseCase = core.ActionHandler[GetCitanduyDASReq, GetCitanduyDASResp]

// Implementation of the use case for Citanduy DAS
func ImplGetCitanduyDASUseCase() GetCitanduyDASUseCase {
	return func(ctx context.Context, req GetCitanduyDASReq) (*GetCitanduyDASResp, error) {
		// Open the JSON file
		file, err := os.Open(filepath.Join("static", "citanduy-das.json"))
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Prepare the response struct to store parsed JSON data
		var response GetCitanduyDASResp

		// Decode the JSON file into the response struct
		if err := json.NewDecoder(file).Decode(&response); err != nil {
			return nil, err
		}

		// Return the parsed data
		return &response, nil
	}
}
