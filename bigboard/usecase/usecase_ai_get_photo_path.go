package usecase

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"os"
	"shared/core"
	minioHelper "shared/helper/minio"
	"strings"
)

type AiGetPhotoPathReq struct{}

type AiGetPhotoPathRes struct {
	Files []FilePath `json:"files"`
}

type FilePath struct {
	RelativePath string `json:"relative_path"`
}

type GetPhotoPathUseCase = core.ActionHandler[AiGetPhotoPathReq, AiGetPhotoPathRes]

func ImplGetPhotoPath() GetPhotoPathUseCase {
	return func(ctx context.Context, req AiGetPhotoPathReq) (*AiGetPhotoPathRes, error) {
		minioClient, err := minioHelper.InitMinioClient()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize MinIO client: %w", err)
		}

		bucketName := os.Getenv("MINIO_BUCKET_NAME")

		exists, err := minioClient.BucketExists(ctx, bucketName)
		if err != nil {
			return nil, fmt.Errorf("error checking if bucket exists: %w", err)
		}
		if !exists {
			return nil, fmt.Errorf("bucket %s does not exist", bucketName)
		}

		objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
			Recursive: true,
		})

		var files []FilePath
		for object := range objectCh {
			if object.Err != nil {
				return nil, fmt.Errorf("error listing objects: %w", object.Err)
			}

			if isImageFile(object.Key) {
				// Remove the "/assets" prefix
				relativePath := strings.TrimPrefix(object.Key, "assets/")
				files = append(files, FilePath{
					RelativePath: "/" + relativePath,
				})
			}
		}

		return &AiGetPhotoPathRes{Files: files}, nil
	}
}
