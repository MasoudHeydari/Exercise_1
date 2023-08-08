package v1

import (
	"context"
	"net/http"
	"strings"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/contract"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/interactor/image"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// DownloadImageHandler gives the ability to user to download an image based on it's name that stored in DB.
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
