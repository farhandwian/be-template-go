package controller

import (
	"fmt"
	iammodel "iam/model"
	"net/http"
	"os"
	"path/filepath"
	"shared/helper"
	"strings"
)

func DownloadFileHandler(mux *http.ServeMux) helper.APIData {

	apiData := helper.APIData{
		Method:  http.MethodGet,
		Url:     "/api/download/{nomor_sk}/{periode_pengambilan_sda}/{filename}",
		Access:  iammodel.DEFAULT_OPERATION,
		Summary: "Download file",
		Tag:     "Laporan Perizinan",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		// Get path parameters
		nomorSK := r.PathValue("nomor_sk")
		periodePengambilan := r.PathValue("periode_pengambilan_sda")
		filename := r.PathValue("filename")

		// Validate all parameters are present
		if nomorSK == "" || periodePengambilan == "" || filename == "" {
			http.Error(w, "Missing required path parameters", http.StatusBadRequest)
			return
		}

		// Get root path from environment
		rootPath := os.Getenv("FILE_UPLOAD")
		if rootPath == "" {
			http.Error(w, "File storage path not configured", http.StatusInternalServerError)
			return
		}

		// Construct the file path
		filePath := filepath.Join(rootPath, nomorSK, periodePengambilan, filename)

		// Security checks

		// 1. Clean the path and ensure it's still under rootPath
		cleanPath := filepath.Clean(filePath)
		if !strings.HasPrefix(cleanPath, filepath.Clean(rootPath)) {
			http.Error(w, "Invalid file path", http.StatusForbidden)
			return
		}

		// 2. Check for any suspicious patterns
		if strings.Contains(filePath, "..") {
			http.Error(w, "Invalid file path", http.StatusForbidden)
			return
		}

		// 3. Verify file exists and is a regular file
		fileInfo, err := os.Stat(cleanPath)
		if err != nil {
			if os.IsNotExist(err) {
				http.Error(w, "File not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Error accessing file", http.StatusInternalServerError)
			return
		}

		// Ensure it's a regular file, not a directory or symlink
		if !fileInfo.Mode().IsRegular() {
			http.Error(w, "Invalid file type", http.StatusForbidden)
			return
		}

		// Open the file
		file, err := os.Open(cleanPath)
		if err != nil {
			http.Error(w, "Error opening file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Get the mime type of the file
		mimeType := http.DetectContentType(make([]byte, 512))

		// Set response headers with more specific content disposition
		w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
		w.Header().Set("Content-Type", mimeType)
		w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

		// Additional headers to prevent caching
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		// Stream the file to the response
		http.ServeContent(w, r, filename, fileInfo.ModTime(), file)
	}

	mux.HandleFunc(apiData.GetMethodUrl(), handler)
	return apiData

}
