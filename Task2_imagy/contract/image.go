package contract

import (
	"context"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/dto"
)

type ImageInteractor interface {
	Download(ctx context.Context, req dto.DownloadImageRequest) (dto.DownloadImageResponse, error)
	List(ctx context.Context, req dto.ListImageRequest) (dto.ListImageResponse, error)
}
