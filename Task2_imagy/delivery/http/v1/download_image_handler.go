package v1

import (
	"context"
	"net/http"
	"strings"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/config"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/contract"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/dto"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/interactor/image"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

func DownloadImageHandler(imageStore contract.ImageStoreInteractor) echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := new(DownloadImagePayload)
		err := c.Bind(payload)
		if err != nil {
			log.Error("DownloadImage - failed to bind input payload - error ", err.Error())
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		req := payload.toDtoDownloadImageRequest()
		ctx := context.Background()
		res, err := image.New(imageStore).Download(ctx, req)
		if err != nil {
			errMsg := err.Error()
			log.Error("DownloadImage - failed to fetch all images - error ", errMsg)
			statusCode := http.StatusInternalServerError
			if strings.Contains(errMsg, "404") {
				statusCode = http.StatusNotFound
			}
			return c.JSON(statusCode, err.Error())
		}
		return c.File(res.ImageAbsPath)
	}
}
