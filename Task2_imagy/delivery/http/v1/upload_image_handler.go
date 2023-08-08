package v1

import (
	"context"
	"net/http"
	"strings"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/contract"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/dto"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/interactor/image"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// UploadImageHandler gives the ability to user to upload an image to Imagy system.
func UploadImageHandler(imageStore contract.ImageStoreInteractor) echo.HandlerFunc {
	return func(c echo.Context) error {
		var response any
		imageFd, err := c.FormFile("image")
		if err != nil {
			log.Error("UploadImage - failed to get image from request body, request must contain a field of 'form-file' type named 'image' - error ", err.Error())
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		req := dto.UploadImageRequest{ImageFile: imageFd}
		ctx := context.Background()
		res, err := image.New(imageStore).Upload(ctx, req)
		if err != nil {
			errMsg := err.Error()
			response = echo.Map{"message": errMsg}
			log.Error("UploadImage - failed to upload image - error ", errMsg)
			statusCode := http.StatusInternalServerError
			if strings.Contains(errMsg, "404") {
				statusCode = http.StatusNotFound
			} else if strings.Contains(errMsg, "409") {
				statusCode = http.StatusConflict
			}
			return c.JSON(statusCode, response)
		}
		return c.JSON(http.StatusOK, res.Image)
	}
}
