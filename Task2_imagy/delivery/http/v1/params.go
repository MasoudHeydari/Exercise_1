package v1

import (
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/config"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/dto"
)

type DownloadImagePayload struct {
	ImageName string `param:"image_name"`
}

func (p *DownloadImagePayload) toDtoDownloadImageRequest() dto.DownloadImageRequest {
	return dto.DownloadImageRequest{
		ImageName:       p.ImageName,
		RootStoragePath: config.GetRootDownloadPath(),
	}
}
