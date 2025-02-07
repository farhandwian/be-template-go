package usecase

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"os"
	"path/filepath"
	"shared/core"
	minioHelper "shared/helper/minio"
	"strings"
)

type AiGetVideoPathReq struct{}

type AiGetVideoPathRes struct {
	Files []VideoFilePath `json:"files"`
}

type VideoFilePath struct {
	RelativePath string `json:"relative_path"`
}

type GetVideoPathUseCase = core.ActionHandler[AiGetVideoPathReq, AiGetVideoPathRes]

func ImplGetVideoPath() GetVideoPathUseCase {
	return func(ctx context.Context, req AiGetVideoPathReq) (*AiGetVideoPathRes, error) {
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

		var files []VideoFilePath
		for object := range objectCh {
			if object.Err != nil {
				return nil, fmt.Errorf("error listing objects: %w", object.Err)
			}

			if isVideoFile(object.Key) {
				relativePath := strings.TrimPrefix(object.Key, "assets/")
				files = append(files, VideoFilePath{
					RelativePath: "/" + relativePath,
				})
			}
		}

		return &AiGetVideoPathRes{Files: files}, nil
	}
}

func isVideoFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".mp4" || ext == ".avi" || ext == ".mkv" || ext == ".mov" || ext == ".flv" || ext == ".wmv"
}
