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

type AiOpenVideoReq struct {
	VideoPath string `json:"video_path"`
}

type AiOpenVideo struct {
	FileName string `json:"file_name"`
	URL      string `json:"url"`
}

type AiOpenVideoRes struct {
	Videos []AiOpenVideo `json:"videos"`
}

type AiOpenVideoUseCase = core.ActionHandler[AiOpenVideoReq, AiOpenVideoRes]

func ImplAiOpenVideoUseCase(
	sendSSEMessageGateway gateway.SendSSEMessage,
) AiOpenVideoUseCase {
	return func(ctx context.Context, req AiOpenVideoReq) (*AiOpenVideoRes, error) {
		minioClient, err := minioHelper.InitMinioClient()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize MinIO client: %w", err)
		}

		bucketName := os.Getenv("MINIO_BUCKET_NAME")
		objectPrefix := filepath.Join("assets", req.VideoPath)

		var videoList []AiOpenVideo

		objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
			Prefix:    objectPrefix,
			Recursive: true,
		})

		for object := range objectCh {
			if object.Err != nil {
				return nil, fmt.Errorf("error listing objects: %w", object.Err)
			}

			if isVideoFile(object.Key) {
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

				videoList = append(videoList, AiOpenVideo{
					FileName: fileName,
					URL:      presignedURL.String(),
				})
			}
		}

		if len(videoList) == 0 {
			return nil, fmt.Errorf("no images found in path: %s", req.VideoPath)
		}

		if _, err := sendSSEMessageGateway(ctx, gateway.SendSSEMessageReq{
			Subject:      "open-video",
			FunctionName: "openVideo",
			Data:         videoList,
		}); err != nil {
			return nil, err
		}

		return &AiOpenVideoRes{
			Videos: videoList,
		}, nil
	}
}
