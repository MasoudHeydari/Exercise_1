package contract

import (
	"context"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/adapter/store/model"
)

type ImageStoreInteractor interface {
	Create(ctx context.Context, image model.Image) (model.Image, error)
	List(ctx context.Context) ([]model.Image, error)
	DoesExit(ctx context.Context, imageName string) (model.Image, error)
}
