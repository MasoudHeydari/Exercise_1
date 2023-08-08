package delivery

import (
	"fmt"
	"net"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/config"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/contract"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/delivery/http/v1"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Delivery struct {
	imageStore contract.ImageStoreInteractor
}

// New creates new Delivery.
func New(imageStore contract.ImageStoreInteractor) Delivery {
	return Delivery{imageStore: imageStore}
}

// Start initialises the Echo and starts a new web server.
func (d *Delivery) Start(conf config.Config) error {
	e := echo.New()
	e.Use(middleware.Logger())
	imagyHttpAddress := fmt.Sprintf("%s:%s", conf.HttpAddress, conf.Port)
	l, err := net.Listen("tcp4", imagyHttpAddress)
	if err != nil {
		return fmt.Errorf("faield to start http server on %s - error: %v", imagyHttpAddress, err)
	}
	e.Listener = l
	d.setupRoute(e)
	return e.Start(imagyHttpAddress)
}

// setupRoute initialises the accepted routes for Imagy.
func (d *Delivery) setupRoute(e *echo.Echo) {
	apiV1 := e.Group("api/v1")
	apiV1.GET("/images", v1.ListImagesHandler(d.imageStore))
	apiV1.GET("/images/:image_name", v1.DownloadImageHandler(d.imageStore))
	apiV1.POST("/images", v1.UploadImageHandler(d.imageStore))
}
