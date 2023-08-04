package v1

import (
	"context"
	"net/http"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/contract"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/dto"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/interactor/image"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func ListImagesHandler(imageStore contract.ImageStoreInteractor) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		res, err := image.New(imageStore).List(ctx, dto.ListImageRequest{})
		if err != nil {
			log.Error("ListImage - failed to fetch all images - error ", err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, res.Images)
	}
}
