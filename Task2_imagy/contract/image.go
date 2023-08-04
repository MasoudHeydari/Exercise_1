package contract

import (
	"context"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/dto"
)

type ImageInteractor interface {
	DownloadFromURL(ctx context.Context, req dto.DownloadImageFromURLRequest) (dto.DownloadImageFromURLResponse, error)
	List(ctx context.Context, req dto.ListImageRequest) (dto.ListImageResponse, error)
}
