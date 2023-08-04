package dto

import (
	"mime/multipart"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/domain"
)

type DownloadImageFromURLRequest struct {
	URLPath   string
	LocalName string
	DstPath   string
}

type DownloadImageFromURLResponse struct {
	ImageName string
}

type ListImageRequest struct{}

type ListImageResponse struct {
	Images []domain.Image
}

type DownloadImageRequest struct {
	RootStoragePath string
	ImageName       string
}

type DownloadImageResponse struct {
	ImageAbsPath string
}

type UploadImageRequest struct {
	ImageFile *multipart.FileHeader
}

type UploadImageResponse struct {
	Image domain.Image
}
