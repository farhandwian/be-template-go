package usecase

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"shared/core"
	"shared/gateway"
	minioHelper "shared/helper/minio"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
)

type AiOpenPhotoReq struct {
	PhotoPath string `json:"photo_path"`
}

type AiOpenPhoto struct {
	FileName string `json:"file_name"`
	URL      string `json:"url"`
}

type AiOpenPhotoRes struct {
	Photos []AiOpenPhoto `json:"photos"`
}

type AiOpenPhotoUseCase = core.ActionHandler[AiOpenPhotoReq, AiOpenPhotoRes]

func ImplAiOpenPhotoUseCase(
	sendSSEMessageGateway gateway.SendSSEMessage,
) AiOpenPhotoUseCase {
	return func(ctx context.Context, req AiOpenPhotoReq) (*AiOpenPhotoRes, error) {
		// Inisialisasi MinIO client
		minioClient, err := minioHelper.InitMinioClient()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize MinIO client: %w", err)
		}

		bucketName := os.Getenv("MINIO_BUCKET_NAME")
		objectPrefix := filepath.Join("assets", req.PhotoPath)

		var imageList []AiOpenPhoto

		objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
			Prefix:    objectPrefix,
			Recursive: true,
		})

		for object := range objectCh {
			if object.Err != nil {
				return nil, fmt.Errorf("error listing objects: %w", object.Err)
			}

			// Hanya ambil file gambar
			if isImageFile(object.Key) {
				// Generate presigned URL for Minio object
				presignedURL, err := minioClient.PresignedGetObject(
					ctx,
					bucketName,
					object.Key,
					time.Hour*24,
					url.Values{},
				)
				if err != nil {
					return nil, fmt.Errorf("failed to generate presigned URL for %s: %w", object.Key, err)
				}

				fileName := filepath.Base(object.Key)
				fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

				imageList = append(imageList, AiOpenPhoto{
					FileName: fileName,
					URL:      presignedURL.String(),
				})
			}
		}

		if len(imageList) == 0 {
			return nil, fmt.Errorf("no images found in path: %s", req.PhotoPath)
		}

		if _, err := sendSSEMessageGateway(ctx, gateway.SendSSEMessageReq{
			Subject:      "open-photo",
			FunctionName: "openPhoto",
			Data:         imageList,
		}); err != nil {
			return nil, err
		}

		return &AiOpenPhotoRes{
			Photos: imageList,
		}, nil
	}
}

// Helper function to check image file extensions
func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}
