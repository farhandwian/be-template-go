package controller

import (
	"errors"
	"fmt"
	"iam/controller"
	iammodel "iam/model"
	"net/http"
	"os"
	"path/filepath"
	"perizinan/model"
	"perizinan/usecase"
	"shared/helper"
	"strings"
	"time"
)

type Response struct {
	Message     string              `json:"message"`
	Attachments []model.Attachments `json:"attachments"`
}

// ValidFileNames contains the allowed enum values for file_name field
var ValidFileNames = map[string]bool{
	"LaporanHasilPemeriksaanTimVerifikasi": true,
	"LaporanHasilPemeriksaanBuangan":       true,
	"DokumenBuktiBayar":                    true,
	"DokumenKewajibanKeuanganLainnya":      true,
	"BuktiKerusakanSumberAir":              true,
	"BuktiUsahaPengendalianPencemaran":     true,
	"BuktiPenggunaanAir":                   true,
	"FileKeberadaanAlatUkurDebit":          true,
	"FileKeberadaanSistemTelemetri":        true,
	"LaporanPengambilanAirBulanan":         true,
}

func getValidFileNamesString() string {
	// Create a slice to store the keys
	keys := make([]string, 0, len(ValidFileNames))

	// Collect all keys from the map
	for k := range ValidFileNames {
		keys = append(keys, k)
	}

	// Join all keys with comma and space
	return strings.Join(keys, ", ")
}

func UploadFileHandler(mux *http.ServeMux, u usecase.LaporanPerizinanExistUseCase) helper.APIData {
	apiData := helper.APIData{
		Method:  http.MethodPost,
		Url:     "/api/upload",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Upload files",
		Tag:     "Laporan Perizinan",
		MultipartFormParam: []helper.MultipartFormParam{
			{
				Name:        "nomor_sk",
				Type:        "string",
				Description: "Reference number (e.g., 2225/KPTS/M/2024)",
				Required:    true,
			},
			{
				Name:        "periode_pengambilan_sda",
				Type:        "string",
				Description: "Period in YYYY-MM format (must be future date, e.g., 2024-08)",
				Required:    true,
			},
			{
				Name:        "file_name",
				Type:        "string",
				Description: "File name category (enum: " + getValidFileNamesString() + ")",
				Required:    true,
			},
			{
				Name:        "files",
				Type:        "file",
				Description: "Multiple files to upload",
				Required:    true,
			},
		},
		Examples: []helper.ExampleResponse{
			{
				StatusCode: 200,
				Content: Response{
					Message: "files uploaded",
					Attachments: []model.Attachments{
						{
							Label: "hello.txt",
							File:  "2225-KPTS-M-2024/2024-08/KeberadaanAlatUkurDebit_hello.txt",
						},
						{
							Label: "bisa.doc",
							File:  "2225-KPTS-M-2024/2024-08/KeberadaanAlatUkurDebit_bisa.doc",
						},
					},
				},
			},
			{
				StatusCode: 400,
				Content: map[string]string{
					"error": "invalid file_name. Must be one of: " + getValidFileNamesString() + "",
				},
			},
		},
	}

	uploadPath := os.Getenv("FILE_UPLOAD")

	handler := func(w http.ResponseWriter, r *http.Request) {

		// Parse multipart form with 32MB max memory
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			controller.Fail(w, errors.New("error parsing form"))
			return
		}

		// Get and validate nomor_sk
		nomorSk := r.FormValue("nomor_sk")
		if nomorSk == "" {
			controller.Fail(w, errors.New("nomor_sk is required"))
			return
		}

		// Get and validate periode_pengambilan_sda
		periode := r.FormValue("periode_pengambilan_sda")

		// check nomor_sk existance and make sure laporan is not submitted yet
		request := usecase.LaporanPerizinanExistUseCaseReq{
			NomorSK:               model.NomorSK(nomorSk),
			PeriodePengambilanSDA: model.Periode(periode),
			Min:                   time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC),
			Now:                   time.Now(),
		}
		if _, err := u(r.Context(), request); err != nil {
			controller.Fail(w, err)
			return
		}

		// Get and validate file_name
		fileName := r.FormValue("file_name")
		if !ValidFileNames[fileName] {
			controller.Fail(w, fmt.Errorf("invalid file_name. Must be one of: %s", getValidFileNamesString()))
			return
		}

		// Process the normalized nomor_sk for folder name
		normalizedNomorSk := strings.ReplaceAll(nomorSk, "/", "-")

		// Create folder for this nomor_sk
		folderPath := filepath.Join(uploadPath, normalizedNomorSk, periode)
		if err := os.MkdirAll(folderPath, 0755); err != nil {
			controller.Fail(w, fmt.Errorf("error creating directory"))
			return
		}

		// Get uploaded files
		files := r.MultipartForm.File["files"]
		if len(files) == 0 {
			controller.Fail(w, fmt.Errorf("no files uploaded"))
			return
		}

		var attachments []model.Attachments

		// Process each file
		for _, fileHeader := range files {
			originalName := fileHeader.Filename
			// New file name format: file-abc_hello.txt
			newFileName := fmt.Sprintf("%s_%s", fileName, originalName)
			// Full path including the nomor_sk folder
			filePath := filepath.Join(folderPath, newFileName)

			// Open the uploaded file
			file, err := fileHeader.Open()
			if err != nil {
				controller.Fail(w, fmt.Errorf("error processing file"))
				return
			}

			// Create the destination file
			dst, err := os.Create(filePath)
			if err != nil {
				file.Close()
				controller.Fail(w, fmt.Errorf("error saving file"))
				return
			}

			// Copy the uploaded file to the destination file
			if _, err = dst.ReadFrom(file); err != nil {
				file.Close()
				dst.Close()
				controller.Fail(w, fmt.Errorf("error saving file"))
				return
			}

			file.Close()
			dst.Close()

			// Store the relative path from upload directory
			relativePath := filepath.Join(normalizedNomorSk, periode, newFileName)
			attachments = append(attachments, model.Attachments{
				Label: originalName,
				File:  relativePath,
			})
		}

		// Prepare and send response
		response := Response{
			Message:     "files uploaded",
			Attachments: attachments,
		}

		controller.Success(w, response)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData
}
